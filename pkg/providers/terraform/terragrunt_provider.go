package terraform

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"sort"

	"github.com/kballard/go-shellquote"

	"github.com/pkg/errors"

	"github.com/infracost/infracost/pkg/clierror"
	"github.com/infracost/infracost/pkg/config"
	"github.com/infracost/infracost/pkg/logging"
	"github.com/infracost/infracost/pkg/schema"
)

var defaultTerragruntBinary = "terragrunt"
var minTerragruntVer = "v0.28.1"

type TerragruntProvider struct {
	ctx             *config.ProjectContext
	Path            string
	TerragruntFlags string
	*DirProvider
	includePastResources bool
}

type TerragruntInfo struct {
	ConfigPath string
	WorkingDir string
}

type terragruntProjectDirs struct {
	ConfigDir  string
	WorkingDir string
}

func NewTerragruntProvider(ctx *config.ProjectContext, includePastResources bool) schema.Provider {
	dirProvider := NewDirProvider(ctx, includePastResources).(*DirProvider)

	terragruntBinary := ctx.ProjectConfig.TerraformBinary
	if terragruntBinary == "" {
		terragruntBinary = defaultTerragruntBinary
	}

	dirProvider.TerraformBinary = terragruntBinary
	dirProvider.IsTerragrunt = true

	return &TerragruntProvider{
		ctx:                  ctx,
		DirProvider:          dirProvider,
		Path:                 ctx.ProjectConfig.Path,
		TerragruntFlags:      ctx.ProjectConfig.TerragruntFlags,
		includePastResources: includePastResources,
	}
}

func (p *TerragruntProvider) ProjectName() string {
	return config.CleanProjectName(p.ctx.ProjectConfig.Path)
}

func (p *TerragruntProvider) VarFiles() []string {
	return nil
}

func (p *TerragruntProvider) RelativePath() string {
	return p.ctx.ProjectConfig.Path
}

func (p *TerragruntProvider) Context() *config.ProjectContext { return p.ctx }

func (p *TerragruntProvider) Type() string {
	return "terragrunt_cli"
}

func (p *TerragruntProvider) DisplayType() string {
	return "Terragrunt CLI"
}

func (p *TerragruntProvider) AddMetadata(metadata *schema.ProjectMetadata) {
	metadata.ConfigSha = p.ctx.ProjectConfig.ConfigSha

	basePath := p.ctx.ProjectConfig.Path
	if p.ctx.RunContext.Config.ConfigFilePath != "" {
		basePath = filepath.Dir(p.ctx.RunContext.Config.ConfigFilePath)
	}

	modulePath, err := filepath.Rel(basePath, metadata.Path)
	if err == nil && modulePath != "" && modulePath != "." {
		logging.Logger.Debug().Msgf("Calculated relative terraformModulePath for %s from %s", basePath, metadata.Path)
		metadata.TerraformModulePath = modulePath
	}

	metadata.TerraformWorkspace = p.ctx.ProjectConfig.TerraformWorkspace
}

func (p *TerragruntProvider) LoadResources(usage schema.UsageMap) ([]*schema.Project, error) {
	// We want to run Terragrunt commands from the config dirs
	// Terragrunt internally runs Terraform in the working dirs, so we need to be aware of these
	// so we can handle reading and cleaning up the generated plan files.
	projectDirs, err := p.getProjectDirs()
	if err != nil {
		return []*schema.Project{}, err
	}

	var outs [][]byte

	if p.UseState {
		outs, err = p.generateStateJSONs(projectDirs)
	} else {
		outs, err = p.generatePlanJSONs(projectDirs)
	}
	if err != nil {
		return []*schema.Project{}, err
	}

	projects := make([]*schema.Project, 0, len(projectDirs))

	logging.Logger.Debug().Msg("Extracting only cost-related params from terragrunt plan")
	for i, projectDir := range projectDirs {
		projectPath := projectDir.ConfigDir
		// attempt to convert project path to be relative to the top level provider path
		if absPath, err := filepath.Abs(p.ctx.ProjectConfig.Path); err == nil {
			if relProjectPath, err := filepath.Rel(absPath, projectPath); err == nil {
				projectPath = filepath.Join(p.ctx.ProjectConfig.Path, relProjectPath)
			}
		}

		metadata := schema.DetectProjectMetadata(projectPath)
		metadata.Type = p.Type()
		p.AddMetadata(metadata)
		name := p.ctx.ProjectConfig.Name
		if name == "" {
			name = metadata.GenerateProjectName(p.ctx.RunContext.VCSMetadata.Remote, p.ctx.RunContext.IsCloudEnabled())
		}

		project := schema.NewProject(name, metadata)

		parser := NewParser(p.ctx, p.includePastResources)
		j, _ := StripSetupTerraformWrapper(outs[i])
		parsedConf, err := parser.parseJSON(j, usage)
		if err != nil {
			return projects, errors.Wrap(err, "Error parsing Terraform JSON")
		}

		project.AddProviderMetadata(parsedConf.ProviderMetadata)

		project.HasDiff = !p.UseState
		if project.HasDiff {
			project.PartialPastResources = parsedConf.PastResources
		}
		project.PartialResources = parsedConf.CurrentResources

		projects = append(projects, project)
	}

	return projects, nil
}

func (p *TerragruntProvider) getProjectDirs() ([]terragruntProjectDirs, error) {
	logging.Logger.Debug().Msg("Running terragrunt run-all terragrunt-info")

	terragruntFlags, err := shellquote.Split(p.TerragruntFlags)
	if err != nil {
		return []terragruntProjectDirs{}, errors.Wrap(err, "Error parsing terragrunt flags")
	}

	opts := &CmdOptions{
		TerraformBinary: p.TerraformBinary,
		Dir:             p.Path,
		Flags:           terragruntFlags,
	}
	out, err := Cmd(opts, "run-all", "--terragrunt-ignore-external-dependencies", "terragrunt-info")
	if err != nil {
		err = p.buildTerraformErr(err, false)

		msg := "terragrunt run-all terragrunt-info failed"
		return []terragruntProjectDirs{}, clierror.NewCLIError(fmt.Errorf("%s: %s", msg, err), msg)
	}

	var jsons [][]byte

	jsonStart := bytes.IndexByte(out, '{') // ignore anything that comes before the json (e.g. unexpected logging to stdout by tgenv)
	if jsonStart >= 0 {
		out = out[jsonStart:]

		jsons = bytes.SplitAfter(out, []byte{'}', '\n'})
		if len(jsons) > 1 {
			jsons = jsons[:len(jsons)-1]
		}
	}

	dirs := make([]terragruntProjectDirs, 0, len(jsons))

	for _, j := range jsons {
		var info TerragruntInfo
		err = json.Unmarshal(j, &info)
		if err != nil {
			return dirs, fmt.Errorf("error unmarshalling terragrunt-info JSON: %w", err)
		}

		dirs = append(dirs, terragruntProjectDirs{
			ConfigDir:  filepath.Dir(info.ConfigPath),
			WorkingDir: info.WorkingDir,
		})
	}

	// Sort the dirs so they are consistent in the output
	sort.Slice(dirs, func(i, j int) bool {
		return dirs[i].ConfigDir < dirs[j].ConfigDir
	})

	return dirs, nil
}

func (p *TerragruntProvider) generateStateJSONs(projectDirs []terragruntProjectDirs) ([][]byte, error) {
	err := p.checks()
	if err != nil {
		return [][]byte{}, err
	}

	outs := make([][]byte, 0, len(projectDirs))

	for _, projectDir := range projectDirs {
		opts, err := p.buildCommandOpts(projectDir.ConfigDir)
		if err != nil {
			return [][]byte{}, err
		}

		terragruntFlags, err := shellquote.Split(p.TerragruntFlags)
		if err != nil {
			return [][]byte{}, errors.Wrap(err, "Error parsing terragrunt flags")
		}
		opts.Flags = terragruntFlags

		if opts.TerraformConfigFile != "" {
			defer os.Remove(opts.TerraformConfigFile)
		}

		out, err := p.runShow(opts, "", false)
		if err != nil {
			return outs, err
		}

		// ignore anything that comes before the json (e.g. unexpected logging to stdout by tgenv)
		jsonStart := bytes.IndexByte(out, '{')
		if jsonStart >= 0 {
			out = out[jsonStart:]
		}

		outs = append(outs, out)
	}

	return outs, nil
}

func (p *TerragruntProvider) generatePlanJSONs(projectDirs []terragruntProjectDirs) ([][]byte, error) {
	err := p.checks()
	if err != nil {
		return [][]byte{}, err
	}

	opts, err := p.buildCommandOpts(p.Path)
	if err != nil {
		return [][]byte{}, err
	}

	terragruntFlags, err := shellquote.Split(p.TerragruntFlags)
	if err != nil {
		return [][]byte{}, errors.Wrap(err, "Error parsing terragrunt flags")
	}
	opts.Flags = terragruntFlags

	if opts.TerraformConfigFile != "" {
		defer os.Remove(opts.TerraformConfigFile)
	}

	logging.Logger.Debug().Msg("Running terragrunt run-all plan")

	planFile, planJSON, err := p.runPlan(opts, true)
	defer func() {
		err := cleanupPlanFiles(projectDirs, planFile)
		if err != nil {
			logging.Logger.Warn().Msgf("Error cleaning up plan files: %v", err)
		}
	}()

	if err != nil {
		return [][]byte{}, err
	}

	if len(planJSON) > 0 {
		return [][]byte{planJSON}, nil
	}

	outs := make([][]byte, 0, len(projectDirs))
	logging.Logger.Debug().Msg("Running terragrunt show")

	for _, projectDir := range projectDirs {
		opts, err := p.buildCommandOpts(projectDir.ConfigDir)
		if err != nil {
			return [][]byte{}, err
		}
		if opts.TerraformConfigFile != "" {
			defer os.Remove(opts.TerraformConfigFile)
		}

		out, err := p.runShow(opts, filepath.Join(projectDir.WorkingDir, planFile), false)
		if err != nil {
			return outs, err
		}

		// ignore anything that comes before the json (e.g. unexpected logging to stdout by tgenv)
		jsonStart := bytes.IndexByte(out, '{')
		if jsonStart >= 0 {
			out = out[jsonStart:]
		}

		outs = append(outs, out)
	}

	return outs, nil
}

func cleanupPlanFiles(projectDirs []terragruntProjectDirs, planFile string) error {
	if planFile == "" {
		return nil
	}

	for _, projectDir := range projectDirs {
		err := os.Remove(filepath.Join(projectDir.WorkingDir, planFile))
		if err != nil {
			return err
		}
	}

	return nil
}

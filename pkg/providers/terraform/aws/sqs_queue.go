package aws

import (
	"github.com/infracost/infracost/pkg/resources/aws"
	"github.com/infracost/infracost/pkg/schema"
)

func getSQSQueueRegistryItem() *schema.RegistryItem {
	return &schema.RegistryItem{
		Name:      "aws_sqs_queue",
		CoreRFunc: NewSQSQueue,
	}
}

func NewSQSQueue(d *schema.ResourceData) schema.CoreResource {
	r := &aws.SQSQueue{
		Address:   d.Address,
		Region:    d.Get("region").String(),
		FifoQueue: d.Get("fifo_queue").Bool(),
	}
	return r
}

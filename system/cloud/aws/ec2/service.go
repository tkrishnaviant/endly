package ec2

import (
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/viant/endly"
	"github.com/viant/endly/system/cloud/aws"
	"log"
)

const (
	//ServiceID aws iam service id.
	ServiceID = "aws/ec2"
)

//no operation service
type service struct {
	*endly.AbstractService
}

func (s *service) registerRoutes() {
	client := &ec2.EC2{}
	routes, err := aws.BuildRoutes(client, getClient)
	if err != nil {
		log.Printf("unable register service %v actions: %v\n", ServiceID, err)
		return
	}
	for _, route := range routes {
		route.OnRawRequest = setClient
		s.Register(route)
	}
}

//New creates a new AWS IAM service.
func New() endly.Service {
	var result = &service{
		AbstractService: endly.NewAbstractService(ServiceID),
	}
	result.AbstractService.Service = result
	result.registerRoutes()
	return result
}

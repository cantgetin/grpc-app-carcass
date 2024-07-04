package service

import (
	"context"
	"grpc-carcass/pkg/api"
)

type ServiceImpl struct {
	api.UnimplementedSampleServiceServer
}

func New() *ServiceImpl {
	return &ServiceImpl{}
}

func (s *ServiceImpl) SampleHandler(ctx context.Context, req *api.SampleRequest) (*api.SampleResponse, error) {
	return nil, nil
}

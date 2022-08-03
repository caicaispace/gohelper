package hello

import (
	context "context"
)

type Service struct{}

func New() *Service {
	return &Service{}
}

func (h *Service) SayHello(ctx context.Context, in *SayHelloReq) (*SayHelloRsp, error) {
	return &SayHelloRsp{Message: "Hello " + in.Name}, nil
}

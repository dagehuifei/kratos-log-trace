package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"

	v1 "kratos-log-trace/api/helloworld/v1"
	"kratos-log-trace/internal/biz"
)

// GreeterService is a greeter service.
type GreeterService struct {
	v1.UnimplementedGreeterServer

	uc *biz.GreeterUsecase
}

// NewGreeterService new a greeter service.
func NewGreeterService(uc *biz.GreeterUsecase) *GreeterService {
	return &GreeterService{uc: uc}
}

// SayHello implements helloworld.GreeterServer.
func (s *GreeterService) SayHello(ctx context.Context, in *v1.HelloRequest) (*v1.HelloReply, error) {
	g, err := s.uc.CreateGreeter(ctx, &biz.Greeter{Hello: in.Name})
	if err != nil {
		return nil, err
	}
	log.Context(ctx).Info("打印TraceID")

	return &v1.HelloReply{Message: "Hello " + g.Hello}, nil
}

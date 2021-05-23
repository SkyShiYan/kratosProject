package service

import (
	"context"
	"strconv"

	v1 "spaco_go/api/helloworld/v1"
	"spaco_go/internal/biz"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
)

// GreeterService is a greeter service.
type GreeterService struct {
	v1.UnimplementedGreeterServer

	uc  *biz.GreeterUsecase
	log *log.Helper
}

// NewGreeterService new a greeter service.
func NewGreeterService(uc *biz.GreeterUsecase, logger log.Logger) *GreeterService {
	return &GreeterService{uc: uc, log: log.NewHelper("service/greeter", logger)}
}

// SayHello implements helloworld.GreeterServer
func (s *GreeterService) SayHello(ctx context.Context, in *v1.HelloRequest) (*v1.HelloReply, error) {
	age, err := s.uc.Get(ctx, in.GetName())
	if err != nil {
		if errors.IsNotFound(err) {
			return &v1.HelloReply{Message: "Sorry not registed."}, nil
		}
	}

	s.log.Infof("SayHello Received: %v", in.GetName())
	if in.GetName() == "error" {
		return nil, errors.NotFound("greeter", v1.ErrorReason_USER_NOT_FOUND.String(), in.GetName())
	}
	return &v1.HelloReply{Message: "Hello " + in.GetName() + "is " + strconv.Itoa(age)}, nil
}

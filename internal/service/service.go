package service

import (
	v1 "kratos-realworld/api/realworld/v1"
	"kratos-realworld/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewRealWorldService)

// RealWorldService is a RealWorld service.
type RealWorldService struct {
	v1.UnimplementedRealworldServer
	uc  *biz.UserUseCase
	log *log.Helper
}

// NewRealWorldService new RealWorld service.
func NewRealWorldService(uc *biz.UserUseCase, logger log.Logger) *RealWorldService {
	return &RealWorldService{uc: uc, log: log.NewHelper(logger)}
}

package service

import (
	"github.com/google/wire"

	"github.com/superproj/superproj-examples/miniusercenter/internal/usercenter/biz"
	v1 "github.com/superproj/superproj-examples/miniusercenter/pkg/api/miniusercenter/v1"
)

// ProviderSet is a set of service providers, used for dependency injection.
var ProviderSet = wire.NewSet(NewUserCenterService)

// UserCenterService is a struct that implements the v1.UnimplementedUserCenterServer interface
// and holds the business logic, represented by a IBiz instance.
type UserCenterService struct {
	v1.UnimplementedUserCenterServer          // Embeds the generated UnimplementedUserCenterServer struct.
	biz                              biz.IBiz // A factory for creating business logic components.
}

// NewUserCenterService is a constructor function that takes a IBiz instance
// as an input and returns a new UserCenterService instance.
func NewUserCenterService(biz biz.IBiz) *UserCenterService {
	return &UserCenterService{biz: biz}
}

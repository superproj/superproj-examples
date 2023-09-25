package biz

import (
	"github.com/google/wire"

	"github.com/superproj/superproj-examples/miniusercenter/internal/usercenter/store"
)

// ProviderSet contains providers for creating instances of the biz struct.
var ProviderSet = wire.NewSet(NewBiz, wire.Bind(new(IBiz), new(*biz)))

// IBiz defines a set of methods for returning interfaces that the biz struct implements.
type IBiz interface {
	Users() UserBiz
}

type biz struct {
	ds store.IStore
}

// NewBiz returns a pointer to a new instance of the biz struct.
func NewBiz(ds store.IStore) *biz {
	return &biz{ds: ds}
}

// Users returns a new instance of the UserBiz interface.
func (b *biz) Users() UserBiz {
	return newUserBiz(b.ds)
}

package store

import (
	"fmt"
	"sync"

	"github.com/superproj/superproj-examples/miniusercenter/internal/usercenter/model"
)

var (
	once sync.Once
	us   *userStore
)

type userCache struct {
	mux   sync.Mutex
	cache map[string]*model.UserM
}

func newUserCache() *userCache {
	return &userCache{cache: make(map[string]*model.UserM)}
}

func (c *userCache) Save(u *model.UserM) error {
	c.mux.Lock()
	defer c.mux.Unlock()

	if _, ok := c.cache[u.Username]; ok {
		return fmt.Errorf("%s already exist", u.Username)
	}

	c.cache[u.Username] = u

	return nil
}

package store

import (
	"github.com/google/wire"
	"gorm.io/gorm"
)

// ProviderSet is a Wire provider set that initializes new datastore instances
// and binds the IStore interface to the actual datastore type.
var ProviderSet = wire.NewSet(NewStore, wire.Bind(new(IStore), new(*datastore)))

// IStore is an interface that represents methods
// required to be implemented by a Store implementation.
type IStore interface {
	Users() UserStore
}

// datastore is an implementation of IStore that provides methods
// to perform operations on a database using gorm library.
type datastore struct {
	// core is the main database instance.
	// The `core` name indicates this is the main database.
	db *gorm.DB
}

// Ensure datastore implements IStore.
var _ IStore = (*datastore)(nil)

// NewStore initializes a new datastore instance using the provided DB gorm instance.
// It also creates a singleton instance for the datastore.
func NewStore(db *gorm.DB) *datastore {
	return &datastore{db}
}

// Users returns an initialized instance of UserStore.
func (ds *datastore) Users() UserStore {
	return newUserStore(ds.db)
}

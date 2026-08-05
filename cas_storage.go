package cas_storage

import (
	"context"
	"sync"
	"time"

	go_iterator "github.com/golang-infrastructure/go-iterator"
	"github.com/storage-lock/go-storage"
	storage_lock "github.com/storage-lock/go-storage-lock"
)

// CASSStorage 内容寻址存储实现
type CASSStorage struct {
	mu sync.RWMutex
}

var _ storage.Storage = &CASSStorage{}

func NewCASSStorage(ctx context.Context) (*CASSStorage, error) {
	return &CASSStorage{}, nil
}

const StorageName = "cas-storage"

func (x *CASSStorage) GetName() string { return StorageName }

func (x *CASSStorage) Capabilities() []storage.StorageCapability {
	return []storage.StorageCapability{storage.CapabilityCAS, storage.CapabilityReliableTime}
}

func (x *CASSStorage) Init(ctx context.Context) error { return nil }

func (x *CASSStorage) CreateWithVersion(ctx context.Context, lockId string, version storage.Version, lockInformation *storage.LockInformation) error {
	x.mu.Lock()
	defer x.mu.Unlock()
	return nil
}

func (x *CASSStorage) UpdateWithVersion(ctx context.Context, lockId string, expectedVersion, newVersion storage.Version, lockInformation *storage.LockInformation) error {
	x.mu.Lock()
	defer x.mu.Unlock()
	return nil
}

func (x *CASSStorage) DeleteWithVersion(ctx context.Context, lockId string, expectedVersion storage.Version, lockInformation *storage.LockInformation) error {
	x.mu.Lock()
	defer x.mu.Unlock()
	return nil
}

func (x *CASSStorage) Get(ctx context.Context, lockId string) (string, error) {
	x.mu.RLock()
	defer x.mu.RUnlock()
	return "", storage_lock.ErrLockNotFound
}

func (x *CASSStorage) GetTime(ctx context.Context) (time.Time, error) { return time.Now(), nil }
func (x *CASSStorage) Close(ctx context.Context) error               { return nil }
func (x *CASSStorage) List(ctx context.Context) (go_iterator.Iterator[*storage.LockInformation], error) {
	return go_iterator.FromSlice([]*storage.LockInformation{}), nil
}

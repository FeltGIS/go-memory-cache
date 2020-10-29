package memory_cache

import (
	"math"
	"time"
)

const Infinite = math.MaxInt32
const Forever = 87660 * time.Hour

type (
	Option struct {
		MaxEntrySize        int
		MaxEntriesKey       int
		MaxEntriesInWindow  int64
		MaxPercentageMemory float64
		OnRemove            func(key string, value interface{})
		OnRemoveWithReason  func(key string, reason string)
		OnMemoryExceed      func(memoryUsedPercentage float64, maxMemoryPercentage float64, memoryUsed float64)
	}
	Cache interface {
		Set(key string, value interface{}, ttl *time.Duration) error
		Get(key string) (interface{}, error)
		Remove(key string) error
		Truncate() error
		Len() int
		Size() uintptr
		Keys() []string
	}
)

func initializeOption(option Option) Option {
	if option.MaxEntriesKey == 0 {
		option.MaxEntriesKey = Infinite
	}
	if option.MaxEntriesInWindow == 0 {
		option.MaxEntriesInWindow = 2 * 1024 * 1024 * 1024
	}
	if option.MaxEntrySize == 0 {
		option.MaxEntrySize = 1024 * 1024
	}
	if option.MaxPercentageMemory == 0 {
		option.MaxPercentageMemory = 95
	}
	return option
}

func NewWithOption(option Option) (Cache, error) {
	return &cache{
		option: initializeOption(option),
		data:   make(map[string]interface{}),
		size:   uintptr(0),
	}, nil
}

func New() (Cache, error) {
	return NewWithOption(Option{})
}

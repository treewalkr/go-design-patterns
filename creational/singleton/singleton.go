package singleton

import (
	"sync"
)

// Singleton is the type that will have a single instance.
type Singleton struct {
	Value string
}

// Idiomatic Singleton (using sync.Once)
var onceInstance *Singleton
var once sync.Once

// GetOnceInstance provides access to the single instance using sync.Once.
func GetOnceInstance() *Singleton {
	once.Do(func() {
		onceInstance = &Singleton{Value: "Initial Value (Once)"}
	})
	return onceInstance
}

// Manual Singleton (using sync.Mutex)
var manualInstance *Singleton
var mu sync.Mutex

// GetManualInstance provides access to the single instance using sync.Mutex.
func GetManualInstance() *Singleton {
	mu.Lock()
	defer mu.Unlock()

	if manualInstance == nil {
		manualInstance = &Singleton{Value: "Initial Value (Manual)"}
	}
	return manualInstance
}

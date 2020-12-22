package pattern2

import (
	"fmt"
	"sync"
)
// Expose only the public methods that work on a package local singleton

// Singleton
var instance = &myStore{
	kvMap: make(map[string]string),
}

type myStore struct {
	lock sync.RWMutex
	kvMap map[string]string
}

func GetValue(key string) string {
	instance.lock.RLock()
	defer instance.lock.RUnlock()
	return instance.kvMap[key]
}

func SetValue(key, value string) {
	instance.lock.Lock()
	defer instance.lock.Unlock()
	instance.kvMap[key] = value
}

func Print() {
	instance.lock.RLock()
	defer instance.lock.RUnlock()
	for k, v := range instance.kvMap {
		fmt.Printf("Key = %s, Value = %s\n", k, v)
	}
}
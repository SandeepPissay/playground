package pattern1

import (
	"fmt"
	"sync"
)

var lock sync.Mutex
var singleton *myFooImpl

// Hide the singleton instance behind this interface.
type Foo interface {
	GetValue(key string) string
	SetValue(key, value string)
	Print()
}

// Return the singleton instance (create if not created already).
func GetFoo() Foo {
	if singleton != nil {
		return singleton
	}
	lock.Lock()
	defer lock.Unlock()

	singleton = & myFooImpl{
		kvMap: make(map[string]string),
	}
	return singleton
}

type myFooImpl struct {
	lock sync.RWMutex
	kvMap map[string]string
}

func (f *myFooImpl) GetValue(key string) string {
	f.lock.RLock()
	defer f.lock.RUnlock()
	return f.kvMap[key]
}

func (f *myFooImpl) SetValue(key, value string) {
	f.lock.Lock()
	defer f.lock.Unlock()
	f.kvMap[key] = value
}

func (f *myFooImpl) Print() {
	f.lock.RLock()
	defer f.lock.RUnlock()
	for k, v := range f.kvMap {
		fmt.Printf("Key = %s, Value = %s\n", k, v)
	}
}
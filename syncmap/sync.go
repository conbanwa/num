package syncmap

import (
	"fmt"
	"sync"
)

func Len(sm *sync.Map) (l int) {
	sm.Range(func(key, value any) bool { l++; return true })
	return
}

func ValueStrict[T any](sm *sync.Map, k any) T {
	value, ok := sm.Load(k)
	if !ok {
		panic(fmt.Sprintf("no such key: %v for %T map", k, *new(T)))
	}
	return value.(T)
}

func Value[T any](sm *sync.Map, k any) (res T) {
	value, ok := sm.Load(k)
	if !ok {
		return res
	}
	return value.(T)
}

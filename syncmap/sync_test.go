package syncmap

import (
	"sync"
)

func Len(sm *sync.Map) (l int) {
	sm.Range(func(key, value any) bool { l++; return true })
	return
}

func ValueStrict[T any](sm *sync.Map, k any) (res T) {
	value, ok := sm.Load(k)
	if !ok {
		panic("inferred wrong type")
	}
	res, ok = value.(T)
	if !ok {
		panic("inferred wrong type")
	}
	return
}

func Value[T any](sm *sync.Map, k any) (res T) {
	value, _ := sm.Load(k)
	return value.(T)
}

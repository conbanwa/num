package syncmap

import (
	"sync"
)

func Len(sm *sync.Map) (l int) {
	sm.Range(func(key, value any) bool { l++; return true })
	return
}

func ValueStrict[T any](sm *sync.Map, k any) T {
	value, ok := sm.Load(k)
	if !ok {
		panic("no such key")
	}
	res, ok := value.(T)
	if !ok {
		panic("inferred wrong type")
	}
	return res
}

func Value(sm *sync.Map, k any) any {
	value, _ := sm.Load(k)
	return value
}

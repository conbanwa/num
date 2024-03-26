package syncmap

import (
	"reflect"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSyncMapValue(t *testing.T) {
	var sm sync.Map
	assert.Equal(t, 0, Len(&sm))
	sm.Store(1, []int{2})
	sm.Store(3, "55")
	assert.Equal(t, 2, Len(&sm))
	t.Run("0", func(t *testing.T) {
		if gotRes := ValueStrict[[]int](&sm, 1); !reflect.DeepEqual(gotRes, []int{2}) || !reflect.DeepEqual(Value[[]int](&sm, 1), []int{2}) {
			t.Errorf("ValueStrict() = %v, want %v", gotRes, []int{})
		}
		if gotRes := Value[[]int](&sm, 11); !reflect.DeepEqual(gotRes, []int(nil)) {
			t.Errorf("ValueStrict() = %v, want %v", gotRes, []int(nil))
		}
		if gotRes := ValueStrict[string](&sm, 3); !reflect.DeepEqual(gotRes, "55") {
			t.Errorf("ValueStrict() = %v, want %v", gotRes, "")
		}
	})
}

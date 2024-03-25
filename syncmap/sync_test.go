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
	assert.Equal(t, 1, Len(&sm))
	t.Run("tt.name", func(t *testing.T) {
		if gotRes := ValueStrict[[]int](&sm, 1); !reflect.DeepEqual(gotRes, []int{2}) || !reflect.DeepEqual(Value(&sm, 1), []int{2}) {
			t.Errorf("ValueStrict() = %v, want %v", gotRes, []int{})
		}
	})
}

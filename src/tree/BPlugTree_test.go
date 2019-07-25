package tree

import (
	"reflect"
	"testing"
)
//插入新数据到空B+树上面
func TestInsertNilRoot(t *testing.T) {
	tree := NewBPlugTree()
	key := 1
	value := []byte("test")
	err := tree.Insert(key, value)
	if err != nil {
		t.Errorf("%s", err)
	}
	r, err := tree.Find(key, false)
	if err != nil {
		t.Errorf("%s\n", err)
	}
	if r == nil {
		t.Errorf("returned nil \n")
	}
	if !reflect.DeepEqual(r.Value, value) {
		t.Errorf("expected %v and got %v \n", value, r.Value)
	}
}



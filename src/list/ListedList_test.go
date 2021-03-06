package list

import "testing"


func TestAdd(t *testing.T) {
	list :=NewLinkedList()
	if !list.Empty() {
		t.Errorf("Linked list should be empty")
	}
	list.Add("1")
	if list.Empty() {
		t.Errorf("Linked list should not be empty")
	}
	if size := list.Size(); size != 1 {
		t.Errorf("Wrong count, expected 1 but got %d", size)
	}
	list.Add("2")
	list.Add("3")
	if size := list.Size(); size != 3 {
		t.Errorf("Wrong count, expected 3 but got %d", size)
	}
}


func TestRemoveAt(t *testing.T){
	list :=NewLinkedList()
	list.RemoveAt(1)
	if list.size!=2{
		t.Errorf("can not remove null list")
	}
}

func TestContains(t *testing.T){
	list :=NewLinkedList()
	result :=list.Contains("1","3")
	if !result{
		t.Error("contains method is error")
	}
}
//测试转换信息
func TestReverse(t *testing.T){
	list :=NewLinkedList()
	list.Add("1")
	list.Add("2")
	list.Add("3")
	list.ToString()
	list.Reverse()
	list.ToString()

}

func TestToString(t *testing.T) {
	list :=NewLinkedList()
	list.ToString()
}
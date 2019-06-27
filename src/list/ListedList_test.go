package list

import "testing"



var list LinkedList

func TestAdd(t *testing.T) {
	if !list.isEmpty() {
		t.Errorf("Linked list should be empty")
	}
	list.addObj(1)
	if list.isEmpty() {
		t.Errorf("Linked list should not be empty")
	}
	if size := list.getSize(); size != 1 {
		t.Errorf("Wrong count, expected 1 but got %d", size)
	}
	list.addObj(2)
	list.addObj(3)
	if size := list.getSize(); size != 3 {
		t.Errorf("Wrong count, expected 3 but got %d", size)
	}
}


func TestRemoveAt(t *testing.T){

	list.RemoveAt(1)
	if list.size!=2{
		t.Errorf("can not remove null list")
	}
}

func TestContains(t *testing.T){
	result :=list.contains(1,3)
	if !result{
		t.Error("contains method is error")
	}
}

func TestToString(t *testing.T) {
	list.toString()
}
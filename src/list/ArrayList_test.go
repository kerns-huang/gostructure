package list

import "testing"

func TestArrayAdd(t *testing.T){
	list :=new(ArrayList)
	list.Add(1)
	if !list.contains(1){
		t.Error("wrong add ,list need contain 1")
	}
	list.Remove(1)

}



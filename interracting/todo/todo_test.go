package todo_test

import (
	"testing"

	"github.com/didate/interracting/todo"
)

func TestAdd(t *testing.T){

	l :=todo.List{}
	task := "Here is a task"
	l.Add(task)

	if l[0].Task != task{
		t.Errorf("Expected => %v, got =>  %v instead", task, l[0].Task)
	}

	if(len(l) !=1){
		t.Errorf("Expected Len 1, got %d", len(l))
	}
}

func TestCompleted(t *testing.T){
	l :=todo.List{}
	l.Add("Here is a task")
	l.Add("Here is a second task")
	i:=2
	l.Complete(i)
	if !l[i-1].Done{
		t.Errorf("Expected true, got %v instead", l[i-1].Done)
	}
}

func TestDelete(t *testing.T){
	l :=todo.List{}
	l.Add("Here is a task")
	l.Add("Here is a second task")
	third :="Here is a third task"
	l.Add(third)
	length := len(l)
	i:=2
	l.Delete(i)

	if len(l) >= length {
		t.Errorf("len after delete (%d) have to be less than inital len (%d)", len(l), length)
	}

	
}


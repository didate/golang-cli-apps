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
	taskName := "New Task"
	l.Add(taskName)
	if l[0].Task != taskName{
		t.Errorf("Expected %q, got %q instead.", taskName, l[0].Task)
	}

	if l[0].Done {
		t.Errorf("New task should not be completed")
	}

	l.Complete(1)
	if !l[0].Done{
		t.Errorf("Expected true, got %v instead", l[0].Done)
	}
}

func TestDelete(t *testing.T){
	l :=todo.List{}
	tasks := []string{"Task 1", "Task 2", "Task 3"}
	for _,task := range tasks {
		l.Add(task)
	}
	if l[0].Task != tasks[0]{
		t.Errorf("Expected %q, got %q instead.", tasks[0], l[0].Task)
	}

	l.Delete(2)

	if len(l) != 2 {
		t.Errorf("Expected %q, got %q instead.", 2,len(l))
	}

	if l[1].Task != tasks[2]{
		t.Errorf("Expected %q, got %q instead", tasks[2], l[1].Task)
	}

	
}


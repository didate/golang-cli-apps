package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/didate/interracting/todo"
)

const todoFileName = ".todo.json"

func main() {

	flag.Usage = func()  {
		fmt.Fprintf(flag.CommandLine.Output(), "%s tool. Developed for the pragmatic Bookshelf\n", os.Args[0])
		fmt.Fprintln(flag.CommandLine.Output(), "Copyright 2022")
		fmt.Fprintln(flag.CommandLine.Output(), "Usage Information :")
		flag.PrintDefaults()
	}

	task := flag.String("task","","Task to be included in the ToDo List")
	list := flag.Bool("list",false,"List of all tasks")
	complete := flag.Int("complete",0,"Item to be completed")

	flag.Parse()

	l := &todo.List{}

	if err := l.Get(todoFileName); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	switch {
	case *list:
		for _, item := range *l {
			fmt.Println(item.Task)
		}
	case *complete >0:
		if err:= l.Complete(*complete); err!=nil{
			fmt.Fprintln(os.Stderr, err)
		}
		if err := l.Save(todoFileName); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	case *task != "":
		l.Add(*task)
		if err := l.Save(todoFileName); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	default:
		fmt.Fprintln(os.Stderr, "Invalid option")
		os.Exit(1)
	}

}

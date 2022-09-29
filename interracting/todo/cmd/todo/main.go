package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/didate/interracting/todo"
)

var todoFileName = ".todo.json"

func main() {

	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "%s tool. Developed for the pragmatic Bookshelf\n", os.Args[0])
		fmt.Fprintln(flag.CommandLine.Output(), "Copyright 2022")
		fmt.Fprintln(flag.CommandLine.Output(), "Usage Information :")
		flag.PrintDefaults()
	}

	add := flag.Bool("add", false, "Add Task to the ToDo List")
	list := flag.Bool("list", false, "List of all tasks")
	complete := flag.Int("complete", 0, "Item to be completed")
	completedTask := flag.Bool("completedTask", false, "List only completed task")
	verbose := flag.Bool("verbose", false, "Item to be completed")

	del := flag.Int("del", 0, "Item to be deleted")

	flag.Parse()

	if os.Getenv("TODO_FILENAME") != "" {
		todoFileName = os.Getenv("TODO_FILENAME")
	}

	l := &todo.List{}

	if err := l.Get(todoFileName); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	switch {
	case *list:
		printList(l,*verbose,*completedTask)
		//fmt.Print(l)
	case *verbose:
		fmt.Println("-verbose flag must be combinated with -list flag")
	case *completedTask:
		fmt.Println("-completedTask flag must be combinated with -list flag")
	case *complete > 0:
		if err := l.Complete(*complete); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
		if err := l.Save(todoFileName); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	case *add:
		t, err := getTask(os.Stdin, flag.Args()...)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		l.Add(t)
		if err := l.Save(todoFileName); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	case *del > 0:
		if err:= l.Delete(*del); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		if err := l.Save(todoFileName); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	default:
		fmt.Fprintln(os.Stderr, "Invalid option")
		os.Exit(1)
	}

}

func printList(l *todo.List, verbose bool, completedTask bool)  {
	formatted := ""

	for k, t := range *l {
		if completedTask && !t.Done{
			continue
		}
		datetime :=""
		prefix := "  "
		if t.Done {
			prefix = "X "
		}
		if verbose {
			datetime = "--created :"+ t.CreatedAt.Format("2006-01-02 15:04")
		}

		formatted += fmt.Sprintf("%s%d: %s %v\n", prefix, k+1, t.Task, datetime)
	}
	fmt.Print(formatted)
}

func getTask(r io.Reader, args ...string) (string, error) {
	if len(args) > 0 {
		return strings.Join(args, " "), nil
	}
	s := bufio.NewScanner(r)
	s.Scan()
	if err := s.Err(); err != nil {
		return "", err
	}
	if len(s.Text()) == 0 {
		return "", fmt.Errorf("task cannot be blank")
	}
	return s.Text(), nil
}

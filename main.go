package main

import (
	"fmt"
	"os"

	errors "github.com/arknable/goseries-errors"
)

func main() {
	if err := doSomethingGreat(); err != nil {
		fmt.Println(err)

		me := err.(*errors.MyError)
		fmt.Println(me.NotesInString())

		os.Exit(1)
	}
}

func doSomethingGreat() error {
	if err := doSomethingUseful(); err != nil {
		return errors.Wrap(err, "doSomethingGreat")
	}
	return nil
}

func doSomethingUseful() error {
	if err := doSomethingCaller(); err != nil {
		return errors.Wrap(err, "doSomethingUseful")
	}
	return nil
}

func doSomethingCaller() error {
	if err := doSomething(); err != nil {
		return errors.Wrap(err, "doSomethingCaller")
	}
	return nil
}

func doSomething() error {
	return fmt.Errorf("ooops, something error!!")
}

package main

import "fmt"

type friut struct {
	Name, Type string
}

func (f friut) print() {
	fmt.Printf("The fruit name is %v and the fruit tyoe is %v", f.Name, f.Type)
}

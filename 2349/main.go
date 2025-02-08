package main

import "fmt"

func main() {
	fmt.Sprintln("Hello World")
}

type NumberContainers struct {
}

func Constructor() NumberContainers {
	return NumberContainers{}
}

func (this *NumberContainers) Change(index int, number int) {

}

func (this *NumberContainers) Find(number int) int {
	return number
}

/**
 * Your NumberContainers object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Change(index,number);
 * param_2 := obj.Find(number);
 */

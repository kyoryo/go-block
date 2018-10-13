package main

import "fmt"

type Block struct {
	Try     func()
	Catch   func(Exception)
	Finally func()
}

type Exception interface{}

func Throw(up Exception) {
	panic(up)
}
func (tcf Block) Do() {
	if tcf.Finally != nil {
		defer tcf.Finally()
	}
	if tcf.Catch != nil {
		defer func() {
			if r := recover(); r != nil {
				tcf.Catch(r)
			}
		}()
	}
	tcf.Try()
}

func main() {
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)

	//cikan dynamic array
	//in go im not sure if i can try catching
	//but this line will do I think
	///Below code will give you a panic (error in go)
	Block{
		Try: func() {
			var arr []int
			for index := 0; index < 10; index++ {
				arr[index] = index
				fmt.Println(arr[index])
			}
		},
		Catch: func(e Exception) {
			fmt.Printf("Caught %v\n", e)
		},
		Finally: func() {
			fmt.Printf("wew... lad")
		},
	}.Do()
}

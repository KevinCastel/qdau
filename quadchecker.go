package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// arg := os.Args[0]

	read, _ := bufio.NewReader(os.Stdin).ReadString(100)

	// ./QuadA 6 5 | go run .

	// out, _, _ := bufio.NewReader(os.Stdout).ReadLine()
	// fmt.Println(arg)
	arr := []string{}
	r := []rune(read)

	start := 0

	for i, ru := range r {
		if ru == '\n' {
			arr = append(arr, string(r[start:i]))
			start = i + 1
		}
	}

	fmt.Printf("Square of size [%d;%d]\n\n", len(arr[0]), len(arr))

	fmt.Println(string(read))
	// fmt.Println(err)

	/*
	   out, err := exec.Command("/home/student/zone01/quadchecker/word").Output()

	   fmt.Println(out)
	   fmt.Println(err)
	*/
}

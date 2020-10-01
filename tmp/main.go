package main

import "fmt"

type Stmt struct {
	Id string
}

func (s *Stmt) Close() {

}

func main() {

	mp := map[string]interface{}{
		"processed":      1,
		"buffer_size":    2,
		"thread_working": 3,
		"agent":          "dfa",
	}

	fmt.Println(mp["thread_working"].(int))
}

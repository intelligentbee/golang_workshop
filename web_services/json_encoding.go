package main

import "encoding/json"
import "fmt"

type StructA struct {
	A string `json:"ab"`
	B int    `json:"b"`
	C string `json:"c,omitempty"`
}

func main() {
	strA, _ := json.Marshal("gopher")
	fmt.Println(string(strA)) // "gopher"

	mapA := map[string]int{
		"apple":   5,
		"lettuce": 7,
	}
	mapB, _ := json.Marshal(mapA)
	fmt.Println(string(mapB)) // {"apple":5, "lettuce": 7}

	structA := StructA{A: "test", B: 5}
	structB, _ := json.Marshal(structA)
	fmt.Println(string(structB))
}

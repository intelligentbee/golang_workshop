package main

import (
	"encoding/json"
	"fmt"
)

type StructA struct {
	A string `json:"ab"`
	B int    `json:"b"`
	C string `json:"c, omitempty"`
}

func main() {
	strA := []string{}
	_ = json.Unmarshal([]byte(`["gopher", "con"]`), &strA)
	fmt.Println(strA) // [gopher, con]

	mapA := map[string]int{}
	_ = json.Unmarshal([]byte(`{"apple": 5, "lettuce": 7}`), &mapA)
	fmt.Println(mapA) // map[apple:5, lettuce:7]

	structA := StructA{}
	_ = json.Unmarshal([]byte(`{"ab":"test", "b":5}`), &structA)
	fmt.Println(structA) // {test 5 }
}

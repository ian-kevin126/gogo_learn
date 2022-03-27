package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	m := make(map[string]interface{}, 4)
	m["company"] = "zhczGO"
	m["subjects"] = []string{"go", "fabric", "python", "Test"}
	m["isok"] = true
	m["price"] = 99

	// result,err :=json.Marshal(m)
	result, err := json.MarshalIndent(m, "", "	")
	if err != nil {
		fmt.Println("err=", err)
		return
	}

	fmt.Println("result=", string(result))

	/*
		result= {
		        "company": "zhczGO",
		        "isok": true,
		        "price": 99,
		        "subjects": [
		                "go",
		                "fabric",
		                "python",
		                "Test"
		        ]
		}
	*/
}

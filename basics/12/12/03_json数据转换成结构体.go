package main

import (
	"encoding/json"
	"fmt"
)

type IT1 struct {
	Company  string   `json:"company"`
	Subjects []string `json:"subjects"`
	IsOk     bool     `json:"isok"`
	Price    float64  `json:"price"`
}

type IT2 struct {
	Company string `json:"company"`
}

func main() {
	jsonbuff := `{
		"company": "zhczGO",
		"isok": true,
		"price": 99,
		"subjects": [
			"go",
			"fabric",
			"python",
			"Test"
		]
	}`

	var temp IT1
	err := json.Unmarshal([]byte(jsonbuff), &temp)

	if err != nil {
		fmt.Println("err=", err)
		return
	}
	//fmt.Println("temp=",temp)
	fmt.Printf("tmp=%+v\n", temp) // tmp={Company:zhczGO Subjects:[go fabric python Test] IsOk:true Price:99}

	var temp2 IT2
	err = json.Unmarshal([]byte(jsonbuff), &temp2)
	if err != nil {
		fmt.Println("err=", err)
		return
	}

	fmt.Printf("tmp2=%+v\n", temp2) // tmp2={Company:zhczGO}

}

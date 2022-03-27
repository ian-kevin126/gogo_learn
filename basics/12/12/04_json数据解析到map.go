package main

import (
	"encoding/json"
	"fmt"
)

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

	//创建一个map
	m := make(map[string]interface{}, 4)
	err := json.Unmarshal([]byte(jsonbuff), &m) //一定要是地址

	if err != nil {
		fmt.Println("err=", err)
		return
	}
	fmt.Printf("m=%+v/n", m)
	// m=map[company:zhczGO isok:true price:99 subjects:[go fabric python Test]]/nmap[company]的值类型为string，value=zhczGO

	//var str string
	//str = m["company"]//err 无法转换

	//类型断言
	for key, value := range m {
		switch data := value.(type) {
		case string:
			//str = data
			fmt.Printf("map[%s]的值类型为string，value=%s\n", key, data)
		case bool:
			fmt.Printf("map[%s]的值类型为bool，value=%v\n", key, data)
		case float64:
			fmt.Printf("map[%s]的值类型为float64，value=%f\n", key, data)
		case []string:
			fmt.Printf("map[%s]的值类型为[]string，value=%v\n", key, data)
		case []interface{}:
			fmt.Printf("map[%s]的值类型为interface{}，value=%v\n", key, data)
		}
	}

	//map[isok]的值类型为bool，value=true
	//map[price]的值类型为float64，value=99.000000
	//map[subjects]的值类型为interface{}，value=[go fabric python Test]

}

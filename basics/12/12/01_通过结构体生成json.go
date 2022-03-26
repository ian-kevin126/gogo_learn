package main

import (
	"encoding/json"
	"fmt"
)

/**
{

  “company”：“zhczGO”,

 "isok":true,

 "price":99.00

"subjects":["go","docker","Test"]

}
*/
/* type IT struct {
	Company string
	Subjects []string
	IsOk   bool
	Price  float64
}*/
/*type IT struct {
	Company string `json:"company"` //二次编码
	Subjects []string `json:"subjects"`
	IsOk   bool `json:"isok",string`
	Price  float64 `json:"price,string"`
}*/

type IT struct {
	Company  string   `json:"-"` //-此字段不会输入到屏幕上面
	Subjects []string `json:"subjects"`
	IsOk     bool     `json:"isok",string`
	Price    float64  `json:"price,string"`
}

func main() {

	s := IT{"zhczGO", []string{"go", "docker", "fabric", "Test"}, true, 999}

	//内置方法
	//buf,err:=json.Marshal(s)
	buf, err := json.MarshalIndent(s, "", "	") //格式化代码
	if err != nil {
		fmt.Println("err=", err)
		return
	}
	fmt.Println("buf=", string(buf))
}

package main

import "fmt"

type Student11 struct {
	name string
	id   int
}

func main() {
	i := make([]interface{}, 3)
	i[0] = 1
	i[1] = "hello go"
	i[2] = Student11{"mike", 7777}

	/*//类型的查询，类型的断言
	 //index 是返回下标  data 对应的值 data【0】 data【1】 data【2】
	for index,data :=range  i{
		//第一个返回时值，ok返回时真还是假
		if value,ok := data.(int);ok==true{
			fmt.Printf("i[%d] 类型为int，内容为%d\n",index,value)
		}else if value,ok:=data.(string);ok==true{
			fmt.Printf("i[%d] 类型为string，内容为%s\n",index,value)
		}else  if value,ok:=data.(Student11);ok==true{
			fmt.Printf("i[%d]类型为student，内容为name=%s,id=%d\n",index,value.name,value.id)
		}
	}*/

	for index, data := range i {
		switch value := data.(type) {
		case int:
			fmt.Printf("i[%d] 类型为int，内容为%d\n", index, value)
		case string:
			fmt.Printf("i[%d] 类型为string，内容为%s\n", index, value)
		case Student11:
			fmt.Printf("i[%d]类型为student，内容为name=%s,id=%d\n", index, value.name, value.id)
		}
	}

}

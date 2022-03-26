![循环](https://gitee.com/ian_kevin126/picturebed/raw/master/mac/algorithms/%E5%BE%AA%E7%8E%AF.png)



内容回顾

printf字符串进行格式化
println 回车换行      
print 是否换行

多重赋值
--------------------
```go
i,j := 10,20

i,j= j,i
```



_隐藏
-----------------------
```go
i,_ =j,i

func test()(int int int){

  return 1,2,3
}

func main(){

  i,j,z := test()

  i,j,_ := test()
}
```




常量
----------------
常量：程序运行期间，不可以改变的量，常量声明需要const


iota
-----------------
准对是常量

```go
const（
i = iota //0
j= iota //1
z= iota //2

）

const（
i = iota //0
j  //1
z //2

）

const d = iota //0
```




fmt输出格式
-------------------------
```go
func main()  {
  a := 10
  b := "abc"
  c :='a'
  d := 3.14

  fmt.Printf("%T,%T,%T,%T\n",a,b,c,d)

  //%d整形格式 %s字符串格式 %c字符 %f浮点
  fmt.Printf("a=%d,b=%s,c=%c,d=%f\n",a,b,c,d)

  fmt.Printf("a=%v,b=%v,c=%v,d=%v\n",a,b,c,d)

}
```




fmt的输入格式
------------------------
```go
package main

import "fmt"

func main()  {
  a := 10
  b := "abc"
  c :='a'
  d := 3.14

  fmt.Printf("%T,%T,%T,%T\n",a,b,c,d)

  //%d整形格式 %s字符串格式 %c字符 %f浮点
  fmt.Printf("a=%d,b=%s,c=%c,d=%f\n",a,b,c,d)

  fmt.Printf("a=%v,b=%v,c=%v,d=%v\n",a,b,c,d)


}
```




类型转换
-----------------------
```go
package main

import "fmt"

func main(){

  var ch byte
  ch = 'A'
  var t  int
  t = int(ch)//不支持隐式转换
  fmt.Println("t=",t)

  var flag bool
  flag = true
  fmt.Printf("flag=%t\n",flag)
  //这种不能转换成int类型 因为不兼容类型 bool不能转换int
  //fmt.Printf("flag=%d\n",int(flag))

}
```



运算符操作
--------------------
```go
func main()  {

  fmt.Println("1+2 结果：",1+2)

  A := 20
  A++  //++只能放在后面
  //++A
  fmt.Println("A++ 的结果",A)

  fmt.Println("5 > 3 结果：",5>3)

  fmt.Println("!(4>3) 结果：",!(4>3))

  //&& 与 并且 左边右边都为真，结果才为真，有一个为假 都为假
  fmt.Println("true&&true 结果", true && true)
  fmt.Println("true&&false 结果", true&&false)

  //|| 或 左边和右边只要有个一个为真都为真

  fmt.Println("true||false 结果", true||false )
  fmt.Println("false||false 结果", false||false)

  a := 20

  fmt.Println("0<=a && a<=10 结果：",0<=a && a<=10)

}
```




流程控制
-------------------
顺序结构 1----2-----3
选择结构
循环结构

if使用
----------------
```go
// a  := 10 //:= 10初始化并赋值给a
if a ==10 { //==  判断 true fasle
  fmt.Println("a==10")
}
```



if else
------------
```go
if a ==10{

} else {

}
```



初始化并判断
------------------


```go
//if语句进行初始化 然后在赋值
if a :=10;a==10{
  fmt.Println("a==10")
}else{
  fmt.Println("a!=10")
}

//if语句进行初始化 然后在赋值
if a =10;a==10{
  fmt.Println("a==10")
}else{
  fmt.Println("a!=10")
}

a :=10 只是if 这个范围里面的变量 a=10 main里面整体变量


if （）{

}else if（）{
} else{

}
 // ---------------------

if a := 8; a ==10{
  fmt.Println("a==10")
}else if a > 10{
  fmt.Println("a > 10")
}else  if a<10{
  fmt.Println("a <10")
}else {
  fmt.Println("这是不可能走到的")
}
```

switch用法
--------------------
Go里面switch默认相当于每个case最后带有break，匹配成功后不会自动向下执行其他case，
而是跳出整个switch, 但是可以使用fallthrough强制执行后面的case代码：



```go
switch b {  //swich case 默认在最后一行加入break，一旦满足条件就不会在往下运行
  //如果想让执行语句满足条件之后继续运行 就必须加fallthrough
  case 1 :
  fmt.Println("这个数字是1")
  //break
  case 2:
  fmt.Println("这个数字是2")
  //break
  case 3:
  fmt.Println("这个数字是3")
  fallthrough  //只要你不跳出switch语句，后面无条件执行
  //break
  case 4:
  fmt.Println("这个数字是4")
  fallthrough
  //break
  default:
  fmt.Println("按下的是xxx")
}
```


swich后面没有条件，case语句里面要放条件
-------------------------------


```go
switch  { //可以没有条件
  case num >60://case 后面要跟上条件语句
  fmt.Println(">60的数字")
  case num>90:
  fmt.Println(">90的数字")
  case num==100: //相同的语句我们可以合并
  fmt.Println("100的数字")
  default:
  fmt.Println("其他")
}
```



for的基本用法
------------------------


```go
func main()  {
  //for 初始化条件 ；判断条件；条件变化{
  //  循环体
  // }
  //1)初始化条件 i:=1
  //2)初始化条件变化是否为真，i<=100,如果为真，执行循环体，如果为假，跳出循环
  //3)条件变化 i++
  //4）重复执行2,3,4
  sum := 0
  for i :=1;i<=100;i++{
    sum = sum +i
  }

  fmt.Println("sum=",sum)
}
```



range 迭代打每个元素，默认返回2个值,第一个是元素的位置，一个是元素的本身

-----------------------------------------------------------------------
```go
 for i,data := range  str{
 	fmt.Printf("str[%d]=%c\n",i,data)
 }

 for i := range str{//第二个返回值，默认丢弃，返回元素的位置（下标）
 	fmt.Printf("str[%d]=%c\n",i,str[i])
 }

 for i,_ :=range str{ //第二个返回值，默认丢弃，_
 	fmt.Printf("str[%d]=%c\n",i,str[i])
 }
```



break和continue
-----------------


```go
func main()  {
  i := 0
  for{
    i++
    time.Sleep(time.Second) //睡觉1s
    if i ==5{
      //break //跳出循环，如果嵌套多个循环，跳出最近的那个内循环，本层循环
      continue //跳出本次循环，继续下一次循环
    }
    fmt.Println("i=",i)
  }
}
```


goto
--------------
goto是关键字，END是用户起的名字，他叫标签，goto不能夸函数调用

```go
func main()  {
  fmt.Println("11111111")
  goto END //goto是关键字，END是用户起的名字，他叫标签，goto不能夸函数调用
  fmt.Println("22222222")
  END:
  fmt.Println("333333")
}
```










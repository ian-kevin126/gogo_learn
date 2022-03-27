package main

import "fmt"

func main() {
	a := [8]int{24, 69, 80, 57, 13, 1, 100}
	//算出a的长度
	//n :=len(a)
	//0 1 2 3
	//冒泡排序
	/*for i:=0;i<5-1;i++ {
	       for j:=0;j<5-1-i;j++{
	       	  if a[j] > a[j+1]{
	       	  	 a[j],a[j+1] = a[j+1],a[j]
			  }
		   }
		}

		fmt.Printf("\n培训后：\n")
		for i:=0;i< 5;i++{
			fmt.Printf("%d\n",a[i])
		}
		fmt.Println()*/

	//升序
	/* n := len(a)
	for i:=0;i<n-1;i++ {
		for j:=0;j<n-1-i;j++{
			if a[j] > a[j+1]{
				a[j],a[j+1] = a[j+1],a[j]
			}
		}
	}*/

	//降序
	n := len(a)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-1-i; j++ {
			if a[j] < a[j+1] {
				a[j], a[j+1] = a[j+1], a[j]
			}
		}
	}

	fmt.Printf("\n排序后：\n")
	for i := 0; i < n; i++ {
		fmt.Printf("%d\n", a[i])
	}

	fmt.Println()
}

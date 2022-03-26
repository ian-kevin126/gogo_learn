package main

import "fmt"

func main() {
	switch num := 4; num {
	case 1:
		fmt.Println("1的数字")
	case 2:
		fmt.Println("2的数字")
	case 3:
		fmt.Println("zz的数字")
	case 4:
		fmt.Println("zz的数字")
	case 5:
		fmt.Println("zz的数字")
	case 6:
		fmt.Println("6的数字")
	default:
		fmt.Println("xxx的数字")

	}
	fmt.Println("--------------------------")

	switch num := 4; num {
	case 1:
		fmt.Println("1的数字")
	case 2:
		fmt.Println("2的数字")
	case 3, 4, 5: //相同的语句我们可以合并
		fmt.Println("zz的数字")

	case 6:
		fmt.Println("6的数字")
	default:
		fmt.Println("xxx的数字")

	}

	num := 4
	switch { //可以没有条件
	case num > 60: //case 后面要跟上条件语句
		fmt.Println(">60的数字")
	case num > 90:
		fmt.Println(">90的数字")
	case num == 100: //相同的语句我们可以合并
		fmt.Println("100的数字")

	default:
		fmt.Println("其他")

	}

}

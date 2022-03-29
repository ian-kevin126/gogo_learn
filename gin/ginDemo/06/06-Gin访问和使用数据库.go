package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/prometheus/common/log"
)

func main() {

	// 1，连接数据库
	connStr := "root:ljy760155@tcp(127.0.0.1:3306)/gin_demo"
	db, err := sql.Open("mysql", connStr)
	if err != nil {
		log.Fatal(err.Error())
		return
	}

	// 2，创建一个表table，只需要执行一次就行了，不能创建同名表
	//_, err = db.Exec("create table person(" +
	//	"id int auto_increment primary key," +
	//	"name varchar(12) not null," +
	//	"age int default 1" +
	//	");")
	//if err != nil {
	//	log.Fatal(err.Error())
	//	return
	//}

	// 3，插入数据到数据库
	_, err = db.Exec("insert into person(name,age)"+
		"values(?,?);", "Ian", 23)
	if err != nil {
		log.Fatal(err.Error())
		return
	} else {
		fmt.Println("数据插入成功")
	}

	// 4，查询数据库
	rows, err := db.Query("select id,name,age from person")
	if err != nil {
		log.Fatal(err.Error())
		return
	}
	/*
		可能读出多行：
		Davie 18
		Kevin 22
		Ian 23
	*/
	// 使用Next方法，每读一行，就走向下一行
scan:
	if rows.Next() {
		person := new(Person)
		// 逐行读取表中的数据，然后放入到person中
		err := rows.Scan(&person.Id, &person.Name, &person.Age)

		if err != nil {
			log.Fatal(err.Error())
			return
		}

		fmt.Println(person.Id, person.Name, person.Age)
		// 通过goto语句来实现逐行读取
		goto scan
	}

	/*
		数据插入成功
		1 Davie 18
		2 Kevin 22
		3 Ian 23
		4 Ian 23
	*/

}

type Person struct {
	Id   int
	Name string
	Age  int
}

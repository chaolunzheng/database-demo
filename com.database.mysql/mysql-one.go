package main

import (
	"database/sql"
	"fmt"
	//加载驱动
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	//数据库的连接信息
	conn, err := sql.Open("mysql", "root:root@/test?charset=utf8")
	check(err)

	//TODO 增加数据
	//insert:="insert into image(url,user_id,created_date) values('http://www.xxx.com',2,'2019-03-15 20:53:14')"
	//stmt, err := conn.Prepare(insert)
	//check(err)
	//result, e := stmt.Exec()
	//check(e)
	//record, e2 := result.LastInsertId()
	//check(e2)
	//fmt.Println(record)

	//TODO 查询数据
	query := "select * from image order by id desc limit 3"
	//stmt, e := conn.Prepare(sql)
	//check(e)
	rows, e3 := conn.Query(query)
	check(e3)

	for rows.Next() {
		var id int32
		var url string
		var user_id int32
		var created_date string

		scan := rows.Scan(&id, &url, &user_id, &created_date)
		check(scan)

		fmt.Println(id, url, user_id, created_date)
	}

	//修改数据
	updateSql := "update image set created_date='2019-03-03' where id =1"
	result, e := conn.Exec(updateSql)
	check(e)

	//删除数据

}

//如果有错误的话，go中关于错误的处理
func check(err error) {
	if err != nil {
		fmt.Printf("发生错误了，err is not nil:%v", err)
	}
}

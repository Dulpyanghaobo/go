package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	_ "./Models"
)



func connectMysql() {
	database,_ := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/demo")
	database.SetConnMaxIdleTime(100)
	database.SetMaxIdleConns(10)
	if err := database.Ping(); err != nil {
		fmt.Println("open database fail")
		return
	}
	fmt.Println("connect success")
	var user User
}
func main() {
	connectMysql()
	//r := gin.Default()
	//r.GET("/ping", func(c *gin.Context) {
	//	c.JSON(200, gin.H{
	//		"message": "pong",
	//	})
	//})
	//r.Run() // listen and serve on 0.0.0.0:8080
}
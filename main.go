package main

import (
	"fmt"
	"gdback/src/mysqldb"
	"gdback/src/request"
)

func main() {
	fmt.Println("config start successed...")
	mysqldb.Start()
	fmt.Println("mysqldb start successed...")
	// crypto.Test1()
	// jwt.Test()
	request.Start()
}

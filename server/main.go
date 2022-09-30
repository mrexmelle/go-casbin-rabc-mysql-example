package main

import (
	"fmt"
	"os"

	"github.com/casbin/casbin/v2"
	"github.com/casbin/xorm-adapter/v2"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	if len(os.Args) < 4 {
		fmt.Fprintf(os.Stderr, "Usage: %s <user> <resource> <action>\n", os.Args[0])
		os.Exit(1)
	}

	adapter, _ := xormadapter.NewAdapter("mysql", "root:123456@tcp(127.0.0.1:3306)/")
	enforcer, _ := casbin.NewEnforcer("model.conf", adapter)
	enforcer.LoadPolicy()

	enforcer.AddNamedPolicy("p", "alice", "data1", "read")
	enforcer.AddNamedPolicy("p", "data2_admin", "data2", "read")
	enforcer.AddNamedPolicy("p", "data2_admin", "data2", "write")
	enforcer.AddNamedPolicy("g", "alice", "data2_admin")

	res, _ := enforcer.Enforce(os.Args[1], os.Args[2], os.Args[3])
	if res {
		fmt.Println("Access allowed")
	} else {
		fmt.Println("Access denied")
	}

	enforcer.SavePolicy()
}

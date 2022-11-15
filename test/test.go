package main

import (
	"fmt"
	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	a, _ := gormadapter.NewAdapter("mysql", "root:root@tcp(127.0.0.1:3306)/health", true)
	e, err := casbin.NewEnforcer("./config/rbac_model.conf", a)
	//e, err := casbin.NewEnforcer("./config/rbac_model.conf", "./config/rbac_policy.csv")
	if err != nil {
		panic(err)
	}
	e.LoadPolicy()

	//e.AddPolicy("alice", "data1", "write")
	//var rule [][]string
	//rule = append(rule, []string{"alice", "data2", "write"})
	//rule = append(rule, []string{"alice", "data2", "read"})
	//e.AddPolicies(rule)
	//ok, err := e.UpdatePolicy([]string{"alice", "data1", "write"}, []string{"alice2", "data1", "read"})
	//if err != nil {
	//	panic(err)
	//}
	//if !ok {
	//	fmt.Println("修改失败")
	//}
	e.AddGroupingPolicy("alice", "fb")
	e.AddPolicy("fb", "data1", "read")
	sub := "alice"
	obj := "data1"
	act := "read"
	ok, err := e.Enforce(sub, obj, act)
	if err != nil {
		return
	}
	if ok {
		fmt.Println("通过")
	} else {
		fmt.Println("拒绝")
	}

}

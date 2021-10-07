package main

import (
	"fmt"
	"github.com/casbin/casbin"
)

func main() {
	// 通过策略文件和模型配置穿件一个casbin访问控制实例
	var e = casbin.NewEnforcer("./perm.conf", "./policy.csv")

	// 定义各种sub，obj和act的数组
	subs := []string{"bob", "zeta"}
	objs := []string{"data1", "data2"}
	acts := []string{"read", "write"}

	fmt.Println(e.Enforce("zeta", "data2", "write"))

	for _, sub := range subs {
		for _, obj := range objs {
			for _, act := range acts{
				fmt.Println(sub, "/", obj, "/", act, "=", e.Enforce(sub, obj, act))
			}
		}
	}
}

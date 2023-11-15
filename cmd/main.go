package main

import (
	"fmt"
	_ "github.com/bytedance/sonic"
	_ "github.com/go-shafaq/go1.20-encoding-json"
	json "github.com/go-shafaq/go1.20-encoding-json"
	"github.com/go-shafaq/go1.20-sonic"
	"runtime"
	"strings"
)

func main() {
	fmt.Println(runtime.Version())
	_ = sonic.Marshal
	d, _ := json.Marshal(User{13})
	println(string(d))
	d, _ = sonic.Marshal(User{13})
	println(string(d))
}

type User struct {
	UserId int
}

func init() {
	json.DefCase(func(tag string) func(pkgPath string, name string) string {
		return func(pkgPath string, name string) string {
			return strings.ToLower(name)
		}
	})
}

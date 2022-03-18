package main

import (
	"code/route"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("[参数缺失]")
		return
	}
	fmt.Println(route.Handler(os.Args))
}

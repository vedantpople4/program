package main

import (
	"fmt"
	"strings"
)

func main() {
	var str string
	fmt.Scanf("%s", &str)
	ans := strings.ToUpper(str)
	fmt.Println(ans)
}

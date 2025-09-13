package main

import (
	"fmt"
)

func main() {
	var str string
	str = "hey what's going on bro"

	preColon(str, 2)
	fmt.Printf("%s", string(str[2]))

}

func preColon(s string, i int) {
	fmt.Printf("String is %s. Index was :%d. Result is:'%s' > '%s' \n", s, i, s[:i], s[i:])
}

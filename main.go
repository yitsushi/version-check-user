package main

import (
	"fmt"

	base "github.com/yitsushi/version-check-base/v2"
)

func main() {
	b := &base.MyStuff{Field: "nice"}
	fmt.Println(b.String())
}

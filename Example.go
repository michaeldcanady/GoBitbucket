package main

import (
	"fmt"

	"github.com/michaeldcanady/gobitbucket/gobitbucket"
)

func main() {
	fmt.Println(gobitbucket.New("User", "Password", "Url"))
}

package hello

import (
	"fmt"
	"rsc.io/quote"
)

func Hello() {
	fmt.Println("Hello by hello package")
	fmt.Println(quote.Hello())
}

package main

import (
	"fmt"
	"time"

	"github.com/araddon/dateparse"
)

func main() {
	dateExample := "21:34"
	t, err := dateparse.ParseLocal(dateExample)
	fmt.Println(t)
	fmt.Println(err)

	form := "21:34"
	t2, e := time.Parse(form, "20:10")
	fmt.Println(t2)
	fmt.Println(e)
}

package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	fmt.Println(time.Now().Unix(),   " ",  fmt.Sprintf("%s_%s", "asfdasd", t.Format("200601021504")))

	test1 :=  "\"2019-12-04T11:05:23.000+0000\""

	ed :=  &ExpiryDate{}

	err := ed.UnmarshalJSON([]byte(test1))
	fmt.Println("err ",err,  ed.Time)
}

type ExpiryDate struct {
	time.Time
}
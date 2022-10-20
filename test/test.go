package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	str := "2022-10-11T00:00:00+08:00"
	t, _, _ := strings.Cut(str, "T")
	fmt.Println(t)
	fmt.Println(time.Parse("2006-01-02", t))
	fmt.Println(time.ParseDuration(str))
}

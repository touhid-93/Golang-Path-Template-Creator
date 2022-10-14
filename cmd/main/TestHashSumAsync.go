package main

import (
	"fmt"
	"time"

	"gitlab.com/evatix-go/pathhelper/checksummer"
	"gitlab.com/evatix-go/pathhelper/hashas"
)

func TestHashSumAsync() {
	start := time.Now()
	c := checksummer.NewAsync(true, "D:\\vm", hashas.Md5)
	elapsed := time.Since(start)
	prettyPrint(c.GetMap())
	fmt.Printf("Elapsed Async: %s\n", elapsed)
}

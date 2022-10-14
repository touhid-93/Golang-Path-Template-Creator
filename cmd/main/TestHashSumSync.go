package main

import (
	"fmt"
	"time"

	"gitlab.com/evatix-go/pathhelper/checksummer"
	"gitlab.com/evatix-go/pathhelper/hashas"
)

func TestHashSumSync() {
	start := time.Now()
	c := checksummer.NewSync(true, "D:\\vm", hashas.Md5)
	elapsed := time.Since(start)
	prettyPrint(c.GetMap())
	fmt.Printf("Elapsed Sync: %s\n", elapsed)
}

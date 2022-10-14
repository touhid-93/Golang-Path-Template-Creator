package main

import (
	"fmt"
	"sort"
)

func prettyPrint(m map[string][]byte) {
	names := make([]string, 0, len(m))
	for name := range m {
		names = append(names, name)
	}
	sort.Strings(names)

	for _, key := range names {
		fmt.Printf("%x: %s\n", m[key], key)
	}
}

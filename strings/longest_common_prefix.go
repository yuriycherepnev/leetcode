package main

import (
	"fmt"
	"sort"
)

func main() {
	strs := []string{"flraower", "flr", "flaow"}

	prefix := longestCommonPrefix(strs)

	fmt.Println(prefix)
}

func longestCommonPrefix(strs []string) string {
	sort.Strings(strs)

	fmt.Println(strs)
	//first := strs[0]
	//last := strs[len(strs)-1]
	//
	//i := 0
	//for i < len(first) && first[i] == last[i] {
	//	i++
	//}

	return ""
}

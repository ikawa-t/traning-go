package main

import (
	"fmt"
	"sort"
)

type RuneSlice []rune
func (r RuneSlice) Len() int {return len(r)}
func (r RuneSlice) Less(i, j int) bool {return r[i] < r[j]}
func (r RuneSlice) Swap(i, j int) {r[i], r[j] = r[j], r[i]}

func anagram(str1, str2 string) bool {
	r1 := []rune(str1)
	sort.Sort(RuneSlice(r1))
	r2 := []rune(str2)
	sort.Sort(RuneSlice(r2))
	return string(r1) == string(r2)
}

func main() {
	fmt.Println(anagram("abcdeabcde", "edcbaedcba")) // true
	fmt.Println(anagram("abc", "abcd")) // false
}


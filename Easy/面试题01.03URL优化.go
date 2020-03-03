package main

import "strings"
//strings.Replace
//返回将s中前n个不重叠old子串都替换为new的新字符串，如果n<0会替换所有old子串。
func replaceSpaces(S string, length int) string {
	return strings.Replace(S," ","%20",-1)
}
func main() {
	
}

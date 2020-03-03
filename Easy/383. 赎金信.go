package main
import "fmt"
//给定一个赎金信 (ransom) 字符串和一个杂志(magazine)字符串，判断第一个字符串ransom能不能由第二个字符串magazines里面的字符构成。如果可以构成，返回 true ；否则返回 false。
//(题目说明：为了不暴露赎金信字迹，要从杂志上搜索各个需要的字母，组成单词来表达意思。)
//
//注意：
//你可以假设两个字符串均只含有小写字母。
//canConstruct("a", "b") -> false
//canConstruct("aa", "ab") -> false
//canConstruct("aa", "aab") -> true

//类似01.02 运用map，遍历一个+1一个-1


func canConstruct(ransomNote string, magazine string) bool {
	hh := make(map[byte]int)
	for i := range ransomNote {
		hh[ransomNote[i]] ++
	}
	for j := range magazine {
		//fmt.Println(j)
		hh[magazine[j]] --
	}

	for i := range ransomNote {
		if hh[ransomNote[i]] > 0 {
			return false
		}
	}
	return true
}
func main() {
	a:= "bg"
	b := "bdg"
	fmt.Println(canConstruct(a,b))
	
}

package main

import "fmt"

//func lengthOfLongestSubstring(s string) int {
//
//	max := 0
//
//	temp := []byte{}
//	temp = append(temp, s[0])
//	for i := 1 ; i < len(s) ; i ++{
//		for  j := len(temp)-1; j >= 0 ; j --{
//			if temp[j] == s[i]{
//				temp = temp[j+1:]
//				break
//			}
//		}
//		temp = append(temp, s[i])
//		if len(temp) > max{
//			max = len(temp)
//		}
//	}
//	return max
//}
//func lengthOfLongestSubstring(s string) int{
//	mp := map[byte]int{}
//
//	res := 1
//	if s ==""{
//		return 0
//	}
//
//	left ,right := 0,0
//
//	mp[s[0]] = 0
//
//	for i := 1 ; i < len(s) ; i++{
//		if a,ok := mp[s[i]] ;ok{
//			if a < left{
//			}else{
//				left = mp[s[i]] + 1
//			}
//
//			mp[s[i]] = i
//			right++
//		}else{
//			mp[s[i]] = i
//			right ++
//		}
//		if right - left + 1> res{
//			res = right - left + 1
//		}
//	}
//	return res
//
//}
//func lengthOfLongestSubstring(s string) int {
//
//	if len(s) == 0{
//		return 0
//	}
//	temp := []byte{}
//	temp = append(temp,s[0])
//
//	for i := 1 ; i < len(s) ; i ++{
//		for  j := len(temp) -1  ; j >= 0 ; j--{
//			if temp[j] =
//		}
//	}
//
//
//}
func main() {

	fmt.Println(lengthOfLongestSubstring("a"))
	
}

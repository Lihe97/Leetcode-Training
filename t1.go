package main

import "fmt"

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 * 对输入集合做hash join，输出关联后的下标的二元组集合
 * @param setA int整型一维数组 输入集合1
 * @param setB int整型一维数组 输入集合2
 * @return int整型二维数组
 */
func hashJoin( setA []int ,  setB []int ) [][]int {
	// write code here

	res := [][]int{}

	mp := map[int][]int{}

	if len(setA) <= len(setB){
		for i := 0 ; i < len(setA) ; i ++{
			mp[setA[i]] = append(mp[setA[i]],i)
		}
		for i := 0 ; i < len(setB) ; i ++{
			for a,b := range mp{
				if a == setB[i]{
					for j := 0 ; j < len(b) ; j ++{
						res = append(res,[]int{b[j],i})
					}
				}
			}
		}
	}else{
		for i := 0 ; i < len(setB) ; i ++{
			mp[setB[i]] = append(mp[setB[i]],i)
		}
		for i := 0 ; i < len(setA) ; i ++{
			for a,b := range mp{
				if a == setA[i]{
					for j := 0 ; j < len(b) ; j ++{
						res = append(res,[]int{i,b[j]})
					}
				}
			}
		}
	}

	return res

}

func main(){

	a := []int{1,2,3,3,4,4,5}
	b := []int{1,2,4,4,7,8}
	fmt.Println(hashJoin(a,b))
}
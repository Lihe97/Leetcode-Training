package main
//给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那 两个 整数，并返回他们的数组下标。
//
//你可以假设每种输入只会对应一个答案。但是，你不能重复利用这个数组中同样的元素。
//
//示例:
//
//给定 nums = [2, 7, 11, 15], target = 9
//
//因为 nums[0] + nums[1] = 2 + 7 = 9
//所以返回 [0, 1]
//
//用hash 存储的元素是[数组元素，数组索引]
import "fmt"

//func twoSum(nums []int, target int) []int {
//
//	hashMap := map[int]int{}
//	res := []int{}
//	for i:= range nums{
//		t := target - nums[i]
//		if _, ok := hashMap [t];ok==true{
//			res = append(res, hashMap[t])
//			res = append(res, i)
//			return res
//		}
//		hashMap[nums[i]] = i
//	}
//	return nil
//}
func twoSum(nums []int, target int) []int {

	mp := map[int]int{}
	res := []int{}
	for i := 0 ; i < len(nums) ; i ++{


		if _,ok := mp[target-nums[i]];ok{
			res = append(res,mp[target-nums[i]])
			res = append(res,i)

			return res
		}
		mp[nums[i]] = i

	}
	return nil
}
func main() {
	a := []int{3,2,4}
	fmt.Println(twoSum(a,6))
}

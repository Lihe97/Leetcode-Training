package main


func merge(nums1 []int, m int, nums2 []int, n int)  {

	i,j := m-1,n-1

	length := m + n - 1

	for j >= 0{
		if i >= 0 && nums1[i] >= nums2[j]{
			nums1[length] = nums1[i]
			i --
		}else{
			nums1[length] = nums2[j]
			j --
		}
		length --
	}
}
func main() {

	merge([]int{1,2,3,0,0,0},3,[]int{2,5,6},3)
	
}

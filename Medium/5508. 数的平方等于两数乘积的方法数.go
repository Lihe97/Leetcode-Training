package main
func numTriplets(nums1 []int, nums2 []int) int {

	res := 0

	mp1 := map[int]int{}
	mp2 := map[int]int{}

	for i := 0 ; i < len(nums1); i ++{
		mp1[nums1[i]*nums1[i]] ++
	}
	for i := 0 ; i < len(nums2) - 1; i ++{
		for j := i + 1 ; j < len(nums2) ; j ++{
			res += mp1[nums2[i]*nums2[j]]
		}
	}

	for i := 0 ; i < len(nums2); i ++{
		mp2[nums2[i]*nums2[i]] ++
	}
	for i := 0 ; i < len(nums1) - 1; i ++{
		for j := i + 1 ; j < len(nums1) ; j ++{
			res += mp2[nums1[i]*nums1[j]]
		}
	}



	return res

}

func main() {
	
}

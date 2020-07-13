package main
func intToRoman(num int) string {

	mp := map[int]string{}
	res := ""

	nums := []int{1,4,5,9,10,40,50,90,100,400,500,900,1000}

	mp[1] = "I"
	mp[4] = "IV"
	mp[5] = "V"
	mp[9] = "IX"

	mp[10] = "X"
	mp[40] = "XL"
	mp[50] = "L"
	mp[90] = "XC"

	mp[100] = "C"
	mp[400] = "CD"
	mp[500] = "D"
	mp[900] = "CM"

	mp[1000] = "M"

	for i := len(nums) - 1 ; num != 0 && i >= 0 ; i --{
		for num >= nums[i]{
			res += mp[nums[i]]
			num -= nums[i]
		}
	}
	return res

}
func main() {
	
}

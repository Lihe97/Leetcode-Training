package main

import "fmt"

type Twitter struct {
	user map[int][]int
	twi [][2]int
	//0 发的人id
	//2 内容
}

func Constructor() Twitter {
	twi := [][2]int{}
	return Twitter{
		user: map[int][]int{},
		twi:  twi,
	}
}
func (this *Twitter) PostTweet(userId int, tweetId int)  {

	if len(this.user[userId]) == 0{
		this.user[userId] = append(this.user[userId], userId)
	}
	if len(this.twi) == 0{
		temp := [2]int{userId,tweetId}
		this.twi = append(this.twi, temp)
	}else{
		temp := [2]int{userId,tweetId}
		this.twi = append(this.twi, temp)
	}
}
func (this *Twitter) GetNewsFeed(userId int) []int {

	if len(this.user[userId]) == 0{
		this.user[userId] = append(this.user[userId], userId)
	}
	res := []int{}
	count := 0
	for i := len(this.twi)-1; i >= 0 && count < 10 ; i--{

		for _,a := range this.user[userId]{
			if a == this.twi[i][0]{
				res = append(res, this.twi[i][1])
				count ++
			}
		}
	}
	return res
}
func (this *Twitter) Follow(followerId int, followeeId int)  {
	if followeeId == followerId{
		return
	}
	if len(this.user[followerId]) == 0{
		this.user[followerId] = append(this.user[followerId], followerId)
	}
	for _,v := range this.user[followerId]{
		if v == followeeId{
			return
		}
	}
	this.user[followerId] = append(this.user[followerId], followeeId)
}
func (this *Twitter) Unfollow(followerId int, followeeId int)  {
	if followeeId == followerId || len(this.user[followerId]) == 0 {
		return
	}
	if this.user[followerId][len(this.user[followerId])-1] == followeeId{
		this.user[followerId] = this.user[followerId][:len(this.user[followerId])-1]
	}else{

		for i := 1 ; i < len(this.user[followerId])-1 ; i++{
			if this.user[followerId][i] == followeeId{
				this.user[followerId] = append(this.user[followerId][0:i],this.user[followerId][i+1:]...)
			}
		}
	}


}


func main() {

	obj := Constructor()
	obj.PostTweet(2,5)
	obj.Follow(1,2)
	obj.Follow(1,2)
	//obj.PostTweet(2,6)

	fmt.Println(obj.GetNewsFeed(1))
	//obj.Unfollow(1, 2)
	////fmt.Println(obj.user,obj.twi)
	//fmt.Println(obj.GetNewsFeed(1))

}

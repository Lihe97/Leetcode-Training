package main

import (
	"fmt"
	"strconv"
)

type UndergroundSystem struct {
	station map[string]map[int]int
}

func Constructor() UndergroundSystem {
	obj := make(map[string]map[int]int)
	return UndergroundSystem{station:obj}
}


func (this *UndergroundSystem) CheckIn(id int, stationName string, t int)  {
	if this.station[stationName] == nil{
		this.station[stationName] = make(map[int]int)
	}
	this.station[stationName][id] = t
}


func (this *UndergroundSystem) CheckOut(id int, stationName string, t int)  {
	if this.station[stationName] == nil{
		this.station[stationName] = make(map[int]int)
	}
	this.station[stationName][id] = t
}


func (this *UndergroundSystem) GetAverageTime(startStation string, endStation string) float64 {
	sum := 0.0
	count := 0.0
	mpstart := this.station[startStation]
	mpend := this.station[endStation]
	for id,t := range mpstart{

		if mpend[id] != 0{
			count ++
			sum = sum - float64(t) + float64(mpend[id])
		}
	}
	value := float64(sum/count)
	value, _ = strconv.ParseFloat(fmt.Sprintf("%.1f", value), 64)

	return value
}


/**
 * Your UndergroundSystem object will be instantiated and called as such:
 * obj := Constructor();
 * obj.CheckIn(id,stationName,t);
 * obj.CheckOut(id,stationName,t);
 * param_3 := obj.GetAverageTime(startStation,endStation);
 */
func main() {
	obj := Constructor()
	obj.CheckIn(45,"a",2)
	obj.CheckIn(46,"a",4)
	obj.CheckOut(45,"b",6)
	obj.CheckOut(46,"b",8)
	fmt.Println(obj.GetAverageTime("a","b"))
	
}

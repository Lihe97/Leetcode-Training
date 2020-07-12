package main

import (
	"fmt"
	"strings"
)

func reformatDate(date string) string {

	res := ""
	s := strings.Split(date," ")
	res += s[2]
	res += "-"
	if s[1] == "Jan"{
		res += "01"
	}else if s[1] == "Feb"{
		res += "02"
	}else if s[1] == "Mar"{
		res += "03"
	}else if s[1] == "Apr"{
		res += "04"
	}else if s[1] == "May"{
		res += "05"
	}else if s[1] == "Jun"{
		res += "06"
	}else if s[1] == "Jul"{
		res += "07"
	}else if s[1] == "Aug"{
		res += "08"
	}else if s[1] == "Sep"{
		res += "09"
	}else if s[1] == "Oct"{
		res += "10"
	}else if s[1] == "Nov"{
		res += "11"
	}else if s[1] == "Dec"{
		res += "12"
	}
	res += "-"
	if len(s[0]) == 3{
		res += "0"
		res += string(s[0][0])
	}else{
		res += string(s[0][0])
		res += string(s[0][1])
	}
	return res




}

func main() {

	s := strings.Split("20th Oct 2052"," ")
	fmt.Println(s)
}

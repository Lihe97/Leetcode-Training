package main

type ThroneInheritance struct {

	kingname string
	mp map[string][]string

}


func Constructor(kingName string) ThroneInheritance {


	mp := map[string][]string{}
	mp[kingName] = append(mp[kingName],"1")

	return ThroneInheritance{
		kingname:kingName,
		mp:    mp,
	}
}


func (this *ThroneInheritance) Birth(parentName string, childName string)  {

	this.mp[parentName] = append(this.mp[parentName],childName)

	this.mp[childName] = []string{"1"}

}


func (this *ThroneInheritance) Death(name string)  {

	this.mp[name][0] = "0"

}


func (this *ThroneInheritance) GetInheritanceOrder() []string {

	res := []string{}



	getOrder(this.mp,&res,this.kingname)
	return res

}

func getOrder(mp map[string][]string,res *[]string,cur string){

	if mp[cur][0] == "1"{
		*res = append(*res,cur)

	}
	if len(mp[cur]) == 1 {
		return
	}

	for i := 1 ; i < len(mp[cur]) ; i ++{
		getOrder(mp,res,mp[cur][i])
	}

}
func main() {
	
}

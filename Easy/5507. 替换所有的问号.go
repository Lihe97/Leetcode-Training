package main
func modifyString(s string) string {



	arr := []byte(s)

	for i := 0; i < len(arr); i++{
		if arr[i] == '?'{
			for  j := 0; j < 26; j++{
				if((i == 0 || i > 0 && j != int(arr[i - 1] - 'a')) && (i + 1 == len(s) || (i + 1 < len(s) && (arr[ i + 1] == '?' || j != int(arr[i + 1] - 'a'))))){
					arr[i] = byte(j + 'a')
					break
				}
			}
		}
	}
	str := string(arr)
	return str
}

func main() {
	
}

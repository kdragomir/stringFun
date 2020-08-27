package main

import(
  "fmt"
  "strings"
)
func isPalindrome(str string)(result bool){
  var reverseStr string = ""
	var length = len(str)

	for i := length - 1; i >= 0; i -- {
		reverseStr = reverseStr + string(str[i])
	}
	// Case in-sensitive comparision
	if strings.ToLower(str) == strings.ToLower(reverseStr) {
    return true
	}
  return false
}

func main(){
fmt.Println(isPalindrome("racecar"))
}

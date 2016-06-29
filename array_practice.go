package main

import (
	"fmt"
	"sort"
)

var myNums = []int{6, 500, 12, 9, 8, 92, 1, 1300}

func main() {

	for ii := 0; ii < len(myNums)/2 ; ii++ {
		jj := len(myNums) - ii - 1
		tmp:= myNums[ii]
		myNums[ii] = myNums[jj]
		myNums[jj] = tmp
		fmt.Println(jj,myNums)
	}
	fmt.Println(myNums)
	sort.Sort(myNums)
	fmt.Print(myNums)
}

package main

import "fmt"

func main() {
	arr :=[2]int{3,5}
	cont := 0
	for i:=1;i<1000;i++{
		for j:=0;j<len(arr);j++{
			if(i%arr[j]==0){
				// fmt.Println(arr[j],i)
				cont+=i
				break
			}
		}
	}
	fmt.Println(cont)
}


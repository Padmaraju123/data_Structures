/*Given an array of random numbers, Push all the zero’s of a given array to the end of the array. 
For example, if the given arrays is {1, 9, 8, 4, 0, 0, 2, 7, 0, 6, 0}, 
it should be changed to {1, 9, 8, 4, 2, 7, 6, 0, 0, 0, 0}. The order of all other elements should be same. 
Expected time complexity is O(n) and extra space is O(1).

Example: 

Input :  arr[] = {1, 2, 0, 4, 3, 0, 5, 0};
Output : arr[] = {1, 2, 4, 3, 5, 0, 0, 0};

Input : arr[]  = {1, 2, 0, 0, 0, 3, 6};
Output : arr[] = {1, 2, 3, 6, 0, 0, 0}; */

package main

import "fmt"

func zero_end(sli []int){
	s1 := []int{}
	s2 := []int{}

	for _,v := range sli{
		if v == 0{
			s2 = append(s2, v)
		}else{
			s1 = append(s1, v)
		}
	}
	s3 := append(s1,s2...)
	fmt.Println(s3)

}

func main(){
	sli := []int{1, 2, 0, 4, 3, 0, 5, 0}
	zero_end(sli)
	
}


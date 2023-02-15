package main

import "fmt"

type heap struct{
	maxsize int   
	sli []int
	csize int  
}

func (obj *heap)length()int{
	return len(obj.sli)
}

func (obj *heap)isempty()bool{
	if obj.length()==0{
		return true
	}
	return false
}

func (obj *heap)insert(val int){
	if obj.maxsize == obj.csize{
		fmt.Println("No space is available")
		return
	}

	obj.csize = obj.csize+1
	hi := obj.csize  

	for hi>1 && val > obj.sli[hi/2]{
		obj.sli[hi] = obj.sli[hi/2]
		hi = hi/2
	}  

	obj.sli[hi] = val    
}

func (obj *heap)max()int{
	if obj.csize==0{
		fmt.Println("heap is empty")
		return 0
	}
	return obj.sli[1]
}

func (obj *heap)deleteMax()int{
	val := obj.sli[1]  //21
	obj.sli[1] = obj.sli[obj.csize]  //2
	obj.sli[obj.csize]=0
	obj.csize--

	i:=1   //1
	j:=i*2 //2

	for j<=obj.csize{
		if obj.sli[j] < obj.sli[j+1]{
			j = j+1
		}
		if obj.sli[i] < obj.sli[j]{
			temp := obj.sli[i]
			obj.sli[i] = obj.sli[j]
			obj.sli[j] = temp
			i = j
			j = i*2
		}else{
			break
		}
	}
	return val
}


func main(){
	obj := heap{}

	fmt.Println("enter the length of the heap")
	fmt.Scanln(&obj.maxsize)

	obj.sli = make([]int,obj.maxsize)
	fmt.Println(obj.sli)

	sl := []int{20,14,2,15,10,21}

	for _,v := range sl{
		obj.insert(v)
	}

	fmt.Println(obj.sli)

	fmt.Println(obj.max())

	fmt.Println(obj.deleteMax())

	fmt.Println(obj.sli)
}
package main

//Write a function rotate(ar[], d, n) that rotates arr[] of size n by d elements.

func rotate(a []int, d int) []int{
		return append(a[d:],a[:d]...)
}

func main()  {
	printArray(rotate([]int{1,2,3,4,6,7,8,9}, 2))
}

package main

import "fmt"

// slice 本身没有数据是对底层数据的一个view --> view of array
//slice <ptr开始指针,len,cap从开始指针到结束的arr>
func main() {
	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7}
	fmt.Println("arr[2:6] =", arr[2:6])
	fmt.Println("arr[:6] =", arr[:6])

	s1 := arr[2:]
	fmt.Println(s1)
	s2 := arr[:]
	fmt.Println(s2)
	fmt.Println("After updateSlice(s1)")
	updateSlice(s1)
	fmt.Println(s1)
	fmt.Println(arr)
	fmt.Println("After updateSlice(s2)")
	updateSlice(s2)
	fmt.Println(s2)
	fmt.Println(arr)

	fmt.Println("Reslice")
	fmt.Println(s2)
	s2 = s2[:5]
	fmt.Println(s2)
	s2 = s2[2:]
	fmt.Println(s2)

	fmt.Println("slice ----")
	arr = [...]int{0, 1, 2, 3, 4, 5, 6, 7}
	s1 = arr[2:6]
	fmt.Println(s1)
	s2 = s1[3:5]
	fmt.Println(s2)
	fmt.Printf("s1=%v,len(s1)=%d,cap(s1)=%d\n", s1, len(s1), cap(s1))
	s3 := append(s2, 10)
	s4 := append(s3, 11)
	s5 := append(s4, 12)
	// s4 & s5 no longer view arr.
	// 添加元素添加超过cap，系统会重新分配更大的底层数组
	fmt.Println(s3, s4, s5)
	fmt.Println(arr)


}

func updateSlice(s []int) {
	s[0] = 100
}

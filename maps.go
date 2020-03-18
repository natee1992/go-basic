package main

import "fmt"

func main() {
	m := map[string]string{
		"name":    "ccmouse",
		"course":  "golang",
		"site":    "imooc",
		"quality": "notbad",
	}
	m2 := make(map[string]int)
	var m3 map[string]int

	fmt.Println(m, m2, m3)

	fmt.Println("Traversing map")
	for k, v := range m {
		fmt.Println(k, v)
	}

	fmt.Println("Getting values")
	courseName, ok := m["course"]
	fmt.Println(courseName, ok)
	courseName, ok = m["cause"]
	fmt.Println(courseName, ok)
	fmt.Println("Deleting values")
	name, ok := m["name"]
	fmt.Println(name, ok)
	delete(m,"name")
	name, ok = m["name"]
	fmt.Println(name, ok)
}

package basic

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	fmt.Println(
		converToBin(5),
		converToBin(12),
		converToBin(13))
	printFile("abc.txt")
	a, b := 3, 4
	swap(&a, &b)
	fmt.Println(a, b)
}

func converToBin(n int) string {
	result := ""
	for ; n > 0; n /= 2 {
		lsb := n % 2
		result = strconv.Itoa(lsb) + result
	}
	return result
}

func printFile(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

}

func swap(a, b *int) {
	*b, *a = *a, *b
}

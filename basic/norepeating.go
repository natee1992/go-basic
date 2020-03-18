package basic

import "fmt"

/*
寻找最长不含有重复字符的子串
对于每一个字母x：
	- lastOccurred[x] 不存在，或者< start 无需操作
	- lastOccurred[x] >= start   更新start
	- 更新lastOccurred[x] 更新maxLength
 */
func main() {
	fmt.Println(lengthOfNonRepeatingSubStr("abab"))
	fmt.Println(lengthOfNonRepeatingSubStr("你好"))
}

func lengthOfNonRepeatingSubStr(s string) int {
	lastOccured := make(map[byte]int)
	start := 0
	maxLength := 0
	for i, ch := range []byte(s) {
		lastI, ok := lastOccured[ch]
		if ok && lastI >= start {
			start = lastOccured[ch] + 1
		}
		if i-start+1 > maxLength {
			maxLength = i - start + 1
		}
		lastOccured[ch] = i
	}
	return maxLength
}

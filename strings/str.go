package strutils

import (
	"bytes"
	"fmt"
	"strings"
)

// func main() {
// 	var s = []byte{84, 104, 101, 32, 97, 110, 115, 119, 101, 114, 32, 105, 115, 32, 53, 48, 33}
// 	fmt.Printf("%s\n", string(s))
// 	s1 := "12345.123"
// 	dot := strings.LastIndex(s1, ".")
// 	fmt.Printf("s1: %v\n", s1)
// 	fmt.Printf("s1: %v\n", s1[dot:])
// 	fmt.Printf("s1: %v\n", s1[:dot])
// 	fmt.Printf("comma(s1): %v\n", comma(s1))
// 	fmt.Printf("comma(\"12234423.12312.231\"): %v\n", comma("12234423.12312.231"))

// 	fmt.Printf("intsToString([]int{1, 2, 3, 4, 5, 6}): %v\n", intsToString([]int{1, 2, 3, 4, 5, 6}))

// 	fmt.Printf("isAnagram(\"daniel\", \"soler\"): %v\n", isAnagram("daniel", "soler"))
// 	fmt.Printf("isAnagram(\"daniel\", \"leinad\"): %v\n", isAnagram("daniel", "leinad"))
// 	fmt.Printf("isAnagram(\"daniel\", \"daniell\"): %v\n", isAnagram("daniel", "daniell"))
// 	fmt.Printf("isAnagram(\"daniell\", \"daniel\"): %v\n", isAnagram("daniell", "daniel"))
// 	fmt.Printf("isAnagram(\"aee\", \"aae\"): %v\n", isAnagram("aee", "aae"))
// 	fmt.Printf("isAnagram(\"qwer\", \"qwer\"): %v\n", isAnagram("qwer", "qwer"))
// 	fmt.Printf("isAnagram(\"ee\", \"ea\"): %v\n", isAnagram("ee", "ea"))
// 	fmt.Printf("isAnagram(string([]byte{51, 0, 56}), string([]byte{51, 32, 56})): %v\n", isAnagram(string([]byte{51, 0, 56}), string([]byte{51, 32, 56})))
// }

func testRunes() {
	s1 := "aẞẠẪ"
	fmt.Println("")
	fmt.Printf("% x\n", s1) // "e3 83 97 e3 83 ad e3 82 b0 e3 83 a9 e3 83 a0"
	r := []rune(s1)
	fmt.Printf("%x\n", r) // "[30d7 30ed 30b0 30e9 30e0]"

	// fmt.Println(string(0x4eac))
}

// comma inserts commas in a non-negative decimal integer string.
func comma(s string) string {
	var sdot string
	if dot := strings.LastIndex(s, "."); dot >= 0 {
		sdot = s[dot:]
		s = s[:dot]
	}
	n := len(s)
	if n <= 3 {
		return s
	}
	return comma(s[:n-3]) + "," + s[n-3:] + sdot
}

func iterComma(s string) string {
	var b bytes.Buffer
	fmt.Printf("b: %v\n", b)
	return b.String()
}

// intsToString is like fmt.Sprint(values) but adds commas.
func intsToString(values []int) string {
	var buf bytes.Buffer
	buf.WriteByte('[')
	for i, v := range values {
		if i > 0 {
			buf.WriteString(", ")
		}
		fmt.Fprintf(&buf, "%d", v)
	}
	buf.WriteByte(']')
	return buf.String()
}

func IsAnagram(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	bs2 := []byte(s2)
	bs1 := []byte(s1)
	for _, b := range bs1 {
		if i := bytes.IndexByte(bs2, b); i >= 0 {
			bs2[i] = bs2[len(bs2)-1]
			bs2 = bs2[:len(bs2)-1]
		} else {
			return false
		}
	}
	return true
}

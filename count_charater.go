package main
import (
	"fmt"
	"unicode/utf8"
)

func main() {
	string1 := "asSASA ddd dsjkdsjs dk"
	fmt.Printf("string1 lens %d:\n", len(string1))
	fmt.Printf("a %d\n", utf8.RuneCountInString(string1))
	string2 := "asSASA ddd dsjkdsjsこん dk"
	fmt.Printf("string1 lens %d:\n", len(string2))
	fmt.Printf("a %d\n", utf8.RuneCountInString(string2))
}

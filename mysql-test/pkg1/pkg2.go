package pkg1

import "fmt"

func PkgTest() {
	fmt.Println("PkgTest")
}

func ForInStr(input string) {
	for po, char := range input {
		fmt.Println(po, string(char))
	}
}

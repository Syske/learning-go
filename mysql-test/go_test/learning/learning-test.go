package learning

import "fmt"

func Func_args(args ...string) {
	fmt.Println(args)
	fmt.Println(len(args))
	fmt.Println(args[2])

}

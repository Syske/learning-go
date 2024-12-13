package learning

import "fmt"

func PrintTest(str string) {
	fmt.Println("PrintTest\t" + str)
}

func CallBack(index string, f func(string)) {
	f(index)
}

func CallBackTest() {
	CallBack("callBack", PrintTest)
}

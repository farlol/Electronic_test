package main

import (
	"fmt"
	"moduleelectronic/electronic"
)

func main() {
	Apple := electronic.NewApple("11Pro", "ios14.7")
	Android := electronic.NewAndroid("Samsung", "Galaxy12", "android9")
	Radio := electronic.NewRadio("Panasonic", "xm8", 12)
	printCharacteristics(Apple)
	//fmt.Println("OS:", Apple.OS())
	printCharacteristics(Android)
	//fmt.Println("OS:", Android.OS())
	printCharacteristics(Radio)
}

func printCharacteristics(a electronic.Phone) {
	fmt.Println("Brand:", a.Brand(), "Model:", a.Model(), "Type:", a.Type())
	switch a.(type) {
	case electronic.SmartPhone:
		temp := a.(electronic.SmartPhone)
		fmt.Println("OS:", temp.OS())
	case electronic.StationPhone:
		temp := a.(electronic.StationPhone)
		fmt.Println("Buttonscount:", temp.ButtonsCount())
	default:
		fmt.Println("Unknown interface")
	}

}

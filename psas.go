package main

import "fmt"

type music struct {
	instrument string
	wood       string
}

func main() {

	myguitar := [3]string{"guitar", "rosewood"}
	fmt.Println(myguitar)
	s := myguitar[:1]
	fmt.Println(s)
	romelu := music{"Classical guitar", "maple"}
	fmt.Println(romelu)
	p := &romelu
	p.wood = "Spruce"
	p.instrument = s[0]
	fmt.Println(romelu)

}

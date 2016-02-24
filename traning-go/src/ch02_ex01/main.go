package main

import (
	"fmt"
	"ch02_ex01/tmpconv"
)
func main() {
	fmt.Printf("AbsoluteZeroC:%v\n", tmpconv.AbsoluteZeroC)
	fmt.Printf("FreazingC:%v\n", tmpconv.FreazingC)
	fmt.Printf("BoilingC:%v\n", tmpconv.BoilingC)
	
	fmt.Println(tmpconv.CToF(tmpconv.AbsoluteZeroC))
	fmt.Println(tmpconv.CToF(tmpconv.FreazingC))
	fmt.Println(tmpconv.CToF(tmpconv.BoilingC))
	
	fmt.Printf("AbsoluteZeroF:%v\n", tmpconv.AbsoluteZeroF)
	fmt.Printf("FreazingF:%v\n", tmpconv.FreazingF)
	fmt.Printf("BoilingF:%v\n", tmpconv.BoilingF)
	
	fmt.Println(tmpconv.FToC(tmpconv.AbsoluteZeroF))
	fmt.Println(tmpconv.FToC(tmpconv.FreazingF))
	fmt.Println(tmpconv.FToC(tmpconv.BoilingF))
}
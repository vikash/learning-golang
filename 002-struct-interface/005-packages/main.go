package main

import S "github.com/vikash/learning-golang/002-struct-interface/005-packages/Shapes"

func main() {

	var a, b S.Shape

	a = S.Circle{1}
	b = S.Rectangle{1, 2}

	S.PrintArea(a, b)
}

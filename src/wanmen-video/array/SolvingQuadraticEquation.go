package array

import (
	"fmt"
	"math"
)

/**
求解一元二次方程，输入a，b，c三个系数，根据判别式b^2-4ac求解ax^2+bx+c=0
 */
func solvingQuadraticEquations( a,b,c int){
	k := b*b -4*a*c
	var x1, x2 int
	if k > 0 {
		x1 = (-b + int(math.Sqrt(float64(k))))/2*a
		x2 = (-b - int(math.Sqrt(float64(k))))/2*a
		fmt.Println("方程：",a,"x^2+",b,"x+",c,"=0","求解结果为","x1=",x1,";x2=",x2)
	}else if k == 0 {
		x1 = (-b + int(math.Sqrt(float64(k))))/2*a
		x2 = x1
		fmt.Println("方程：",a,"x^2+",b,"x+",c,"=0","求解结果为","x1=",x1,";x2=",x2)
	}else{
		fmt.Println("方程：",a,"x^2+",b,"x+",c,"=0","在实数范围内无解")
	}


}

package array

import (
	. "testing"
	"fmt"
)

/**
求解一元二次方程，输入a，b，c三个系数，根据判别式求解ax^2+bx+c=0
 */
func TestSolvingQuadraticEquations( test *T){
	 abc := [][]int{
	 	{1,2,4},
	 	{1,5,6},
	 	{1,2,1},
	 	{1,-5,6},
	 	{1,2,3},
	 }

	 for _,k := range abc {
	 	fmt.Println(k[0],k[1],k[2])
		 solvingQuadraticEquations(k[0],k[1],k[2],)
	 }



}

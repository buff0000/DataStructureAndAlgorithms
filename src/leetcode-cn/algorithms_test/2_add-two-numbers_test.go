package algorithms

import (
	. "leetcode-cn/algorithms"
	"testing"
	"fmt"
)

func TestAddTwoNumbers(t *testing.T)  {
	//var l1 *ListNode = &ListNode{Val:1, Next:&ListNode{Val:4, Next:&ListNode{Val:7, Next:&ListNode{Val:9, Next:nil}}}}
	//var l2 *ListNode = &ListNode{Val:1, Next:&ListNode{Val:4, Next:&ListNode{Val:7, Next:&ListNode{Val:9, Next:nil}}}}

	l1 := ListNode{Val:1, Next:nil}
	l1_2 := ListNode{Val:5, Next:nil}
	l1.Next = &l1_2
	l1_3 := ListNode{Val:3, Next:nil}
	l1_2.Next = &l1_3
	l1_4 := ListNode{Val:8, Next:nil}
	l1_3.Next = &l1_4
	l1_5 := ListNode{Val:1, Next:nil}
	l1_4.Next = &l1_5


	l2 := ListNode{Val:1, Next:nil}
	l2_2 := ListNode{Val:5, Next:nil}
	l2.Next = &l2_2
	l2_3 := ListNode{Val:3, Next:nil}
	l2_2.Next = &l2_3
	l2_4 := ListNode{Val:8, Next:nil}
	l2_3.Next = &l2_4
	l2_5 := ListNode{Val:1, Next:nil}
	l2_4.Next = &l2_5

	var l3 = AddTwoNumbers(&l1, &l2, 0)

	temp := l3
	for  {
		fmt.Println(temp.Val)
		if temp.Next == nil {
			break
		}else{
			temp = temp.Next
		}
	}
}

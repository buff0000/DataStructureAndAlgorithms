package algorithms

/*
给出两个 非空 的链表用来表示两个非负的整数。其中，它们各自的位数是按照 逆序 的方式存储的，并且它们的每个节点只能存储 一位 数字。

如果，我们将这两个数相加起来，则会返回一个新的链表来表示它们的和。

您可以假设除了数字 0 之外，这两个数都不会以 0 开头。

示例：

输入：(2 -> 4 -> 3) + (5 -> 6 -> 4)
输出：7 -> 0 -> 8
原因：342 + 465 = 807

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/add-two-numbers
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

/**
 * Definition for singly-linked list.
*/
type ListNode struct {
	Val  int
	Next *ListNode
}

func AddTwoNumbers(l1 *ListNode, l2 *ListNode, carry int) *ListNode {
	var v3 int
	var n1 *ListNode
	var n2 *ListNode
	if l1 == nil && l2 == nil {
		if carry == 0 {
			return nil
		}else{
			return &ListNode{Val:carry,Next:nil}
		}
		return nil
	} else if l1 == nil {
		v3 = l2.Val + carry
		n1 = nil
		n2 = l2.Next
	} else if l2 == nil {
		v3 = l1.Val + carry
		n2 = nil
		n1 = l1.Next
	} else {
		v3 = l1.Val + l2.Val + carry
		n1 = l1.Next
		n2 = l2.Next
	}

	var node = &ListNode{
		Val:  v3 % 10,
		Next: AddTwoNumbers(n1, n2, v3 / 10),
	}

	return node
}

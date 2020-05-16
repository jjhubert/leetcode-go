package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type ListNode struct {
	Val int
	Next *ListNode
}

/**
变量声明：
	a 保存的是l1链表中节点的value
	b 保存的是l2链表中节点的value
	c 保存的是a+b是否进1的值，默认是0
	必须在开头新建节点和保存指向开头的指针，以便返回
思路：
	1、判断链表l1和l2是否为空且进位标志是否为1，若都满足退出循环
	2、迭代l1、l2分别取出每个节点的值并计算
	3、计算结果存入链表l3
*/
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	l3 := new(ListNode)
	res := l3
	var a, b, c, tmpVal int
	for  {
		if l1 == nil && l2 == nil && c == 0 {
			break
		}
		l3.Next = new(ListNode)
		l3 = l3.Next
		if l1 != nil {
			a = l1.Val
			l1 = l1.Next
		} else {
			a = 0
		}
		if l2 != nil {
			b = l2.Val
			l2 = l2.Next
		} else {
			b = 0
		}
		l3.Val = (a + b + c) % 10
		tmpVal = (a + b + c) / 10
		if tmpVal == 1 {
			c = 1
		} else {
			c = 0
		}
	}
	return res.Next
}

func main() {
	l1 := &ListNode{2, nil}
	l2 := &ListNode{4, nil}
	l3 := &ListNode{3, nil}
	l1.Next = l2
	l2.Next = l3
	n1 := &ListNode{5, nil}
	n2 := &ListNode{6, nil}
	n3 := &ListNode{4, nil}
	n1.Next = n2
	n2.Next = n3
	addTwoNumbers(l1, n1)
}
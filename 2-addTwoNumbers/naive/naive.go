package addTwoNumbers_naive

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	l3 := &ListNode{}
	l3Curr := l3
	for l1 != nil || l2 != nil {
		if l1 != nil {
			l3Curr.Val += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			l3Curr.Val += l2.Val
			l2 = l2.Next
		}
		if l3Curr.Val >= 10 {
			l3Curr.Val = l3Curr.Val % 10
			l3Curr.Next = &ListNode{Val: 1}
		} else {
			l3Curr.Next = &ListNode{}
		}
		l3Curr = l3Curr.Next
	}
	return l3
}

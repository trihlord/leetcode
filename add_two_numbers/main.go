package add_two_numbers

import (
	"strconv"
	"strings"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func (n *ListNode) String() string {
	var sb strings.Builder
	sb.WriteString("[")
	c := n
	for c != nil {
		if c != n {
			sb.WriteString(" ")
		}
		sb.WriteString(strconv.Itoa(c.Val))
		c = c.Next
	}
	sb.WriteString("]")
	return sb.String()
}

func NewListNode(vals ...int) *ListNode {
	h := &ListNode{Val: vals[0], Next: nil}
	c := h
	for _, vv := range vals[1:] {
		t := &ListNode{Val: vv, Next: nil}
		c.Next = t
		c = t
	}
	return h
}

func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	h := new(ListNode)
	c := h
	s := 0
	for l1 != nil || l2 != nil || s > 0 {
		var v1, v2 int
		if l1 != nil {
			v1 = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			v2 = l2.Val
			l2 = l2.Next
		}
		v := s + v1 + v2
		s = v / 10
		c.Next = &ListNode{Val: v % 10, Next: nil}
		c = c.Next
	}
	return h.Next
}

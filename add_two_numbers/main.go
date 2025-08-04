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
	c1, c2 := l1, l2
	s := c1.Val + c2.Val
	v := s
	if v > 9 {
		v %= 10
	}
	h := &ListNode{Val: v, Next: nil}
	c := h
	s /= 10
	for c1.Next != nil || c2.Next != nil {
		c1, c2 = c1.Next, c2.Next
		if c1 == nil {
			c1 = new(ListNode)
		}
		if c2 == nil {
			c2 = new(ListNode)
		}
		s += c1.Val + c2.Val
		v := s
		if v > 9 {
			v %= 10
		}
		t := &ListNode{Val: v, Next: nil}
		c.Next = t
		c = t
		s /= 10
	}
	if s > 0 {
		t := &ListNode{Val: s, Next: nil}
		c.Next = t
		c = t
	}
	return h
}

package addtwonumbers

import (
	"testing"
)

func Test_add_two_numbers(t *testing.T) {

	l1 := ListNode{Val: 2}
	l2 := ListNode{Val: 4}
	l3 := ListNode{Val: 3}
	l1.Next = &l2
	l2.Next = &l3

	r1 := ListNode{Val: 5}
	r2 := ListNode{Val: 6}
	r3 := ListNode{Val: 4}
	r1.Next = &r2
	r2.Next = &r3

	r := addTwoNumbers(&l1, &r1)
	if !(r.Val == 7 && r.Next.Val == 0 && r.Next.Next.Val == 8) {
		t.Error("add_two_numbers has err")
	}

}

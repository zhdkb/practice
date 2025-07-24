package list

func ChangePretoNext(p *ListNode) *ListNode {
	if p == nil || p.Next == p {
		return p
	}

	pre := p
	for pre.Next.Next != p {
		pre = pre.Next
	}

	ppre, next := pre.Next, p.Next
	ppre.Next = next
	p.Next = ppre
	pre.Next = p
	return p

}

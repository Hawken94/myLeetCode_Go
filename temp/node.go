package temp

// Node ...
type Node struct {
	Val  int
	Next *Node
}

//插入一个节点
//h: 头结点
//d:要插入的节点
//p:要插入的位置
func Insert(h, d *Node, p int) bool {
	if h.Next == nil {
		h.Next = d
		return true
	}

	i := 0
	n := h
	for n.Next != nil {
		i++
		if i == p {
			if n.Next.Next == nil {
				n.Next = d
				return true
			} else {
				d.Next = n.Next
				n.Next = d.Next
				return true
			}
		}
		n = n.Next
		if n.Next == nil {
			n.Next = d
			return true
		}
	}
	return false
}

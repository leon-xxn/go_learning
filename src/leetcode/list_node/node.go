package list_node

//单链表
type ListNode struct {
	val  int
	next *ListNode
}

//双链表
type Node[T any] struct {
	Val  T
	next *Node[T]
	prev *Node[T]
}

func NewNode[T any](prev *Node[T], next *Node[T], val T) *Node[T] {
	return &Node[T]{
		Val:  val,
		next: next,
		prev: prev,
	}

}

//数组转链表
func creteLinkedList(arr []int) *ListNode {
	if arr == nil || len(arr) == 0 {
		return nil
	}
	head := &ListNode{val: arr[0]}
	cur := head
	for i := 1; i < len(arr); i++ {
		cur.next = &ListNode{val: arr[i]}
		cur = cur.next
	}
	return head
}

func AddLinkFirst() {
	head := &ListNode{val: 1}
	newHead := &ListNode{val: 0}
	newHead.next = head
	head = newHead
}

func AddLinkLast() {
	head := &ListNode{val: 1}
	cur := head
	for cur.next != nil {
		cur = cur.next
	}
	cur.next = &ListNode{val: 2}
}

func AddLink() {
	head := &ListNode{val: 1}
	p := head
	index := 2
	//index的前驱节点
	for i := 0; i < index-1; i++ {
		p = p.next
	}
	newNode := &ListNode{val: 3}
	newNode.next = p.next
	p.next = newNode

}

func RemoveLinkNode() {
	head := creteLinkedList([]int{1, 2, 3, 4, 5})
	p := head
	//考虑边界情况 1的时候
	index := 2
	//index的前驱节点
	for i := 0; i < index-2; i++ {
		p = p.next
	}
	p.next = p.next.next
}

func RemoveLastLinkNode() {
	head := creteLinkedList([]int{1, 2, 3, 4, 5})
	p := head
	for p.next.next != nil {
		p = p.next
	}
	p.next = nil
}

type DoubleListNode struct {
	Val        int
	Prev, Next *DoubleListNode
}

func NewDoublyListNode(x int) *DoubleListNode {
	return &DoubleListNode{Val: x}
}

func CreateDoublyListNode(arr []int) *DoubleListNode {
	if arr == nil || len(arr) == 0 {
		return nil
	}
	head := &DoubleListNode{Val: arr[0]}
	cur := head
	for i := 1; i < len(arr); i++ {
		cur.Next = &DoubleListNode{Val: arr[i]}
		cur.Next.Prev = cur
		cur = cur.Next
	}
	return head
}

func AddDoublyListFirstNode() {
	head := CreateDoublyListNode([]int{1, 2, 3, 4, 5})
	//头部插入
	newHead := &DoubleListNode{Val: 0}
	newHead.Next = head
	head.Prev = newHead
	head = newHead
}

func AddDoublyListLastNode() {
	head := CreateDoublyListNode([]int{1, 2, 3, 4, 5})
	cur := head
	for cur.Next != nil {
		cur = cur.Next
	}
	newNode := &DoubleListNode{Val: 6}
	cur.Next = newNode
	newNode.Prev = cur
	cur = newNode
}

func AddDoublyListNode() {
	head := CreateDoublyListNode([]int{1, 2, 3, 4, 5})
	index := 3
	p := head
	for i := 0; i < index-1; i++ {
		p = p.Next
	}
	newNode := &DoubleListNode{Val: 66}
	newNode.Prev = p
	newNode.Next = p.Next

	//p的下一个节点的前驱节点
	p.Next.Prev = newNode
	p.Next = newNode

}

func RemoveDoublyListNode() {
	head := CreateDoublyListNode([]int{1, 2, 3, 4, 5})
	index := 4
	//找到index的前驱节点
	p := head
	for i := 0; i < index-2; i++ {
		p = p.Next
	}
	//删除节点
	toDelete := p.Next
	p.Next = toDelete.Next
	if toDelete.Next != nil {
		toDelete.Next.Prev = p
	}
	toDelete.Next = nil
	toDelete.Prev = nil
}

func RemoveDoublyListFirstNode() {
	head := CreateDoublyListNode([]int{1, 2, 3, 4, 5})
	toDetelte := head
	head = head.Next
	head.Prev = nil
	toDetelte.Next = nil
}

func RemoveDoublyListLastNode() {
	head := CreateDoublyListNode([]int{1, 2, 3, 4, 5})
	tail := head
	for tail.Next != nil {
		tail = tail.Next
	}
	tail.Prev.Next = nil
	tail.Prev = nil
}

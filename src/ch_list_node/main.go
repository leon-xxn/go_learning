package ch_list_node

type ListNode struct {
	Val  int
	Next *ListNode
}

func createLinkedList(arr []int) *ListNode {
	if arr == nil || len(arr) == 0 {
		return nil
	}
	head := &ListNode{Val: arr[0]}
	cur := head
	for i := 1; i < len(arr); i++ {
		cur.Next = &ListNode{Val: arr[i], Next: head}
		cur = cur.Next
	}
	return head
}

func addHeadNode() {
	head := createLinkedList([]int{1, 2, 3, 4, 5})
	cur := &ListNode{Val: 1}
	cur.Next = head
	head = cur
}

type DoublyListNode struct {
	Val        int
	Prev, Next *DoublyListNode
}

func NewDoublyListNode(x int) *DoublyListNode {
	return &DoublyListNode{Val: x}
}

func CreateDoublyLinkedList(arr []int) *DoublyListNode {
	if arr == nil || len(arr) == 0 {
		return nil
	}
	head := NewDoublyListNode(arr[0])
	cur := head
	for i := 1; i < len(arr); i++ {
		newNode := NewDoublyListNode(arr[i])
		cur.Next = newNode
		newNode.Prev = cur
		cur = cur.Next
	}
	return head
}

func addDoublyLinkedList() {
	head := CreateDoublyLinkedList([]int{1, 2, 3, 4, 5})
	newHead := &DoublyListNode{Val: 0}
	newHead.Next = head
	head.Prev = newHead
	head = newHead

	tail := head
	for tail.Next != nil {
		tail = tail.Next
	}
	newNode := NewDoublyListNode(0)
	tail.Next = newNode
	newNode.Prev = tail
	tail = newNode

	p := head
	for i := 0; i < 2; i++ {
		p = p.Next
	}
	newNode = NewDoublyListNode(66)
	newNode.Next = p.Next
	newNode.Prev = p

	p.Next.Prev = newNode
	p.Next = newNode

	toDelete := p.Next
	p.Next = toDelete.Next
	if toDelete.Next != nil {
		toDelete.Next.Prev = p
	}

	toDelete.Next = nil
	toDelete.Prev = nil
}

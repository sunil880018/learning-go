package demo

type Node struct {
	Val  int
	Next *Node
}

func insertNodeAtLast() {

}

func main() {
	head := &Node{}
	head.insertNodeAtLast(10)
	head.insertNodeAtLast(150)
	head.insertNodeAtLast(120)
	head.insertNodeAtLast(30)
	head.insertNodeAtLast(35)
	head.insertNodeAtLast(19)
}

package main

/*
给出两个 非空 的链表用来表示两个非负的整数。其中，它们各自的位数是按照 逆序 的方式存储的，并且它们的每个节点只能存储 一位 数字。
如果，我们将这两个数相加起来，则会返回一个新的链表来表示它们的和。
您可以假设除了数字 0 之外，这两个数都不会以 0 开头。

示例：
输入：(2 -> 4 -> 3) + (5 -> 6 -> 4)
输出：7 -> 0 -> 8
原因：342 + 465 = 807
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var sum int
	var addOne int
	result := new(ListNode)
	head := result

	for l1 != nil || l2 != nil || addOne > 0 {
		result.Next = new(ListNode)
		result = result.Next
		if l1 != nil {
			sum = sum + l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum = sum + l2.Val
			l2 = l2.Next
		}
		sum = sum + addOne
		result.Val = sum % 10
		addOne = sum / 10
		sum = 0
	}
	return head.Next
}

/*
1. 单链表
1.1. 定义
单向链表（单链表）是链表的一种，其特点是链表的链接方向是单向的，对链表的访问要通过顺序读取从头部开始；链表是使用指针进行构造的列表；又称为结点列表，因为链表是由一个个结点组装起来的；其中每个结点都有指针成员变量指向列表中的下一个结点；
列表是由结点构成，head指针指向第一个成为表头结点，而终止于最后一个指向nuLL的指针。

1.2. 优点
单个结点创建非常方便，普通的线性内存通常在创建的时候就需要设定数据的大小
结点的删除非常方便，不需要像线性结构那样移动剩下的数据
结点的访问方便，可以通过循环或者递归的方法访问到任意数据，但是平均的访问效率低于线性表。
*/

/*
//链表的相关操作
//单链表 -- 线性表
//定义节点
type Node struct {
	Data int
	Next *Node
}

//返回第一个节点
//h 头结点
func GetFirst(h *Node) *Node {
	if h.Next == nil {
		return nil
	}
	return h.Next
}

//返回最后一个节点
//h 头结点
func GetLast(h *Node) *Node {
	if h.Next == nil {
		return nil
	}
	i := h
	for i.Next != nil {
		i = i.Next
		if i.Next == nil {
			return i
		}
	}
	return nil
}

//取长度
func GetLength(h *Node) int {
	var i int = 0
	n := h
	for n.Next != nil {
		i++
		n = n.Next
	}
	return i
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
			n.Next = d
			return true
		}
	}
	n = n.Next
	if n.Next == nil {
		n.Next = d
		return true
	}
	return false
}

//取出指定节点
func GetLoc(h *Node, p int) *Node {
	if p < 0 || p > GetLength(h) {
		return nil
	}
	var i int = 0
	n := h
	for n.Next != nil {
		i++
		n = n.Next
		if i == p {
			return n
		}
	}
	return nil
}

func main() {
	//初始化一个头结点
	var h Node
	//往链表插入10个元素
	for i := 1; i <= 10; i++ {
		var d Node
		d.Data = i
		Insert(&h, &d, i)
		fmt.Println(GetLoc(&h, i))
	}
	fmt.Println(GetLength(&h))
	fmt.Println(GetFirst(&h))
	fmt.Println(GetLast(&h))
	fmt.Println(GetLoc(&h, 6))
}
*/

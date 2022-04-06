package main

import "fmt"

//定义单链表结构体

type ListNode struct {
	data interface{} //数据
	next *ListNode   //下个指针
}
type ListTable struct {
	length   int //储存链表的长度
	headNode *ListNode
}

/*单链表的初始化*/
func InitList() *ListTable {
	//即构造一个空的单链表L(包含头指针)
	node := new(ListNode)
	L := new(ListTable)
	L.headNode = node
	return L
}

func (list *ListTable) LocateElem(v interface{}) bool {
	if list.IsNull() {
		fmt.Println("err")
		return false
	} else {
		pre := list.headNode
		for pre != nil {
			if pre.data == v {
				return true
			}
			pre = pre.next
		}
		return false
	}
}

func (list *ListTable) DeleteElem(index int) {
	if index <= 0 || index > list.length {
		fmt.Println("删除位置非法")
		return
	} else {
		pre := list.headNode
		if index == 1 {
			list.headNode = pre.next
		} else {
			pre := list.headNode
			for count := 1; count < index-1; count++ {
				pre = pre.next
			}
			pre.next = pre.next.next
		}
		list.length--
	}
}

func (list *ListTable) IsNull() bool {
	if list.length == 0 {
		return true
	} else {
		return false
	}
}

func (list *ListTable) AppendElem(v interface{}) {
	node := &ListNode{data: v}
	if list.IsNull() {
		list.headNode.next = node
	} else {
		cur := list.headNode
		for cur.next != nil {
			cur = cur.next
		}
		cur.next = node
	}
	list.length++
	return
}

func (list *ListTable) ShowList() {
	if !list.IsNull() {
		cur := list.headNode
		for {
			fmt.Printf("\t%v", cur.data)
			if cur.next != nil {
				cur = cur.next
			} else {
				break
			}
		}
		fmt.Printf("\n")
	}
}

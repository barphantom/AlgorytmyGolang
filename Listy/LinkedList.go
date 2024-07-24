package main

import (
	"fmt"
)

type node struct {
	value                  int
	nextNode, previousNode *node
}

func main() {
	var decision int = 999
	var head *node

	fmt.Println("Witam w liście wiązanej!\n1. Stwórz listę.\n2. Dodaj element.\n3. Usuń element.\n4. Wypisz elementy.\n0. Wyjście")
	for decision != 0 {
		fmt.Scanln(&decision)
		switch decision {
		case 1:
			head = initializeList()
		case 2:
			var value int
			fmt.Scanln(&value)
			addElement(head, value)
		case 3:
			// deleteElement()
		case 4:
			printList(head)
		default:
			fmt.Println("Nie rozpoznano wyboru.")
		}
	}
}

func initializeList() *node {
	return &node{
		value:        888,
		nextNode:     nil,
		previousNode: nil,
	}
}

func addElement(head *node, value int) {
	if head == nil {
		fmt.Println("Nie można dodać elementu. List nie istnieje.")
	} else if head.nextNode == nil {
		newNode := node{value: value, nextNode: nil, previousNode: nil}
		head.nextNode = &newNode
	} else {
		var tempNode *node = head
		for tempNode.nextNode != nil {
			tempNode = tempNode.nextNode
		}
		newNode := node{value: value, nextNode: nil, previousNode: tempNode}
		tempNode.nextNode = &newNode
	}
}

func printList(head *node) {
	if head == nil {
		fmt.Println("Nie można wyświetlić listy. List nie istnieje.")
	} else if head.nextNode == nil {
		fmt.Println("Lista jest pusta.")
	} else {
		var tempNode *node = head
		for tempNode.nextNode != nil {
			fmt.Println(tempNode.value)
			tempNode = tempNode.nextNode
		}
	}

}

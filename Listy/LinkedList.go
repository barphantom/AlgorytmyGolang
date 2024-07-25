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
	// head = initializeList()

	for decision != 0 {
		fmt.Println("\nWitam w algorytmie listy wiązanej!\n1. Stwórz listę.\n2. Dodaj element.\n3. Usuń element.\n4. Wypisz elementy.\n5. Usuń listę.\n0. Wyjście")
		fmt.Scanln(&decision)
		switch decision {
		case 1:
			if head != nil {
				fmt.Println("Nie można zainicjalizować listy. Lista już istnieje.")
			} else {
				head = initializeList()
			}
		case 2:
			var value int
			fmt.Scanln(&value)
			addElement(head, value)
		case 3:
			var value int
			fmt.Scanln(&value)
			deleteElement(head, value)
		case 4:
			printList(head)
		case 5:
			head = nil
			fmt.Println("\nUsunięto listę.")
		case 0:
			fmt.Println("\nWyjście z programu.")
		default:
			fmt.Println("Nie rozpoznano wyboru.")
		}
	}
}

func initializeList() *node {
	fmt.Println("\nUtworzono listę.")

	return &node{
		value:        888,
		nextNode:     nil,
		previousNode: nil,
	}
}

func addElement(head *node, value int) {
	if head == nil {
		fmt.Println("\nNie można dodać elementu. Lista nie istnieje.")
	} else if head.nextNode == nil {
		newNode := node{value: value, nextNode: nil, previousNode: nil}
		head.nextNode = &newNode
		fmt.Print("\nDodano element: ", value, "\n")
	} else {
		var tempNode *node = head.nextNode
		for tempNode.nextNode != nil {
			tempNode = tempNode.nextNode
		}
		newNode := node{value: value, nextNode: nil, previousNode: tempNode}
		tempNode.nextNode = &newNode
		fmt.Print("\nDodano element: ", value, "\n")
	}
}

func deleteElement(head *node, value int) {
	fmt.Print("\nSzukam węzeła o wartości: ", value, "\n")
	if head == nil {
		fmt.Println("Nie można usunąć elementu z listy. Lista nie istnieje.")
	} else if head.nextNode == nil {
		fmt.Println("Lista jest pusta. Nie ma elementów do usunięcia.")
	} else {
		tempNode := head.nextNode
		for tempNode.value != value && tempNode.nextNode != nil {
			tempNode = tempNode.nextNode
		}
		if tempNode.value == value && tempNode.nextNode == nil {
			if tempNode == head.nextNode {
				head.nextNode = nil
			} else {
				tempNode.previousNode.nextNode = nil
			}
			fmt.Print("Usunięto węzeł: ", value, "\n")
		} else if tempNode.value == value {
			if tempNode == head.nextNode {
				head.nextNode = tempNode.nextNode
				tempNode.nextNode.previousNode = nil
			} else {
				tempNode.previousNode.nextNode = tempNode.nextNode
				tempNode.nextNode.previousNode = tempNode.previousNode
			}
			fmt.Print("Usunięto węzeł: ", value, "\n")
		} else {
			fmt.Println("Nie znaleziono takiego węzła.")
		}
	}

}

func printList(head *node) {
	if head == nil {
		fmt.Println("\nNie można wyświetlić listy. Lista nie istnieje.")
	} else if head.nextNode == nil {
		fmt.Println("\nLista jest pusta.")
	} else {
		fmt.Print("\nWartości listy: ")
		var tempNode *node = head.nextNode
		for tempNode.nextNode != nil {
			fmt.Print(tempNode.value, " ")
			tempNode = tempNode.nextNode
		}
		fmt.Println(tempNode.value)

	}

}

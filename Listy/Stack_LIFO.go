// package pckLists

// import (
// 	"fmt"
// )

// type Node struct {
// 	value                  int
// 	nextNode, previousNode *Node
// }

// type Queue struct {
// 	head          *Node
// 	tail          *Node
// 	currentLength int
// }

// func LifoList() {
// 	var decision int = 999
// 	var queue *Queue

// 	for decision != 0 {
// 		fmt.Println("\nWitam w algorytmie listy wiązanej!\n1. Stwórz listę.\n2. Dodaj element.\n3. Usuń element.\n4. Wypisz elementy.\n5. Usuń listę.\n0. Wyjście")
// 		fmt.Scanln(&decision)
// 		switch decision {
// 		case 1:
// 			if queue != nil {
// 				fmt.Println("Nie można zainicjalizować listy. Lista już istnieje.")
// 			} else {
// 				queue = initializeList()
// 			}
// 		case 2:
// 			var value int
// 			fmt.Println("Wpisz nowy element listy.")
// 			fmt.Scanln(&value)
// 			addElement(queue, value)
// 		case 3:
// 			var value int
// 			fmt.Scanln(&value)
// 			deleteElement(queue, value)
// 		case 4:
// 			printList(queue)
// 		case 5:
// 			queue = nil
// 			fmt.Println("\nUsunięto listę.")
// 		case 0:
// 			fmt.Println("\nWyjście z programu.")
// 		default:
// 			fmt.Println("Nie rozpoznano wyboru.")
// 		}
// 	}
// }

// func initializeList() *Queue {
// 	fmt.Println("\nUtworzono listę.")

// 	return &Queue{
// 		head:          nil,
// 		tail:          nil,
// 		currentLength: 0,
// 	}
// }

// func addElement(queue *Queue, value int) {
// 	if queue == nil {
// 		fmt.Println("\nNie można dodać elementu. Lista nie istnieje.")
// 	} else if queue.currentLength == 0 {
// 		newNode := Node{value: value, nextNode: nil, previousNode: nil}
// 		queue.head = &newNode
// 		queue.tail = &newNode
// 		queue.currentLength += 1
// 		fmt.Print("\nDodano element: ", value, "\n")
// 	} else {
// 		newNode := Node{value: value, nextNode: nil, previousNode: queue.tail}
// 		queue.tail.nextNode = &newNode
// 		queue.tail = &newNode
// 		queue.currentLength += 1
// 		fmt.Print("\nDodano element: ", value, "\n")
// 	}
// }

// func deleteElement(queue *Queue, value int) {
// 	fmt.Print("\nSzukam węzeła o wartości: ", value, "\n")
// 	if queue.currentLength == 0 {
// 		fmt.Println("Lista jest pusta. Nie ma elementów do usunięcia.")
// 		return
// 	}

// 	tempNode := queue.head
// 	for tempNode != nil && tempNode.value != value {
// 		tempNode = tempNode.nextNode
// 	}

// 	if tempNode == nil {
// 		fmt.Printf("Nie znaleziono węzła o wartości: %d\n", value)
// 		return
// 	}

// 	if tempNode == queue.head {
// 		queue.head = tempNode.nextNode
// 		if tempNode.nextNode != nil {
// 			tempNode.nextNode.previousNode = nil
// 		}
// 	} else if tempNode == queue.tail {
// 		queue.tail = tempNode.previousNode
// 		if tempNode.previousNode != nil {
// 			tempNode.previousNode.nextNode = nil
// 		}
// 	} else {
// 		tempNode.nextNode.previousNode = tempNode.previousNode
// 		tempNode.previousNode.nextNode = tempNode.nextNode
// 	}
// 	queue.currentLength--
// 	fmt.Printf("Usunięto węzeł: %d\n", value)
// }

// func printList(queue *Queue) {
// 	if queue == nil {
// 		fmt.Println("\nNie można wyświetlić listy. Lista nie istnieje.")
// 	} else if queue.currentLength == 0 {
// 		fmt.Println("\nLista jest pusta.")
// 	} else {
// 		fmt.Print("\nWartości listy: ")
// 		var tempNode *Node = queue.head
// 		for tempNode.nextNode != nil {
// 			fmt.Print(tempNode.value, " ")
// 			tempNode = tempNode.nextNode
// 		}
// 		fmt.Println(tempNode.value)
// 	}
// }
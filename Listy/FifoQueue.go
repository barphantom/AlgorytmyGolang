package listy

import (
	"fmt"
)

type fifoNode struct {
	value                  int
	nextNode, previousNode *fifoNode
}

type queue struct {
	head          *fifoNode
	tail          *fifoNode
	currentLength int
}

func StartFifoQueue() {
	var decision int = 999
	var myQueue *queue

	for decision != 0 {
		fmt.Println("\nWitam w algorytmie listy wiązanej!\n1. Stwórz listę.\n2. Dodaj element.\n3. Usuń element.\n4. Wypisz elementy.\n5. Usuń listę.\n6. Wyświetl pierwszy element listy.\n7. Wyczyść kolejkę.\n8. Zdejmij pierwszy element.\n0. Wyjście")
		fmt.Scanln(&decision)
		switch decision {
		case 1:
			if myQueue != nil {
				fmt.Println("Nie można zainicjalizować listy. Lista już istnieje.")
			} else {
				myQueue = initializeList()
			}
		case 2:
			var value int
			fmt.Println("Wpisz nowy element listy.")
			fmt.Scanln(&value)
			enqueue(myQueue, value)
		case 3:
			var value int
			fmt.Scanln(&value)
			deleteElement(myQueue, value)
		case 4:
			printList(myQueue)
		case 5:
			myQueue = nil
			fmt.Println("\nUsunięto listę.")
		case 6:
			if myQueue == nil {
				fmt.Println("Nie można wyświetlić pierwszego elementu. Lista nie istnieje.")
			} else {
				value := front(myQueue)
				fmt.Println("Pierwszy element: ", value)
			}
		case 7:
			makeNull(myQueue)
			fmt.Println("Wyczyszczono listę.")
		case 8:
			if myQueue == nil {
				fmt.Println("Nie można zdjąć pierwszego elementu. Lista nie istnieje.")
			} else if empty(myQueue) {
				fmt.Println("Nie można zdjąć pierwszego elementu. Lista jest pusta.")
			} else {
				value := dequeue(myQueue)
				fmt.Println("Zdjęto element: ", value)
			}
		case 0:
			fmt.Println("\nWyjście z programu.")
		default:
			fmt.Println("Nie rozpoznano wyboru.")
		}
	}
}

func initializeList() *queue {
	fmt.Println("\nUtworzono listę.")

	return &queue{
		head:          nil,
		tail:          nil,
		currentLength: 0,
	}
}

func deleteElement(queue *queue, value int) {
	if queue == nil {
		fmt.Println("Nie można usunąć elementu. Lista nie istnieje.")
		return
	}
	fmt.Print("\nSzukam węzeła o wartości: ", value, "\n")
	if queue.currentLength == 0 {
		fmt.Println("Lista jest pusta. Nie ma elementów do usunięcia.")
		return
	}

	tempNode := queue.head
	for tempNode != nil && tempNode.value != value {
		tempNode = tempNode.nextNode
	}

	if tempNode == nil {
		fmt.Printf("Nie znaleziono węzła o wartości: %d\n", value)
		return
	}

	if tempNode == queue.head {
		queue.head = tempNode.nextNode
		if tempNode.nextNode != nil {
			tempNode.nextNode.previousNode = nil
		}
	} else if tempNode == queue.tail {
		queue.tail = tempNode.previousNode
		if tempNode.previousNode != nil {
			tempNode.previousNode.nextNode = nil
		}
	} else {
		tempNode.nextNode.previousNode = tempNode.previousNode
		tempNode.previousNode.nextNode = tempNode.nextNode
	}
	queue.currentLength--
	fmt.Printf("Usunięto węzeł: %d\n", value)
}

func printList(queue *queue) {
	if queue == nil {
		fmt.Println("\nNie można wyświetlić listy. Lista nie istnieje.")
	} else if queue.currentLength == 0 {
		fmt.Println("\nLista jest pusta.")
	} else {
		fmt.Print("\nWartości listy: ")
		var tempNode *fifoNode = queue.head
		for tempNode.nextNode != nil {
			fmt.Print(tempNode.value, " ")
			tempNode = tempNode.nextNode
		}
		fmt.Println(tempNode.value)
	}
}

func makeNull(queue *queue) {
	if queue == nil {
		fmt.Println("Nie można wyczyścić kolejki. Lista nie istnieje.")
		return
	}
	queue.head = nil
	queue.tail = nil
	queue.currentLength = 0
}

func empty(queue *queue) bool {
	return queue.currentLength == 0
}

func front(queue *queue) int {
	return queue.head.value
}

func enqueue(queue *queue, value int) {
	if queue == nil {
		fmt.Println("\nNie można dodać elementu. Lista nie istnieje.")
	} else if empty(queue) {
		newNode := &fifoNode{value: value, previousNode: nil, nextNode: nil}
		queue.head = newNode
		queue.tail = newNode
		queue.currentLength++
		fmt.Print("\nDodano element: ", value, "\n")
	} else {
		newNode := &fifoNode{value: value, previousNode: queue.tail, nextNode: nil}
		queue.tail.nextNode = newNode
		queue.tail = newNode
		queue.currentLength++
		fmt.Print("\nDodano element: ", value, "\n")
	}
}

func dequeue(queue *queue) int {
	if queue.currentLength == 1 {
		temp := queue.head.value
		queue.head = nil
		queue.tail = nil
		queue.currentLength--
		return temp
	} else {
		temp := queue.head.value
		queue.head = queue.head.nextNode
		queue.currentLength--
		return temp
	}
}

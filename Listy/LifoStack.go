package listy

import (
	"fmt"
)

type lifoNode struct {
	value                  int
	nextNode, previousNode *lifoNode
}

type stack struct {
	top           *lifoNode
	currentLength int
}

func StartLifoStack() {
	var decision int = 999
	var myQueue *queue

	for decision != 0 {
		fmt.Println("\nWitam w algorytmie stosu!\n1. Stwórz stos.\n2. Dodaj element.\n3. Zdejmij element.\n4. Wypisz elementy.\n5. Usuń stos.\n0. Wyjście")
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

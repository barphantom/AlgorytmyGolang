package Drzewa

// import (
// 	"fmt"
// )

// type node struct {
// 	key                 int
// 	left, right, parent *node
// }

// func BinaryTree() {
// 	// var intArr [3]int16 = [3]int16{1, 2, 3}
// 	// fmt.Println(intArr)
// 	// var myMap map[string]uint8 = make(map[string]uint8)
// 	// myMap["Adam"] = 30
// 	// myMap["Bartek"] = 22
// 	// fmt.Println(myMap)
// 	var root *node = nil
// 	var decision int = 999

// 	for decision != 0 {
// 		fmt.Println("\nWitam w algorytmie drzewa BST!\n1. Wyszykaj element.\n2. Dodaj element.\n3. Usuń element.\n4. Wypisz drzewo.\n0. Wyjście")
// 		fmt.Scanln(&decision)
// 		switch decision {
// 		case 1:
// 			addNode(root, decision)
// 		case 2:
// 			var value int
// 			fmt.Println("Wpisz nowy element listy.")
// 			fmt.Scanln(&value)
// 			enqueue(queue, value)
// 		case 3:
// 			var value int
// 			fmt.Scanln(&value)
// 			deleteElement(queue, value)
// 		case 4:
// 			printList(queue)
// 		case 5:
// 			queue = nil
// 			fmt.Println("\nUsunięto listę.")
// 		case 6:
// 			if queue == nil {
// 				fmt.Println("Nie można wyświetlić pierwszego elementu. Lista nie istnieje.")
// 			} else {
// 				value := front(queue)
// 				fmt.Println("Pierwszy element: ", value)
// 			}
// 		case 7:
// 			makeNull(queue)
// 			fmt.Println("Wyczyszczono listę.")
// 		case 8:
// 			if queue == nil {
// 				fmt.Println("Nie można zdjąć pierwszego elementu. Lista nie istnieje.")
// 			} else if empty(queue) {
// 				fmt.Println("Nie można zdjąć pierwszego elementu. Lista jest pusta.")
// 			} else {
// 				value := dequeue(queue)
// 				fmt.Println("Zdjęto element: ", value)
// 			}
// 		case 0:
// 			fmt.Println("\nWyjście z programu.")
// 		default:
// 			fmt.Println("Nie rozpoznano wyboru.")
// 		}
// 	}
// }

// func addNode(root **node, value int) {
// 	var now *node = *root
// 	var newNode *node = &node{key: value, left: nil, right: nil, parent: nil}

// 	if *root == nil {
// 		*root = newNode
// 		return
// 	}

// 	for {
// 		if value < now.key {
// 			if now.left != nil {
// 				now = now.left
// 			} else {
// 				newNode.parent = now
// 				now.left = newNode
// 				break
// 			}
// 		} else {
// 			if now.right != nil {
// 				now = now.right
// 			} else {
// 				newNode.parent = now
// 				now.right = newNode
// 				break
// 			}
// 		}
// 	}

// }

// func searchNode(root **node, value int) bool {

// }

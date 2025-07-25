package main

import (
	"bufio"
	"fmt"
	commands "main/Commands"
	"os"
	"strings"
)

func main() {
	for {
		scanner := bufio.NewScanner(os.Stdin)   // создает сканер с буфером сразу
		scanner.Scan()                          // ожидает ввод слова
		array := strings.Fields(scanner.Text()) // делит строку на слова
		cmd := array[0]                         // первое слово команда

		if cmd == "add" { // можно было и switch использовать
			str := ""
			for i := 1; i < len(array); i++ {
				str += array[i]
				if i < len(array)-1 {
					str += " "
				}
			}
			commands.Add(str)
			fmt.Println()
		} else if cmd == "delete" {
			if len(array) == 1 {
				fmt.Println("delete 'ID'")
			} else {
				commands.Delete(array[1])
				fmt.Println()
			}
		} else if cmd == "list" {
			commands.List()
			fmt.Println()
		} else if cmd == "update" {
			str := ""
			for i := 2; i < len(array); i++ {
				str += array[i]
				if i < len(array)-1 {
					str += " "
				}
			}
			commands.Update(array[1], str)
			fmt.Println()
		} else if cmd == "mark-in-progress" {
			if len(array) == 1 {
				fmt.Println("mark-in-progress 'ID'")
			} else {
				commands.MarkInProgress(array[1])
				fmt.Println()
			}
		} else if cmd == "mark-done" {
			if len(array) == 1 {
				fmt.Println("mark-done 'ID'")
			} else {
				commands.MarkDone(array[1])
				fmt.Println()
			}
		} else if cmd == "exit" {
			fmt.Println("Thanks for using!")
			break
		} else {
			fmt.Println("Command not found")
		}
	}

}

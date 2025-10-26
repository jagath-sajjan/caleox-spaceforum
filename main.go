package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"caleox-spaceforum/commands"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Welcome to CaleoX-SpaceForum 🌌")
	fmt.Println("Type 'help' to see commands.")

	for {
		fmt.Print("> ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		args := strings.Split(input, " ")

		switch args[0] {
		case "help":
			fmt.Println("Commands: signup, login, threads, post, view [thread] [page], reply [thread], deletepost [thread] [post], deleteaccount, exit")
		case "signup":
			commands.Signup()
		case "login":
			commands.Login()
		case "threads":
			commands.ListThreads()
		case "post":
			commands.CreateThread()
		case "view":
			if len(args) < 2 {
				fmt.Println("Usage: view [thread_number] [page_number]")
				continue
			}
			page := 1
			if len(args) >= 3 {
				fmt.Sscanf(args[2], "%d", &page)
			}
			commands.ViewThread(args[1], page)
		case "reply":
			if len(args) < 2 {
				fmt.Println("Usage: reply [thread_number]")
				continue
			}
			commands.ReplyThread(args[1])
		case "deletepost":
			if len(args) < 3 {
				fmt.Println("Usage: deletepost [thread_number] [post_number]")
				continue
			}
			commands.DeletePost(args[1], args[2])
		case "deleteaccount":
			commands.DeleteAccount()
		case "exit":
			fmt.Println("Goodbye!")
			return
		default:
			fmt.Println("Unknown command! Type 'help' to see commands.")
		}
	}
}

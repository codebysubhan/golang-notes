package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"cms-cmdline-v2/contact"
	// "cms-cmdline-v2/validation"
)

func displayHelp (){
	fmt.Println("Available commands:")
	fmt.Println("  add <name> <phone_number> <email> - Adds a new contact")
	fmt.Println("  view <name> - View contact details")
	fmt.Println("  delete <name> - Delete a contact")
	fmt.Println("  list - List all the contacts")
	fmt.Println("  help - Displays this help message")
	fmt.Println("  quit - Exists the program")
}



func main(){
	fmt.Println("Welcome to the Contacts Management System!")
	displayHelp()

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("\nEnter a command: ")
		scanner.Scan()
		command := scanner.Text()

		args := strings.Fields(command) // splits on a space

		if len(args) == 0{
			fmt.Println("Please enter a valid command!")
			continue
		}

		switch args[0]{
		case "add":
			if len(args) < 4{
				fmt.Println("Usage: add <name> <phone_number> <email>")
				continue
			}

			name := args[1]
			phoneNumber := args[2]
			email := args[3]

			contact := contact.Contact {
				Name: name,
				PhoneNumber: phoneNumber,
				Email: email,
			}

			err := contact.AddContact(contact)

			if err != nil {
				fmt.Println("Err: ", err)
				continue
			}
			fmt.Println("Contact added successfully!")

		case "view":
			if len(args) < 2{
				fmt.Println("Usage: view <name>")
				continue
			}
			name := args[1]
			contact, err := contact.ViewContact(name)

			if err != nil{
				fmt.Println("Error: ", err)
			} else {
				fmt.Printf("Name: %s\nPhone: %s\nEmail: %s", contact.Name, contact.PhoneNumber, contact.Email)
			}

		case "delete":
			if len(args) < 1{
				fmt.Println("Usage: delete <name>")
				continue
			}
			name := args[1]
			err := contact.DeleteContact(name)

			if err != nil{
				fmt.Println("Err: ", err)
			}else {
				fmt.Println("Contact deleted successfully!")
			}

		case "list":
			contacts := contact.GetAllContacts()
			if len(contacts) == 0 {
				fmt.Println("No contacts found!")
			} else {
				fmt.Println("Contacts details:")
				for _, c := range contacts{
					fmt.Printf("Name: %s\nPhone: %s\nEmail: %s\n\n", c.Name, c.PhoneNumber, c.Email)
				} 
			}

		case "help":
			displayHelp()

		case "quit":
			fmt.Println("Exiting the program...")
			os.Exit(0)

		default:
			fmt.Println("Please enter a valid command!")
		}
	}
}
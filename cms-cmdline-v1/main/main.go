package main

import (
	"fmt"
	cm "cms-cmdline/contact"
)

func main() {
	// create an empty contact slice as a db
	db := [][3]string{}
	fmt.Println("Adding a contact...")
	cm.CreateContact(&db, "Contact1", "0333123456", "contact1@gmail.com")
	fmt.Println("Contact created... listing contacts from main...", db)
	fmt.Println("Listing contact from modules...")
	i,_ := cm.GetContact(db, "Contact1")
	fmt.Println("Contact fetched... printing...", db[i])
	// deleting contact
	fmt.Println("Deleting contact...")
	_ = cm.DeleteContact(&db, "Contact1")
	fmt.Println("After deletion...", db)

}
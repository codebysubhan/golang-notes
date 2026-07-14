package contact

import (
	"cms-cmdline-v2/validation"
	"errors"
)

type Contact struct {
	Name string
	PhoneNumber string
	Email string
}

var contactsDB map[string]Contact

func init(){
	contactsDB = make(map[string]Contact)
}

func (c Contact) AddContact(contact Contact) error {
	if contact.Name == "" {
		return errors.New("Contact name cannot be empty!")
	}

	if _, exists := contactsDB[contact.Name]; exists {
		return errors.New("Contact already exists!")
	}

	if !validation.IsValidPhoneNumber(contact.PhoneNumber){
		return errors.New("Phone Number is not valid!")
	}

	if !validation.IsValidEmail(contact.Email){
		return errors.New("Email is not valid!")
	}

	contactsDB[contact.Name] = contact
	return nil
}


func ViewContact(name string) (Contact, error){
	contact, exists := contactsDB[name]
	if !exists {
		return Contact{}, errors.New("Contact not found!")
	}

	return contact, nil
}

func DeleteContact(name string) error{
	_, exists := contactsDB[name]

	if !exists {
		return errors.New("Contact not found")
	}

	delete(contactsDB, name)
	return nil
}

func GetAllContacts() []Contact {
	contacts := make([]Contact, 0, len(contactsDB))

	for _, contact := range contactsDB{ // _ is getting the index from the range function
		contacts = append(contacts, contact)
	}

	return contacts
}

package main

import "fmt"

type Contact struct {
	Name  string
	Phone string
	Email string
}

const (
	text          = "contact"
	add           = "1) Add " + text
	search        = "2) Search " + text
	deleteContact = "3) Delete " + text
	showAll       = "4) Show all " + text
	closeApp      = "5) Exit"
)

func addContact(contacts map[string]Contact) (map[string]Contact, error) {
	contact := Contact{}
	fmt.Println("Add contact:")
	fmt.Print("Name: ")
	_, err := fmt.Scan(&contact.Name)
	if err != nil {
		return contacts, err
	}

	fmt.Print("\nPhone: ")
	_, err = fmt.Scan(&contact.Phone)
	if err != nil {
		return contacts, err
	}

	fmt.Print("\nEmail: ")
	_, err = fmt.Scan(&contact.Email)
	if err != nil {
		return contacts, err
	}

	_, ok := contacts[contact.Name]
	if ok {
		return contacts, fmt.Errorf("contact %s already exists", contact.Name)
	}

	contacts[contact.Name] = contact
	fmt.Println("Contact added successfully")

	return contacts, nil
}

func searchContact(contacts map[string]Contact) (Contact, error) {
	search := ""
	fmt.Println("Search for Name: ")
	_, err := fmt.Scan(&search)
	if err != nil {
		return Contact{}, err
	}
	contact, ok := contacts[search]
	if !ok {
		return Contact{}, fmt.Errorf("Contact not found")
	}
	return contact, nil
}

func deleteContactByKey(contacts map[string]Contact) (map[string]Contact, error) {
	deleteByKey := ""
	fmt.Println("Type the name of the contact you want to delete:")
	_, err := fmt.Scan(&deleteByKey)
	if err != nil {
		return contacts, err
	}

	_, ok := contacts[deleteByKey]
	if !ok {
		return contacts, fmt.Errorf("Contact not found")
	}

	delete(contacts, deleteByKey)

	return contacts, nil
}

func showAllContacts(contacts map[string]Contact) error {
	if len(contacts) > 0 {
		for _, v := range contacts {
			fmt.Printf("| Name: %s | Phone: %s | Email: %s", v.Name, v.Phone, v.Email)
		}
	} else {
		return fmt.Errorf("No contacts found")
	}
	return nil
}

func main() {
	contacts := make(map[string]Contact)

	runCli := true
	for runCli {
		fmt.Println("\nPhonebook - CLI")
		fmt.Println("Menu:")
		fmt.Printf("\n %s \n %s \n %s \n %s \n %s \n", add, search, deleteContact, showAll, closeApp)
		voiceSelected := 0
		fmt.Print("Selected: ")
		_, err := fmt.Scan(&voiceSelected)
		if err != nil {
			fmt.Println("Error:", err)
			break
		}

		switch voiceSelected {
		case 1:
			contacts, err = addContact(contacts)
			if err != nil {
				fmt.Println("Error: ", err)
			}
		case 2:
			contact, err := searchContact(contacts)
			if err != nil {
				fmt.Println("Error: ", err)
			} else {
				fmt.Printf("Name: %s\n Phone: %s \n Email: %s \n", contact.Name, contact.Phone, contact.Email)
			}
		case 3:
			contacts, err = deleteContactByKey(contacts)
			if err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Println("Contact deleted successfully")
			}
		case 4:
			err = showAllContacts(contacts)
			if err != nil {
				fmt.Println("Error: ", err)
			}
		case 5:
			runCli = false
			fmt.Println("Exit success")
		default:
			fmt.Println("Input not valid")
		}
	}
}

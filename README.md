# Phonebook CLI

A CLI application in Go to emulate a phonebook.

## Operations
- add 
- search
- delete 
- show all

## Setup
```bash
mkdir phonebook-cli
go mod init phonebook-cli
touch main.go
```

## Usage
```bash
go run main.go
```

## Example Output
```text
Phonebook - CLI
Menu:

 1) Add contact 
 2) Search contact 
 3) Delete contact 
 4) Show all contact 
 5) Exit 
Selected: 1
Add contact:
Name: Jhon

Phone: 3457684949

Email: jhon@dhdh.com
Contact added successfully
```
## What I Learned
- How to use structs to represent real-world data
- How to use maps to store and retrieve data by key
- CRUD operations on a map (add, search, delete, iterate)
- Error handling with multiple return values
- How to check if a key exists in a map with two-value assignment
- How to prevent duplicate entries using map key lookup
package main

import (
	"fmt"
	"strings"
)

type User struct {
	Name     string
	Age      int
	isNormal bool
}

// Decalre a slice for storage
var users []User

func main() {
	//Declare a variable to store user choice
	var choice int

	//while loop
	for {
		fmt.Println("===============================")
		fmt.Println("WELCOME")
		fmt.Println("Please choose what you want to do:")
		fmt.Println("1.Add new user")
		fmt.Println("2.Delete user")
		fmt.Println("3.Update user")
		fmt.Println("4.Show all user")
		fmt.Println("5.Exit")
		fmt.Println("===============================")

		//get input
		fmt.Print("Enter your choice: ")
		_, err := fmt.Scanf("%d", &choice)
		if err != nil {
			fmt.Println("Invalid choice! Please choose a number from menu above!!!")
			continue
		}

		switch choice {
		case 1:
			AddNewUser()
		case 2:
			DeleteUser()
		case 3:
			UpdateUser()
		case 4:
			ShowUsers()
		case 5:
			fmt.Println("Exit the program. GOODBYE!!!")
			return
		default:
			fmt.Println("Invalid choice! Please select a valid option from the menu.")
		}
	}

}

func ShowUsers() {
	if len(users) == 0 {
		fmt.Println("No users found!")
		return
	}

	fmt.Println("List of users:")
	for _, user := range users {
		fmt.Printf("Name: %s, Age: %d, Normal: %t\n", user.Name, user.Age, user.isNormal)
	}
}

func AddNewUser() {
	var name string
	var age int
	var isNormal string

	fmt.Println("Enter your name: ")
	_, err := fmt.Scanf("%s", &name)
	if err != nil {
		fmt.Println("Invalid name!")
		return
	}

	fmt.Println("Enter your age: ")
	_, err = fmt.Scanf("%d", &age)
	if err != nil || age <= 0 {
		fmt.Println("Invalid age!")
		return
	}

	fmt.Println("Is this person normal? (y/n): ")
	_, err = fmt.Scanf("%s", &isNormal)
	if err != nil {
		fmt.Println("Invalid isNormal!")
		return
	}

	isNormalInput := strings.ToLower(isNormal) == "y"

	newUser := User{
		Name:     name,
		Age:      age,
		isNormal: isNormalInput,
	}

	users = append(users, newUser)

	fmt.Println("New user was added successfully!!!")
}

func DeleteUser() {
	var name string

	fmt.Println("Enter name of the user you want to delete: ")
	_, err := fmt.Scanf("%s", &name)
	if err != nil {
		fmt.Println("Invalid name!")
		return
	}

	for i, user := range users {
		if user.Name == name {
			users = append(users[:i], users[i+1:]...)
			fmt.Println("User deleted successfully!!!")
			return // Exit after deletion
		}
	}

	// If user not found
	fmt.Println("User not found!!!")
}

func UpdateUser() {
	var name string
	// Ask for the user's name to update
	fmt.Println("Enter the name of the user you want to update: ")
	_, err := fmt.Scanf("%s", &name)
	if err != nil {
		fmt.Println("Error reading input!")
		return
	}

	// Find the user by name and update details
	for i, user := range users {
		if user.Name == name {
			// Update the user's details
			var newAge int
			var newIsNormal string
			// Update age
			fmt.Println("Enter the new age for the user: ")
			_, err = fmt.Scanf("%d", &newAge)
			if err != nil || newAge <= 0 {
				fmt.Println("Invalid age!")
				return
			}

			// Update normal status
			fmt.Println("Is this person normal? (y/n): ")
			_, err = fmt.Scanf("%s", &newIsNormal)
			if err != nil {
				fmt.Println("Invalid isNormal input!")
				return
			}

			// Convert the "y/n" input to boolean
			isNormalInput := strings.ToLower(newIsNormal) == "y"

			// Update user in the slice
			users[i].Age = newAge
			users[i].isNormal = isNormalInput

			fmt.Println("User updated successfully!")
			return
		}
	}

	// If user not found
	fmt.Println("User not found!")
}

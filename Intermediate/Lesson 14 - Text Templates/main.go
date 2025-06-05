package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"text/template"
)

func main() {

	// // Creating a new text template
	// // tmpl := template.New("example")
	// // Parsing the template

	// // Either
	// // tmpl, err := template.New("example").Parse("Welcome, {{.name}}! How are you doing?\n")
	// // if err != nil {
	// // 	panic(err)
	// // }

	// // OR --> will panic if error from parse function
	// tmpl := template.Must(template.New("example").Parse("Welcome, {{.name}}! How are you doing?\n"))

	// // Define data for the welcome message template
	// data := map[string]any{
	// 	"name": "John",
	// }

	// err := tmpl.Execute(os.Stdout, data)
	// if err != nil {
	// 	panic(err)
	// }

	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Enter your name: ")
	name, _ := reader.ReadString('\n') // takes delimeter that will read until dealimeter is found
	name = strings.TrimSpace(name)

	// Define a named template for different types of
	templates := map[string]string{
		"welcome":      "Welcome, {{.name}}! We're glad you joined.",
		"notification": "{{.name}}, you have a new notification: {{.notification}}",
		"error":        "Opps! An error occured: {{.errorMessage}}",
	}

	// Parse and store templates
	parsedTemplates := make(map[string]*template.Template)
	for name, tmpl := range templates {
		parsedTemplates[name] = template.Must(template.New(name).Parse(tmpl))
	}

	for {
		// Show the menu
		fmt.Println("\nMenu")
		fmt.Println("1. Join")
		fmt.Println("2. Get Notification")
		fmt.Println("3. Get Error")
		fmt.Println("4. Exit")
		fmt.Print("Choose an option: ")
		choice, _ := reader.ReadString('\n')
		choice = strings.TrimSpace(choice)

		var data map[string]any
		var tmpl *template.Template

		switch choice {
		case "1":
			tmpl = parsedTemplates["welcome"]
			data = map[string]any{
				"name": name,
			}
		case "2":
			fmt.Print("Enter your notification message: ")
			notification, _ := reader.ReadString('\n')
			notification = strings.TrimSpace(notification)
			tmpl = parsedTemplates["notification"]
			data = map[string]any{
				"name":         name,
				"notification": notification,
			}
		case "3":
			fmt.Print("Enter your error message: ")
			errorMessage, _ := reader.ReadString('\n')
			errorMessage = strings.TrimSpace(errorMessage)
			tmpl = parsedTemplates["error"]
			data = map[string]any{
				"errorMessage": errorMessage,
			}
		case "4":
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid choice. Please select a valid option.")
			continue
		}

		err := tmpl.Execute(os.Stdout, data)
		if err != nil {
			fmt.Println("Error executing template:", err)
		}

	}

}

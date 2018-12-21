package main

import "fmt"

func main() {

	initialized("Openfile")
	defer cleanup("CloseFile")

	fmt.Println("Doing some work......")
	fmt.Println("More work.....")

	saveToDatabase()
	sendEmail()
}

func initialized(v string) {
	fmt.Println("Init.....", v)
}

func cleanup(v string) {
	fmt.Println("Clean...", v)
}

func saveToDatabase() {

	initialized("Open DB")
	defer cleanup("CloseDB")

	fmt.Println("Writing to DB......")
	fmt.Println("DB write done")

}

func sendEmail() {

	initialized("Open Connection to SMTP server")
	defer cleanup("Close SMTP connection")

	fmt.Println("Sending email......")
	fmt.Println("Email sent")

}

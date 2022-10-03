package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/bootcamp-go/hackaton-go-bases/internal/file"
	"github.com/bootcamp-go/hackaton-go-bases/internal/service"
)

const filePath = "./tickets.csv"

func menuOptions(t []service.Ticket, b service.Bookings) {
	var menuOption int
	for menuOption == 6 {
		fmt.Println("\n\t Hackathon Go Bases💥\n")
		fmt.Println("Menu:")
		fmt.Println("1 - See all tickets")
		fmt.Println("2 - Create a new ticket")
		fmt.Println("3 - Get a ticket")
		fmt.Println("4 - Update a ticket")
		fmt.Println("5 - Delete a ticket")
		fmt.Println("6 - Exit")
		fmt.Print("Select a option: ")
		fmt.Scanf("%d", &menuOption)
		optionSelected(menuOption, b)
	}
}

func optionSelected(option int, b service.Bookings) {
	switch option {
	case 1:
		fmt.Printf("Option %d -> See all tickets", option)
		data, err := os.ReadFile(filePath)
		if err != nil {
			panic(err)
		}
		fileData := strings.Split(string(data), "\n")
		count := 1
		for _, line := range fileData {
			if len(line) > 0 {
				fline := strings.Split(line, ",")
				fmt.Printf("--------------\nTicket number %d\n--------------\n", count)
				fmt.Printf(
					"ID:%s\nName: %s\nEmail: %s\nDestination: %s\nHour: %s\nPrice:%s\n",
					fline[0], fline[1], fline[2], fline[3], fline[4], fline[5])
			}
			count++
		}
	case 2:
		fmt.Printf("Option %d -> Create a new ticket", option)

		newTicket := service.Ticket{
			Id:          1234,
			Names:       "Name",
			Email:       "email@email.com",
			Destination: "Rio cuarto",
			Date:        "12:34",
			Price:       10000,
		}

		_, err := b.Create(newTicket)
		if err != nil {
			panic(err)
		}
		fmt.Println("The ticket has been added!")
	case 3:
		fmt.Printf("Option %d -> Get a ticket", option)

		var ticketId int
		fmt.Println("Insert ticket ID: ")
		fmt.Scanf("%d", &ticketId)
		ticket, err := b.Read(ticketId)
		if err != nil {
			panic(err)
		}
		fmt.Printf("Ticket %d: \n %+v", ticket.Id, ticket)
	case 4:
		fmt.Printf("Option %d -> Update a ticket", option)

		var ticketId int
		fmt.Println("Insert ticket ID: ")
		fmt.Scanf("%d", &ticketId)
		ticket, err := b.Read(ticketId)
		if err != nil {
			panic(err)
		}
		ticket.Price = 1000
		_, err = b.Update(ticket.Id, ticket)
		if err != nil {
			panic(err)
		}

		fmt.Printf("Ticket %d updated\n %+v", ticket.Id, ticket)
	case 5:
		fmt.Printf("Option %d -> Delete a ticket", option)

		var ticketId int
		fmt.Println("Insert ticket ID: ")
		fmt.Scanf("%d", &ticketId)
		_, err := b.Delete(ticketId)
		if err != nil {
			panic(err)
		}
		fmt.Printf("Ticket %d deleted", ticketId)

	default:
		fmt.Printf("The option %d is invalid.", option)
	}
}

func main() {
	var tickets []service.Ticket
	// Funcion para obtener tickets del archivo csv
	service.NewBookings(tickets)
	file := file.File{path: filePath}
	data, err := file.Read()
	if err != nil {
		panic(err)
	}
	tickets = append(tickets, data...)

	bookings := service.NewBookings(tickets)

	menuOptions(tickets, bookings)

	err = file.Write(&tickets)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Changes have been saved!")
	fmt.Println("........")
}

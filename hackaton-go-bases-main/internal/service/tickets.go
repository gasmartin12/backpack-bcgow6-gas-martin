package service

import (
	"errors"
)

type Bookings interface {
	// Create create a new Ticket
	Create(t Ticket) (Ticket, error)
	// Read read a Ticket by id
	Read(id int) (Ticket, error)
	// Update update values of a Ticket
	Update(id int, t Ticket) (Ticket, error)
	// Delete delete a Ticket by id
	Delete(id int) (int, error)
}

type bookings struct {
	Tickets []Ticket
}

type Ticket struct {
	Id                              int
	Names, Email, Destination, Date string
	Price                           int
}

// NewBookings creates a new bookings service
func NewBookings(Tickets []Ticket) Bookings {
	return &bookings{Tickets: Tickets}
}

// Un método para crear un nuevo ticket añadir al registro.
func (b *bookings) Create(t Ticket) (Ticket, error) {
	for _, value := range b.Tickets {
		if value.Id == t.Id {
			return Ticket{}, errors.New("error: the tickect id is already in use")
		}
	}
	b.Tickets = append(b.Tickets, t)
	return t, nil
}

// Un método para traer un ticket a través de su campo id.
func (b *bookings) Read(id int) (Ticket, error) {
	for _, ticket := range b.Tickets {
		if ticket.Id == id {
			return ticket, nil
		}
	}
	return Ticket{}, errors.New("error: not found")
}

// Un método para actualizar los campos de un ticket.
func (b *bookings) Update(id int, t Ticket) (Ticket, error) {
	if t.Id != id {
		return Ticket{}, errors.New("error: not found")
	}

	for index, value := range b.Tickets {
		if value.Id == id {
			b.Tickets[index] = t
			return t, nil
		}
	}
	return Ticket{}, errors.New("error: not found")
}

// Un método para eliminar un campo por su id.
func (b *bookings) Delete(id int) (int, error) {
	for index, value := range b.Tickets {
		if value.Id == id {
			b.Tickets = append(b.Tickets[:index], b.Tickets[index+1:]...)
			return id, nil
		}
	}
	return 0, errors.New("error: not found")
}

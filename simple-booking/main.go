package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	totalTickets  = 100
	bookingAgents = 5
)

type Ticket struct {
	ID int
}

type BookingAgent struct {
	ID int
}

func main() {
	ticketChan := make(chan Ticket, totalTickets)
	bookingChan := make(chan Ticket, totalTickets)

	var wg sync.WaitGroup

	go generateTickets(ticketChan)

	for i := 1; i <= bookingAgents; i++ {
		wg.Add(1)
		go bookTicketsByAgent(BookingAgent{ID: i}, ticketChan, bookingChan, &wg)
	}

	go func() {
		wg.Wait()
		close(bookingChan)
	}()

	bookedTickets := 0
	for range bookingChan {
		bookedTickets++
	}

	fmt.Printf("Booking complete! %d tickets out of %d tickets booked.\n", bookedTickets, totalTickets)
}

func generateTickets(ticketChan chan<- Ticket) {
	for i := 1; i <= totalTickets; i++ {
		ticketChan <- Ticket{ID: i}
	}
	close(ticketChan)
}

func bookTicketsByAgent(agent BookingAgent, ticketChan <-chan Ticket, bookingChan chan<- Ticket, wg *sync.WaitGroup) {
	defer wg.Done()

	for ticket := range ticketChan {
		time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)

		if rand.Float32() < 0.2 {
			fmt.Printf("Agent %d failed to book ticket %d\n", agent.ID, ticket.ID)
			continue
		}

		fmt.Printf("Agent %d booked ticket %d\n", agent.ID, ticket.ID)
		bookingChan <- ticket
	}
}

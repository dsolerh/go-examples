package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/fatih/color"
)

// variables
var seatingCapacity = 10
var arrivalRate = 100 // ms
var cutDuration = 1000 * time.Millisecond
var timeOpen = 10 * time.Second

func main() {
	// seed our random number generator
	rand.Seed(time.Now().UnixNano())

	// print welcome message
	color.Yellow("The sleeping barber problem")
	color.Yellow("---------------------------")

	// create channel if needed
	clientChan := make(chan string, seatingCapacity)
	doneChan := make(chan bool)

	// create the barbershop
	shop := BarberShop{
		ShopCapacity:    seatingCapacity,
		HairCutDuration: cutDuration,
		NumberOfBarbers: 0,
		ClientsChan:     clientChan,
		BarbersDoneChan: doneChan,
		Open:            true,
	}

	color.Green("The shop is open for the day!")

	// add barbers
	shop.addBarber("Frank")
	shop.addBarber("Jane")
	shop.addBarber("Kely")
	shop.addBarber("Joe")
	shop.addBarber("Polly")

	// start the barbershop as a goroutine
	shopClosing := make(chan bool)
	closed := make(chan bool)

	go func() {
		<-time.After(timeOpen)
		shopClosing <- true

		shop.closeShop()

		closed <- true
	}()

	// add cliets
	i := 1
	go func() {
		for {
			// get a random number with an average of arrival rate
			randomMillisecond := rand.Int() % (2 * arrivalRate)

			select {
			case <-shopClosing:
				return
			case <-time.After(time.Millisecond * time.Duration(randomMillisecond)):
				shop.addClient(fmt.Sprintf("Client #%d", i))
				i++
			}
		}
	}()

	// wait for shop to close
	<-closed
}

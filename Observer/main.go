// Package main demonstrates the Observer design pattern in Go.
//
// The program models a simple inventory system where a customer
// is notified when an item becomes available.
//
// Reference: https://refactoring.guru/design-patterns/observer
package main

import "fmt"

// main creates an item, registers two customers as observers and
// triggers an availability update which causes all observers to be
// notified through their update() method.
func main() {
	shirtItem := newItem("Nike Shirt")

	observerFirst := &Customer{id: "abc@gmail.com"}
	observerSecond := &Customer{id: "xyz@gmail.com"}

	shirtItem.register(observerFirst)
	shirtItem.register(observerSecond)

	shirtItem.updateAvailability()
}

// Customer represents an observer that receives updates when
// a monitored item changes state.  In a real system this could
// be an email address, a push‑notification token, etc.
type Customer struct {
	id string // customer's unique identifier (e.g., e‑mail address)
}

// update implements the Observer interface’s update method.  It
// simply prints out a simulated email notification.
func (c *Customer) update(itemName string) {
	fmt.Printf("Sending email to customer %s for item %s\n", c.id, itemName)
}

// getID returns the customer's id.  It is used by the
// removal logic to determine whether two observers are the same.
func (c *Customer) getID() string {
	return c.id
}

// Observer defines the contract that all observers must fulfil.
// It contains methods for receiving updates and identifying
// themselves.
type Observer interface {
	update(string)
	getID() string
}

// Item is the subject – it maintains a list of observers and
// stores its own state (name, stock status, etc.).
type Item struct {
	observerList []Observer // slice of registered observers
	name         string     // human‑readable product name
	inStock      bool       // true if the product is available
}

// newItem allocates a fresh Item.  Only the name is supplied –
// additional fields such as observerList and inStock are zero–valued.
func newItem(name string) *Item {
	return &Item{
		name: name,
	}
}

// updateAvailability changes the subject’s internal state to show
// the item is now in stock and then notifies all observers.
func (i *Item) updateAvailability() {
	fmt.Printf("Item %s is now in stock\n", i.name)
	i.inStock = true
	i.notifyAll()
}

// register adds an observer to the item’s observer list.
func (i *Item) register(o Observer) {
	i.observerList = append(i.observerList, o)
}

// deregister removes a previously registered observer from the
// list.  It does so by delegating to removeFromslice which
// performs the slice manipulation.
func (i *Item) deregister(o Observer) {
	i.observerList = removeFromslice(i.observerList, o)
}

// notifyAll iterates over all observers and calls their update()
// method to inform them of the new state.
func (i *Item) notifyAll() {
	for _, observer := range i.observerList {
		observer.update(i.name)
	}
}

// removeFromslice removes the supplied observer from the slice
// by swapping it with the last … element and slicing it off.
// It returns the new slice with the observer omitted.
func removeFromslice(observerList []Observer, observerToRemove Observer) []Observer {
	observerListLength := len(observerList)
	for i, observer := range observerList {
		if observerToRemove.getID() == observer.getID() {
			observerList[observerListLength-1], observerList[i] = observerList[i], observerList[observerListLength-1]
			return observerList[:observerListLength-1]
		}
	}
	return observerList
}

// Subject is the interface that the concrete Subject type (Item)
// implements.  It exposes the public operations that
// observers can register, deregister and that allow
// notifications to be triggered.
type Subject interface {
	register(observer Observer)
	deregister(observer Observer)
	notifyAll()
}
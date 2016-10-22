package main

import (
	"errors"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Customer represents a customer document we
// store in our database.
type Customer struct {
	ID   int
	Name string
}

// db represents our internal database system.
type db struct {
	customers map[int]Customer
	lock      sync.Mutex
}

// SaveCustomer stores a customer document in the database.
func (db *db) SaveCustomer(c Customer) (int, error) {
	db.lock.Lock()
	defer db.lock.Unlock()

	// How many customers do we have. Customer Id's
	// are generated in order.
	maxID := len(db.customers)

	// If this customer id is out of the range
	// then we have an integrity issue.
	if c.ID > maxID {
		return 0, errors.New("Invalid customer id")
	}

	// If no id is provided this is a new customer.
	// Generate a new id.
	if c.ID == 0 {
		c.ID = maxID + 1
	}

	// Save the customer in the database.
	db.customers[c.ID] = c

	// Return the customer id.
	return c.ID, nil
}

// FindCustomer locates a customer by id in the database.
func (db *db) FindCustomer(id int) (Customer, error) {

	// Locate the customer in the database.
	c, found := db.customers[id]
	if !found {
		return Customer{}, fmt.Errorf("customer with ID %d does not exist", id)
	}

	return c, nil
}

// AllCustomers returns the full database of customers.
func (db *db) AllCustomers() []Customer {

	// Allocate enough elements for the customers.
	all := make([]Customer, len(db.customers))

	// Range over the map storing each customer
	// in their ordered position.
	for _, c := range db.customers {
		all[c.ID-1] = c
	}

	// Return the slice exlcusing index 0.
	return all
}

// DB is an instance of the database.
var DB = db{
	customers: map[int]Customer{},
}

// Initalize the database with some values.
func init() {
	rand.Seed(time.Now().UnixNano())

	DB.SaveCustomer(Customer{Name: "Mary Jane"})
	DB.SaveCustomer(Customer{Name: "Bob Smith"})
}

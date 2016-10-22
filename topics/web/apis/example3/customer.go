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
	ID   string
	Name string
}

// db represents our internal database system.
type db struct {
	customers map[string]Customer
	lock      sync.Mutex
}

// SaveCustomer stores a customer document in the database.
func (db *db) SaveCustomer(c Customer) error {
	if c.ID == "" {
		return errors.New("invalid customer, missing id")
	}

	db.lock.Lock()
	{
		db.customers[c.ID] = c
	}
	db.lock.Unlock()

	return nil
}

// FindCustomer locates a customer by id in the database.
func (db *db) FindCustomer(id string) (Customer, error) {
	if c, ok := db.customers[id]; ok {
		return c, nil
	}

	return Customer{}, fmt.Errorf("Could not find Customer with ID %s", id)
}

// AllCustomers returns the full database of customers.
func (db *db) AllCustomers() []Customer {
	all := make([]Customer, 0, len(db.customers))

	for _, c := range db.customers {
		all = append(all, c)
	}

	return all
}

// DB is an instance of the database.
var DB db

// Initalize the database with some values.
func init() {
	rand.Seed(time.Now().UnixNano())

	DB.SaveCustomer(Customer{ID: "1", Name: "Mary Jane"})
	DB.SaveCustomer(Customer{ID: "2", Name: "Bob Smith"})
}

// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

package customer

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
	Name string `form:"name"`
}

// db represents our internal database system.
var db = struct {
	customers map[int]Customer
	lock      sync.Mutex
}{
	customers: map[int]Customer{},
}

// Initalize the database with some values.
func init() {
	rand.Seed(time.Now().UnixNano())

	Save(Customer{Name: "Mary Jane"})
	Save(Customer{Name: "Bob Smith"})
}

// Save stores a customer document in the database.
func Save(c Customer) (int, error) {
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

// Update updates the customer in the database.
func Update(c Customer) error {
	db.lock.Lock()
	defer db.lock.Unlock()

	if _, ok := db.customers[c.ID]; !ok {
		return fmt.Errorf("customer with ID %d does not exist", c.ID)
	}

	db.customers[c.ID] = c
	return nil
}

// Delete removes the customer from the database.
func Delete(c Customer) error {
	db.lock.Lock()
	defer db.lock.Unlock()

	if _, ok := db.customers[c.ID]; !ok {
		return fmt.Errorf("customer with ID %d does not exist", c.ID)
	}
	delete(db.customers, c.ID)
	return nil
}

// Find locates a customer by id in the database.
func Find(id int) (Customer, error) {
	db.lock.Lock()
	defer db.lock.Unlock()

	// Locate the customer in the database.
	c, found := db.customers[id]
	if !found {
		return Customer{}, fmt.Errorf("customer with ID %d does not exist", id)
	}

	return c, nil
}

// All returns the full database of customers.
func All() []Customer {
	db.lock.Lock()
	defer db.lock.Unlock()

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

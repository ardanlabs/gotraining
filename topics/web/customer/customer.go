// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

package customer

import (
	"errors"
	"fmt"
	"sync"
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
	maxID     int
	lock      sync.Mutex
}{
	customers: map[int]Customer{},
}

// Initalize the database with some values.
func init() {
	Save(Customer{Name: "Mary Jane"})
	Save(Customer{Name: "Bob Smith"})
}

// Save stores a customer document in the database.
func Save(c Customer) (int, error) {
	db.lock.Lock()
	defer db.lock.Unlock()

	// If this customer id is out of the range
	// then we have an integrity issue.
	if c.ID > db.maxID {
		return 0, errors.New("Invalid customer id")
	}

	// If no id is provided this is a new customer.
	// Generate a new id.
	if c.ID == 0 {
		c.ID = db.maxID + 1
		db.maxID = c.ID
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

	all := []Customer{}

	// Range over the map storing each customer
	// in their ordered position.
	for i := 1; i <= db.maxID; i++ {
		if c, ok := db.customers[i]; ok {
			all = append(all, c)
		}
	}

	return all
}

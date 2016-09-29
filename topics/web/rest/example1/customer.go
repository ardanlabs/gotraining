package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"sync"
	"time"
)

type customers map[string]Customer

var Customers = customers{}
var lock = &sync.Mutex{}

func init() {
	rand.Seed(time.Now().UnixNano())

	Customers.Save(NewCustomer("Mary Jane"))
	Customers.Save(NewCustomer("Bob Smith"))
}

type Customer struct {
	ID   string
	Name string
}

func NewCustomer(name string) Customer {
	id := strconv.Itoa(rand.Int())
	return Customer{
		ID:   id,
		Name: name,
	}
}

func (db customers) Save(c Customer) {
	lock.Lock()
	defer lock.Unlock()
	db[c.ID] = c
}

func (db customers) Find(id string) (Customer, error) {
	if c, ok := db[id]; ok {
		return c, nil
	}
	return Customer{}, fmt.Errorf("Could not find Customer with ID %s", id)
}

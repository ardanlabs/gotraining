package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"sync"
	"time"
)

type customers map[string]*Customer

var Customers = customers{}
var lock = &sync.Mutex{}

func init() {
	rand.Seed(time.Now().UnixNano())

	Customers.Save(&Customer{ID: "1", Name: "Mary Jane"})
	Customers.Save(&Customer{ID: "2", Name: "Bob Smith"})
}

type Customer struct {
	ID   string
	Name string `form:"name"`
}

func (db customers) Save(c *Customer) {
	lock.Lock()
	defer lock.Unlock()
	if c.ID == "" {
		c.ID = strconv.Itoa(rand.Int())
	}
	db[c.ID] = c
}

func (db customers) Find(id string) (*Customer, error) {
	if c, ok := db[id]; ok {
		return c, nil
	}
	return nil, fmt.Errorf("Could not find Customer with ID %s", id)
}

func (db customers) All() []*Customer {
	all := []*Customer{}
	for _, v := range db {
		all = append(all, v)
	}
	return all
}

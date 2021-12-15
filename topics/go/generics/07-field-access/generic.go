package main

import (
	"database/sql"
	"log"
)

// =============================================================================

// The User and Customer types each have a specialized insert function. These
// functions use value semantic mutation by accepting a value of the specific
// type and returning a new value of the type modified with the id for the newly
// inserted data. The query and arguments required are bound inside the function.

type User struct {
	ID    int64
	Name  string
	Email string
}

func InsertUser(db *sql.DB, u User) (User, error) {
	const query = "insert into users (name, email) values ($1, $2)"
	result, err := ExecuteQuery(query, u.Name, u.Email)
	if err != nil {
		return User{}, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return User{}, err
	}

	u.ID = id
	return u, nil
}

type Customer struct {
	ID    int64
	Name  string
	Email string
}

func InsertCustomer(db *sql.DB, c Customer) (Customer, error) {
	const query = "insert into customers (name, email) values ($1, $2)"
	result, err := ExecuteQuery(query, c.Name, c.Email)
	if err != nil {
		return Customer{}, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return Customer{}, err
	}

	c.ID = id
	return c, nil
}

// =============================================================================

// This version of a generic insert function is cleaner. It accepts the query
// and args as parameters and still implements the value semantic mutation like
// the concrete version. There is no setter needed since the entities interface
// provides the compiler with the information that values of type T will contain
// an ID field. This is because both these types share this common field.
//
// NOTE: I am not sure I want anyone writing code like this.
//       This is an experiment.

type entities interface {
	User | Customer
}

func Insert[T entities](db *sql.DB, entity T, query string, args ...interface{}) (T, error) {
	var zero T

	result, err := ExecuteQuery(query, args...)
	if err != nil {
		return zero, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return zero, err
	}

	// The entities interface provides this support. This has not be implemented
	// by the tooling though it is supported by the draft.
	entity.ID = id
	return entity, nil
}

func InsertUser2(db *sql.DB, u User) (User, error) {
	const query = "insert into users (name, email) values ($1, $2)"	
	u, err := Insert(db, u, query, u.Name, u.Email)
	if err != nil {
		return User{}, err
	}
	return u, nil
}

// =============================================================================

type Result struct{}

func (r Result) LastInsertId() (int64, error) { return 1, nil }
func (r Result) RowsAffected() (int64, error) { return 1, nil }

func ExecuteQuery(query string, args ...interface{}) (sql.Result, error) {
	return Result{}, nil
}

// =============================================================================

func main() {
	var db *sql.DB

	var u User
	query := "insert into users (name, email) values ($1, $2)"
	u, err := Insert(db, u, query, u.Name, u.Email)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(u)

	var c Customer
	query = "insert into customers (name, email) values ($1, $2)"
	c, err = Insert(db, c, query, u.Name, u.Email)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(c)
}

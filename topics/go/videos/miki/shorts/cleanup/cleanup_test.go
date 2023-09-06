// Use t.Cleanup to close resources after tests.

package main

import "testing"

func createDB(t testing.TB) *DB {
	db, err := ConnectDB()
	if err != nil {
		t.Fatalf("db.New() -> %s", err)
	}
	t.Log("connected to database")

	t.Cleanup(func() {
		db.Close()
		t.Log("closed database")
	})

	return db
}

func TestDB(t *testing.T) {
	db := createDB(t)
	t.Logf("db = %v", db)
	// TODO: work with db
}

// ---

type DB struct{}

func ConnectDB() (*DB, error) {
	return &DB{}, nil

}
func (db *DB) Close() error {
	return nil
}

var db *DB

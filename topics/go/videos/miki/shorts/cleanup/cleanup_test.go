// Use t.Cleanup to close resources after tests.

package main

import "testing"

func createDB(t testing.TB) *DB {
	db, err := db.New()
	if err != nil {
		t.Fatalf("db.New() -> %s", err)
	}

	t.Cleanup(func() {
		db.Close()
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

func (db *DB) New() (*DB, error) {
	return db, nil

}
func (db *DB) Close() error {
	return nil
}

var db *DB

// High-level modules should not depend on low-level modules. Both should depend on abstractions.
// Dependency injection is commonly used to follow this principle.

package main

import (
	"fmt"
)

type DataStore interface {
	Save(data string) error
}

type DatabaseStore struct{}

func (db *DatabaseStore) Save(data string) error {
	fmt.Println("Saving to database:", data)
	return nil
}

type FileStore struct{}

func (fs *FileStore) Save(data string) error {
	fmt.Println("Saving to file:", data)
	return nil
}

type DataService struct {
	store DataStore
}

func (ds *DataService) SaveData(data string) {
	ds.store.Save(data)
}

func main() {
	// Dependency Injection - we can switch between stores easily
	dbStore := &DatabaseStore{}
	fileStore := &FileStore{}

	dataService := &DataService{store: dbStore}
	dataService.SaveData("Important data")

	dataService.store = fileStore
	dataService.SaveData("Important data")
}

// DataService depends on the DataStore interface, not a specific implementation,
// allowing easy swapping between DatabaseStore and FileStore.

package main

import "fmt"

func LogOutput(message string) {
	fmt.Println(message)
}

type SimpleDataStore struct {
	userData map[string]string
}

func (sds SimpleDataStore) UserNameForID(userId string) (string, bool) {
	userName, ok := sds.userData[userId]
	return userName, ok
}

// create an instance of a SimpleDataStore

func NewSimpleDataStore() SimpleDataStore {
	dataStore := SimpleDataStore{userData: map[string]string{
		"1": "zakir",
		"2": "asif",
		"3": "Pat",
		"4": "rakib",
		"5": "rakin",
	}}
	return dataStore
}

type DataStore interface {
	UserNameForID(userID string) (string, bool)
}
type Logger interface {
	Log(message string)
}

type LoggerAdapter func(message string)

func (lg LoggerAdapter) Log(message string) {
	lg(message)
}

func main() {
	store := NewSimpleDataStore()

	name, ok := SimpleDataStore.UserNameForID(store, "2222")
	if !ok {
		fmt.Println("User does not exist")
	} else {
		fmt.Printf("User Name is %s\n", name)
	}
}

package main

import (
	"fmt"
	"math/rand"
)

var (
	lastUserID int64
)

/*
	Request la DB dupa uID
	Formeaza un user pregenerat
*/
func getUser(uID int64) (User, error) {
	return User{
		UID:              uID,
		InstitutionEmail: "mihai.indreias@unibuc.ro",
		YearOfStudy:      1,
		Facultate:        "Politehnica",
		University:       "Politehnica Bucuresti",
		Serie:            "123",
		Major:            "Inginerie",
		PersonalEmail:    "mihai.indreias@gmail.com",
		Username:         "g0g05arui",
	}, nil
}

func generate16DigitID() int64 {
	return 1e16 + rand.Int63n(1e15)
}

func findUserByUsername(userName string) bool {
	return false
}

func findUserByID(userID int64) bool {
	return false
}

/*
	Hide password field at marshal
*/

func (passwordType) MarshalJSON() ([]byte, error) {
	return []byte(`""`), nil
}

func addUser(newUser User) HTTPResponse {
	if findUserByUsername(newUser.Username) {
		return HTTPResponse{Response: "User already exists", Code: 409}
	}
	/*
		Add User to DB
	*/
	userID := generate16DigitID()
	for findUserByID(userID) {
		userID = generate16DigitID()
	}
	newUser.UID = userID
	fmt.Println(newUser)
	return HTTPResponse{Response: "User added", Code: 200}
}

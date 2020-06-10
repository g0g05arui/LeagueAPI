package main

var (
	lastUserID int64
)

/*
	Request la DB dupa uID
	Formeaza un user pregenerat
*/
func getUser(uID int64) User {
	return User{UID: uID, Mail: "mihai.indreias@gmail.com", Username: "g0g05arui"}
}

func findUserByUsername(newUser string) bool {
	return true
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
	newUser.UID = lastUserID
	lastUserID++
	return HTTPResponse{Response: "User added", Code: 200}
}

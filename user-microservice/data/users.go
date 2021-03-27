package data

type User struct {
	Username string `json:"username"`
	DOB      string `json:"dob"`
	Age      int    `json:"age"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
}

func GetUsers() map[string]*User {
	users := make(map[string]*User)
	users["user1"] = &User{
		Username: "user1",
		DOB:      "1990-01-1",
		Age:      31,
		Email:    "user1@sample.org",
		Phone:    "12345677890",
	}
	users["user2"] = &User{
		Username: "user2",
		DOB:      "1991-02-2",
		Age:      30,
		Email:    "user2@sample.org",
		Phone:    "0987654321",
	}
	return users
}

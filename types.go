package main

type User struct {
	FirstName string `json:"firstname" bson:"firstname"`
	LastName  string `json:"lastname" bson:"lastname"`
	Email     string `json:"email" bson:"email"`
	Password  string `json:"password" bson:"password"`
}

type Salt struct {
	Salt string `json:"salt"`
}

type TCard struct {
	Id     string `json:"id"`
	Name   string `json:"name"`
	Onsale bool   `json:"onSale"`
	Age    int    `json:"age"`
}

type SCard struct {
	Card TCard `json:"card"`
}

type TUser struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
type SignInInput struct {
	Input TUser `json:"input"`
}

type CUser struct {
	Slug    string `json:"slug"`
	Balance int    `json:"coinBalance"`
}

type TError struct {
	Message string `json:"message"`
}

type SignInResp struct {
	CurrentUser CUser    `json:"currentUser"`
	Err         []TError `json:"errors"`
}

type SignIn struct {
	SignIn SignInResp `json:"signIn"`
}

//coger variables os.getenv()

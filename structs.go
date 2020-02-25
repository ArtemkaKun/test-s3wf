package main

type User struct {
	Uuiduser     string `json:"id"`
	Avatar_image string `json:"avatar_img"`
	Avatar_type  string `json:"avatar_type"`
	Name         string `json:"name"`
	Surname      string `json:"surname"`
	Datastart    string `json:"datastart"`
	Login        string `json:"login"`
	Pass         string `json:"pass"`
}

type LogErr struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

type LogSucc struct {
	Status   string `json:"status"`
	Message  string `json:"message"`
	Token    string `json:"tokenjwt"`
	Exp      string `json:"expires"`
	TokenMsg string `json:"tokenmsg"`
	Login    User
}

package main

import "testing"

func TestLoginProcess(t *testing.T) {
	testpairs := []LoginTestpair{
		{Test_login: "0000", Test_pass: "0000", Err_status: "error", Login_status: ""},
		{Test_login: "0000", Test_pass: "12345", Err_status: "", Login_status: "success"},
		{Test_login: "1111", Test_pass: "0000", Err_status: "error", Login_status: ""},
		{Test_login: "agams1987", Test_pass: "12345", Err_status: "", Login_status: "success"},
		{Test_login: "agam1987", Test_pass: "12345", Err_status: "error", Login_status: ""},
		{Test_login: "hack1337", Test_pass: "loght_Pa$$!", Err_status: "error", Login_status: ""},
		{Test_login: "hack1337", Test_pass: "light_Pa$$!", Err_status: "", Login_status: "success"},
		{Test_login: "admin", Test_pass: "", Err_status: "error", Login_status: ""},
		{Test_login: "", Test_pass: "", Err_status: "error", Login_status: ""},
		{Test_login: "-1", Test_pass: "-1", Err_status: "error", Login_status: ""},
	}

	for _, test_pair := range testpairs {
		exp_err_status := test_pair.Err_status
		exp_login_status := test_pair.Login_status
		err, succ := LoginProcess(test_pair.Test_login, test_pair.Test_pass)
		if exp_err_status != err.Status || exp_login_status != succ.Status {
			t.Error("Expect: ", exp_err_status, exp_login_status,
				"Got: ", err.Status, succ.Status)
		}
	}
}

func TestGetUserInfo(t *testing.T) {
	testpairs := []UserInfoTestpair{
		{Test_uuid: "1", Test_user: User{
			Uuiduser:     "1",
			Avatar_image: "image.png",
			Avatar_type:  "image",
			Name:         "Junior",
			Surname:      "S3WF",
			Datastart:    "today",
			Login:        "0000",
			Pass:         "12345"},
			Err_status: ""},
		{Test_uuid: "2", Test_user: User{
			Uuiduser:     "2",
			Avatar_image: "image.png",
			Avatar_type:  "",
			Name:         "Middle",
			Surname:      "Mid",
			Datastart:    "",
			Login:        "agams1987",
			Pass:         "12345"},
			Err_status: ""},
		{Test_uuid: "3", Test_user: User{
			Uuiduser:     "3",
			Avatar_image: "image.png",
			Avatar_type:  "image",
			Name:         "Senior",
			Surname:      "Pomidor",
			Datastart:    "today",
			Login:        "hack1337",
			Pass:         "light_Pa$$!"},
			Err_status: ""},
	}

	for _, pair := range testpairs {
		err, user := GetUserInfo(pair.Test_uuid)

		if err.Status != pair.Err_status || user != pair.Test_user {
			t.Error("Expect: ", pair.Err_status, pair.Test_user,
				"Got: ", err.Status, user)
		}
	}
}

//func TestAddNewUser(t *testing.T) {
//	new_users := []User{
//		{Uuiduser: "4",
//			Avatar_image: "image.png",
//			Avatar_type:  "image",
//			Name:         "Artem",
//			Surname:      "Yurchenko",
//			Datastart:    "today",
//			Login:        "yammil",
//			Pass:         "passWord"},
//		{Uuiduser: "5",
//			Avatar_image: "image.png",
//			Avatar_type:  "image",
//			Name:         "Artem",
//			Surname:      "NotYurchenko",
//			Datastart:    "today",
//			Login:        "notyammil",
//			Pass:         "notpassWord"},
//	}
//	testpairs := []UserInfoTestpair{
//		{Test_uuid: "4", Test_user: User{
//			Uuiduser:     "4",
//			Avatar_image: "image.png",
//			Avatar_type:  "image",
//			Name:         "Artem",
//			Surname:      "Yurchenko",
//			Datastart:    "today",
//			Login:        "yammil",
//			Pass:         "passWord"},
//			Err_status: ""},
//		{Test_uuid: "5", Test_user: User{
//			Uuiduser:     "5",
//			Avatar_image: "image.png",
//			Avatar_type:  "image",
//			Name:         "Artem",
//			Surname:      "NotYurchenko",
//			Datastart:    "today",
//			Login:        "notyammil",
//			Pass:         "notpassWord"},
//			Err_status: ""},
//	}
//
//	for _, one_user := range new_users {
//		AddNewUser(one_user)
//	}
//
//	for _, pair := range testpairs {
//		err, user := GetUserInfo(pair.Test_uuid)
//
//		if err.Status != pair.Err_status || user != pair.Test_user {
//			t.Error("Expect: ", pair.Err_status, pair.Test_user,
//				"Got: ", err.Status, user)
//		}
//	}
//}

func TestUpdateUserInfo(t *testing.T) {
	testpairs := []UserInfoTestpair{
		{Test_uuid: "4", Test_user: User{
			Uuiduser:     "4",
			Avatar_image: "image.png",
			Avatar_type:  "image",
			Name:         "Artem",
			Surname:      "Yurchenko",
			Datastart:    "today",
			Login:        "yammil123456",
			Pass:         "passWord"},
			Err_status: ""},
		{Test_uuid: "6", Test_user: User{},
			Err_status: "error"},
	}

	for _, pair := range testpairs {
		err := UpdateUserInfo(pair.Test_uuid, pair.Test_user)
		if err.Status != pair.Err_status {
			t.Error("Expect: ", pair.Err_status,
				"Got: ", err.Status)
		}
	}

	for _, pair := range testpairs {
		err, user := GetUserInfo(pair.Test_uuid)

		if err.Status != pair.Err_status || user != pair.Test_user {
			t.Error("Expect: ", pair.Err_status, pair.Test_user,
				"Got: ", err.Status, user)
		}
	}
}

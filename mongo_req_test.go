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
		if exp_err_status != err.Status && exp_login_status != succ.Status {
			t.Error("Expect: ", exp_err_status, exp_login_status,
				"Got: ", err.Status, succ.Status)
		}
	}
}

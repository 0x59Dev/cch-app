package user

import (
	"errors"
	"regexp"
	"strings"
)

type User struct {
	Login        string `json:"login"`
	Mail         string `json:"mail"`
	Password     string `json:"password"`
	ConfPassword string `json:"confPassword"`
}

//AddUser uses verify Function and inserts user into database as an result of register
func (r *User) AddUser() error {
	if err := r.verifyData(); err != nil {
		return err
	} else {
		return nil
	}
}

func (r *User) verifyMail() bool {
	mailRegexp := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,10}$`)
	return mailRegexp.MatchString(r.Mail)
}

func (r *User) verifyPassword() (bool, err) {
	passCompare := strings.Compare(r.Pass, r.ConfPassword)
	if passCompare != -1 {
		err := errors.New("Password's dont match")
		return false, err
	}
	return true
}

func (r *User) verifyData() error {
	if !r.verifyMail() {
		return errors.New("MailError")
	} else if !r.verifyPassword() {
		return errors.New("Pass Strength Error")
	} else {
		return nil
	}
}

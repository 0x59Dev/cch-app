package auth

import (
	"errors"
	"fmt"
	"regexp"
)

type RegisterData struct {
	Login        string `json:"login"`
	Mail         string `json:"mail"`
	Password     string `json:"password"`
	ConfPassword string `json:"confPassword"`
}

func (r *RegisterData) AddUser() error {
	if err := r.verifyData(); err != nil {
		return err
	} else {
		return nil
	}
}

func (r *RegisterData) verifyMail() bool {
	mailRegexp := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,10}$`)
	return mailRegexp.MatchString(r.Mail)
}

func (r *RegisterData) verifyPassword() bool {
	return true
}

func (r *RegisterData) verifyData() error {
	if !r.verifyMail() {
		fmt.Println(r.verifyMail())
		return errors.New("MailError")
	} else if !r.verifyPassword() {
		return errors.New("Pass Strength Error")
	} else if r.Password != r.ConfPassword {
		return errors.New("Password Match Error")
	} else {
		return nil
	}
}

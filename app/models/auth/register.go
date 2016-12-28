package auth

import (
	"errors"
	"fmt"
	"regexp"
)

type RegisterData struct {
	login        string `json:"login"`
	mail         string `json:"mail"`
	password     string `json:"password"`
	confPassword string `json:"confPassword"`
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
	return mailRegexp.MatchString(r.mail)
}

func (r *RegisterData) verifyPassword() bool {
	return true
}

func (r *RegisterData) verifyData() error {
	if !r.verifyMail() {
		fmt.Println(r.verifyMail())
		return errors.New("MailError")
	} else if !r.verifyPassword() {
		return errors.New("PassError")
	} else if r.password != r.confPassword {
		return errors.New("MatchError")
	} else {
		return nil
	}
}

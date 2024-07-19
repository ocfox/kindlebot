package main

import (
	"os"
)

type Envs struct {
	SendMail      string
	Password      string
	Server        string
	RecipientMail string
}

func GetEnvs() Envs {
	return Envs{
		SendMail:      os.Getenv("SEND_MAIL"),
		Password:      os.Getenv("PASSWORD"),
		Server:        os.Getenv("SERVER"),
		RecipientMail: os.Getenv("RECIPIENT_MAIL"),
	}
}

func IsEnvsEmpty(envs Envs) bool {
	return envs.Server == "" ||
		envs.Password == "" ||
		envs.SendMail == "" ||
		envs.RecipientMail == ""
}

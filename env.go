package main

import (
	"os"
)

type Envs struct {
	SendMail string
	Password string
	Server   string
}

func GetEnvs() Envs {
	return Envs{
		SendMail: os.Getenv("SEND_MAIL"),
		Password: os.Getenv("PASSWORD"),
		Server:   os.Getenv("SERVER"),
	}
}

func IsEnvsEmpty(envs Envs) bool {
	return envs.Server == "" ||
		envs.Password == "" ||
		envs.SendMail == ""
}

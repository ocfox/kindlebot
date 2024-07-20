package main

import (
	"os"

	"encoding/json"
)

type UserMap map[int64]string

func CreateUserMap(file string) {
	// Check if the file exists
	if _, err := os.Stat(file); os.IsNotExist(err) {
		// File does not exist, create it
		f, err := os.Create(file)
		if err != nil {
			panic("Failed to create file: " + err.Error())
		}
		defer f.Close()

		// Initialize an empty map
		userMap := make(UserMap)

		// Marshal the map into JSON format
		data, err := json.Marshal(userMap)
		if err != nil {
			panic("Failed to marshal user map: " + err.Error())
		}

		// Write the JSON data to the file
		_, err = f.Write(data)
		if err != nil {
			panic("Failed to write to file: " + err.Error())
		}
	}
}

func ReadUserMap(file string) UserMap {
	// Open the file for reading
	f, err := os.Open(file)
	if err != nil {
		panic("Failed to open file: " + err.Error())
	}
	defer f.Close()

	// Read the JSON data from the file
	data := make([]byte, 1024)
	n, err := f.Read(data)
	if err != nil {
		panic("Failed to read from file: " + err.Error())
	}

	// Unmarshal the JSON data into a map
	userMap := make(UserMap)
	err = json.Unmarshal(data[:n], &userMap)
	if err != nil {
		panic("Failed to unmarshal user map: " + err.Error())
	}

	return userMap
}

func AddUser(id int64, mail string) {
	// Add the user to the map
	userMap[id] = mail
	UpdateUserFile("./users.json", userMap)
}

func UpdateUserFile(file string, userMap UserMap) {
	// Marshal the map into JSON format
	data, err := json.Marshal(userMap)
	if err != nil {
		panic("Failed to marshal user map: " + err.Error())
	}

	// Write the JSON data to the file
	err = os.WriteFile(file, data, 0644)
	if err != nil {
		panic("Failed to write to file: " + err.Error())
	}
}

func GetUserMail(id int64) string {
	if mail, ok := userMap[id]; ok {
		return mail
	}
	return ""
}

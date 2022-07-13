package main

import (
	"encoding/json"
	"fmt"
)

type user struct {
	FullName string     `json:"full_name"`
	UserName string     `json:"username"`
	Email    string     `json:"email"`
	Password string     `json:"-"`
	Status   string     `json:"status"`
	UserMeta []userMeta `json:"user_meta ,omitempty"`
}

type userMeta struct {
	MetaKey   string `json:"meta_key"`
	MetaValue string `json:"meta_value"`
}

func main() {
	newUser := []user{
		{
			FullName: "John Doe",
			UserName: "john_doe",
			Email:    "john@doe.com",
			Password: "password",
			Status:   "active",
			UserMeta: []userMeta{
				{MetaKey: "country", MetaValue: "USA"},
				{MetaKey: "city", MetaValue: "New York"},
				{MetaKey: "address", MetaValue: "123 Main St"},
				{MetaKey: "company", MetaValue: "Acme Inc"},
				{MetaKey: "phone", MetaValue: "555-555-5555"},
			},
		},
		{
			FullName: "Ratul Hasan",
			UserName: "ratul_hasan",
			Email:    "ratul@hasan.com",
			Password: "password",
			Status:   "active",
			UserMeta: []userMeta{
				{MetaKey: "country", MetaValue: "Bangladesh"},
				{MetaKey: "city", MetaValue: "Dhaka"},
				{MetaKey: "address", MetaValue: "27/13 Pallabi, Mirpur, Dhaka"},
				{MetaKey: "company", MetaValue: "weDevs"},
				{MetaKey: "phone", MetaValue: "555-555-5555"},
				{MetaKey: "profile_image", MetaValue: "https://www.gravatar.com/avatar/e64c7d89f26bd1972efa854d13d7dd61?s=80&d=mm&r=g"},
				{MetaKey: "facebook", MetaValue: "https://www.facebook.com/ratuljh"},
				{MetaKey: "twitter", MetaValue: "https://twitter.com/ratuljh"},
				{MetaKey: "google", MetaValue: "https://plus.google.com/+RatulHasan"},
				{MetaKey: "linkedin", MetaValue: "https://www.linkedin.com/in/ratulhasan"},
				{MetaKey: "github", MetaValue: "https://github.com/ratulhasan"},
				{MetaKey: "website", MetaValue: "https://ratulhasan.com"},
				{MetaKey: "about", MetaValue: "I am a web developer"},
				{MetaKey: "website", MetaValue: "https://www.ratulhasan.com"},
				{MetaKey: "github", MetaValue: "https://github.com/ratulhasan"},
			},
		},
	}

	// Convert the user struct to a JSON string
	encodeToJSON, err := EncodeToJSON(newUser)
	if err != nil {
		return
	}

	// Print the JSON string
	fmt.Println(encodeToJSON)
}

func EncodeToJSON(user []user) (string, error) {
	jsonData, err := json.MarshalIndent(user, "", "\t")
	if err != nil {
		return "", err
	}
	return string(jsonData), nil
}

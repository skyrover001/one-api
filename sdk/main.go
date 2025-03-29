package main

import (
	"fmt"
	"log"
	"one-api/sdk/api"
)

func main() {
	// for demo
	config := api.Config{
		Host: "http://localhost",
		Port: 8300,
		Key:  "12345678901234567890",
	}
	client := api.OneClient{
		Config: &config,
	}

	// user api
	// add user
	user := api.User{
		Username:    "test1",
		DisplayName: "test1",
		Password:    "test@123_%6",
	}
	err := user.AddUser(&client)
	if err != nil {
		log.Fatal(err)
	}
	// list users
	users := api.Users{}
	err = users.ListUsers(&client)
	for _, u := range users.Users {
		if u.Username == "test1" {
			user = *u
		}
	}
	//update user
	user.Quota = 500000000
	err = user.UpdateUser(&client)
	err = users.ListUsers(&client)
	for _, u := range users.Users {
		if u.Username == "test1" {
			user = *u
		}
	}
	// delete user
	err = user.DeleteUser(&client)
	err = users.ListUsers(&client)
	for _, u := range users.Users {
		fmt.Println(u)
	}

	// channel api
	channel := api.Channel{
		Name:    "ch1",
		BaseUrl: "",
		ChannelConfig: api.ChannelConfig{
			Region: "",
			Sk:     "",
			Ak:     "",
		},
		Group:        "default",
		Models:       "deepseek-r1",
		ModelMapping: "",
		Other:        "",
		SystemPrompt: "",
		Type:         1,
		Key:          "12345678901234567890",
	}
	err = channel.AddChannel(&client)
	if err != nil {
		log.Fatal(err)
	}

	// list channels
	channels := api.Channels{}
	err = channels.ListChannels(&client)
	if err != nil {
		log.Fatal(err)
	}
	updateChannel := api.Channel{}
	for _, c := range channels.Channels {
		if c.Name == "ch1" {
			updateChannel = *c
		}
	}
	// update channel
	updateChannel.Name = "ch1-updated"
	err = updateChannel.UpdateChannel(&client)
	err = channels.ListChannels(&client)
	for _, c := range channels.Channels {
		fmt.Println("c=", c)
	}

	//delete channel
	err = updateChannel.DeleteChannel(&client)
	if err != nil {
		log.Fatal(err)
	}
	err = channels.ListChannels(&client)
	for _, c := range channels.Channels {
		fmt.Println("c=", c)
	}
}

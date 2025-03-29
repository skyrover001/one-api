package api

import (
	"encoding/json"
	"fmt"
)

type Channel struct {
	ID                 int           `json:"id"`
	Type               int           `json:"type"`
	Key                string        `json:"key"`
	Status             int           `json:"status"`
	Name               string        `json:"name"`
	Weight             int           `json:"weight"`
	CreatedTime        int           `json:"created_time"`
	TestTime           int           `json:"test_time"`
	ResponseTime       int           `json:"response_time"`
	BaseUrl            string        `json:"base_url"`
	Other              string        `json:"other"`
	Balance            int           `json:"balance"`
	BalanceUpdatedTime int           `json:"balance_updated_time"`
	Models             string        `json:"models"`
	Group              string        `json:"group"`
	UsedQuota          int           `json:"used_quota"`
	ModelMapping       string        `json:"model_mapping"`
	Priority           int           `json:"priority"`
	Config             string        `json:"config"`
	SystemPrompt       string        `json:"system_prompt"`
	ChannelConfig      ChannelConfig `json:"channel_confi"`
}

type ChannelConfig struct {
	Region            string `json:"region"`
	Sk                string `json:"sk"`
	Ak                string `json:"ak"`
	UserId            string `json:"user_id"`
	VertexAiProjectId string `json:"vertex_ai_project_id"`
	VertexAiAdc       string `json:"vertex_ai_adc"`
}
type NewChannel struct {
	BaseUrl      string   `json:"base_url"`
	Config       string   `json:"config"`
	Group        string   `json:"group"`
	Groups       []string `json:"groups"`
	Key          string   `json:"key"`
	ModelMapping string   `json:"model_mapping"`
	Models       string   `json:"models"`
	Name         string   `json:"name"`
	Other        string   `json:"other"`
	SystemPrompt string   `json:"system_prompt"`
	Type         int      `json:"type"`
}

type Channels struct {
	Channels []*Channel
}

type ChannelRespData struct {
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
	Success bool        `json:"success"`
}

type ChannelImpl interface {
	AddChannel(channel *Channel) error
	GetChannel(id int) error
	UpdateChannel(channel *Channel) error
	DeleteChannel(id int) error
}

// list channel
func (channels *Channels) ListChannels(client *OneClient) error {
	client.Url = "/api/channel/?p=0"
	resp, err := client.get()
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	data := ChannelRespData{Data: []*Channel{}, Message: "", Success: false}
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return err
	}
	for _, v := range data.Data.([]interface{}) {
		channel := &Channel{}
		channelData, _ := json.Marshal(v)
		err = json.Unmarshal(channelData, channel)
		if err != nil {
			return err
		}
		channels.Channels = append(channels.Channels, channel)
	}
	return nil
}

// add channel
func (channel *Channel) AddChannel(client *OneClient) error {
	client.Url = "/api/channel/"
	channelConfigData, err := json.Marshal(channel.ChannelConfig)
	newChannel := NewChannel{
		BaseUrl:      channel.BaseUrl,
		Config:       string(channelConfigData),
		Group:        channel.Group,
		Groups:       []string{channel.Group},
		Key:          channel.Key,
		ModelMapping: channel.ModelMapping,
		Models:       channel.Models,
		Name:         channel.Name,
		Other:        channel.Other,
		SystemPrompt: channel.SystemPrompt,
		Type:         channel.Type,
	}
	data, err := json.Marshal(newChannel)
	if err != nil {
		return err
	}
	fmt.Println("eeeeeeeeeeeeeeeeeeeeeeeeerr=", err)
	return client.post(data)
}

// update channel
func (channel *Channel) UpdateChannel(client *OneClient) error {
	client.Url = "/api/channel/"
	data, err := json.Marshal(channel)
	if err != nil {
		return err
	}
	return client.put(data)
}

// delete channel
func (channel *Channel) DeleteChannel(client *OneClient) error {
	client.Url = "/api/channel/" + fmt.Sprintf("%d", channel.ID) + "/"
	fmt.Println("url=", client.Url)
	return client.delete(nil)
}

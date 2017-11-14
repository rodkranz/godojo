package main

import (
	"net/http"
	"fmt"
	"os"
	"encoding/json"
	"bytes"
)

var (
	BotName   = "Tomilio"
	BotAvatar = ":gopher:"
	
	GithubBaseUrl = "https://api.github.com/users/"
	SlackBaseUrl  = "https://hooks.slack.com/services/XXXXXXXXX/XXXXXXXXX/000000000000000000"
)

type GithubUser struct {
	Login   string `json:"login"`
	HtmlUrl string `json:"html_url"`
	
	Name      string `json:"name"`
	AvatarURL string `json:"avatar_url"`
	Company   string `json:"company"`
}

type SlackMessage struct {
	Channel     string       `json:"channel"`
	Username    string       `json:"username"`
	IconEmoji   string       `json:"icon_emoji"`
	Attachments []Attachment `json:"attachments"`
}

type Attachment struct {
	Color      string `json:"color"`
	AuthorName string `json:"author_name"`
	AuthorLink string `json:"author_link"`
	AuthorIcon string `json:"author_icon"`
	Title      string `json:"title"`
	TitleLink  string `json:"title_link"`
	Text       string `json:"text"`
}

func main() {
	githubUser, err := fetchUserInfoFromGithub("rodkranz")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	
	slackMessage := hydrateMessage(githubUser)
	
	err = sendMessageToSlack(slackMessage)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	
	os.Exit(0)
}

func (s SlackMessage) Bytes() []byte {
	bs, _ := json.Marshal(s)
	return bs
}

func hydrateMessage(github GithubUser) SlackMessage {
	atm := Attachment{
		Color:      "#36a64f",
		AuthorIcon: github.AvatarURL,
		AuthorName: github.Name,
		AuthorLink: github.HtmlUrl,
		TitleLink:  github.HtmlUrl,
		Title:      github.Login,
		Text:       fmt.Sprintf("%s is working at %s company.", github.Name, github.Company),
	}
	
	return SlackMessage{
		Channel:     "#rodrigo",
		Username:    BotName,
		IconEmoji:   BotAvatar,
		Attachments: []Attachment{atm},
	}
}

func fetchUserInfoFromGithub(username string) (gu GithubUser, err error) {
	r, err := http.Get(GithubBaseUrl + username)
	if err != nil {
		return gu, err
	}
	
	if r.StatusCode != http.StatusOK {
		return gu, fmt.Errorf("user not found")
	}
	
	err = json.NewDecoder(r.Body).Decode(&gu)
	if err != nil {
		return
	}
	
	return gu, nil
}

func sendMessageToSlack(message SlackMessage) error {
	buff := bytes.NewBuffer(message.Bytes())
	res, err := http.Post(SlackBaseUrl, "application/json", buff)
	if err != nil {
		return err
	}
	
	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("the status is not ok")
	}
	
	return nil
}

package main

import (
	"testing"
	"net/http/httptest"
	"net/http"
	"fmt"
	"bytes"
	"io/ioutil"
)

// Uncomment this test at Step 3
func TestFetchUserInfoFromGithub(t *testing.T) {
	t.Run("StatusCodeInvalid", func(t *testing.T) {
		// Fake server for tests
		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(`{"message": "Not Found","documentation_url": "https://developer.github.com/v3/users/#get-a-single-user"}`))
		}))
		
		// Define base url to fake server url
		GithubBaseUrl = ts.URL + "/"
		
		// execute main function
		_, err := fetchUserInfoFromGithub("rodkranz")
		if err == nil {
			t.Fatalf("Expected error, got none")
		}
		
		expected := "User not found"
		if err.Error() != expected {
			t.Fatalf("Expected error message %s, got %v", expected, err.Error())
		}
	})
	t.Run("Valid", func(t *testing.T) {
		// should be returned from fetch.
		expected := GithubUser{
			Login:     "rodkranz",
			Name:      "Rodrigo Lopes",
			Company:   "OLX",
			HtmlUrl:   "https://github.com/rodkranz",
			AvatarURL: "https://avatars2.githubusercontent.com/u/16897636?v=4",
		}
		
		// Fake server for tests
		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte(`{"login":"rodkranz","html_url":"https://github.com/rodkranz","name":"Rodrigo Lopes","avatar_url":"https://avatars2.githubusercontent.com/u/16897636?v=4","company":"OLX"}`))
		}))
		
		// Define base url to fake server url
		GithubBaseUrl = ts.URL + "/"
		
		// execute main function
		actual, err := fetchUserInfoFromGithub("rodkranz")
		if err != nil {
			t.Fatalf("Expected none error, got %s", err)
		}
		
		t.Run("Login", func(t *testing.T) {
			if expected.Login != actual.Login {
				t.Errorf("[Login] Expected %s, got %s", expected.Login, actual.Login)
			}
		})
		
		t.Run("Name", func(t *testing.T) {
			if expected.Name != actual.Name {
				t.Errorf("[Name] Expected %s, got %s", expected.Name, actual.Name)
			}
		})
		
		t.Run("Company", func(t *testing.T) {
			if expected.Company != actual.Company {
				t.Errorf("[Company] Expected %s, got %s", expected.Company, actual.Company)
			}
		})
		
		t.Run("HtmlUrl", func(t *testing.T) {
			if expected.HtmlUrl != actual.HtmlUrl {
				t.Errorf("[HtmlUrl] Expected %s, got %s", expected.HtmlUrl, actual.HtmlUrl)
			}
		})
		
		t.Run("AvatarURL", func(t *testing.T) {
			if expected.AvatarURL != actual.AvatarURL {
				t.Errorf("[AvatarURL] Expected %s, got %s", expected.AvatarURL, actual.AvatarURL)
			}
		})
	})
}

// Uncomment this test at Step 5
func TestHydrateMessage(t *testing.T) {
	tst := GithubUser{
		Login:     "rodkranz",
		Name:      "Rodrigo Lopes",
		Company:   "OLX",
		HtmlUrl:   "https://github.com/rodkranz",
		AvatarURL: "https://avatars2.githubusercontent.com/u/16897636?v=4",
	}
	
	actual := hydrateMessage(tst)
	expected := SlackMessage{
		Channel:   "#rodrigo",
		Username:  BotName,
		IconEmoji: BotAvatar,
		Attachments: []Attachment{
			{
				Color:      "#36a64f",
				AuthorIcon: "https://avatars2.githubusercontent.com/u/16897636?v=4",
				AuthorName: "Rodrigo Lopes",
				AuthorLink: "https://github.com/rodkranz",
				TitleLink:  "https://github.com/rodkranz",
				Title:      "rodkranz",
				Text:       fmt.Sprintf("%s is working at %s company.", "Rodrigo Lopes", "OLX"),
			},
		},
	}
	
	t.Run("Channel", func(t *testing.T) {
		if expected.Channel != actual.Channel {
			t.Errorf("[Channel] Expected %s, got %s", expected.Channel, actual.Channel)
		}
	})
	
	t.Run("Username", func(t *testing.T) {
		if expected.Username != actual.Username {
			t.Errorf("[Username] Expected %s, got %s", expected.Username, actual.Username)
		}
	})
	
	t.Run("IconEmoji", func(t *testing.T) {
		if expected.IconEmoji != actual.IconEmoji {
			t.Errorf("[IconEmoji] Expected %s, got %s", expected.IconEmoji, actual.IconEmoji)
		}
	})
	
	t.Run("Attachments", func(t *testing.T) {
		
		if len(expected.Attachments) != len(actual.Attachments) {
			t.Fatalf("[Attachments] Expected %d, got %d", len(expected.Attachments), len(actual.Attachments))
		}
		
		if len(actual.Attachments) == 0 {
			t.Fatalf("[Attachments] Expected more then 0, got %d", len(actual.Attachments))
		}
		
		actualAttac := expected.Attachments[0]
		expectedAttac := actual.Attachments[0]
		
		t.Run("Color", func(t *testing.T) {
			if actualAttac.Color != expectedAttac.Color {
				t.Errorf("[Attachment.Color] Expected %s, got %s", expectedAttac.Color, actualAttac.Color)
			}
		})
		
		t.Run("AuthorIcon", func(t *testing.T) {
			if actualAttac.AuthorIcon != expectedAttac.AuthorIcon {
				t.Errorf("[Attachment.AuthorIcon] Expected %s, got %s", expectedAttac.AuthorIcon, actualAttac.AuthorIcon)
			}
		})
		
		t.Run("AuthorName", func(t *testing.T) {
			if actualAttac.AuthorName != expectedAttac.AuthorName {
				t.Errorf("[Attachment.AuthorName] Expected %s, got %s", expectedAttac.AuthorName, actualAttac.AuthorName)
			}
		})
		
		t.Run("AuthorLink", func(t *testing.T) {
			if actualAttac.AuthorLink != expectedAttac.AuthorLink {
				t.Errorf("[Attachment.AuthorLink] Expected %s, got %s", expectedAttac.AuthorLink, actualAttac.AuthorLink)
			}
		})
		
		t.Run("TitleLink", func(t *testing.T) {
			if actualAttac.TitleLink != expectedAttac.TitleLink {
				t.Errorf("[Attachment.TitleLink] Expected %s, got %s", expectedAttac.TitleLink, actualAttac.TitleLink)
			}
		})
		
		t.Run("Title", func(t *testing.T) {
			if actualAttac.Title != expectedAttac.Title {
				t.Errorf("[Attachment.Title] Expected %s, got %s", expectedAttac.Title, actualAttac.Title)
			}
		})
		
		t.Run("Text", func(t *testing.T) {
			if actualAttac.Text != expectedAttac.Text {
				t.Errorf("[Attachment.Text] Expected %s, got %s", expectedAttac.Text, actualAttac.Text)
			}
		})
		
	})
}

// Uncomment this test at Step 6
func TestSlackMessage_Bytes(t *testing.T) {
	sm := SlackMessage{
		Username:  "Test",
		Channel:   "#rodrigo",
		IconEmoji: ":ghost:",
	}
	
	expected := []byte(`{"channel":"#rodrigo","username":"Test","icon_emoji":":ghost:","attachments":null}`)
	
	if len(expected) != len(sm.Bytes()) {
		t.Errorf("Expected %d, got %d", len(expected), len(sm.Bytes()))
	}
	
	if !bytes.Equal(expected, sm.Bytes()) {
		t.Errorf("Expected %s, got %s", expected, sm.Bytes())
	}
}

// Uncomment this test at Step 7
func TestSendMessageToSlack(t *testing.T) {
	sm := SlackMessage{
		Username:  "Test",
		Channel:   "#rodrigo",
		IconEmoji: ":ghost:",
	}
	
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		actual, err := ioutil.ReadAll(r.Body)
		if err != nil {
			t.Fatalf("Expected none error, got %v", err)
		}
		
		expected := []byte(`{"channel":"#rodrigo","username":"Test","icon_emoji":":ghost:","attachments":null}`)
		if len(expected) != len(sm.Bytes()) {
			t.Errorf("Expected %d, got %d", len(expected), len(sm.Bytes()))
		}
		
		if !bytes.Equal(expected, actual) {
			t.Errorf("Expected %s, got %s", expected, sm.Bytes())
		}
	}))
	
	SlackBaseUrl = ts.URL
	err := sendMessageToSlack(sm)
	if err != nil {
		t.Errorf("Expected none error, got %v", err)
	}
}

package twitch

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

var DefaultTwitchAPIEndpoint = "https://api.twitch.tv/helix"

type TwitchClient struct {
	ClientID    string
	BearerToken string
}

func NewTwitchClient(ClientID string, BearerToken string) *TwitchClient {
	return &TwitchClient{ClientID: ClientID, BearerToken: BearerToken}
}

func (t TwitchClient) FindUser(StreamerUsername string) (response UsersResponse, err error) {
	twitchURL := fmt.Sprintf("%s/users?login=%s", DefaultTwitchAPIEndpoint, StreamerUsername)
	log.Printf("Twitch URL: %s", twitchURL)
	req, err := http.NewRequest("GET", twitchURL, nil)
	if err != nil {
		return response, err
	}

	req.Header.Set("Client-ID", t.ClientID)
	req.Header.Set("Authorization", "Bearer "+t.BearerToken)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return response, err
	}
	defer res.Body.Close()

	err = json.NewDecoder(res.Body).Decode(&response)

	return response, err
}

func (t TwitchClient) GetFollowersForStreamer(TwitchID string) (response FollowersResponse, err error) {
	twitchFollowersURL := fmt.Sprintf("%s/channels/followers?user_id=%s", DefaultTwitchAPIEndpoint, TwitchID)
    log.Printf("Twitch URL: %s", twitchFollowersURL)

    reqFollowers, err := http.NewRequest("GET", twitchFollowersURL, nil)
    if err != nil {
        log.Printf("Error creating request: %v", err)
        return
    }
    reqFollowers.Header.Set("Client-ID", t.ClientID)
    reqFollowers.Header.Set("Authorization", "Bearer "+t.BearerToken)

    resFollowers, err := http.DefaultClient.Do(reqFollowers)
    if err != nil {
        log.Printf("Error making API request: %v", err)
        return
    }
    defer resFollowers.Body.Close()

    err = json.NewDecoder(resFollowers.Body).Decode(&response)
    if err != nil {
        log.Printf("Error decoding response: %v", err)
        return
    }

    log.Printf("Followers Response: %v", response)
    return
}

func (t TwitchClient) GetStreamInformation(StreamerUsername string) (response StreamInformationResponse, err error) {
	twitchURL := fmt.Sprintf("%s/streams?user_login=%s", DefaultTwitchAPIEndpoint, StreamerUsername)
	log.Printf("Twitch URL: %s", twitchURL)
	req, err := http.NewRequest("GET", twitchURL, nil)
	if err != nil {
		log.Printf("Failed to create request: %v", err)
		return
	}

	req.Header.Set("Client-ID", t.ClientID)
	req.Header.Set("Authorization", "Bearer "+t.BearerToken)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Printf("Failed to get stream data: %v", err)
		return
	}
	defer res.Body.Close()

	err = json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		log.Printf("Failed to decode response: %v", err)
		return
	}
	return
}

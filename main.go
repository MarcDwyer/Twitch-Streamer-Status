package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"runtime"
	"time"

	"github.com/joho/godotenv"
)

type Response struct {
	Stream struct {
		ID          *int64     `json:"id"`
		Game        *string    `json:"game"`
		Viewers     *int       `json:"viewers"`
		VideoHeight *int       `json:"videoheight"`
		AverageFps  *int       `json:"averggefps"`
		Delay       *int       `json:"delay"`
		CreatedAt   *time.Time `json:"createdAt"`
		IsPlaylist  *bool      `json:"isPlaylist"`
		Preview     struct {
			Small    *string `json:"small"`
			Medium   *string `json:"medium"`
			Large    *string `json:"large"`
			Template *string `json:"template"`
		}
		Channel struct {
			Mature                       bool        `json:"mature"`
			Status                       *string     `json:"status"`
			BroadcasterLanguage          string      `json:"broadcaster_language"`
			DisplayName                  string      `json:"display_name"`
			Game                         string      `json:"game"`
			Language                     string      `json:"language"`
			ID                           int         `json:"_id"`
			Name                         string      `json:"name"`
			CreatedAt                    time.Time   `json:"created_at"`
			UpdatedAt                    time.Time   `json:"updated_at"`
			Partner                      bool        `json:"partner"`
			Logo                         string      `json:"logo"`
			VideoBanner                  string      `json:"video_banner"`
			ProfileBanner                string      `json:"profile_banner"`
			ProfileBannerBackgroundColor interface{} `json:"profile_banner_background_color"`
			URL                          string      `json:"url"`
			Views                        int         `json:"views"`
			Followers                    int         `json:"followers"`
		} `json:"channel"`
	} `json:"stream"`
}

func init() {
	fmt.Println(runtime.NumCPU())
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	fmt.Println("starting...")
	url := fmt.Sprintf("https://api.twitch.tv/kraken/streams/richardlewisreports?client_id=%v", os.Getenv("CLIENT"))
	fmt.Println(url)
	var resp Response
	rz, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	}
	body, err := ioutil.ReadAll(rz.Body)

	json.Unmarshal(body, &resp)

	fmt.Println(resp)
	pr := fmt.Sprintf("status is %v", *resp.Stream.Channel.Status)
	fmt.Println(pr)

	if resp.Stream.Channel.Status == nil {
		fmt.Println("offline")
		return
	}
	fmt.Println("online")
}

package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	"github.com/ChimeraCoder/anaconda"
	"github.com/jasonlvhit/gocron"
)

// JSONData holds the array of Tweets
type JSONData struct {
	Tweets []Tweet
}

// Tweet is a scheduled Twitter status
type Tweet struct {
	Time   string
	Status string
}

var api *anaconda.TwitterApi

func main() {
	// Load the Twitter API keys

	consumerKey := os.Getenv("CONSUMER_KEY")

	if consumerKey == "" {
		log.Fatal("Environment variable CONSUMER_KEY not set. See https://apps.twitter.com/app/new for more info.")
	}

	consumerSecret := os.Getenv("CONSUMER_SECRET")

	if consumerSecret == "" {
		log.Fatal("Environment variable CONSUMER_SECRET not set. See https://apps.twitter.com/app/new for more info.")
	}

	accessToken := os.Getenv("ACCESS_TOKEN")

	if accessToken == "" {
		log.Fatal("Environment variable ACCESS_TOKEN not set. See https://apps.twitter.com/app/new for more info.")
	}

	accessTokenSecret := os.Getenv("ACCESS_TOKEN_SECRET")

	if accessTokenSecret == "" {
		log.Fatal("Environment variable ACCESS_TOKEN_SECRET not set. See https://apps.twitter.com/app/new for more info.")
	}

	// Read the config file

	configFile := os.Args[1]

	if _, err := os.Stat(configFile); os.IsNotExist(err) {
		log.Fatalf("File does not exist: %s", configFile)
		return
	}

	file, err := ioutil.ReadFile(configFile)
	if err != nil {
		log.Fatalf("Eror opening file %s: %v", configFile, err)
		return
	}

	// Deserialize the JSON data

	var tweets []Tweet

	err = json.Unmarshal(file, &tweets)
	if err != nil {
		log.Fatalf("Error while parsing file: %v", err)
		return
	}

	log.Println("Guffer is queueing the following tweets: ")

	// Init the Twitter API
	anaconda.SetConsumerKey(consumerKey)
	anaconda.SetConsumerSecret(consumerSecret)
	api = anaconda.NewTwitterApi(accessToken, accessTokenSecret)

	// Queue up the Tweets

	for _, tweet := range tweets {
		log.Printf("[%v] %s", tweet.Time, tweet.Status)
		gocron.Every(1).Day().At(tweet.Time).Do(postTweet, tweet.Status)
	}

	<-gocron.Start()
}

func postTweet(status string) {
	log.Println("Tweeting: ", status)
	api.PostTweet(status, nil)
}

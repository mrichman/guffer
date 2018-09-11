package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	"fmt"
	"github.com/BurntSushi/toml"
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

// Twitter Auth Keys
type TwitterAuthKeys struct {
	ConsumerKey       string
	ConsumerSecret    string
	AccessToken       string
	AccessTokenSecret string
}

var api *anaconda.TwitterApi
var auth TwitterAuthKeys

func main() {
	// Load the Twitter API keys

	// Check if the second argument with toml auth file was given, if not load auth keys from env variables
	if os.Args[2] == "" {
		auth.loadFromEnvVariables()
	} else {
		auth.loadFromTomlFile(os.Args[2])
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
	anaconda.SetConsumerKey(auth.ConsumerKey)
	anaconda.SetConsumerSecret(auth.ConsumerSecret)
	api = anaconda.NewTwitterApi(auth.AccessToken, auth.AccessTokenSecret)

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

func (t *TwitterAuthKeys) loadFromEnvVariables() {
	t.ConsumerKey = os.Getenv("CONSUMER_KEY")

	if t.ConsumerKey == "" {
		log.Fatal("Environment variable CONSUMER_KEY not set. See https://apps.twitter.com/app/new for more info.")
	}

	t.ConsumerSecret = os.Getenv("CONSUMER_SECRET")

	if t.ConsumerSecret == "" {
		log.Fatal("Environment variable CONSUMER_SECRET not set. See https://apps.twitter.com/app/new for more info.")
	}

	t.AccessToken = os.Getenv("ACCESS_TOKEN")

	if t.AccessToken == "" {
		log.Fatal("Environment variable ACCESS_TOKEN not set. See https://apps.twitter.com/app/new for more info.")
	}

	t.AccessTokenSecret = os.Getenv("ACCESS_TOKEN_SECRET")

	if t.AccessTokenSecret == "" {
		log.Fatal("Environment variable ACCESS_TOKEN_SECRET not set. See https://apps.twitter.com/app/new for more info.")
	}
}

func (t *TwitterAuthKeys) loadFromTomlFile(filename string) {
	// Load file to []byte
	data, err := ioutil.ReadFile(filename)
	// Check for error during loading
	if err != nil {
		log.Fatal(fmt.Sprintf("The %s file does not exists.", filename))
	}
	// Decode file contents
	if _, err := toml.Decode(string(data), t); err != nil {
		log.Fatal(fmt.Sprintf("Failed to decode %s file:%s", filename, err.Error()))
	}
}

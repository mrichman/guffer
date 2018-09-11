# Guffer

[![guffer Build Status](https://travis-ci.org/mrichman/guffer.svg?branch=master)](https://travis-ci.org/mrichman/guffer)&nbsp;[![GoDoc](https://godoc.org/github.com/mrichman/guffer?status.svg)](https://godoc.org/github.com/mrichman/guffer) [![Go Report Card](https://goreportcard.com/badge/github.com/mrichman/guffer)](https://goreportcard.com/report/github.com/mrichman/guffer) [![Join the chat at https://gitter.im/mrichman-guffer](https://badges.gitter.im/mrichman/guffer.svg)](https://gitter.im/mrichman-guffer) [![MIT License](https://img.shields.io/badge/license-MIT-blue.svg)](https://github.com/mrichman/guffer/blob/master/LICENSE) [![Issues](http://img.shields.io/github/issues/mrichman/guffer.svg)]( https://github.com/mrichman/guffer/issues )


Guffer tweets based on a daily schedule. If you've used tools like [Buffer](https://buffer.com) or [HootSuite](https://hootsuite.com), this will be familiar to you.

## Creating Twitter auth keys

Visit [https://apps.twitter.com/app/new](https://apps.twitter.com/app/new)

Set the following environment variables in either your user profile or at the command line:

* CONSUMER_KEY
* CONSUMER_SECRET
* ACCESS_TOKEN
* ACCESS_TOKEN_SECRET

You can also save these auth keys in `auth.toml` file (see below).
## Defining guffer.json

Guffer looks for a config file, for example `guffer.json`, which defines the schedule and status message (the tweet). Here's an example:

```
[  
  {
    "time": "15:22",
    "status": "Runs at 15:22 every day"
  },
  {
    "time": "15:23",
    "status": "Runs at 15:23 every day"
  },
  {
    "time": "15:26",
    "status": "Runs at 15:26 every day"
  }
]
```
## Defining auth.toml
If you prefer to save twitter auth keys in a file, you can create a `.toml` file with this data
```
ConsumerKey = "YOUR_CONSUMER_KEY"
ConsumerSecret = "YOUR_CONSUMER_SECRET"
AccessToken = "YOUR_ACCES_TOKEN"
AccessTokenSecret = "YOUR_TOKEN_SECRET"
```
## Running guffer

From source:

```
go get github.com/mrichman/guffer
cd $GOPATH/src/github.com/mrichman/guffer
CONSUMER_KEY=xxxxx CONSUMER_SECRET=xxxxx ACCESS_TOKEN=xxxxx ACCESS_TOKEN_SECRET=xxxxx go run main.go guffer.json
# With auth.toml
go run main.go guffer.json auth.toml
```

Binary:

```
CONSUMER_KEY=xxxxx CONSUMER_SECRET=xxxxx ACCESS_TOKEN=xxxxx ACCESS_TOKEN_SECRET=xxxxx guffer guffer.json
```
or if you're using the auth.toml file  
```
guffer guffer.json auth.toml
```
Guffer will print out a summary of the queued tweets, and log each tweet to the console. Quit with `Ctrl+C`.

# Contributing

If you find any bugs, please report them! I am also happy to accept pull requests from anyone.

You can use the [GitHub issue tracker](https://github.com/mrichman/guffer/issues) to report bugs, ask questions, or suggest new features.

For a more informal setting to discuss this project, you can join the [Gitter chat](https://gitter.im/mrichman/guffer).

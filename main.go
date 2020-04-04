package main

import (
	"bufio"
	"bytes"
	"crypto/rand"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"math/big"
	"net/http"
	"os"
	"strings"
)

var config struct {
	Welcome  []string `json:"welcome"`
	Username string   `json:"username"`
	Hostname string   `json:"hostname"`
	HomeDir  string   `json:"homeDir"`
	Webhook  string   `json:"webhook"`
}

type discordEmbed struct {
	Description string `json:"description"`
}

type discordMessage struct {
	Content string         `json:"content"`
	Embeds  []discordEmbed `json:"embeds"`
}

var currentDirectory = "~"

func main() {
	cfgPath := flag.String("cfg", "config.json", "the path to the config file")
	flag.Parse()

	{
		// Config parsing
		data, err := ioutil.ReadFile(*cfgPath)
		if err != nil {
			panic(err)
		}

		if err = json.Unmarshal(data, &config); err != nil {
			panic(err)
		}
	}

	sessID, err := rand.Int(rand.Reader, big.NewInt(1000000))
	if err != nil {
		panic(err)
	}

	fd := bufio.NewReader(os.Stdin)

	for _, msg := range config.Welcome {
		fmt.Println(msg)
	}

	// IO flow
	for {
		fmt.Printf("[%s@%s]:%s$ ", config.Username, config.Hostname, currentDirectory)
		input, err := fd.ReadString('\n')
		if err != nil {
			panic(err)
		}

		commands := strings.Split(strings.TrimSpace(input), " ")
		if len(commands) < 1 {
			continue
		}
		switch commands[0] {
		case "exit":
			os.Exit(0)
		case "pwd":
			fmt.Println(config.HomeDir)
		default:
			fmt.Printf("%s: command not found\n", commands[0])
		}

		if len(input) > 2048 {
			input = input[:2047]
		}

		cl := &http.Client{}

		var body discordMessage
		body.Content = fmt.Sprintf("new command from session %d", sessID)
		body.Embeds = []discordEmbed{{
			Description: input,
		}}

		data, err := json.Marshal(&body)
		if err != nil {
			panic(err)
		}

		req, err := http.NewRequest("POST", config.Webhook, bytes.NewReader(data))
		if err != nil {
			panic(err)
		}
		req.Header.Set("Content-Type", "application/json")

		resp, err := cl.Do(req)
		if err != nil {
			panic(err)
		}

		err = resp.Body.Close()
		if err != nil {
			panic(err)
		}
	}
}

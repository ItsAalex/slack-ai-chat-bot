package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/krognol/go-wolfram"
	"github.com/shomali11/slacker"
	"github.com/tidwall/gjson"
	witai "github.com/wit-ai/wit-go/v2"
)

var wolframClient *wolfram.Client

// func that takes parameter of type channel which is a pointer to a Command Event.
func printCommandEvents(analyticsChannel <-chan *slacker.CommandEvent) {
	for event := range analyticsChannel {
		fmt.Println("Command Events")
		fmt.Println(event.Timestamp)
		fmt.Println(event.Command)
		fmt.Println(event.Parameters)
		fmt.Println(event.Event)
		fmt.Println()
	}
}

func main() {
	//loading the enviroment file
	godotenv.Load(".env")

	// Taking tokens from env file which are neccessary to access the Slack API so we can
	// create client to interact with slack API
	bot := slacker.NewClient(os.Getenv("SLACK_BOT_TOKEN"), os.Getenv("SLACK_APP_TOKEN"))

	//taking tokens creating client interacting with wit.ai
	client := witai.NewClient(os.Getenv("WIT_AI_TOKEN"))
	wolframClient := &wolfram.Client{AppID: os.Getenv("WOLFRAM_APP_ID")}
	//this func prints out the events that bot subscribes to
	go printCommandEvents(bot.CommandEvents())

	//when u pass a parameter to a slack u need to pass it like message
	bot.Command("query - <message>", &slacker.CommandDefinition{
		Description: "send any question to wolfram",
		Examples: []string{
			"Who is the president of Serbia",
		}, //anonymous func
		Handler: func(bc slacker.BotContext, request slacker.Request, response slacker.ResponseWriter) {
			query := request.Param("message")

			msg, _ := client.Parse(&witai.MessageRequest{
				Query: query,
			})
			//marshals into json
			data, _ := json.MarshalIndent(msg, "", "    ")
			//converts into string and stores it into rough var
			rough := string(data[:])

			// .0.value means I want to access the first value from the json file
			value := gjson.Get(rough, "entities.wit$wolfram_search_query:wolfram_search_query.0.value")
			//converts the value into the string
			answer := value.String()
			//it uses client to interact with wolfram alphaAPI. It's passing the answer
			res, err := wolframClient.GetSpokentAnswerQuery(answer, wolfram.Metric, 1000)
			if err != nil {
				fmt.Println("There is an error")
			}
			fmt.Println(value)
			response.Reply(res)
		},
	})

	//stops the program. CTX = CONTEXT
	ctx, cancel := context.WithCancel(context.Background())
	//call this func after everything is done
	defer cancel()

	err := bot.Listen(ctx)

	if err != nil {
		log.Fatal(err)
	}
}

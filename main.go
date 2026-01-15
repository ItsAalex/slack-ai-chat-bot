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
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file:", err)
	}
	fmt.Println("✓ .env file loaded successfully")

	// Taking tokens from env file which are neccessary to access the Slack API so we can
	// create client to interact with slack API
	botToken := os.Getenv("SLACK_BOT_TOKEN")
	appToken := os.Getenv("SLACK_APP_TOKEN")

	if botToken == "" {
		log.Fatal("SLACK_BOT_TOKEN is empty")
	}
	if appToken == "" {
		log.Fatal("SLACK_APP_TOKEN is empty")
	}
	fmt.Println("✓ Slack tokens loaded")

	bot := slacker.NewClient(botToken, appToken)

	//taking tokens creating client interacting with wit.ai
	witToken := os.Getenv("WIT_AI_TOKEN")
	if witToken == "" {
		log.Fatal("WIT_AI_TOKEN is empty")
	}
	fmt.Println("✓ Wit.ai token loaded")

	client := witai.NewClient(witToken)
	wolframID := os.Getenv("WOLFRAM_APP_ID")
	if wolframID == "" {
		log.Fatal("WOLFRAM_APP_ID is empty")
	}
	fmt.Println("✓ Wolfram App ID loaded")

	wolframClient := &wolfram.Client{AppID: wolframID}
	//this func prints out the events that bot subscribes to
	go printCommandEvents(bot.CommandEvents())
	fmt.Println("✓ Command event listener started")

	//when u pass a parameter to a slack u need to pass it like message
	bot.Command("query - <message>", &slacker.CommandDefinition{
		Description: "send any question to wolfram",
		Examples: []string{
			"Who is the president of Serbia",
		}, //anonymous func
		Handler: func(bc slacker.BotContext, request slacker.Request, response slacker.ResponseWriter) {
			fmt.Println("🔵 [HANDLER] Command received!")

			query := request.Param("message")
			fmt.Println("📝 [HANDLER] Query extracted:", query)

			fmt.Println("🔄 [HANDLER] Parsing query with Wit.ai...")
			msg, parseErr := client.Parse(&witai.MessageRequest{
				Query: query,
			})
			if parseErr != nil {
				fmt.Println("❌ [HANDLER] Wit.ai parse error:", parseErr)
				response.Reply("Error parsing query with Wit.ai")
				return
			}

			//marshals into json
			data, marshalErr := json.MarshalIndent(msg, "", "    ")
			if marshalErr != nil {
				fmt.Println("❌ [HANDLER] JSON marshal error:", marshalErr)
				response.Reply("Error marshaling response")
				return
			}

			//converts into string and stores it into rough var
			rough := string(data[:])
			fmt.Println("📋 [HANDLER] Full JSON response:", rough)

			// .0.value means I want to access the first value from the json file
			value := gjson.Get(rough, "entities.wit$wolfram_search_query:wolfram_search_query.0.value")
			//converts the value into the string
			answer := value.String()
			fmt.Println("💡 [HANDLER] Extracted answer:", answer)

			if answer == "" {
				fmt.Println("⚠️ [HANDLER] No answer extracted from Wit.ai")
				response.Reply("Could not extract a question from your query")
				return
			}

			//it uses client to interact with wolfram alphaAPI. It's passing the answer
			fmt.Println("🔄 [HANDLER] Querying Wolfram Alpha...")
			res, err := wolframClient.GetSpokentAnswerQuery(answer, wolfram.Metric, 1000)
			if err != nil {
				fmt.Println("❌ [HANDLER] Wolfram error:", err)
				response.Reply("Error querying Wolfram Alpha: " + err.Error())
				return
			}

			fmt.Println("✅ [HANDLER] Response ready:", res)
			response.Reply(res)
		},
	})

	//stops the program. CTX = CONTEXT
	ctx, cancel := context.WithCancel(context.Background())
	//call this func after everything is done
	defer cancel()

	fmt.Println("✓ Bot is listening for commands...")
	listenErr := bot.Listen(ctx)

	if listenErr != nil {
		log.Fatal("Bot listen error:", listenErr)
	}
}

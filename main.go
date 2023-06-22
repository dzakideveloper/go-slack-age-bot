package main

import(
	"github.com/shomali11/slacker"
	"context"
	"strconv"
	"os"
	"fmt"
	"log"
)

func printCommandEvent(analysticsChannel <-chan *slacker.CommandEvent){
	for event := range analysticsChannel{
		println("Command Events")
		fmt.Println(event.Timestamp)
		fmt.Println(event.Command)
		fmt.Println(event.Parameters)
		fmt.Println(event.Event)
		fmt.Println()
	}
}

func main(){
	os.Setenv("SLACK_BOT_TOKEN", "xoxb-5467434298180-5462092795493-wkmjiLnjpK3JbgRya0BU2bED")
	os.Setenv("SLACK_APP_TOKEN", "xapp-1-A05DL20CY2Z-5450468980903-83fccf590c1bc7d83938a45f8c5d0e53f117466b2e6b416d96c64872565bc78b")

	bot := slacker.NewClient(os.Getenv(os.Getenv("SLACK_BOT_TOKEN")), os.Getenv("SLACK_APP_TOKEN"))

	go printCommandEvent(bot.CommandEvents())

	bot.Command("my yob is <year>", &slacker.CommandDefinition{
		Description: "yob calculator",
		Examples: []string{"my yob is 2000", "my yob is 1995"},
		Handler: func(bc slacker.BotContext, r slacker.Request, w slacker.ResponseWriter) {
			year := r.Param("year")
			yob, err := strconv.Atoi(year)
			if err != nil {
				println("error")
			}
			age := 2023 - yob
			re := fmt.Sprintf("age is %d", age)
			w.Reply(re)
		},
	})
	

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err := bot.Listen(ctx)
	if err != nil{
		log.Fatal(err)
	}
}
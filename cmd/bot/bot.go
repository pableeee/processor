package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Lukaesebrot/dgc"
	"github.com/bwmarrin/discordgo"
)

const (
	url = "http://127.0.0.1:8000/game/%s"
)

func main() {
	Main()
}

func Main() {
	// Open a simple Discord session
	token := os.Getenv("TOKEN")
	session, err := discordgo.New("Bot " + token)
	if err != nil {
		panic(err)
	}
	err = session.Open()
	if err != nil {
		panic(err)
	}

	// Wait for the user to cancel the process
	defer func() {
		sc := make(chan os.Signal, 1)
		signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
		<-sc
	}()
	session.Channel("")
	// Create a dgc router
	// NOTE: The dgc.Create function makes sure all the maps get initialized
	router := dgc.Create(&dgc.Router{
		// We will allow '!' and 'example!' as the bot prefixes
		Prefixes: []string{
			"!",
			"alan!",
		},

		// We will ignore the prefix case, so 'eXaMpLe!' is also a valid prefix
		IgnorePrefixCase: true,

		// We don't want bots to be able to execute our commands
		BotsAllowed: false,

		// We may initialize our commands in here, but we will use the corresponding method later on
		Commands: []*dgc.Command{},

		// We may inject our middlewares in here, but we will also use the corresponding method later on
		Middlewares: []dgc.Middleware{},

		// This handler gets called if the bot just got pinged (no argument provided)
		PingHandler: func(ctx *dgc.Ctx) {
			err := ctx.RespondText("Pong!")
			if err != nil {
				fmt.Printf("error Writing pong: %s", err.Error())
			}
		},
	})

	// Register the default help command
	router.RegisterDefaultHelpCommand(session, nil)

	// Register a simple middleware that injects a custom object
	router.RegisterMiddleware(func(next dgc.ExecutionHandler) dgc.ExecutionHandler {
		return func(ctx *dgc.Ctx) {
			// is this really working?
			log.Printf("Middleware log: %s", ctx.Command.Name)

			// Call the next execution handler
			next(ctx)
		}
	})

	// Register a simple command that responds with our custom object
	router.RegisterCmd(&dgc.Command{
		// We want to use 'obj' as the primary name of the command
		Name: "get",

		// We also want the command to get triggered with the 'object' alias
		Aliases: []string{
			"get",
		},

		// These fields get displayed in the default help messages
		Description: "Gets all games asociated with a user",
		Usage:       "get",
		Example:     "get pableeeee",

		// You can assign custom flags to a command to use them in middlewares
		Flags: []string{},

		// We want to ignore the command case
		IgnoreCase: true,

		// You may define sub commands in here
		SubCommands: []*dgc.Command{},

		// We want the user to be able to execute this command once in five seconds and the cleanup interval shpuld be one second
		RateLimiter: dgc.NewRateLimiter(5*time.Second, 1*time.Second, func(ctx *dgc.Ctx) {
			err := ctx.RespondText("You are being rate limited!")
			if err != nil {
				fmt.Printf("error limitin rate: %s", err.Error())
			}
		}),

		// Now we want to define the command handler
		Handler: objCommand,
	})

	router.Initialize(session)
}

func objCommand(ctx *dgc.Ctx) {

	write := func(ctx *dgc.Ctx, msg string) {
		err := ctx.RespondText(msg)
		if err != nil {
			fmt.Printf("error writing get response: %s", err.Error())
		}
	}
	// Respond with the just set custom object
	if ctx.Arguments.Amount() != 1 {
		write(ctx, "invalid number of arguments")
	}

	userID := ctx.Arguments.Get(0).Raw()
	if len(userID) == 0 {
		log.Fatal("invalid username")
		return
	}

	resp, err := http.Get(fmt.Sprintf(url, userID))
	if err != nil {
		log.Fatal("invalid number of arguments")
		return
	}

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("unable to read server response")
		return
	}

	write(ctx, string(b))
}

# tbauth
A simple Telegram bot authentication library
Authenticates user temporarily using a hardocded password

### Example

```
package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	. "github.com/logrusorgru/aurora"
	"github.com/mrl33t/tbauth"
	tb "gopkg.in/tucnak/telebot.v2"
)

func main() {

	c := make(chan os.Signal, 2)
	signal.Notify(c, os.Interrupt, syscall.SIGINT)
	go func() {
		<-c
		fmt.Println(Bold(Red("\r[-] Exiting")))
		os.Exit(0)
	}()

	b, err := tb.NewBot(tb.Settings{
		Token:  "YOUR-TOKEN-HERE",
		Poller: &tb.LongPoller{Timeout: 10 * time.Second},
	})

	if err != nil {
		log.Fatal(err)
		return
	}

	b.Handle("/start", func(m *tb.Message) {

		b.Send(m.Sender, "This is tbauth example, how can i help you?")
	})

	b.Handle("/help", func(m *tb.Message) {

		helpPrompt := `
/auth -> Authenticate

/dosomething -> Do something if authenticated

/help -> Help

		`
		b.Send(m.Sender, helpPrompt)
	})

	b.Handle("/auth", func(m *tb.Message) {
		splittedText := strings.Split(m.Text, " ")
		if len(splittedText) != 2 {
			b.Send(m.Sender, "Usage: /auth passphrase")
			return
		}

		resp := tbauth.Authenticate(*m.Sender, splittedText[1])
		b.Send(m.Sender, resp)
	})

    // Check if authenticated, otherwise exit
	b.Handle("/dosomething", func(m *tb.Message) {
		if tbauth.IsAuthenticated(m.Sender) == false {
			b.Send(m.Sender, "You're not authorized!")
			return
		}

        b.Send(c.Sender, "Hi! You're authorized")
	})

	b.Start()

}
```

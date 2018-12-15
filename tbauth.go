package tbauth

import (
	tb "gopkg.in/tucnak/telebot.v2"
)

var (
	authenticatedUserIDs []int
	Passphrase           string = "1337"
)

func Authenticate(sender *tb.User, pass string) int {
	if pass != Passphrase {
		// Wrong passphrase
		return 2
	}
	for _, val := range authenticatedUserIDs {
		if val == sender.ID {
			// Already authenticated
			return 1
		}
	}

	authenticatedUserIDs = append(authenticatedUserIDs, sender.ID)

	return 0
}

func IsAuthenticated(sender *tb.User) bool {
	for _, val := range authenticatedUserIDs {
		if val == sender.ID {
			return true
		}
	}
	return false
}

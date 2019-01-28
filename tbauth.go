package tbauth

import (
	tb "gopkg.in/tucnak/telebot.v2"
)

var (
	AuthenticatedUsers []*tb.User
	Passphrase         string = "1337"
)

func Authenticate(sender *tb.User, pass string) int {
	if pass != Passphrase {
		// Wrong passphrase
		return 2
	}
	for _, u := range AuthenticatedUsers {
		if u == sender {
			// Already authenticated
			return 1
		}
	}

	AuthenticatedUsers = append(AuthenticatedUsers, sender)

	return 0
}

func IsAuthenticated(sender *tb.User) bool {
	for _, u := range AuthenticatedUsers {
		if u == sender {
			return true
		}
	}
	return false
}

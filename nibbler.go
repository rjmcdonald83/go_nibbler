package go_nibbler

import "strings"

type state struct {
	InQuotes, PreviousDot, PreviousSlash, InDomain, PreviousAt, PreviousDash bool
}

const (
	ATEXT    = "!#$%&'*+-/=?^_`.{|}~@\"abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789ÀÁÂÃÄÅÆÇÈÉÊËÌÍÎÏÐÑÒÓÔÕÖ×ØÙÚÛÜÝÞSSÀÁÂÃÄÅÆÇÈÉÊËÌÍÎÏÐÑÒÓÔÕÖ÷ØÙÚÛÜÝÞŸàáâãäåæçèéêëìíîïðñòóôõö×øùúûüýþssàáâãäåæçèéêëìíîïðñòóôõö÷øùúûüýþÿ"
	SPECIAL  = "(),:;<>[\\] "
	HOSTNAME = ":-.abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789ÀÁÂÃÄÅÆÇÈÉÊËÌÍÎÏÐÑÒÓÔÕÖ×ØÙÚÛÜÝÞSSÀÁÂÃÄÅÆÇÈÉÊËÌÍÎÏÐÑÒÓÔÕÖ÷ØÙÚÛÜÝÞŸàáâãäåæçèéêëìíîïðñòóôõö×øùúûüýþssàáâãäåæçèéêëìíîïðñòóôõö÷øùúûüýþÿı"
)

func ParseEmail(email string) (bool, string) {
	valid := true
	var address string
	currentState := state{}

	// Max length of emails is 254 chars http://stackoverflow.com/a/574698
	if len(email) > 254 {
		return false, email[:254]
	}
	if strings.HasSuffix(email, "@") {
		return false, email[:len(email)-1]
	}
	if strings.HasSuffix(email, "-") {
		return false, email[:len(email)-1]
	}
	if strings.HasSuffix(email, ".") {
		return false, email[:len(email)-1]
	}

	for offset, character := range email {
		// Local part
		if !currentState.InDomain {
			if character == ' ' && !currentState.InQuotes {
				valid = false
				break
			} else if character == '\\' {
				if currentState.InQuotes {
					// Check if slash was backslashed within quotes
					if currentState.PreviousSlash {
						currentState.PreviousSlash = false
					} else {
						currentState.PreviousSlash = true
					}
					// \ can only occur within slashes
				} else {
					valid = false
					break
				}
			} else if character == '"' {
				if currentState.InQuotes {
					// Ignore if it was preceded by a backslash
					if !currentState.PreviousSlash {
						currentState.InQuotes = false
					} else {
						currentState.PreviousSlash = false
					}
				} else {
					// Quotes must happen as the first character or after a dot
					if offset != 0 && !currentState.PreviousDot {
						valid = false
						break
					} else {
						currentState.InQuotes = true
					}
				}
			} else if character == '.' {
				// We can't start with a dot
				if offset == 0 {
					valid = false
					break
				}
				// We can't have two consecutive dots
				if !currentState.InQuotes && currentState.PreviousDot {
					valid = false
					break
				}
				if !currentState.InQuotes {
					currentState.PreviousDot = true
				}
			} else if character == '@' && !currentState.InQuotes {
				// can't start with a @
				if offset == 0 {
					valid = false
					break
				}

				// can't have a dot before the @
				if currentState.PreviousDot {
					valid = false
					break
				}
				currentState.InDomain = true
				currentState.PreviousAt = true
			}
			if strings.ContainsRune(SPECIAL, character) {
				// These characters must only occur in quotes
				if strings.ContainsRune(SPECIAL, character) {
					if !currentState.InQuotes {
						valid = false
						break
					} else {
						address += string(character)
					}
				}
			} else {
				if !strings.ContainsRune(ATEXT, character) {
					valid = false
					break
				} else {
					address += string(character)
				}
			}
			// Check states and clear them if necessary
			if currentState.PreviousSlash && character != '\\' {
				currentState.PreviousSlash = false
			}
			if currentState.PreviousDot && character != '.' {
				currentState.PreviousDot = false
			}
		} else {
			if !strings.ContainsRune(HOSTNAME, character) {
				valid = false
				break
			}

			if character == '.' {
				// We can't have the domain start with a dot
				if currentState.PreviousAt {
					valid = false
					break
				}

				// We can't have two consecutive dots, even in the domain
				if currentState.PreviousDot {
					valid = false
					break
				} else {
					currentState.PreviousDot = true
				}

				// We can't have a dot follow a dash
				if currentState.PreviousDash {
					valid = false
					break
				}
			}

			if character == '-' {
				// We can't have the domain start with a dash
				if currentState.PreviousAt {
					valid = false
					break
				}

				// We can't have two consecutive dashes, even in the domain
				if currentState.PreviousDash {
					valid = false
					break
				} else {
					currentState.PreviousDash = true
				}

				// We can't have a dash follow a dot
				if currentState.PreviousDot {
					valid = false
					break
				}
			}

			address += string(character)
			// Check states and clear them if necessary
			if currentState.PreviousSlash && character != '\\' {
				currentState.PreviousSlash = false
			}
			if currentState.PreviousDot && character != '.' {
				currentState.PreviousDot = false
			}
			if currentState.PreviousDash && character != '-' {
				currentState.PreviousDash = false
			}
			if currentState.PreviousAt && character != '@' {
				currentState.PreviousAt = false
			}
		}
	}
	if !currentState.InDomain {
		valid = false
	}
	return valid, address
}

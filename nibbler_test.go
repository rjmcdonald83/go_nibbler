package go_nibbler

import "testing"

type example struct {
	Input   string
	Valid   bool
	Address string
}

var examples = []example{
	example{"woo", false, "woo"},
	example{"woo@woot.com", true, "woo@woot.com"},
	example{"woo/+@blah.com", true, "woo/+@blah.com"},
	example{"#!$%&'*+-/=?^_`{}|~@example.org", true, "#!$%&'*+-/=?^_`{}|~@example.org"},
	example{"\"Bob\" <bobthebuilder@dlc.com>", false, "\"Bob\""},
	// example{"bobthebuilder@176.2.0.234", false, "bobthebuilder@176.2.0.234"},
}

func TestNibbler(t *testing.T) {
	for _, e := range examples {
		valid, address := ParseEmail(e.Input)
		if valid != e.Valid || address != e.Address {
			t.Errorf("Failed for \"%s\": wanted {%t, %s}, got {%t, %s}", e.Input, e.Valid, e.Address, valid, address)
		}
	}
}

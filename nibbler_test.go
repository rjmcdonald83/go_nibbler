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
	example{"hoot@hoot-hoot.com", true, "hoot@hoot-hoot.com"},
	example{"woo/+@blah.com", true, "woo/+@blah.com"},
	example{"#!$%&'*+-/=?^_`{}|~@example.org", true, "#!$%&'*+-/=?^_`{}|~@example.org"},
	example{"üñîçøðé@üñîçøðéı.com", true, "üñîçøðé@üñîçøðéı.com"},
	example{"\"Bob\" <bobthebuilder@dlc.com>", false, "\"Bob\""},
}

func TestNibbler(t *testing.T) {
	for _, e := range examples {
		valid, address := ParseEmail(e.Input)
		if valid != e.Valid || address != e.Address {
			t.Errorf("Failed for \"%s\": wanted {%t, %s}, got {%t, %s}", e.Input, e.Valid, e.Address, valid, address)
		}
	}
}

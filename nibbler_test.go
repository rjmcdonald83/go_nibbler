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
	example{"üñîçøðé@üñîçøðé.com", true, "üñîçøðé@üñîçøðé.com"},
	example{"\"Bob\" <bobthebuilder@dlc.com>", false, "\"Bob\""},
	example{"really.long.but.vaild.address@example.com", true, "really.long.but.vaild.address@example.com"},
	example{"Bob bobthebuilder@dlc.com", false, "Bob"},
	example{"A@b@c@example.com", false, "A@b"},
	example{"a\"b(c)d,e:f;g<h>i[j\\k]l@example.com", false, "a"},
	example{"[j\\k]l@example.com", false, ""},
	example{"john..doe@example.com", false, "john."},
	example{"john......doe@example.com", false, "john."},
	example{"john.doe@example..com", false, "john.doe@example."},
	example{" johndoe@example.com", false, ""},
	example{"johndoe@example.com ", false, "johndoe@example.com"},
}

func TestNibbler(t *testing.T) {
	for _, e := range examples {
		valid, address := ParseEmail(e.Input)
		if valid != e.Valid || address != e.Address {
			t.Errorf("Failed for \"%s\": wanted {%t, %s}, got {%t, %s}", e.Input, e.Valid, e.Address, valid, address)
		}
	}
}

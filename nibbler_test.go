package go_nibbler

import (
	"fmt"
	"strings"
	"testing"
)

type example struct {
	Input   string
	Valid   bool
	Address string
}

var validExamples = []example{
	example{"hoot@hoot-hoot.com", true, "hoot@hoot-hoot.com"},
	example{"woo/+@blah.com", true, "woo/+@blah.com"},
	example{"üñîçøðé@üñîçøðéı.com", true, "üñîçøðé@üñîçøðéı.com"},
	example{"johndoe@example.co.uk", true, "johndoe@example.co.uk"},
	example{"niceandsimple@example.com", true, "niceandsimple@example.com"},
	example{"very.common@example.com", true, "very.common@example.com"},
	example{"a.little.lengthy.but.fine@dept.example.com", true, "a.little.lengthy.but.fine@dept.example.com"},
	example{"disposable.style.email.with+symbol@example.com", true, "disposable.style.email.with+symbol@example.com"},
	example{"other.email-with-dash@example.com", true, "other.email-with-dash@example.com"},
	example{`"much.more unusual"@example.com`, true, `"much.more unusual"@example.com`},
	example{`"very.unusual.@.unusual.com"@example.com`, true, `"very.unusual.@.unusual.com"@example.com`},
	example{`"very.(),:;<>[]\".VERY.\"very@\\ \"very\".unusual"@strange.example.com`, true, `"very.(),:;<>[]\".VERY.\"very@\\ \"very\".unusual"@strange.example.com`},
	example{"postbox@com", true, "postbox@com"},
	example{"admin@mailserver1", true, "admin@mailserver1"},
	example{"!#$%&'*+-/=?^_`{}|~@example.org", true, "!#$%&'*+-/=?^_`{}|~@example.org"},
	example{`"()<>[]:,;@\\\"!#$%&'*+-/=?^_` + "`{}| ~.a\"@example.org", true, `"()<>[]:,;@\\\"!#$%&'*+-/=?^_` + "`{}| ~.a\"@example.org"},
	example{`" "@example.org`, true, `" "@example.org`},
	example{`abc."defghi".xyz@example.com`, true, `abc."defghi".xyz@example.com`},
	example{`"two..dot"@example.com`, true, `"two..dot"@example.com`},
	example{`"test."@example.com`, true, `"test."@example.com`},
	example{"_somename@example.com", true, "_somename@example.com"},
	example{"customer/department=shipping@example.com", true, "customer/department=shipping@example.com"},
	example{"!def!xyz.c%abc+@example.com", true, "!def!xyz.c%abc+@example.com"},
	example{"1234567890.0987654321@EXAMPLE.WHATEVER.ABC.123.CAT", true, "1234567890.0987654321@EXAMPLE.WHATEVER.ABC.123.CAT"},
	example{"relay.mil%john.smith@EXAMPLE-WHATEVER.ARPA", true, "relay.mil%john.smith@EXAMPLE-WHATEVER.ARPA"},
	example{"joe@123.45.67.89", true, "joe@123.45.67.89"},
	example{"joe@2001:0db8::1428:57ab", true, "joe@2001:0db8::1428:57ab"},
	example{`"first".second@employs.allowable.trick`, true, `"first".second@employs.allowable.trick`},
	example{`""@the-void.example.com`, true, `""@the-void.example.com`},
}

var invalidExamples = []example{
	example{"woo", false, "woo"},
	example{"\"Bob\" <bobthebuilder@dlc.com>", false, "\"Bob\""},
	example{"[j\\k]l@example.com", false, ""},
	example{" johndoe@example.com", false, ""},
	example{"johndoe@example.com ", false, "johndoe@example.com"},
	example{"Abc.example.com", false, "Abc.example.com"},
	example{"A@b@c@example.com", false, "A@b"},
	example{`a"b(c)d,e:f;g<h>i[j\\k]l@example.com`, false, `a`},
	example{`just"not"right@example.com`, false, `just`},
	example{`this is"not\\allowed@example.com`, false, `this`},
	example{`this\\ still\\"not\\\\allowed@example.com`, false, `this`},
	example{`abc"defghi"xyz@example.com`, false, `abc`},
	example{"invalid@", false, "invalid"},
	example{"john@example...com", false, "john@example."},
	example{"test@.com", false, "test@"},
	example{"test@-.com", false, "test@"},
	example{"test@example--thing.com", false, "test@example-"},
	example{"test@example.com-", false, "test@example.com"},
	example{".invalid@example.com", false, ""},
	example{"invalid.@example.com", false, "invalid."},
	example{"@example.com", false, ""},
	example{"hello world@example.com", false, "hello"},
	example{"bird@test example.com", false, "bird@test"},
	example{`\$A12345@example.com`, false, ""},
	example{`"quote@example.com`, false, `"quote@example.com`},
	example{`quote\"@example.com`, false, `quote`},
	example{"two..dot@example.com", false, "two."},
	example{"user@???", false, "user@"},
	example{"user@example-.com", false, "user@example-"},
	example{"user@test.-example.com", false, "user@test."},
	example{"test@example.com.", false, "test@example.com"},
	example{"test@[example.com]", false, "test@"},
	example{strings.Replace(fmt.Sprintf(fmt.Sprintf("%%%ds", 255), "x"), " ", "x", -1), false, strings.Replace(fmt.Sprintf(fmt.Sprintf("%%%ds", 254), "x"), " ", "x", -1)},
}

func TestValidEmail(t *testing.T) {
	for _, e := range validExamples {
		valid, address := ParseEmail(e.Input)
		if valid != e.Valid || address != e.Address {
			t.Errorf("Failed for \"%s\": wanted {%t, %s}, got {%t, %s}", e.Input, e.Valid, e.Address, valid, address)
		}
	}
}

func TestInvalidEmail(t *testing.T) {
	for _, e := range invalidExamples {
		valid, address := ParseEmail(e.Input)
		if valid != e.Valid || address != e.Address {
			t.Errorf("Failed for \"%s\": wanted {%t, %s}, got {%t, %s}", e.Input, e.Valid, e.Address, valid, address)
		}
	}
}

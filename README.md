# GoNibbler

Validate email addresses

Based on https://github.com/sendgridlabs/nibbler-python and https://github.com/sendgrid/ruby_nibbler

## Installation

```
go get -u github.com/sendgrid/go_nibbler
```

## Usage

```
import (
	nibbler "github.com/sendgrid/go_nibbler"
)
valid, parsedEmail := nibbler.ParseEmail(email)
```

## Contributing

1. Fork it ( https://github.com/[my-github-username]/go_nibbler/fork )
2. Create your feature branch (`git checkout -b my-new-feature`)
3. Commit your changes (`git commit -am 'Add some feature'`)
4. Push to the branch (`git push origin my-new-feature`)
5. Create a new Pull Request

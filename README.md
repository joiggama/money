# Money

[![travis](https://travis-ci.org/joiggama/money.svg)](https://travis-ci.org/joiggama/money)
[![godoc](https://godoc.org/github.com/joiggama/money?status.svg)](https://godoc.org/github.com/joiggama/money)

A golang library to deal with money and currency representation.
Inspired by [ruby money library](http://rubymoney.github.io/money).

## Installation

    $ go get github.com/joiggama/money

## Usage

First, import the package adding:

```go
import "github.com/joiggama/money"
```

Examples:

```go
money.Format(10)                                                     // "$10.00"
money.Format(10, money.Options{"currency": "EUR"})                   // "€10.00"
money.Format(10, money.Options{"with_cents": false})                 // "$10"
money.Format(10, money.Options{"with_currency:" true })              // "$10.00 USD"
money.Format(10, money.Options{"with_symbol": false})                // "10.00"
money.Format(10, money.Options{"with_symbol_space":true})            // "$ 10.00"
money.Format(1000)                                                   // "$1,000.00"
money.Format(1000, money.Options{"with_thousands_separator": false}) // "$1000.00"
```

For more detailed documentation refer to [godoc](http://godoc.org/github.com/joiggama/money)

## Contributing

1. [Fork it](https://github.com/joiggama/money/fork)
2. Create your feature branch (`git checkout -b my-new-feature`)
3. Commit your changes (`git commit -am 'Add some feature'`)
4. Push to the branch (`git push origin my-new-feature`)
5. Create a new Pull Request

## License

[The MIT licence](LICENSE.md)

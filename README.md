# mqtt-wildcard

[![License][mit-badge]][mit-url]

> Match a MQTT Topic against Wildcards

## Install and Usage

``` bash
go get github.com/panicneo/mqtt-wildcard
```

## API

``` go
type MatchResult struct {
    Result []string
}

func Match(topic, wildcard string) *MatchResult{}
```

Returns `nil` if not matched, otherwise an pointer to `MatchResult` containing the wildcards contents will be returned.

Examples:

``` go
mqttwildcard.Match('test/foo/bar', 'test/foo/bar'); // &MatchResult{Result: []string{}}
mqttwildcard.Match('test/foo/bar', 'test/+/bar'); // &MatchResult{Result: []string{"foo"}}
mqttwildcard.Match('test/foo/bar', 'test/#'); // &MatchResult{Result: []string{"foo/bar"}}]
mqttwildcard.Match('test/foo/bar/baz', 'test/+/#'); // &MatchResult{Result: []string{"foo", "bar/baz"}}]
mqttwildcard.Match('test/foo/bar/baz', 'test/+/+/baz'); // &MatchResult{Result: []string{"foo", "bar"}}]

mqttwildcard.Match('test', 'test/#'); //  &MatchResult{Result: []string{}]
mqttwildcard.Match('test/', 'test/#'); //  &MatchResult{Result: []string{""}]

mqttwildcard.Match('test/foo/bar', 'test/+'); // nil
mqttwildcard.Match('test/foo/bar', 'test/nope/bar'); // nil
```

## License

MIT (c) 2017 [NEO](https://github.com/panicneo)

[mit-badge]: https://img.shields.io/badge/License-MIT-blue.svg?style=flat
[mit-url]: LICENSE

package mqttwildcard

import (
	"reflect"
	"testing"
)

type Case struct {
	Topic    string
	Wildcard string
	Expect   *MatchResult
}

var cases = []Case{
	{
		Topic:    "test/123",
		Wildcard: "test/123",
		Expect:   &MatchResult{Result: []string{}},
	},
	{
		Topic:    "test/test/test",
		Wildcard: "test/test",
		Expect:   nil,
	},
	{
		Topic:    "test/test",
		Wildcard: "test/test/test",
		Expect:   nil,
	},
	{
		Topic:    "test/test",
		Wildcard: "test/test/test/test",
		Expect:   nil,
	},
	// matching #
	{
		Topic:    "test",
		Wildcard: "#",
		Expect:   &MatchResult{Result: []string{"test"}},
	},
	{
		Topic:    "test/test",
		Wildcard: "#",
		Expect:   &MatchResult{Result: []string{"test/test"}},
	},
	{
		Topic:    "test/test",
		Wildcard: "test/#",
		Expect:   &MatchResult{Result: []string{"test"}},
	},
	{
		Topic:    "test/test/test",
		Wildcard: "test/#",
		Expect:   &MatchResult{Result: []string{"test/test"}},
	},
	{
		Topic:    "/",
		Wildcard: "/#",
		Expect:   &MatchResult{Result: []string{""}},
	},
	{
		Topic:    "/test",
		Wildcard: "/#",
		Expect:   &MatchResult{Result: []string{"test"}},
	},
	{
		Topic:    "/test/",
		Wildcard: "/#",
		Expect:   &MatchResult{Result: []string{"test/"}},
	},
	{
		Topic:    "test/test",
		Wildcard: "test/test/#",
		Expect:   &MatchResult{Result: []string{}},
	},
	// mismatching #
	{
		Topic:    "test",
		Wildcard: "/#",
		Expect:   nil,
	},
	{
		Topic:    "test/test",
		Wildcard: "test/#",
		Expect:   &MatchResult{Result: []string{"test"}},
	},
	{
		Topic:    "",
		Wildcard: "test/#",
		Expect:   nil,
	},
	// matching +
	{
		Topic:    "test",
		Wildcard: "+",
		Expect:   &MatchResult{Result: []string{"test"}},
	},
	{
		Topic:    "test/",
		Wildcard: "test/+",
		Expect:   &MatchResult{Result: []string{""}},
	},
	{
		Topic:    "test/test",
		Wildcard: "test/+",
		Expect:   &MatchResult{Result: []string{"test"}},
	},
	{
		Topic:    "test/test/test",
		Wildcard: "test/+/+",
		Expect:   &MatchResult{Result: []string{"test", "test"}},
	},
	{
		Topic:    "test/test/test",
		Wildcard: "test/+/test",
		Expect:   &MatchResult{Result: []string{"test"}},
	},
	// mismatching +
	{
		Topic:    "test",
		Wildcard: "/+",
		Expect:   nil,
	},
	{
		Topic:    "test",
		Wildcard: "test/+",
		Expect:   nil,
	},
	{
		Topic:    "test/test",
		Wildcard: "test/test/+",
		Expect:   nil,
	},
	// matching + #

	{
		Topic:    "test/test",
		Wildcard: "+/#",
		Expect:   &MatchResult{Result: []string{"test", "test"}},
	},
	{
		Topic:    "test/test/",
		Wildcard: "+/test/#",
		Expect:   &MatchResult{Result: []string{"test", ""}},
	},
	{
		Topic:    "test/test/",
		Wildcard: "test/+/#",
		Expect:   &MatchResult{Result: []string{"test", ""}},
	},
	{
		Topic:    "test/test/test",
		Wildcard: "+/test/#",
		Expect:   &MatchResult{Result: []string{"test", "test"}},
	},
	{
		Topic:    "test/test/test",
		Wildcard: "test/+/#",
		Expect:   &MatchResult{Result: []string{"test", "test"}},
	},
	{
		Topic:    "test/test/test",
		Wildcard: "+/+/#",
		Expect:   &MatchResult{Result: []string{"test", "test", "test"}},
	},
	{
		Topic:    "test/test/test/test",
		Wildcard: "test/+/+/#",
		Expect:   &MatchResult{Result: []string{"test", "test", "test"}},
	},
	{
		Topic:    "test",
		Wildcard: "+/#",
		Expect:   &MatchResult{Result: []string{"test"}},
	},
	{
		Topic:    "test/test",
		Wildcard: "test/+/#",
		Expect:   &MatchResult{Result: []string{"test"}},
	},
	{
		Topic:    "test/test/test",
		Wildcard: "test/+/test/#",
		Expect:   &MatchResult{Result: []string{"test"}},
	},
	// mismatching + #
	{
		Topic:    "test/foo/test",
		Wildcard: "+/test/#",
		Expect:   nil,
	},
	{
		Topic:    "foo/test/test",
		Wildcard: "test/+/#",
		Expect:   nil,
	},
	{
		Topic:    "foo/test/test/test",
		Wildcard: "test/+/+/#",
		Expect:   nil,
	},

	// readme examples
	{
		Topic:    "test/foo/bar",
		Wildcard: "test/foo/bar",
		Expect:   &MatchResult{Result: []string{}},
	},
	{
		Topic:    "test/foo/bar",
		Wildcard: "test/+/bar",
		Expect:   &MatchResult{Result: []string{"foo"}},
	},
	{
		Topic:    "test/foo/bar",
		Wildcard: "test/#",
		Expect:   &MatchResult{Result: []string{"foo/bar"}},
	},

	{
		Topic:    "test/foo/bar/baz",
		Wildcard: "test/+/#",
		Expect:   &MatchResult{Result: []string{"foo", "bar/baz"}},
	},
	{
		Topic:    "test/foo/bar/baz",
		Wildcard: "test/+/+/baz",
		Expect:   &MatchResult{Result: []string{"foo", "bar"}},
	},

	{
		Topic:    "test",
		Wildcard: "test/#",
		Expect:   &MatchResult{Result: []string{}},
	},
	{
		Topic:    "test/",
		Wildcard: "test/#",
		Expect:   &MatchResult{Result: []string{""}},
	},

	{
		Topic:    "test/foo/bar",
		Wildcard: "test/+",
		Expect:   nil,
	},
	{
		Topic:    "test/foo/bar",
		Wildcard: "test/nope/baz",
		Expect:   nil,
	},
}

func TestMatch(t *testing.T) {
	for i := range cases {
		c := cases[i]
		r := Match(c.Topic, c.Wildcard)
		if !reflect.DeepEqual(r, c.Expect) {
			t.Errorf("t: %v, w: %v, r: %v, e: %v", c.Topic, c.Wildcard, r, c.Expect)
		}
	}
}

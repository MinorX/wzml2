package httpserver

import (
	"regexp"

	gomock "github.com/golang/mock/gomock"
)

var _ Logger = (*testLogger)(nil)

type testLogger struct{}

func (t *testLogger) Info(msg string)  {}
func (t *testLogger) Warn(msg string)  {}
func (t *testLogger) Error(msg string) {}

var _ gomock.Matcher = (*regexMatcher)(nil)

type regexMatcher struct {
	regexp *regexp.Regexp
}

func (r *regexMatcher) Matches(x interface{}) bool {
	s, ok := x.(string)
	if !ok {
		return false
	}
	return r.regexp.MatchString(s)
}

func (r *regexMatcher) String() string {
	return "regular expression " + r.regexp.String()
}

func newRegexMatcher(regex string) *regexMatcher {
	return &regexMatcher{
		regexp: regexp.MustCompile(regex),
	}
}

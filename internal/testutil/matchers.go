package testutil

import (
	"github.com/golang/mock/gomock"
)

type matcherFunc struct {
	f func(x interface{}) error
	lastProblem string
}

func (m *matcherFunc) Matches(x interface{}) bool {
	if err := m.f(x); err != nil {
		m.lastProblem = err.Error()
		return false
	}
	return true
}

func (m *matcherFunc) String() string {
	return m.lastProblem
}

// MatcherFunc wraps f as gomock.Matcher. Matcher matches x if
// f(x) == nil. Matcher's String() method returns string representation
// of the last non-nil return value of f.
func MatcherFunc(f func(x interface{}) error) gomock.Matcher {
	return &matcherFunc{
		f: f,
	}
}

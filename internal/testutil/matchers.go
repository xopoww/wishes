package testutil

import (
	"github.com/golang/mock/gomock"
)

type matcherFunc struct {
	f func(x interface{}) bool
}

func (m *matcherFunc) Matches(x interface{}) bool {
	return m.f(x)
}

func (m *matcherFunc) String() string {
	return "matcherFunc"
}

func MatcherFunc(f func(x interface{}) bool) gomock.Matcher {
	return &matcherFunc{
		f: f,
	}
}

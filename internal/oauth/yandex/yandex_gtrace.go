// Code generated by gtrace. DO NOT EDIT.

package yandex

import (
	"net/http"
)

// Compose returns a new Trace which has functional fields composed
// both from t and x.
func (t Trace) Compose(x Trace) (ret Trace) {
	switch {
	case t.OnValidate == nil:
		ret.OnValidate = x.OnValidate
	case x.OnValidate == nil:
		ret.OnValidate = t.OnValidate
	default:
		h1 := t.OnValidate
		h2 := x.OnValidate
		ret.OnValidate = func(o OnValidateStartInfo) func(OnValidateDoneInfo) {
			r1 := h1(o)
			r2 := h2(o)
			switch {
			case r1 == nil:
				return r2
			case r2 == nil:
				return r1
			default:
				return func(o OnValidateDoneInfo) {
					r1(o)
					r2(o)
				}
			}
		}
	}
	return ret
}
func (t Trace) onValidate(o OnValidateStartInfo) func(OnValidateDoneInfo) {
	fn := t.OnValidate
	if fn == nil {
		return func(OnValidateDoneInfo) {
			return
		}
	}
	res := fn(o)
	if res == nil {
		return func(OnValidateDoneInfo) {
			return
		}
	}
	return res
}
func traceOnValidate(t Trace, req *http.Request) func(resp InfoResponse, _ error) {
	var p OnValidateStartInfo
	p.Req = req
	res := t.onValidate(p)
	return func(resp InfoResponse, e error) {
		var p OnValidateDoneInfo
		p.Resp = resp
		p.Error = e
		res(p)
	}
}

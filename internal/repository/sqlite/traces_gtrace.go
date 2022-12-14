// Code generated by gtrace. DO NOT EDIT.

package sqlite

import (
	"database/sql"
)

// Compose returns a new Trace which has functional fields composed
// both from t and x.
func (t Trace) Compose(x Trace) (ret Trace) {
	switch {
	case t.OnQuery == nil:
		ret.OnQuery = x.OnQuery
	case x.OnQuery == nil:
		ret.OnQuery = t.OnQuery
	default:
		h1 := t.OnQuery
		h2 := x.OnQuery
		ret.OnQuery = func(o OnQueryStartInfo) func(OnQueryDoneInfo) {
			r1 := h1(o)
			r2 := h2(o)
			switch {
			case r1 == nil:
				return r2
			case r2 == nil:
				return r1
			default:
				return func(o OnQueryDoneInfo) {
					r1(o)
					r2(o)
				}
			}
		}
	}
	switch {
	case t.OnExec == nil:
		ret.OnExec = x.OnExec
	case x.OnExec == nil:
		ret.OnExec = t.OnExec
	default:
		h1 := t.OnExec
		h2 := x.OnExec
		ret.OnExec = func(o OnExecStartInfo) func(OnExecDoneInfo) {
			r1 := h1(o)
			r2 := h2(o)
			switch {
			case r1 == nil:
				return r2
			case r2 == nil:
				return r1
			default:
				return func(o OnExecDoneInfo) {
					r1(o)
					r2(o)
				}
			}
		}
	}
	switch {
	case t.OnConnect == nil:
		ret.OnConnect = x.OnConnect
	case x.OnConnect == nil:
		ret.OnConnect = t.OnConnect
	default:
		h1 := t.OnConnect
		h2 := x.OnConnect
		ret.OnConnect = func(o OnConnectStartInfo) func(OnConnectDoneInfo) {
			r1 := h1(o)
			r2 := h2(o)
			switch {
			case r1 == nil:
				return r2
			case r2 == nil:
				return r1
			default:
				return func(o OnConnectDoneInfo) {
					r1(o)
					r2(o)
				}
			}
		}
	}
	switch {
	case t.OnTxBegin == nil:
		ret.OnTxBegin = x.OnTxBegin
	case x.OnTxBegin == nil:
		ret.OnTxBegin = t.OnTxBegin
	default:
		h1 := t.OnTxBegin
		h2 := x.OnTxBegin
		ret.OnTxBegin = func(o OnTxBeginStartInfo) func(OnTxBeginDoneInfo) {
			r1 := h1(o)
			r2 := h2(o)
			switch {
			case r1 == nil:
				return r2
			case r2 == nil:
				return r1
			default:
				return func(o OnTxBeginDoneInfo) {
					r1(o)
					r2(o)
				}
			}
		}
	}
	switch {
	case t.OnTxEnd == nil:
		ret.OnTxEnd = x.OnTxEnd
	case x.OnTxEnd == nil:
		ret.OnTxEnd = t.OnTxEnd
	default:
		h1 := t.OnTxEnd
		h2 := x.OnTxEnd
		ret.OnTxEnd = func(o OnTxEndStartInfo) func(OnTxEndDoneInfo) {
			r1 := h1(o)
			r2 := h2(o)
			switch {
			case r1 == nil:
				return r2
			case r2 == nil:
				return r1
			default:
				return func(o OnTxEndDoneInfo) {
					r1(o)
					r2(o)
				}
			}
		}
	}
	return ret
}
func (t Trace) onQuery(o OnQueryStartInfo) func(OnQueryDoneInfo) {
	fn := t.OnQuery
	if fn == nil {
		return func(OnQueryDoneInfo) {
			return
		}
	}
	res := fn(o)
	if res == nil {
		return func(OnQueryDoneInfo) {
			return
		}
	}
	return res
}
func (t Trace) onExec(o OnExecStartInfo) func(OnExecDoneInfo) {
	fn := t.OnExec
	if fn == nil {
		return func(OnExecDoneInfo) {
			return
		}
	}
	res := fn(o)
	if res == nil {
		return func(OnExecDoneInfo) {
			return
		}
	}
	return res
}
func (t Trace) onConnect(o OnConnectStartInfo) func(OnConnectDoneInfo) {
	fn := t.OnConnect
	if fn == nil {
		return func(OnConnectDoneInfo) {
			return
		}
	}
	res := fn(o)
	if res == nil {
		return func(OnConnectDoneInfo) {
			return
		}
	}
	return res
}
func (t Trace) onTxBegin(o OnTxBeginStartInfo) func(OnTxBeginDoneInfo) {
	fn := t.OnTxBegin
	if fn == nil {
		return func(OnTxBeginDoneInfo) {
			return
		}
	}
	res := fn(o)
	if res == nil {
		return func(OnTxBeginDoneInfo) {
			return
		}
	}
	return res
}
func (t Trace) onTxEnd(o OnTxEndStartInfo) func(OnTxEndDoneInfo) {
	fn := t.OnTxEnd
	if fn == nil {
		return func(OnTxEndDoneInfo) {
			return
		}
	}
	res := fn(o)
	if res == nil {
		return func(OnTxEndDoneInfo) {
			return
		}
	}
	return res
}
func traceOnQuery(t Trace, method string, query string, args []interface{}) func(error) {
	var p OnQueryStartInfo
	p.Method = method
	p.Query = query
	p.Args = args
	res := t.onQuery(p)
	return func(e error) {
		var p OnQueryDoneInfo
		p.Error = e
		res(p)
	}
}
func traceOnExec(t Trace, query string, args []interface{}) func(sql.Result, error) {
	var p OnExecStartInfo
	p.Query = query
	p.Args = args
	res := t.onExec(p)
	return func(r sql.Result, e error) {
		var p OnExecDoneInfo
		p.Result = r
		p.Error = e
		res(p)
	}
}
func traceOnConnect(t Trace, dBS string) func() {
	var p OnConnectStartInfo
	p.DBS = dBS
	res := t.onConnect(p)
	return func() {
		var p OnConnectDoneInfo
		res(p)
	}
}
func traceOnTxBegin(t Trace) func(error) {
	var p OnTxBeginStartInfo
	res := t.onTxBegin(p)
	return func(e error) {
		var p OnTxBeginDoneInfo
		p.Error = e
		res(p)
	}
}
func traceOnTxEnd(t Trace, commit bool) func(error) {
	var p OnTxEndStartInfo
	p.Commit = commit
	res := t.onTxEnd(p)
	return func(e error) {
		var p OnTxEndDoneInfo
		p.Error = e
		res(p)
	}
}

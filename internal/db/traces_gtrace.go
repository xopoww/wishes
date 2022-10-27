// Code generated by gtrace. DO NOT EDIT.

package db

// Compose returns a new Trace which has functional fields composed
// both from t and x.
func (t Trace) Compose(x Trace) (ret Trace) {
	switch {
	case t.OnCheckUser == nil:
		ret.OnCheckUser = x.OnCheckUser
	case x.OnCheckUser == nil:
		ret.OnCheckUser = t.OnCheckUser
	default:
		h1 := t.OnCheckUser
		h2 := x.OnCheckUser
		ret.OnCheckUser = func(o OnCheckUserStartInfo) func(OnCheckUserDoneInfo) {
			r1 := h1(o)
			r2 := h2(o)
			switch {
			case r1 == nil:
				return r2
			case r2 == nil:
				return r1
			default:
				return func(o OnCheckUserDoneInfo) {
					r1(o)
					r2(o)
				}
			}
		}
	}
	return ret
}
func (t Trace) onCheckUser(o OnCheckUserStartInfo) func(OnCheckUserDoneInfo) {
	fn := t.OnCheckUser
	if fn == nil {
		return func(OnCheckUserDoneInfo) {
			return
		}
	}
	res := fn(o)
	if res == nil {
		return func(OnCheckUserDoneInfo) {
			return
		}
	}
	return res
}
func traceOnCheckUser(t Trace, username string) func(error) {
	var p OnCheckUserStartInfo
	p.Username = username
	res := t.onCheckUser(p)
	return func(e error) {
		var p OnCheckUserDoneInfo
		p.Error = e
		res(p)
	}
}

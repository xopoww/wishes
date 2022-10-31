package log

import (
	"github.com/rs/zerolog"
	"github.com/xopoww/wishes/internal/handlers"
)

func Handlers(l zerolog.Logger) (t handlers.Trace) {
	t.OnLogin = func(si handlers.OnLoginStartInfo) func(handlers.OnLoginDoneInfo) {
		return func(di handlers.OnLoginDoneInfo) {
			if di.Error != nil {
				l.Error().
					Str("username", si.Username).
					Err(di.Error).
					Msg("login error")
			} else {
				l.Debug().
					Str("username", si.Username).
					Bool("ok", di.Ok).
					Msg("new login")
			}
		}
	}

	t.OnGetUser = func(si handlers.OnGetUserStartInfo) func(handlers.OnGetUserDoneInfo) {
		return func(di handlers.OnGetUserDoneInfo) {
			if di.Error != nil {
				l.Error().
					Int("user_id", si.UserID).
					Str("principal", string(*si.Principal)).
					Err(di.Error).
					Msg("get user error")
			} else {
				l.Debug().
					Int("user_id", si.UserID).
					Str("principal", string(*si.Principal)).
					Msg("get user done")
			}
		}
	}

	t.OnPatchUser = func(si handlers.OnPatchUserStartInfo) func(handlers.OnPatchUserDoneInfo) {
		return func(di handlers.OnPatchUserDoneInfo) {
			if di.Error != nil {
				l.Error().
					Str("username", string(si.User.Username)).
					Str("principal", string(*si.Principal)).
					Err(di.Error).
					Msg("patch user error")
			} else {
				l.Debug().
					Str("username", string(si.User.Username)).
					Str("principal", string(*si.Principal)).
					Msg("patch user done")
			}
		}
	}

	t.OnPostUser = func(si handlers.OnPostUserStartInfo) func(handlers.OnPostUserDoneInfo) {
		return func(di handlers.OnPostUserDoneInfo) {
			if di.Error != nil {
				l.Error().
					Str("username", si.Username).
					Err(di.Error).
					Msg("post user error")
			} else {
				l.Debug().
					Str("username", si.Username).
					Bool("ok", di.Ok).
					Msg("new user registration")
			}
		}
	}

	t.OnKeySecurityAuth = func(si handlers.OnKeySecurityAuthStartInfo) func(handlers.OnKeySecurityAuthDoneInfo) {
		return func(di handlers.OnKeySecurityAuthDoneInfo) {
			if di.Err != nil {
				l.Warn().AnErr("validate-error", di.Err).Msg("invalid key auth")
			} else {
				l.Debug().Str("principal", string(*di.Principal)).Msg("new key auth")
			}
		}
	}

	return t
}

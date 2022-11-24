package log

import (
	"github.com/rs/zerolog"
	"github.com/xopoww/wishes/internal/oauth/yandex"
)

func YandexOAuth(l zerolog.Logger) (t yandex.Trace) {
	l = l.With().
		Str(zerolog.CallerFieldName, "oauth").
		Str("provider", "yandex").
		Logger()

	t.OnValidate = func(si yandex.OnValidateStartInfo) func(yandex.OnValidateDoneInfo) {
		l.Trace().
			Stringer("url", si.Req.URL).
			Str("auth", si.Req.Header.Get("Authorization")).
			Msg("sending http request")
		return func(di yandex.OnValidateDoneInfo) {
			if di.Error != nil {
				l.Error().
					Err(di.Error).
					Msg("validate error")
			} else {
				l.Debug().
					Dict("info", zerolog.Dict().
						Str("login", di.Resp.Login).
						Str("id", di.Resp.ID),
					).
					Msg("validate done")
			}
		}
	}

	return t
}

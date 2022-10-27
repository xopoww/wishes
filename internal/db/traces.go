package db

//go:generate gtrace

type (
	//gtrace:gen
	//gtrace:set shortcut
	Trace struct {
		OnCheckUser func(OnCheckUserStartInfo) func(OnCheckUserDoneInfo)
	}

	OnCheckUserStartInfo struct {
		Username string
	}
	OnCheckUserDoneInfo  struct {
		Error error
	}
)
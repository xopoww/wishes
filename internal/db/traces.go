package db

import "database/sql"

//go:generate gtrace

type (
	//gtrace:gen
	//gtrace:set shortcut
	Trace struct {
		OnQuery func(OnQueryStartInfo) func(OnQueryDoneInfo)
		OnExec  func(OnExecStartInfo) func(OnExecDoneInfo)

		OnConnect   func(OnConnectStartInfo) func(OnConnectDoneInfo)
		OnCheckUser func(OnCheckUserStartInfo) func(OnCheckUserDoneInfo)
	}
	OnQueryStartInfo struct {
		Method string
		Query  string
		Args   []interface{}
	}
	OnQueryDoneInfo struct {
		Error error
	}

	OnExecStartInfo struct {
		Query string
		Args  []interface{}
	}
	OnExecDoneInfo struct {
		Result sql.Result
		Error  error
	}

	OnCheckUserStartInfo struct {
		Username string
	}
	OnCheckUserDoneInfo struct{}

	OnConnectStartInfo struct {
		DBS string
	}
	OnConnectDoneInfo struct{}
)

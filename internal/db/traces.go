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

	OnConnectStartInfo struct {
		DBS string
	}
	OnConnectDoneInfo struct{}
)

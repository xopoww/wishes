package sqlite

import "database/sql"

//go:generate gtrace

type (
	//gtrace:gen
	//gtrace:set shortcut
	Trace struct {
		OnQuery func(OnQueryStartInfo) func(OnQueryDoneInfo)
		OnExec  func(OnExecStartInfo) func(OnExecDoneInfo)

		OnConnect func(OnConnectStartInfo) func(OnConnectDoneInfo)

		OnTxBegin func(OnTxBeginStartInfo) func(OnTxBeginDoneInfo)
		OnTxEnd   func(OnTxEndStartInfo) func(OnTxEndDoneInfo)
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

	OnTxBeginStartInfo struct{}
	OnTxBeginDoneInfo  struct {
		Error error
	}

	OnTxEndStartInfo struct {
		Commit bool
	}
	OnTxEndDoneInfo struct {
		Error error
	}
)

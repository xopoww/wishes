package handlers

//go:generate gtrace

//gtrace:gen
//gtrace:set shortcut
type Trace struct {
	OnLogin func(OnLoginStartInfo) func(OnLoginDoneInfo)

	OnGetUser   func(OnGetUserStartInfo) func(OnGetUserDoneInfo)
	OnPatchUser func(OnPatchUserStartInfo) func(OnPatchUserDoneInfo)
	OnRegister  func(OnRegisterStartInfo) func(OnRegisterDoneInfo)

	OnKeySecurityAuth func(OnKeySecurityAuthStartInfo) func(OnKeySecurityAuthDoneInfo)
}

package handlers

//go:generate gtrace

//gtrace:gen
//gtrace:set shortcut
type Trace struct {
	OnLogin func(OnLoginStartInfo) func(OnLoginDoneInfo)

	OnGetUser   func(OnGetUserStartInfo) func(OnGetUserDoneInfo)
	OnPatchUser func(OnPatchUserStartInfo) func(OnPatchUserDoneInfo)
	OnPostUser func(OnPostUserStartInfo) func(OnPostUserDoneInfo)

	OnKeySecurityAuth func(OnKeySecurityAuthStartInfo) func(OnKeySecurityAuthDoneInfo)
}

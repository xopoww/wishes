package handlers

//go:generate gtrace

//gtrace:gen
//gtrace:set shortcut
type Trace struct {
	OnLogin func(OnLoginStartInfo) func(OnLoginDoneInfo)

	OnGetUser   func(OnGetUserStartInfo) func(OnGetUserDoneInfo)
	OnPatchUser func(OnPatchUserStartInfo) func(OnPatchUserDoneInfo)
	OnRegister  func(OnRegisterStartInfo) func(OnRegisterDoneInfo)

	OnGetList func(OnGetListStartInfo) func(OnGetListDoneInfo)
	OnPostList func(OnPostListStartInfo) func(OnPostListDoneInfo)
	OnPatchList func(OnPatchListStartInfo) func(OnPatchListDoneInfo)
	OnDeleteList func(OnDeleteListStartInfo) func(OnDeleteListDoneInfo)
	OnGetUserLists func(OnGetUserListsStartInfo) func(OnGetUserListsDoneInfo)

	OnKeySecurityAuth func(OnKeySecurityAuthStartInfo) func(OnKeySecurityAuthDoneInfo)
}

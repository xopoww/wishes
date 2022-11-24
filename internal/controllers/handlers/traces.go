package handlers

//go:generate gtrace

//gtrace:gen
//gtrace:set shortcut
type Trace struct {
	OnLogin func(OnLoginStartInfo) func(OnLoginDoneInfo)

	OnGetUser   func(OnGetUserStartInfo) func(OnGetUserDoneInfo)
	OnPatchUser func(OnPatchUserStartInfo) func(OnPatchUserDoneInfo)
	OnRegister  func(OnRegisterStartInfo) func(OnRegisterDoneInfo)

	OnGetList         func(OnGetListStartInfo) func(OnGetListDoneInfo)
	OnGetListItems    func(OnGetListItemsStartInfo) func(OnGetListItemsDoneInfo)
	OnPostList        func(OnPostListStartInfo) func(OnPostListDoneInfo)
	OnPostListItems   func(OnPostListItemsStartInfo) func(OnPostListItemsDoneInfo)
	OnPatchList       func(OnPatchListStartInfo) func(OnPatchListDoneInfo)
	OnDeleteList      func(OnDeleteListStartInfo) func(OnDeleteListDoneInfo)
	OnDeleteListItems func(OnDeleteListItemsStartInfo) func(OnDeleteListItemsDoneInfo)
	OnGetUserLists    func(OnGetUserListsStartInfo) func(OnGetUserListsDoneInfo)
	OnGetListToken    func(OnGetListTokenStartInfo) func(OnGetListTokenDoneInfo)
	OnPostItemTaken   func(OnPostItemTakenStartInfo) func(OnPostItemTakenDoneInfo)
	OnDeleteItemTaken func(OnDeleteItemTakenStartInfo) func(OnDeleteItemTakenDoneInfo)

	OnKeySecurityAuth func(OnKeySecurityAuthStartInfo) func(OnKeySecurityAuthDoneInfo)

	OnOAuthRegister	  func(OnOAuthRegisterStartInfo) func(OnOAuthRegisterDoneInfo)
	OnOAuthLogin	  func(OnOAuthLoginStartInfo) func(OnOAuthLoginDoneInfo)
}

package log

import (
	"github.com/rs/zerolog"
	"github.com/xopoww/wishes/internal/controllers/handlers"
)

func Handlers(l zerolog.Logger) (t handlers.Trace) {
	l = l.With().Str(zerolog.CallerFieldName, "handlers").Logger()

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
					Int64("user_id", si.UserID).
					Dict("client", zerolog.Dict().
						Str("name", si.Client.Name).
						Int64("id", si.Client.ID),
					).
					Err(di.Error).
					Msg("get user error")
			} else {
				l.Debug().
					Int64("user_id", si.UserID).
					Dict("client", zerolog.Dict().
						Str("name", si.Client.Name).
						Int64("id", si.Client.ID),
					).
					Msg("get user done")
			}
		}
	}

	t.OnPatchUser = func(si handlers.OnPatchUserStartInfo) func(handlers.OnPatchUserDoneInfo) {
		return func(di handlers.OnPatchUserDoneInfo) {
			if di.Error != nil {
				l.Error().
					Int64("id", si.User.ID).
					Dict("client", zerolog.Dict().
						Str("name", si.Client.Name).
						Int64("id", si.Client.ID),
					).
					Err(di.Error).
					Msg("patch user error")
			} else {
				l.Debug().
					Int64("id", si.User.ID).
					Dict("client", zerolog.Dict().
						Str("name", si.Client.Name).
						Int64("id", si.Client.ID),
					).
					Msg("patch user done")
			}
		}
	}

	t.OnRegister = func(si handlers.OnRegisterStartInfo) func(handlers.OnRegisterDoneInfo) {
		return func(di handlers.OnRegisterDoneInfo) {
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
				l.Warn().
					AnErr("validate-error", di.Err).
					Msg("invalid key auth")
			} else {
				l.Debug().
					Dict("client", zerolog.Dict().
						Str("name", di.Client.Name).
						Int64("id", di.Client.ID),
					).
					Msg("new key auth")
			}
		}
	}

	t.OnGetList = func(si handlers.OnGetListStartInfo) func(handlers.OnGetListDoneInfo) {
		return func(di handlers.OnGetListDoneInfo) {
			if di.Error != nil {
				l.Error().
					Int64("list_id", si.ListID).
					Dict("client", zerolog.Dict().
						Str("name", si.Client.Name).
						Int64("id", si.Client.ID),
					).
					Bool("token", si.Token != nil).
					Err(di.Error).
					Msg("get list error")
			} else {
				l.Debug().
					Int64("list_id", si.ListID).
					Dict("client", zerolog.Dict().
						Str("name", si.Client.Name).
						Int64("id", si.Client.ID),
					).
					Bool("token", si.Token != nil).
					Msg("get list done")
			}
		}
	}
	t.OnGetListItems = func(si handlers.OnGetListItemsStartInfo) func(handlers.OnGetListItemsDoneInfo) {
		return func(di handlers.OnGetListItemsDoneInfo) {
			if di.Error != nil {
				l.Error().
					Int64("list_id", si.ListID).
					Dict("client", zerolog.Dict().
						Str("name", si.Client.Name).
						Int64("id", si.Client.ID),
					).
					Bool("token", si.Token != nil).
					Err(di.Error).
					Msg("get list items error")
			} else {
				l.Debug().
					Int64("list_id", si.ListID).
					Dict("client", zerolog.Dict().
						Str("name", si.Client.Name).
						Int64("id", si.Client.ID),
					).
					Bool("token", si.Token != nil).
					Msg("get list items done")
			}
		}
	}
	t.OnPostList = func(si handlers.OnPostListStartInfo) func(handlers.OnPostListDoneInfo) {
		return func(di handlers.OnPostListDoneInfo) {
			if di.Error != nil {
				l.Error().
					Dict("client", zerolog.Dict().
						Str("name", si.Client.Name).
						Int64("id", si.Client.ID),
					).
					Err(di.Error).
					Msg("post list error")
			} else {
				l.Debug().
					Dict("client", zerolog.Dict().
						Str("name", si.Client.Name).
						Int64("id", si.Client.ID),
					).
					Int64("list_id", di.ListID).
					Msg("post list done")
			}
		}
	}
	t.OnPatchList = func(si handlers.OnPatchListStartInfo) func(handlers.OnPatchListDoneInfo) {
		return func(di handlers.OnPatchListDoneInfo) {
			if di.Error != nil {
				l.Error().
					Int64("list_id", si.List.ID).
					Dict("client", zerolog.Dict().
						Str("name", si.Client.Name).
						Int64("id", si.Client.ID),
					).
					Err(di.Error).
					Msg("patch list error")
			} else {
				l.Debug().
					Int64("list_id", si.List.ID).
					Dict("client", zerolog.Dict().
						Str("name", si.Client.Name).
						Int64("id", si.Client.ID),
					).
					Msg("patch list done")
			}
		}
	}
	t.OnDeleteList = func(si handlers.OnDeleteListStartInfo) func(handlers.OnDeleteListDoneInfo) {
		return func(di handlers.OnDeleteListDoneInfo) {
			if di.Error != nil {
				l.Error().
					Int64("list_id", si.List.ID).
					Dict("client", zerolog.Dict().
						Str("name", si.Client.Name).
						Int64("id", si.Client.ID),
					).
					Err(di.Error).
					Msg("delete list error")
			} else {
				l.Debug().
					Int64("list_id", si.List.ID).
					Dict("client", zerolog.Dict().
						Str("name", si.Client.Name).
						Int64("id", si.Client.ID),
					).
					Msg("delete list done")
			}
		}
	}
	t.OnGetUserLists = func(si handlers.OnGetUserListsStartInfo) func(handlers.OnGetUserListsDoneInfo) {
		return func(di handlers.OnGetUserListsDoneInfo) {
			if di.Error != nil {
				l.Error().
					Int64("user_id", si.UserID).
					Dict("client", zerolog.Dict().
						Str("name", si.Client.Name).
						Int64("id", si.Client.ID),
					).
					Err(di.Error).
					Msg("get user lists error")
			} else {
				l.Debug().
					Int64("user_id", si.UserID).
					Dict("client", zerolog.Dict().
						Str("name", si.Client.Name).
						Int64("id", si.Client.ID),
					).
					Int("num_lists", len(di.ListIDs)).
					Msg("get user lists done")
			}
		}
	}
	t.OnGetListToken = func(si handlers.OnGetListTokenStartInfo) func(handlers.OnGetListTokenDoneInfo) {
		return func(di handlers.OnGetListTokenDoneInfo) {
			if di.Error != nil {
				l.Error().
					Int64("list_id", si.ListID).
					Dict("client", zerolog.Dict().
						Str("name", si.Client.Name).
						Int64("id", si.Client.ID),
					).
					Err(di.Error).
					Msg("get list token error")
			} else {
				l.Debug().
					Int64("list_id", si.ListID).
					Dict("client", zerolog.Dict().
						Str("name", si.Client.Name).
						Int64("id", si.Client.ID),
					).
					Msg("get list token done")
			}
		}
	}
	t.OnPostListItems = func(si handlers.OnPostListItemsStartInfo) func(handlers.OnPostListItemsDoneInfo) {
		return func(di handlers.OnPostListItemsDoneInfo) {
			if di.Error != nil {
				l.Error().
					Int64("list_id", si.List.ID).
					Int64("revision", si.List.RevisionID).
					Int("num_items", len(si.Items)).
					Dict("client", zerolog.Dict().
						Str("name", si.Client.Name).
						Int64("id", si.Client.ID),
					).
					Err(di.Error).
					Msg("post list items error")
			} else {
				l.Debug().
					Int64("list_id", si.List.ID).
					Int64("revision", si.List.RevisionID).
					Int("num_items", len(si.Items)).
					Dict("client", zerolog.Dict().
						Str("name", si.Client.Name).
						Int64("id", si.Client.ID),
					).
					Msg("post list items done")
			}
		}
	}
	t.OnDeleteListItems = func(si handlers.OnDeleteListItemsStartInfo) func(handlers.OnDeleteListItemsDoneInfo) {
		return func(di handlers.OnDeleteListItemsDoneInfo) {
			if di.Error != nil {
				l.Error().
					Int64("list_id", si.List.ID).
					Int64("revision", si.List.RevisionID).
					Int("num_items", len(si.ItemIDs)).
					Dict("client", zerolog.Dict().
						Str("name", si.Client.Name).
						Int64("id", si.Client.ID),
					).
					Err(di.Error).
					Msg("delete list items error")
			} else {
				l.Debug().
					Int64("list_id", si.List.ID).
					Int64("revision", si.List.RevisionID).
					Int("num_items", len(si.ItemIDs)).
					Dict("client", zerolog.Dict().
						Str("name", si.Client.Name).
						Int64("id", si.Client.ID),
					).
					Msg("delete list items done")
			}
		}
	}
	t.OnPostItemTaken = func(si handlers.OnPostItemTakenStartInfo) func(handlers.OnPostItemTakenDoneInfo) {
		return func(di handlers.OnPostItemTakenDoneInfo) {
			if di.Error != nil {
				l.Error().
					Int64("list_id", si.List.ID).
					Int64("revision", si.List.RevisionID).
					Int64("item_id", si.ItemID).
					Dict("client", zerolog.Dict().
						Str("name", si.Client.Name).
						Int64("id", si.Client.ID),
					).
					Bool("token", si.Token != nil).
					Err(di.Error).
					Msg("post item taken error")
			} else {
				l.Debug().
					Int64("list_id", si.List.ID).
					Int64("revision", si.List.RevisionID).
					Int64("item_id", si.ItemID).
					Dict("client", zerolog.Dict().
						Str("name", si.Client.Name).
						Int64("id", si.Client.ID),
					).
					Bool("token", si.Token != nil).
					Msg("post item taken done")
			}
		}
	}
	t.OnDeleteItemTaken = func(si handlers.OnDeleteItemTakenStartInfo) func(handlers.OnDeleteItemTakenDoneInfo) {
		return func(di handlers.OnDeleteItemTakenDoneInfo) {
			if di.Error != nil {
				l.Error().
					Int64("list_id", si.List.ID).
					Int64("revision", si.List.RevisionID).
					Int64("item_id", si.ItemID).
					Dict("client", zerolog.Dict().
						Str("name", si.Client.Name).
						Int64("id", si.Client.ID),
					).
					Bool("token", si.Token != nil).
					Err(di.Error).
					Msg("delete item taken error")
			} else {
				l.Debug().
					Int64("list_id", si.List.ID).
					Int64("revision", si.List.RevisionID).
					Int64("item_id", si.ItemID).
					Dict("client", zerolog.Dict().
						Str("name", si.Client.Name).
						Int64("id", si.Client.ID),
					).
					Bool("token", si.Token != nil).
					Msg("delete item taken done")
			}
		}
	}

	return t
}

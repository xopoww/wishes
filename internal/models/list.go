package models

import "fmt"

type ListItem struct {
	ID 	  int64
	Title string
	Desc  string
}

type List struct {
	ID         int64
	OwnerID    int64
	Access     ListAccess
	Title      string
	RevisionID int64
	Items      []ListItem
}

type ListAccess int

const (
	PublicAccess ListAccess = iota
	LinkAccess
	PrivateAccess
)

func (a ListAccess) String() string {
	switch a {
	case PublicAccess:
		return "public"
	case LinkAccess:
		return "link"
	case PrivateAccess:
		return "private"
	default:
		return fmt.Sprintf("unknown ListAccess: %d", a)
	}
}

func ListAccessFromString(s string) ListAccess {
	switch s {
	case "public":
		return PublicAccess
	case "link":
		return LinkAccess
	case "private":
		return PrivateAccess
	default:
		panic(fmt.Sprintf("unknown ListAccess string: %s", s))
	}
}

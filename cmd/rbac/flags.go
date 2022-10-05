package rbac

const (
	Subject    = "subject"
	Permission = "permission"
	Resource   = "resource"
	Scope      = "scope"

	SubjectUsage = "The handle of the subject of the permissions. If you're using LDAP, please use the usernameKey configured in StackState"

	ScopeDefault = "id = '-1'"
)

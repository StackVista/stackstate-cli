package rbac

const (
	Subject    = "subject"
	Permission = "permission"
	Resource   = "resource"
	Scope      = "scope"

	SubjectUsage            = "The handle of the subject of the permissions. If you're using LDAP, please use the usernameKey configured in StackState"
	PermissionRevokeUsage   = "The permission to revoke"
	PermissionGrantUsage    = "The permission to grant"
	PermissionDescribeUsage = "Filter the permissions by permission name"
	ResourceDescribeUsage   = "Filter the permissions by a resource identifier (e.g. system or a view name)"
	ResourceGrantUsage      = "The resource to grant the permission to (e.g. \"system\" or a view name)"
	ResourceRevokeUsage     = "The resource to revoke the permission to (e.g. \"system\" or a view name)"
	ScopeUsage              = "The query in STQL that will be prepended to every topology element retrieved in StackState. " +
		"For example, if your scope is \"label = 'A'\", then all STQL executed in StackState" +
		" (e.g. Retrieving topology) will only return elements that have the label A"

	DefaultResource = "system"
)

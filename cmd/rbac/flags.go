package rbac

const (
	Subject    = "subject"
	Permission = "permission"
	Resource   = "resource"

	SubjectUsage            = "The handle of the subject of the permissions. If you're using LDAP, please use the usernameKey configured in StackState"
	PermissionRevokeUsage   = "The permission to revoke"
	PermissionGrantUsage    = "The permission to grant"
	PermissionDescribeUsage = "Filter the permissions by permission name"
	ResourceDescribeUsage   = "Filter the permissions by a resource identifier (e.g. system, a resource name or tag)"
	ResourceGrantUsage      = "The resource to grant the permission to (e.g. \"system\" or a resource name or tag)"
	ResourceRevokeUsage     = "The resource to revoke the permission to (e.g. \"system\" or a resource name or tag)"

	DefaultResource = "system"
)

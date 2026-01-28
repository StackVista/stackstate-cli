package rbac

import (
	"fmt"
	"strings"
)

const (
	Subject    = "subject"
	Source     = "source"
	Permission = "permission"
	Resource   = "resource"
	Scope      = "scope"
	CreateOnly = "create-only"

	SubjectUsage            = "The handle of the subject of the permissions. If you're using LDAP, please use the usernameKey configured in SUSE Observability"
	PermissionRevokeUsage   = "The permission to revoke"
	PermissionGrantUsage    = "The permission to grant"
	PermissionDescribeUsage = "Filter the permissions by permission name"
	ResourceDescribeUsage   = "Filter the permissions by a resource identifier (e.g. system or a view name)"
	ResourceGrantUsage      = "The resource to grant the permission to (e.g. \"system\" or a view name)"
	ResourceRevokeUsage     = "The resource to revoke the permission to (e.g. \"system\" or a view name)"
	ScopeUsage              = "The query in STQL that will be prepended to every topology element retrieved in SUSE Observability. " +
		"For example, if your scope is \"label = 'A'\", then all STQL executed in SUSE Observability" +
		" (e.g. Retrieving topology) will only return elements that have the label A"
	CreateOnlyUsage = "Fail when a subject with the same name already exists"

	DefaultResource = "system"
)

var (
	SourceChoices = []string{"static", "observability", "kubernetes"}
	SourceUsage   = "Get the subject from a specific authorization source" + fmt.Sprintf(" (must be { %s })", strings.Join(SourceChoices, " | "))
)

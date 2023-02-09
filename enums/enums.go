package enums

// ENVIRONMENT run environment
type ENVIRONMENT string

const (
	// PRODUCTION production environment
	PRODUCTION = ENVIRONMENT("PRODUCTION")
	// DEVELOP development environment
	DEVELOP = ENVIRONMENT("DEVELOP")
	// TEST test environment
	TEST = ENVIRONMENT("TEST")
)

const (
	// MONGO mongo as db
	MONGO = "MONGO"
)

// STATUS status update action
type STATUS string

const (
	// ACTIVE user status for active user
	ACTIVE = STATUS("active")
	// INACTIVE user status for inactive user
	INACTIVE = STATUS("inactive")
	// DELETED user status for deleted user
	DELETED = STATUS("deleted")
)

// AUTH_TYPE AuthType update action
type AUTH_TYPE string

const (
	// PASSWORD grand_type of users authentication
	PASSWORD      = AUTH_TYPE("password")
	AUTH_CODE     = AUTH_TYPE("authorization_code")
	REFRESH_TOKEN = AUTH_TYPE("refresh_token")
	SSO           = AUTH_TYPE("sso")
)

// TOKEN_TYPE token type of user
type TOKEN_TYPE string

const (
	// REGULAR_TOKEN refers to limited lifetime token and refresh token
	REGULAR_TOKEN = TOKEN_TYPE("regular")
	// CTL_TOKEN refers to long lifetime token and refresh token
	CTL_TOKEN = TOKEN_TYPE("ctl")
)

// ROLE role string
type ROLE string

const (
	// SUPER_ADMIN refers to super admin role
	SUPER_ADMIN = ROLE("SUPER_ADMIN")
	// ADMIN refers to admin role
	ADMIN = ROLE("ADMIN")
	// VIEWER refers to user role
	VIEWER = ROLE("VIEWER")
)

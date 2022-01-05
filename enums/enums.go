package enums

// ENVIRONMENT run environment
type ENVIRONMENT string

const (
	// PRODUCTION mongo as db
	PRODUCTION = ENVIRONMENT("PRODUCTION")
	// INMEMORY in memory storage as db
	DEV  = ENVIRONMENT("DEV")
	TEST = ENVIRONMENT("TEST")
)

const (
	// MONGO mongo as db
	MONGO = "MONGO"
	// INMEMORY in memory storage as db
	INMEMORY = "INMEMORY"
)

type PERMISION_TYPE string

const (
	PERMISSION_READ   = PERMISION_TYPE("read")
	PERMISSION_UPDATE = PERMISION_TYPE("update")
	PERMISSION_DELETE = PERMISION_TYPE("delete")
	PERMISSION_CREATE = PERMISION_TYPE("create")
)

type ROLE_UPDATE_OPTION string

const (
	APPEND_PERMISSION = ROLE_UPDATE_OPTION("append")
	REMOVE_PERMISSION = ROLE_UPDATE_OPTION("remove")
)

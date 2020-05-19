package constant

const (
	// Environment
	ConfigDirectoryEnvVar  = "CJ_EXP_CONFIG_PATH"
	DefaultConfigDirectory = "/etc/cjexp"

	// Cache
	CachePrefix         = "cjexp-%s"
	CachePrefixModified = "cjexp-%s-modified"
	CacheFileFolderName = "cjexp"

	// Session
	SessionCookie      = "cj-exp-session"
	SessionCachePrefix = "cj-exp-session-%s-%s"
	FlashBagSession    = "cj-exp-flashBag"
)

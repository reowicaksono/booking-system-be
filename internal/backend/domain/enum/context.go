package enum

type ContextKey string

var (
	GormCtxKey   ContextKey = "gorm.ctx.key"
	ConfigCtxKey ContextKey = "config.ctx.key"
	LoggerCtxKey ContextKey = "logger.ctx.key"
)

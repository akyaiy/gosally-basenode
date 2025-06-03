package sessions

type Session struct {
	Description string
	TTL         int64
	SessionID   string
}

type SessionBuilderContract interface {
	BuildSession() *Builder
	WithDescription(description string) *Builder
	WithTTL(ttl int64) *Builder
	WithSessionID(sessionID string) *Builder
	EndSafeBuild() (*Session, error)
	EndBuild() *Session
}

type Builder struct {
	session Session
}

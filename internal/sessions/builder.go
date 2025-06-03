package sessions

func NewSessionBuilder() *Builder {
	return &Builder{
		session: Session{},
	}
}

func (b *Builder) FromSession(session *Session) *Builder {
	b.session = *session
	return b
}

func (b *Builder) WithDescription(description string) *Builder {
	b.session.Description = description
	return b
}

func (b *Builder) WithTTL(ttl int64) *Builder {
	b.session.TTL = ttl
	return b
}

func (b *Builder) WithSessionID(sessionID string) *Builder {
	b.session.SessionID = sessionID
	return b
}

func (b *Builder) EndSafeBuild() (*Session, error) {
	if b.session.SessionID == "" {
		return nil, ErrSessionIDRequired
	}
	if b.session.TTL <= 0 {
		return nil, ErrInvalidTTL
	}
	return &b.session, nil
}

func (b *Builder) EndBuild() *Session {
	if b.session.SessionID == "" {
		panic(ErrSessionIDRequired)
	}
	if b.session.TTL <= 0 {
		panic(ErrInvalidTTL)
	}
	return &b.session
}

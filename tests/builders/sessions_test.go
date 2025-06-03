package builder_test

import (
	"reflect"
	"testing"

	"github.com/akyaiy/gosally-basenode/internal/sessions"
)

var (
	Description = "Test session description"
	TTL         = int64(3600)
	SessionID   = "xxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx"

	expectedSession = &sessions.Session{
		Description: Description,
		TTL:         TTL,
		SessionID:   SessionID,
	}
)

func TestSessionsBuilder_Build(t *testing.T) {
	session := sessions.NewSession().
		WithDescription(Description).
		WithTTL(TTL).
		WithSessionID(SessionID).
		EndBuild()
	if !reflect.DeepEqual(session, expectedSession) {
		t.Errorf("Expected session: %+v, got: %+v", expectedSession, session)
	}
}

func TestSessionsBuilder_BuildSafe(t *testing.T) {
	_, err := sessions.NewSession().
		WithDescription(Description).
		WithTTL(TTL).
		EndSafeBuild()
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
}

func TestSessionsBuilder_BuildUnsafe(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected panic, got nil")
		}
	}()
	_ = sessions.NewSession().
		WithDescription(Description).WithDescription(Description).
		WithTTL(TTL).
		EndBuild()
}

package s2tore

import "time"

type (
	Session interface {
		Token() string
		UserID() string
		Expires() time.Time
	}

	SessionStore interface {
		NewSession(string, time.Time) (Session, error)
		ReadSession(string) (Session, bool)
		RemoveSession(string) error
		RemoveExpiredSessions() (int, error)
	}
)

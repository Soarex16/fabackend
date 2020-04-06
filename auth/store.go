package auth

import (
	"fmt"
	"sync"
)

// SessionStore - in-memory storage for users sessions
// third-party key-value libraries (like pudge) are not used, because the planned load on the service is small
// in-memory cache of a limited size, a queue, and a database. When a new session is registered,
//
// TODO: elegant solution for storing many sessions:
//  https://softwareengineering.stackexchange.com/questions/278276/user-sessions-in-a-web-server-speed-or-persistence
// it stays in the cache and is put into the queue. A background thread keeps reading the queue and
// putting the new sessions into the database. (Another background process would periodically expire them.)
type SessionStore struct {
	AccessTokens map[string]*Session
	atM          sync.RWMutex

	RefreshTokens map[string]*Session
	rtM           sync.RWMutex
}

func NewSessionStore() *SessionStore {
	s := &SessionStore{
		AccessTokens:  make(map[string]*Session),
		RefreshTokens: make(map[string]*Session),
	}

	return s
}

// Add - add session to storage
func (s *SessionStore) Add(session *Session) {
	fmt.Println(session.AccessToken)

	s.atM.Lock()
	s.AccessTokens[session.AccessToken] = session
	s.atM.Unlock()

	s.rtM.Lock()
	s.RefreshTokens[session.RefreshToken] = session
	s.rtM.Unlock()
}

// GetByAccessToken - get session by access token
func (s *SessionStore) GetByAccessToken(tok string) (*Session, bool) {
	s.atM.RLock()
	v, ok := s.AccessTokens[tok]
	s.atM.RUnlock()

	if ok {
		return v, true
	}

	return nil, false
}

// GetByRefreshToken - get session by refresh token
func (s *SessionStore) GetByRefreshToken(tok string) (*Session, bool) {
	s.rtM.RLock()
	v, ok := s.RefreshTokens[tok]
	s.rtM.RUnlock()

	if ok {
		return v, true
	}

	return nil, false
}

// RemoveByAccessToken - remove session by access token
func (s *SessionStore) RemoveByAccessToken(tok string) *Session {
	session, ok := s.GetByAccessToken(tok)

	if !ok {
		return nil
	}

	s.atM.Lock()
	delete(s.AccessTokens, session.AccessToken)
	s.atM.Unlock()

	s.rtM.Lock()
	delete(s.RefreshTokens, session.RefreshToken)
	s.rtM.Unlock()

	return session
}

// RemoveByRefreshToken - remove session by refresh token
func (s *SessionStore) RemoveByRefreshToken(tok string) *Session {
	session, ok := s.GetByRefreshToken(tok)

	if !ok {
		return nil
	}

	s.atM.Lock()
	delete(s.AccessTokens, session.AccessToken)
	s.atM.Unlock()

	s.rtM.Lock()
	delete(s.RefreshTokens, session.RefreshToken)
	s.rtM.Unlock()

	return session
}

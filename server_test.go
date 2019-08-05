// The MIT License (MIT)
// Copyright (C) 2019 Georgy Komarov <jubnzv@gmail.com>

package tmux

import (
	"testing"
    "strings"
)

// Kills sessions that contains namePattern substring in the name.
func sessionsReaper(namePattern string) {
    s := new(Server)
    // Suppose that ListSession works.
    sessions, _ := s.ListSessions()
    for _, ss := range sessions {
        if strings.Contains(ss.Name, namePattern) {
            s.KillSession(ss.Name)
        }
    }
}

func TestListSessions(t *testing.T) {
	s := new(Server)
	if _, err := s.ListSessions(); err != nil {
		t.Errorf("ListSessions: %s", err)
	}
}

func TestSessionNames(t *testing.T) {
	s := new(Server)
    _, err := s.NewSession(".");
	if err == nil {
		t.Errorf("Session with restricted name was created")
	}
    _, err = s.NewSession(":");
	if err == nil {
		t.Errorf("Session with restricted name was created")
	}
    _, err = s.NewSession("111:");
	if err == nil {
		t.Errorf("Session with restricted name was created")
	}
}

func TestNewSession(t *testing.T) {
	s := new(Server)
    session, err := s.NewSession("test-new-session");
	if err != nil {
		t.Errorf("NewSession: %s", err)
	}
    defer sessionsReaper("test-new-session")
    sessions, _ := s.ListSessions()

    // Check created session name
    found := false
    for _, isession := range sessions {
        if isession.Name == session.Name {
            found = true
            break
        }
    }
    if found == false {
        t.Errorf("Can't find created session by name: %s", session.Name)
    }

    // Check created session id
    found = false
    for _, isession := range sessions {
        t.Logf("%d -- %d", isession.Id, session.Id)
        if isession.Id == session.Id {
            found = true
            break
        }
    }
    if found == false {
        t.Errorf("Can't find created session by id: %d", session.Id)
    }
}

func TestHasSession(t *testing.T) {
	s := new(Server)
	s.NewSession("test-has-session")
    defer sessionsReaper("test-has-session")
    has, err := s.HasSession("test-has-session")
	if err != nil {
		t.Errorf("HasSession: %s", err)
	}
    if has == false {
        t.Errorf("Can't find created session: 'test-has-session'")
    }
	s.KillSession("test-has-session")
}

func TestKillSession(t *testing.T) {
	s := new(Server)
	s.NewSession("test-kill-session")
	s.KillSession("test-kill-session")
	if has, _ := s.HasSession("test-kill-session"); has == true {
		t.Errorf("KillSession: Can't kill 'test-kill-session' session!")
	}
}
// The MIT License (MIT)
// Copyright (C) 2019 Georgy Komarov <jubnzv@gmail.com>

package tmux

import (
	"testing"
)

func TestWindowListPanes(t *testing.T) {
	s := createSession()
	defer sessionsReaper(s.Name)
    w, _ := s.NewWindow("test-window")
    panes, _ := w.ListPanes()

    for _, p := range panes {
        if p.SessionId != s.Id {
            t.Errorf("Incorrect session id (expected %d got %d)", s.Id, p.SessionId)
        }
        if p.SessionName != s.Name {
            t.Errorf("Incorrect session name (expected %s got %s)", s.Name, p.SessionName)
        }
        if p.WindowId != w.Id {
            t.Errorf("Incorrect window id (expected %d got %d)", w.Id, p.WindowId)
        }
        if p.WindowName != w.Name {
            t.Errorf("Incorrect window name (expected %s got %s)", w.Name, p.WindowName)
        }
    }
}
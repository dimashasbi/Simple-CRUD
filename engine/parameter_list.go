package engine

import (
	"golang.org/x/net/context"
)

func (p *parameter) List(c context.Context) *SimpleConfigResp {
	// q := NewQuery("greeting").Order("date", Descending).Slice(0, r.Count)
	return &SimpleConfigResp{
		// Greetings: g.repository.List(c, q),
	}
}

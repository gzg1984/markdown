package markdown

import "strings"

type MarkDownDoc struct {
	builder *strings.Builder
}

func NewMarkDown() *MarkDownDoc {
	m := new(MarkDownDoc)
	m.builder = new(strings.Builder)
	return m
}

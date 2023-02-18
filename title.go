package markdown

import "strings"

func (m *MarkDownDoc) WriteLevel1Title(content string) *MarkDownDoc {
	m.WriteTitle(content, 1)
	return m
}

func (m *MarkDownDoc) GetTitle(content string, level int) string {
	return strings.Repeat("#", level) + " " + content
}

func (m *MarkDownDoc) WriteTitle(content string, level int) *MarkDownDoc {
	m.write(m.GetTitle(content, level))
	m.Writeln()
	return m
}

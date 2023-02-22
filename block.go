package markdown

import "strings"

func (m *MarkDownDoc) GetBlockPrefix(content string, level int) string {
	return strings.Repeat("	", level) + "-" + content
}

func (m *MarkDownDoc) WriteBlock(content string, level int) *MarkDownDoc {
	m.write(m.GetBlockPrefix(content, level))
	m.Writeln()
	return m
}
func (m *MarkDownDoc) WriteDefaultBlock(content string) *MarkDownDoc {
	m.write(m.GetBlockPrefix(content, 0))
	m.Writeln()
	return m
}

package markdown

import "fmt"

func (m *MarkDownDoc) GetMultiCode(content, t string) string {
	return fmt.Sprintf("``` %s\n%s\n```\n", t, content)
}

func (m *MarkDownDoc) WriteMultiCode(content, t string) *MarkDownDoc {
	m.write(m.GetMultiCode(content, t))
	return m

}

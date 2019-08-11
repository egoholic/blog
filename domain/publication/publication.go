package publication

type Attrs struct {
	MetaKeywords    []string
	MetaDescription string
	Title           string
	Content         string
	CreatedAt       string
}

type Publication struct {
	attrs *Attrs
}

func New(attrs *Attrs) *Publication {
	return &Publication{attrs}
}
func (p *Publication) MetaKeywords() []string {
	return p.attrs.MetaKeywords
}
func (p *Publication) MetaDescription() string {
	return p.attrs.MetaDescription
}
func (p *Publication) Title() string {
	return p.attrs.Title
}
func (p *Publication) Content() string {
	return p.attrs.Content
}

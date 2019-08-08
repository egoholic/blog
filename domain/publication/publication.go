package publication

type Meta struct {
	Keywords    []string
	Description string
}
type Attrs struct {
	Meta
	Title   string
	Content string
}
type Publication struct {
	attrs *Attrs
}

func New(attrs *Attrs) *Publication {
	return &Publication{attrs}
}
func (p *Publication) MetaKeywords() []string {
	return p.attrs.Keywords
}
func (p *Publication) MetaDescription() string {
	return p.attrs.Description
}
func (p *Publication) Title() string {
	return p.attrs.Title
}
func (p *Publication) Content() string {
	return p.attrs.Content
}

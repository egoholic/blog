package publication

type Attrs struct {
	Slug            string
	MetaKeywords    string
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
func (p *Publication) Attrs() *Attrs {
	return p.attrs
}

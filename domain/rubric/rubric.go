package rubric

type Attrs struct {
	Slug            string
	MetaKeywords    string
	MetaDescription string
	Title           string
	Description     string
}
type Rubric struct {
	attrs *Attrs
}

func New(attrs *Attrs) *Rubric {
	return &Rubric{attrs}
}
func (r *Rubric) Attrs() *Attrs {
	return r.attrs
}

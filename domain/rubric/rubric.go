package rubric

type Attrs struct {
	MetaKeywords    []string
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
func (r *Rubric) MetaKeywords() []string {
	return r.attrs.MetaKeywords
}
func (r *Rubric) MetaDescription() string {
	return r.attrs.MetaDescription
}
func (r *Rubric) Title() string {
	return r.attrs.Title
}
func (r *Rubric) Description() string {
	return r.attrs.Description
}

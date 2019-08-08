package rubric

type Meta struct {
	Keywords    []string
	Description string
}
type Attrs struct {
	Meta
	Title   string
	Content string
}
type Rubric struct {
	attrs *Attrs
}

func New(attrs *Attrs) *Rubric {
	return &Rubric{attrs}
}
func (r *Rubric) MetaKeywords() []string {
	return r.attrs.Keywords
}
func (r *Rubric) MetaDescription() string {
	return r.attrs.Description
}
func (r *Rubric) Title() string {
	return r.attrs.Title
}
func (r *Rubric) Content() string {
	return r.attrs.Content
}

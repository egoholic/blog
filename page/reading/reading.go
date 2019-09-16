package reading

type (
	Page struct {
		Slug    string
		Title   string
		Content string
	}
	Value      struct{}
	PageSource interface {
		BySlug(string) *Page
	}
	Destination interface {
		Deliver(*Page) error
	}
	Form interface {
		Slug() string
	}
)

func (v *Value) Deliver(form Form, psource PageSource, destination Destination) {
	slug := form.Slug()
	page := psource.BySlug(slug)
	destination.Deliver(page)
}

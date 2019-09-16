package reading

type (
	Publication struct {
		Slug    string
		Title   string
		Content string
	}
	Value             struct{}
	PublicationSource interface {
		BySlug(string) *Publication
	}
	Destination interface {
		Deliver(*Publication) error
	}
	Form interface {
		Slug() string
	}
)

func (v *Value) Deliver(form Form, psource PublicationSource, destination Destination) {
	slug := form.Slug()
	publication := psource.BySlug(slug)
	destination.Deliver(publication)
}

package previewing

type (
	Author struct {
		Slug     string
		FullName string
		Bio      string
	}
	Publication struct {
		Slug    string
		Title   string
		Content string
	}
	Deliverable struct {
		Author       *Author
		Publications []*Publication
	}
	Value              struct{}
	PublicationsSource interface {
		GetListByAuthorSlug(string) []*Publication
	}
	AuthorSource interface {
		BySlug(string) *Author
	}
	Destination interface {
		Deliver(*Deliverable) error
	}
	Form interface {
		Slug() string
	}
)

func (v *Value) Deliver(form Form, asource AuthorSource, psource PublicationsSource, destination Destination) {
	slug := form.Slug()
	author := asource.BySlug(slug)
	publications := psource.GetListByAuthorSlug(slug)
	destination.Deliver(&Deliverable{
		Author:       author,
		Publications: publications,
	})
}

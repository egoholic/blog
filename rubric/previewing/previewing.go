package previewing

type (
	Rubric struct {
		Slug        string
		Title       string
		Description string
	}
	Publication struct {
		Slug    string
		Title   string
		Content string
	}
	Value        struct{}
	RubricSource interface {
		BySlug(string) *Rubric
	}
	PublicationsSource interface {
		ByRubricSlug(string) []*Publication
	}
	Deliverable struct {
		Rubric       *Rubric
		Publications []*Publication
	}
	Destination interface {
		Deliver(*Deliverable) error
	}
	Form interface {
		Slug() string
	}
)

func (v *Value) Deliver(form Form, rsource RubricSource, psource PublicationsSource, destination Destination) {
	slug := form.Slug()
	rubric := rsource.BySlug(slug)
	publications := psource.ByRubricSlug(slug)
	destination.Deliver(&Deliverable{
		Rubric:       rubric,
		Publications: publications,
	})
}

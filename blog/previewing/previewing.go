package previewing

type (
	Publication struct {
		Slug       string
		Title      string
		Content    string
		CreatedAt  string
		Popularity int
	}
	Value struct {
		PopularProvider PopularPublicationsProvider
		RecentProvider  RecentPublicationsProvider
	}
	RecentPublicationsProvider interface {
		ProvideRecent() []*Publication
	}
	PopularPublicationsProvider interface {
		ProvidePopular() []*Publication
	}
)

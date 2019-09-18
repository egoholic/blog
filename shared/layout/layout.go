package layout

type PageContent struct {
	Value           interface{}
	PageTitle       string
	BlogTitle       string
	MetaKeywords    string
	MetaDescription string
}

func New(v interface{}, pt, bt, md, mk string) *PageContent {
	return &PageContent{
		Value:           v,
		PageTitle:       pt,
		BlogTitle:       bt,
		MetaDescription: md,
		MetaKeywords:    mk,
	}
}

package homepage

import (
	"html/template"
	"io"

	. "github.com/egoholic/blog/view"
)

type Homepage struct {
	Meta
	Header
	Content
	Footer
}

func New(meta Meta, header Header, content Content, footer Footer) *Homepage {
	return &Homepage{
		Meta:    meta,
		Header:  header,
		Content: content,
		Footer:  footer,
	}
}

func (hp *Homepage) Execute(out io.Writer) {
	t := template.Must(template.ParseFiles("view/homepage/homepage.html.tmpl"))
	err := t.Execute(out, hp)
	if err != nil {
		panic(err)
	}
}

# router - stupid simple routing library

## Install

```
dep ensure github.com/egoholic/router
```

```golang
import (
  "github.com/egoholic/router"
  "github.com/egoholic/node"
)
```

## Define routes
```golang
type articleIDForm struct{}

func (f *articleIDForm) CheckAndPopulate(pattern string, chunk string, prms *params.Params) bool {
	num, err := strconv.ParseInt(chunk, 10, 64)
	if err != nil {
		return false
	}
	prms.Set(pattern, num)
	return true
}

func homeHandler(w http.ResponseWriter, r *http.Request, p *params.Params) {
  // some handler logic here
}

rtr := router.New()
rtr.Root().GET(homeHandler, "renders home page")
articles := rtr.Root().Child("articles", &node.DumbForm{})
articles.GET(articlesHandler, "renders articles list")
articles.POST(createArticleHanler, "creates new article")
article := articles.Child(":article_id", &articleIDForm{})
article.GET(articleHandler, "renders particular article")
```

## Getting handlers

### Dumb way
```golang
_prms := map[string]interface{}{}
_prms["someKey"] = "SomeValue"
prms := params.New("/", node.GET, prms)
// returns handler function: type HandlerFn func(w http.ResponseWriter, r *http.Request, p *params.Params)
rtr.Handler(prms)
```

## Using with net/http lib

```golang
server.ListenAndServe(":8080", rtr)
```
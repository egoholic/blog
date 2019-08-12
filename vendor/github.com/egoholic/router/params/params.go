package params

import (
	"strings"
)

type Params struct {
	path       string
	pathChunks []string
	verb       string
	params     map[string]interface{}
}

func New(path, verb string, params map[string]interface{}) *Params {
	var chunks []string
	if path == "/" {
		chunks = []string{""}
	} else {
		chunks = strings.Split(path, "/")
	}
	return &Params{path, chunks, verb, params}
}
func (p *Params) Path() string {
	return p.path
}
func (p *Params) PathChunks() []string {
	return p.pathChunks
}
func (p *Params) Verb() string {
	return p.verb
}
func (p *Params) Get(pname string) (value interface{}, ok bool) {
	value, ok = p.params[pname]
	return value, ok
}
func (p *Params) Set(pname string, val interface{}) {
	p.params[pname] = val
}
func (p *Params) NewIterator() *PathChunksIterator {
	return &PathChunksIterator{p.pathChunks, 0}
}

type PathChunksIterator struct {
	pathChunks []string
	cursor     int
}

func (i *PathChunksIterator) Next() (string, error) {
	i.cursor++
	return i.Current(), nil
}
func (i *PathChunksIterator) HasNext() bool {
	max := len(i.pathChunks) - 1
	return i.cursor < max
}
func (i *PathChunksIterator) Current() string {
	return i.pathChunks[i.cursor]
}

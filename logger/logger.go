package logger

import (
	"encoding/json"
	"io"
	"time"
)

const (
	DURATION        = "duration"
	DEBUG           = "debug"
	ERROR           = "error"
	REQUESTHANDLING = "request_handling"
)

type (
	logger struct {
		writers        map[string]io.Writer
		fallbackWriter io.Writer
	}
	Proxy struct {
		logger *logger
		meta   *Meta
	}
	Meta struct {
		path []string `json:"p"`
		tags []string `json:"t"`
	}
	DurationRecord struct {
		meta     *Meta         `json:"m"`
		OpName   string        `json:"op"`
		Duration time.Duration `json:"dr"`
	}
	DebugRecord struct {
		meta    *Meta  `json:"m"`
		Message string `json:"msg"`
	}
	ErrorRecord struct {
		meta      *Meta    `json:"m"`
		Message   string   `json:"msg"`
		Location  string   `json:"loc"`
		Backtrace []string `json:"bt"`
	}
	RequestHandlingRecord struct {
		meta            *Meta         `json:"m"`
		ID              string        `json:"id"`
		PerformerID     string        `json:"pid"`
		Duration        time.Duration `json:"dr"`
		Path            string        `json:"p"`
		RequestPayload  interface{}   `json:"rq"`
		ResponsePayload interface{}   `json:"rs"`
	}
)

func (p *Proxy) writer(kind string) io.Writer {
	writer, ok := p.logger.writers[kind]
	if !ok {
		return p.logger.fallbackWriter
	}
	return writer
}

func (p *Proxy) Path() []string {
	return p.meta.path
}

func (p *Proxy) Tags() []string {
	return p.meta.tags
}

func (p *Proxy) Debug(r *DebugRecord) {
	r.meta = p.meta
	writer := p.writer(DEBUG)
	m, _ := json.Marshal(r)
	writer.Write(m)
}

func (p *Proxy) Error(r *ErrorRecord) {
	r.meta = p.meta
	writer := p.writer(ERROR)
	m, _ := json.Marshal(r)
	writer.Write(m)
}

func (p *Proxy) Child(key string, tags []string) *Proxy {
	return &Proxy{
		logger: p.logger,
		meta: &Meta{
			path: append(p.Path(), key),
			tags: append(tags, key),
		},
	}
}

func (p *Proxy) RecordDuration(opName string) (finish func()) {
	startTime := time.Now()
	return func() {
		finishTime := time.Now()
		r := &DurationRecord{
			OpName:   opName,
			Duration: finishTime.Sub(startTime),
			meta:     p.meta,
		}
		writer := p.writer(DURATION)
		m, _ := json.Marshal(r)
		writer.Write(m)
	}
}

func (p *Proxy) RecordRequestHandling(id, path, requestName, performerID string, request interface{}) (succed func(response interface{})) {
	startTime := time.Now()
	return func(response interface{}) {
		finishTime := time.Now()
		r := &RequestHandlingRecord{
			meta:            p.meta,
			ID:              id,
			PerformerID:     performerID,
			Path:            path,
			Duration:        finishTime.Sub(startTime),
			RequestPayload:  request,
			ResponsePayload: response,
		}
		writer := p.writer(DURATION)
		m, _ := json.Marshal(r)
		writer.Write(m)
	}
}

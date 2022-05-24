package fluentd

import (
	"fmt"
	"time"

	_log "log"

	"github.com/athoune/chevillette/conf"
	"github.com/athoune/chevillette/log"
	"github.com/athoune/chevillette/memory"
	"github.com/athoune/fluent-server/defaultreader"
	"github.com/athoune/fluent-server/options"
	"github.com/athoune/fluent-server/server"
)

type FluentdInput struct {
	server *server.Server
	tag    string
	line   log.LineReader
	logKey string
	memory memory.Memory
}

func New(tag string, line log.LineReader, memory *memory.Memory, conf *conf.Fluentd) (*FluentdInput, error) {
	f := &FluentdInput{
		tag:    tag,
		line:   line,
		logKey: "log",
		memory: *memory,
	}
	handler := func(tag string, ts *time.Time, record map[string]interface{}) error {
		_log.Println(" log", tag, ts, record)
		if tag == f.tag {
			keys, err := f.line([]byte(record[f.logKey].(string)))
			if err != nil {
				fmt.Println("error", err)
				return nil
			}
			memory.Set(keys, *ts)
			_log.Println("keys", keys)
		}
		return nil
	}
	cfg := &options.FluentOptions{
		SharedKey:             conf.SharedKey,
		Debug:                 true,
		MessagesReaderFactory: defaultreader.DefaultMessagesReaderFactory(handler),
	}
	s, err := server.New(cfg)
	if err != nil {
		return nil, err
	}
	f.server = s
	return f, nil
}

func (f *FluentdInput) ListenAndServe(listen string) error {
	_log.Println("Starting fluentd", listen)
	return f.server.ListenAndServe(listen)
}

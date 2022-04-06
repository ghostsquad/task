package tpl

import (
	"bytes"
	"fmt"
	"text/template"

	"github.com/go-task/task/sdk/v4/lazy"
	"github.com/go-task/task/sdk/v4/pat"
	"github.com/go-task/task/sdk/v4/taskfile"

	sprig "github.com/go-task/slim-sprig"
)

type Templater interface{}

type LazyTemplate struct {
	lazy lazy.Lazy[string]
}

type LazyTemplateOptions struct {
	lazyData lazy.Lazy[taskfile.Vars]
	funcMap  template.FuncMap
}

func WithVars(vars taskfile.Vars) pat.Option[*LazyTemplateOptions] {
	return func(o *LazyTemplateOptions) error {
		for k, v := range vars {
			o.lazyData[k] = v
		}
		return nil
	}
}

func WithFuncMap(funcMap template.FuncMap) pat.Option[*LazyTemplateOptions] {
	return func(o *LazyTemplateOptions) error {
		for k, v := range funcMap {
			o.funcMap[k] = v
		}
		return nil
	}
}

func NewLazyTemplate(text string, options ...pat.Option[*LazyTemplateOptions]) (*LazyTemplate, error) {
	opts := &LazyTemplateOptions{}

	err := pat.ApplyOptions[pat.Option[*LazyTemplateOptions]](opts, options)
	if err != nil {
		return nil, err
	}

	if opts.funcMap == nil || len(opts.funcMap) == 0 {
		opts.funcMap = template.FuncMap(sprig.FuncMap())
	}

	tpl := template.New("")
	tpl.Funcs(opts.funcMap)
	tpl, err = tpl.Parse(text)
	if err != nil {
		return nil, fmt.Errorf("parsing template: %q: %w", text, err)
	}

	lzt := &LazyTemplate{
		lazy: lazy.Lazy[string]{
			New: func() (string, error) {
				buf := bytes.NewBuffer([]byte{})
				err = tpl.ExecuteTemplate(buf, "", opts.data)
				if err != nil {
					return "", err
				}
				return buf.String(), nil
			},
		},
	}

	return lzt, nil
}

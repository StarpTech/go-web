package i18n

import gotext "gopkg.in/leonelquinteros/gotext.v1"

type I18ner interface {
	Get(string, ...interface{}) string
}

type I18n struct{}

func New() *I18n {
	return &I18n{}
}

func Configure(lib, lang, dom string) {
	gotext.Configure(lib, lang, dom)
}

func (i *I18n) Get(str string, vars ...interface{}) string {
	return gotext.Get(str, vars...)
}

func Get(str string, vars ...interface{}) string {
	return gotext.Get(str, vars...)
}

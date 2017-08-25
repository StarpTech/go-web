package i18n

import gotext "gopkg.in/leonelquinteros/gotext.v1"

func Configure(lib, lang, dom string) {
	gotext.Configure(lib, lang, dom)
}

func Get(str string, vars ...interface{}) string {
	return gotext.Get(str, vars...)
}

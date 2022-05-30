package main 

import (
	"net/http"
)

func (m *MyApi) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/user/create":
		h.wrapperDoSomeJob(w, r)
	default:
		// 404
	}
}

func (m *MyApi) wrapperDoSomeJob() {
	// заполнение структуры params
	// валидирование параметров
	res, err := h.DoSomeJob(ctx, params)
	// прочие обработки
}


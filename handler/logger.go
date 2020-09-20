package handler

import (
	"bytes"
	"io/ioutil"
	"net/http"

	"github.com/geraldofigueiredo/mypage/service"
)

type LoggerHandler struct {
	DebugMode bool
}

func (l *LoggerHandler) Logging(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		services := ctx.Value("services").(*service.Services)
		services.Log.Infof("%s %s %s %s", r.RemoteAddr, r.Method, r.URL, r.Proto)
		services.Log.Infof("User agent : %s", r.UserAgent())
		if l.DebugMode {
			body, err := ioutil.ReadAll(r.Body)
			if err != nil {
				services.Log.Errorf("Reading request body error: %s", err)
			}
			reqStr := ioutil.NopCloser(bytes.NewBuffer(body))
			services.Log.Debugf("Request body : %v", reqStr)
			r.Body = reqStr
		}
		h.ServeHTTP(w, r)
	})
}

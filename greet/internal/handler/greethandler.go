package handler

import (
	"net/http"

	"learnZero/greet/internal/logic"
	"learnZero/greet/internal/svc"
	"learnZero/greet/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func GreetHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.Request
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewGreetLogic(r.Context(), svcCtx)
		// 将请求转给对应的func
		resp, err := l.Greet(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}

func TestHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		httpx.OkJson(w, "hello")
	}
}

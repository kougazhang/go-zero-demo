package handler

import (
	"net/http"

	"blog/api/internal/logic"
	"blog/api/internal/svc"
	"blog/api/internal/types"

	"github.com/tal-tech/go-zero/rest/httpx"
)

func UpdateUserHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ReqUpdateUser
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewUpdateUserLogic(r.Context(), ctx)
		resp, err := l.UpdateUser(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}

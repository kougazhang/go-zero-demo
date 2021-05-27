package handler

import (
	"net/http"

	"blog/api/internal/logic"
	"blog/api/internal/svc"
	"blog/api/internal/types"

	"github.com/tal-tech/go-zero/rest/httpx"
)

func DeleteUserHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ReqUserId
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewDeleteUserLogic(r.Context(), ctx)
		resp, err := l.DeleteUser(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}

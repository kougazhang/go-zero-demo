package handler

import (
	"net/http"

	"blog/api/internal/logic"
	"blog/api/internal/svc"
	"github.com/tal-tech/go-zero/rest/httpx"
)

func GetAllUsersHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		l := logic.NewGetAllUsersLogic(r.Context(), ctx)
		resp, err := l.GetAllUsers()
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}

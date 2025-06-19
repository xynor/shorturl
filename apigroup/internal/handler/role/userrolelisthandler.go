package role

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"shorturl/apigroup/internal/logic/role"
	"shorturl/apigroup/internal/svc"
)

func UserRoleListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := role.NewUserRoleListLogic(r.Context(), svcCtx)
		resp, err := l.UserRoleList()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}

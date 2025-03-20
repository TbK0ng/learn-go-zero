package user

import (
	"net/http"

	"github.com/TbK0ng/learn-go-zero/app/user-api/internal/logic/user"
	"github.com/TbK0ng/learn-go-zero/app/user-api/internal/svc"
	"github.com/TbK0ng/learn-go-zero/app/user-api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// 修改用户信息
func UserUpdateHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UserUpdateReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := user.NewUserUpdateLogic(r.Context(), svcCtx)
		resp, err := l.UserUpdate(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}

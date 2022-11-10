package api

import (
	"context"
	"net/http"

	"github.com/suyuan32/simple-admin-core/api/internal/svc"
	"github.com/suyuan32/simple-admin-core/api/internal/types"
	"github.com/suyuan32/simple-admin-core/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetApiListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	r      *http.Request
}

func NewGetApiListLogic(r *http.Request, svcCtx *svc.ServiceContext) *GetApiListLogic {
	return &GetApiListLogic{
		Logger: logx.WithContext(r.Context()),
		ctx:    r.Context(),
		svcCtx: svcCtx,
		r:      r,
	}
}

func (l *GetApiListLogic) GetApiList(req *types.ApiListReq) (resp *types.ApiListResp, err error) {
	data, err := l.svcCtx.CoreRpc.GetApiList(l.ctx,
		&core.ApiPageReq{
			Page:        req.Page,
			PageSize:    req.PageSize,
			Path:        req.Path,
			Description: req.Description,
			Method:      req.Method,
			Group:       req.Group,
		})
	if err != nil {
		return nil, err
	}
	resp = &types.ApiListResp{}
	resp.Total = data.GetTotal()

	lang := l.r.Header.Get("Accept-Language")

	for _, v := range data.Data {
		resp.Data = append(resp.Data,
			types.ApiInfo{
				BaseInfo: types.BaseInfo{
					Id:        v.Id,
					CreatedAt: v.CreatedAt,
					UpdatedAt: v.UpdatedAt,
				},
				Path:        v.Path,
				Title:       l.svcCtx.Trans.Trans(lang, v.Description),
				Description: v.Description,
				Group:       v.Group,
				Method:      v.Method,
			})
	}
	return resp, nil
}

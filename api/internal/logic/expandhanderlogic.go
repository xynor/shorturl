package logic

import (
	"context"

	"shorturl/api/internal/svc"
	"shorturl/api/internal/types"
	"shorturl/rpc/transform/transform"

	"github.com/zeromicro/go-zero/core/logx"
)

type ExpandHanderLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewExpandHanderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ExpandHanderLogic {
	return &ExpandHanderLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ExpandHanderLogic) ExpandHander(req *types.ExpandReq) (resp *types.ExpandResp, err error) {
	rpcResp, err := l.svcCtx.Transformer.Expand(l.ctx, &transform.ExpandReq{
		Shorten: req.ShortUrl,
	})
	if err != nil {
		return nil, err
	}
	return &types.ExpandResp{
		Url: rpcResp.Url,
	}, nil
}

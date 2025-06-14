package logic

import (
	"context"

	"shorturl/api/internal/svc"
	"shorturl/api/internal/types"
	"shorturl/rpc/transform/transformer"

	"github.com/zeromicro/go-zero/core/logx"
)

type ShortenHanderLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewShortenHanderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ShortenHanderLogic {
	return &ShortenHanderLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ShortenHanderLogic) ShortenHander(req *types.ShortenReq) (resp *types.ShortenResp, err error) {
	rpcResp, err := l.svcCtx.Transformer.Shorten(l.ctx, &transformer.ShortenReq{
		Url: req.Url,
	})
	if err != nil {
		return nil, err
	}

	return &types.ShortenResp{
		Shorten: rpcResp.Shorten,
	}, nil
}

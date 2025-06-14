package logic

import (
	"context"

	"shorturl/rpc-gorm/transform/internal/svc"
	"shorturl/rpc-gorm/transform/model"
	"shorturl/rpc-gorm/transform/transform"

	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/core/logx"
)

type ExpandLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewExpandLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ExpandLogic {
	return &ExpandLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ExpandLogic) Expand(in *transform.ExpandReq) (*transform.ExpandResp, error) {
	shortUrl := model.Shorturl{}
	logc.Info(l.ctx, "called Expand")
	// rest := l.svcCtx.DB.Model(&shortUrl).Where("shorten = ?", in.Shorten).First(&shortUrl)
	// rest := l.svcCtx.DB.WithContext(contextx.ValueOnlyFrom(l.ctx)).Model(&shortUrl).Where("shorten = ?", in.Shorten).First(&shortUrl)
	rest := l.svcCtx.DB.WithContext(l.ctx).Model(&shortUrl).Where("shorten = ?", in.Shorten).First(&shortUrl)
	if rest.Error != nil {
		return nil, rest.Error
	}

	return &transform.ExpandResp{
		Url: shortUrl.Url,
	}, nil
}

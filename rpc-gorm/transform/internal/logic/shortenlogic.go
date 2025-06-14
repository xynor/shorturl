package logic

import (
	"context"

	"shorturl/rpc-gorm/transform/internal/svc"
	"shorturl/rpc-gorm/transform/model"
	"shorturl/rpc-gorm/transform/transform"

	"github.com/zeromicro/go-zero/core/hash"
	"github.com/zeromicro/go-zero/core/logx"
)

type ShortenLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewShortenLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ShortenLogic {
	return &ShortenLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ShortenLogic) Shorten(in *transform.ShortenReq) (*transform.ShortenResp, error) {
	key := hash.Md5Hex([]byte(in.Url))[:6]
	shortUrl := model.Shorturl{
		Shorten: key,
		Url:     in.Url,
	}
	// rest := l.svcCtx.DB.Model(&shortUrl).Create(&shortUrl)
	// rest := l.svcCtx.DB.WithContext(contextx.ValueOnlyFrom(l.ctx)).Model(&shortUrl).Create(&shortUrl)
	rest := l.svcCtx.DB.WithContext(l.ctx).Model(&shortUrl).Create(&shortUrl)
	if rest.Error != nil {
		return nil, rest.Error
	}

	return &transform.ShortenResp{
		Shorten: key,
	}, nil
}

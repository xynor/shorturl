// Code generated by goctl. DO NOT EDIT.
// goctl 1.8.4
// Source: transform.proto

package transformer

import (
	"context"

	"shorturl/rpc/transform/transform"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	ExpandReq   = transform.ExpandReq
	ExpandResp  = transform.ExpandResp
	ShortenReq  = transform.ShortenReq
	ShortenResp = transform.ShortenResp

	Transformer interface {
		Expand(ctx context.Context, in *ExpandReq, opts ...grpc.CallOption) (*ExpandResp, error)
		Shorten(ctx context.Context, in *ShortenReq, opts ...grpc.CallOption) (*ShortenResp, error)
	}

	defaultTransformer struct {
		cli zrpc.Client
	}
)

func NewTransformer(cli zrpc.Client) Transformer {
	return &defaultTransformer{
		cli: cli,
	}
}

func (m *defaultTransformer) Expand(ctx context.Context, in *ExpandReq, opts ...grpc.CallOption) (*ExpandResp, error) {
	client := transform.NewTransformerClient(m.cli.Conn())
	return client.Expand(ctx, in, opts...)
}

func (m *defaultTransformer) Shorten(ctx context.Context, in *ShortenReq, opts ...grpc.CallOption) (*ShortenResp, error) {
	client := transform.NewTransformerClient(m.cli.Conn())
	return client.Shorten(ctx, in, opts...)
}

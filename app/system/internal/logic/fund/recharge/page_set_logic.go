// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.0

package recharge

import (
	"context"

	"ovra/app/system/internal/svc"
	"ovra/app/system/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

// PageSetLogic 充值管理分页查询（占位，待接 fund 充值表）
type PageSetLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPageSetLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PageSetLogic {
	return &PageSetLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// PageSet 分页查询充值管理
// 1. 业务表尚未接入，返回空分页供前端 VxeGrid 联调
// 2. 后续按 keyword/status 及时间范围查询
func (l *PageSetLogic) PageSet(req *types.PageSetFundRechargeReq) (resp *types.PageSetFundRechargeResp, err error) {
	_ = req
	return &types.PageSetFundRechargeResp{Rows: []*types.FundRechargeItem{}, Total: 0}, nil
}

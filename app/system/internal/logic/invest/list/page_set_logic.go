// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.0

package list

import (
	"context"

	"ovra/app/system/internal/svc"
	"ovra/app/system/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

// PageSetLogic 投信列表分页查询（占位，待接 invest 产品表）
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

// PageSet 分页查询投信列表
// 1. 业务表尚未接入，返回空分页供前端 VxeGrid 联调
// 2. 后续按 keyword/status/soldOut 及时间范围查询
func (l *PageSetLogic) PageSet(req *types.PageSetInvestListReq) (resp *types.PageSetInvestListResp, err error) {
	_ = req
	return &types.PageSetInvestListResp{
		Rows:  []*types.InvestListItem{},
		Total: 0,
	}, nil
}

// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.0

package user

import (
	"context"

	"ovra/app/system/internal/svc"
	"ovra/app/system/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

// PageSetLogic 业务用户分页查询（占位，待接 biz 用户表）
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

// PageSet 分页查询业务用户
// 1. 业务表尚未接入，返回空分页供前端 VxeGrid 联调
// 2. 后续按 keyword/authStatus/时间范围/deleted 及时间范围查询
func (l *PageSetLogic) PageSet(req *types.PageSetMemberUserReq) (resp *types.PageSetMemberUserResp, err error) {
	// 2. 后续按 keyword/authStatus/时间范围/deleted 查询 biz 用户表
	_ = req
	return &types.PageSetMemberUserResp{
		Rows:  []*types.MemberUserItem{},
		Total: 0,
	}, nil
}

// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.0

package kyc

import (
	"context"

	"ovra/app/system/internal/svc"
	"ovra/app/system/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

// PageSetLogic 实名认证分页查询（占位，待接 member KYC 表）
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

// PageSet 分页查询实名认证
// 1. 业务表尚未接入，返回空分页供前端 VxeGrid 联调
// 2. 后续按 keyword/authStatus 及时间范围查询
func (l *PageSetLogic) PageSet(req *types.PageSetMemberKycReq) (resp *types.PageSetMemberKycResp, err error) {
	// 2. 后续按 keyword/authStatus/提交时间范围查询并映射 MemberKycItem
	_ = req
	return &types.PageSetMemberKycResp{
		Rows:  []*types.MemberKycItem{},
		Total: 0,
	}, nil
}

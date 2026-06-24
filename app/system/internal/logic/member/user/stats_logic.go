// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.0

package user

import (
	"context"

	"ovra/app/system/internal/svc"
	"ovra/app/system/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

// StatsLogic 业务用户活跃会话统计（占位，待接在线会话/用户表）
type StatsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewStatsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *StatsLogic {
	return &StatsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// Stats 统计当前筛选条件下的活跃会话数
// 1. 业务表尚未接入，返回 0 供前端标题 Tag 联调
// 2. 后续按 keyword/authStatus/deleted/时间范围聚合在线会话 totalActiveSessions
func (l *StatsLogic) Stats(req *types.MemberUserQuery) (resp *types.MemberUserStatsResp, err error) {
	_ = req
	return &types.MemberUserStatsResp{
		TotalActiveSessions: 0,
	}, nil
}

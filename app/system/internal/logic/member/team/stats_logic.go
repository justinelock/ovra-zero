// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.0

package team

import (
	"context"

	"ovra/app/system/internal/svc"
	"ovra/app/system/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

// StatsLogic 团队代理统计（占位，待接 member 团队表）
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

// Stats 聚合各级代理人数与团队总余额
// 1. 业务表尚未接入，返回零值供前端统计卡片联调
// 2. 后续按 keyword/status 聚合 level1～level5、totalMembers、totalBalance
func (l *StatsLogic) Stats(req *types.MemberTeamQuery) (resp *types.MemberTeamStatsResp, err error) {
	_ = req
	return &types.MemberTeamStatsResp{
		Level1Count:  0,
		Level2Count:  0,
		Level3Count:  0,
		Level4Count:  0,
		Level5Count:  0,
		TotalMembers: 0,
		TotalBalance: "0.00",
	}, nil
}

// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.0

package team

import (
	"context"

	"ovra/app/system/internal/dal"
	"ovra/app/system/internal/svc"
	"ovra/app/system/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

// StatsLogic 团队代理统计
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

// Stats 按 fb_users.level 1～5 分层计数并汇总团队钱包余额
func (l *StatsLogic) Stats(req *types.MemberTeamQuery) (resp *types.MemberTeamStatsResp, err error) {
	f := dal.MemberListFilter{
		Keyword: req.Keyword,
		Status:  req.Status,
	}
	levels, totalMembers, totalBalance, err := l.svcCtx.Dal.FbMemberDal.TeamStats(l.ctx, f)
	if err != nil {
		return nil, err
	}
	return &types.MemberTeamStatsResp{
		TotalMembers:     totalMembers,
		Level1Members:    levels[1],
		Level2Members:    levels[2],
		Level3Members:    levels[3],
		Level4Members:    levels[4],
		Level5Members:    levels[5],
		TotalTeamBalance: totalBalance,
	}, nil
}

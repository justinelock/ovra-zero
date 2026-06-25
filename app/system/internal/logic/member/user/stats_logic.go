// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.0

package user

import (
	"context"

	"ovra/app/system/internal/dal"
	"ovra/app/system/internal/svc"
	"ovra/app/system/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

// StatsLogic 业务用户活跃统计
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

// Stats 统计在线用户、今日登录与近 30 分钟活跃设备会话
func (l *StatsLogic) Stats(req *types.MemberUserQuery) (resp *types.MemberUserStatsResp, err error) {
	f := dal.MemberListFilter{
		Keyword:    req.Keyword,
		AuthStatus: req.AuthStatus,
		Deleted:    req.Deleted,
	}
	online, today, sessions, err := l.svcCtx.Dal.FbMemberDal.UserStats(l.ctx, f)
	if err != nil {
		return nil, err
	}
	return &types.MemberUserStatsResp{
		TotalOnlineUsers:    online,
		TodayLogins:         today,
		TotalActiveSessions: sessions,
	}, nil
}

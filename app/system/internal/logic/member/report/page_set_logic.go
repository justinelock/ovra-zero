// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.0

package report

import (
	"context"

	"ovra/app/system/internal/dal"
	"ovra/app/system/internal/svc"
	"ovra/app/system/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

// PageSetLogic 用户报表分页
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

// PageSet 分页查询用户报表（充值/提现/余额/上级）
func (l *PageSetLogic) PageSet(req *types.PageSetMemberReportReq) (resp *types.PageSetMemberReportResp, err error) {
	f := dal.MemberListFilter{
		Keyword:   req.Keyword,
		Level:     req.Level,
		BeginTime: req.BeginTime,
		EndTime:   req.EndTime,
		PageNum:   req.PageNum,
		PageSize:  req.PageSize,
	}
	rows, total, err := l.svcCtx.Dal.FbMemberDal.PageReports(l.ctx, f)
	if err != nil {
		return nil, err
	}
	items := make([]*types.MemberReportItem, 0, len(rows))
	for _, r := range rows {
		item := &types.MemberReportItem{
			Id:             dal.IDStr(r.ID),
			UserId:         dal.IDStr(r.UserID),
			Username:       r.Username,
			Mobile:         r.Mobile,
			RealName:       r.RealName,
			Level:          int64(r.Level),
			Amount:         r.Amount,
			RechargeAmount: r.RechargeAmount,
			WithdrawAmount: r.WithdrawAmount,
			RechargeDiff:   r.RechargeDiff,
			TotalProfit:    r.TotalProfit,
			TeamCount:      int64(r.TeamCount),
			RegisterTime:   dal.FormatFbTimeVal(r.RegisterTime),
			LastLogin:      dal.FormatFbTime(r.LastLogin),
			LoginIp:        r.LoginIP,
		}
		if r.ParentID > 0 {
			item.ParentUser = &types.MemberReportParentUser{
				Id:       dal.IDStr(r.ParentID),
				Username: r.ParentUsername,
			}
		}
		items = append(items, item)
	}
	return &types.PageSetMemberReportResp{Rows: items, Total: total}, nil
}

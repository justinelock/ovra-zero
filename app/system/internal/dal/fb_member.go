package dal

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"time"

	"ovra/toolkit/errx"

	"gorm.io/gorm"
)

// FbMemberDal 业务用户域查询（fb_users 及关联表）
type FbMemberDal struct {
	db *gorm.DB
}

func NewFbMemberDal(db *gorm.DB) *FbMemberDal {
	return &FbMemberDal{db: db}
}

// MemberListFilter 用户管理列表通用筛选
type MemberListFilter struct {
	Keyword    string
	AuthStatus string
	Deleted    string
	BeginTime  string
	EndTime    string
	Level      string
	Status     string
	PageNum    int64
	PageSize   int64
}

type memberUserRow struct {
	ID                   int64
	Username             string
	RealName             string
	IDCard               string
	AgentLevel           int32
	InviteCode           string
	CommissionRate       float64
	TotalCommission      float64
	Status               string
	IsOnline             string
	LastLogin            *time.Time
	CreatedAt            time.Time
	TotalBalance         float64
	FundPositionAmount   float64
	FundPositionDividend float64
}

// PageUsers 分页查询业务用户列表（含钱包余额与投信持仓聚合）
func (d *FbMemberDal) PageUsers(ctx context.Context, f MemberListFilter) (rows []memberUserRow, total int64, err error) {
	where, args := d.userWhere(f)
	base := `
FROM fb_users u
LEFT JOIN (
  SELECT user_id, SUM(balance) AS total_balance FROM fb_user_wallets GROUP BY user_id
) w ON w.user_id = u.id
LEFT JOIN (
  SELECT user_id, SUM(amount) AS position_amount, SUM(profit) AS position_dividend
  FROM fb_fund_position WHERE status = 1 GROUP BY user_id
) p ON p.user_id = u.id
WHERE ` + where

	countSQL := "SELECT COUNT(*) " + base
	if err = d.db.WithContext(ctx).Raw(countSQL, args...).Scan(&total).Error; err != nil {
		return nil, 0, errx.GORMErr(err)
	}

	offset := (f.PageNum - 1) * f.PageSize
	if offset < 0 {
		offset = 0
	}
	if f.PageSize <= 0 {
		f.PageSize = 10
	}
	listSQL := `SELECT u.id, u.username, u.real_name, u.id_card, u.agent_level, u.invite_code,
  u.commission_rate, u.total_commission, u.status, u.is_online, u.last_login, u.created_at,
  COALESCE(w.total_balance, 0) AS total_balance,
  COALESCE(p.position_amount, 0) AS fund_position_amount,
  COALESCE(p.position_dividend, 0) AS fund_position_dividend ` + base +
		" ORDER BY u.created_at DESC LIMIT ? OFFSET ?"
	listArgs := append(append([]any{}, args...), f.PageSize, offset)
	if err = d.db.WithContext(ctx).Raw(listSQL, listArgs...).Scan(&rows).Error; err != nil {
		return nil, 0, errx.GORMErr(err)
	}
	return rows, total, nil
}

// UserStats 用户列表顶部统计
func (d *FbMemberDal) UserStats(ctx context.Context, f MemberListFilter) (online, todayLogins, activeSessions int64, err error) {
	where, args := d.userWhere(f)
	var onlineCnt int64
	if err = d.db.WithContext(ctx).Raw(
		"SELECT COUNT(*) FROM fb_users u WHERE "+where+" AND u.is_online = '1'", args...,
	).Scan(&onlineCnt).Error; err != nil {
		return 0, 0, 0, errx.GORMErr(err)
	}
	today := time.Now().Format("2006-01-02")
	var todayCnt int64
	if err = d.db.WithContext(ctx).Raw(`
SELECT COUNT(*) FROM fb_device_login_log l
INNER JOIN fb_users u ON u.id = l.user_id
WHERE `+where+` AND l.login_result = 'SUCCESS' AND l.login_time >= ? AND l.login_time < ? + INTERVAL 1 DAY`,
		append(args, today, today)...,
	).Scan(&todayCnt).Error; err != nil {
		return 0, 0, 0, errx.GORMErr(err)
	}
	// 活跃会话：近 30 分钟内有成功登录记录的去重设备数
	var sessionCnt int64
	if err = d.db.WithContext(ctx).Raw(`
SELECT COUNT(DISTINCT l.device_id) FROM fb_device_login_log l
INNER JOIN fb_users u ON u.id = l.user_id
WHERE `+where+` AND l.login_result = 'SUCCESS' AND l.login_time >= NOW() - INTERVAL 30 MINUTE`,
		args...,
	).Scan(&sessionCnt).Error; err != nil {
		return 0, 0, 0, errx.GORMErr(err)
	}
	return onlineCnt, todayCnt, sessionCnt, nil
}

type memberKycRow struct {
	ID           int64
	UserID       int64
	Username     string
	Mobile       string
	RealName     string
	IDCardNo     string
	IDCardFront  string
	IDCardBack   string
	Status       string
	RejectReason string
	VerifiedAt   *time.Time
	CreatedAt    time.Time
	UpdatedAt    *time.Time
}

// PageKyc 实名认证分页
func (d *FbMemberDal) PageKyc(ctx context.Context, f MemberListFilter) (rows []memberKycRow, total int64, err error) {
	where := []string{"1=1"}
	var args []any
	if f.Keyword != "" {
		where = append(where, "(u.username LIKE ? OR u.mobile LIKE ? OR v.real_name LIKE ? OR CAST(v.id AS CHAR) LIKE ?)")
		kw := "%" + f.Keyword + "%"
		args = append(args, kw, kw, kw, kw)
	}
	if f.AuthStatus != "" {
		where = append(where, "v.status = ?")
		args = append(args, f.AuthStatus)
	}
	if f.BeginTime != "" {
		where = append(where, "v.created_at >= ?")
		args = append(args, f.BeginTime)
	}
	if f.EndTime != "" {
		where = append(where, "v.created_at <= ?")
		args = append(args, f.EndTime)
	}
	w := strings.Join(where, " AND ")
	base := `FROM fb_identity_verify v LEFT JOIN fb_users u ON u.id = v.user_id WHERE ` + w
	if err = d.db.WithContext(ctx).Raw("SELECT COUNT(*) "+base, args...).Scan(&total).Error; err != nil {
		return nil, 0, errx.GORMErr(err)
	}
	offset := (f.PageNum - 1) * f.PageSize
	listSQL := `SELECT v.id, v.user_id, u.username, u.mobile, v.real_name, v.id_card_no,
  v.id_card_front, v.id_card_back, v.status, v.reject_reason, v.verified_at, v.created_at, v.updated_at ` +
		base + " ORDER BY v.created_at DESC LIMIT ? OFFSET ?"
	listArgs := append(append([]any{}, args...), f.PageSize, offset)
	if err = d.db.WithContext(ctx).Raw(listSQL, listArgs...).Scan(&rows).Error; err != nil {
		return nil, 0, errx.GORMErr(err)
	}
	return rows, total, nil
}

type memberWalletRow struct {
	ID           int64
	UserID       int64
	Username     string
	Mobile       string
	RealName     string
	AccountType  string
	Balance      float64
	FrozenAmount float64
	Frozen       bool
	Version      int64
	Currency     string
	DrawTicket   int32
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

// PageWallets 钱包分页
func (d *FbMemberDal) PageWallets(ctx context.Context, f MemberListFilter, accountType, currency, frozenStatus string) (rows []memberWalletRow, total int64, err error) {
	where := []string{"1=1"}
	var args []any
	if f.Keyword != "" {
		where = append(where, "(u.username LIKE ? OR u.mobile LIKE ? OR u.real_name LIKE ? OR CAST(w.id AS CHAR) LIKE ?)")
		kw := "%" + f.Keyword + "%"
		args = append(args, kw, kw, kw, kw)
	}
	if accountType != "" {
		where = append(where, "w.account_type = ?")
		args = append(args, accountType)
	}
	if currency != "" {
		where = append(where, "w.currency = ?")
		args = append(args, currency)
	}
	if frozenStatus != "" {
		where = append(where, "w.frozen = ?")
		args = append(args, frozenStatus == "1" || frozenStatus == "true")
	}
	if f.BeginTime != "" {
		where = append(where, "w.created_at >= ?")
		args = append(args, f.BeginTime)
	}
	if f.EndTime != "" {
		where = append(where, "w.created_at <= ?")
		args = append(args, f.EndTime)
	}
	w := strings.Join(where, " AND ")
	base := `FROM fb_user_wallets w LEFT JOIN fb_users u ON u.id = w.user_id WHERE ` + w
	if err = d.db.WithContext(ctx).Raw("SELECT COUNT(*) "+base, args...).Scan(&total).Error; err != nil {
		return nil, 0, errx.GORMErr(err)
	}
	offset := (f.PageNum - 1) * f.PageSize
	listSQL := `SELECT w.id, w.user_id, u.username, u.mobile, u.real_name, w.account_type,
  w.balance, w.frozen_amount, w.frozen, w.version, w.currency, w.draw_ticket, w.created_at, w.updated_at ` +
		base + " ORDER BY w.created_at DESC LIMIT ? OFFSET ?"
	listArgs := append(append([]any{}, args...), f.PageSize, offset)
	if err = d.db.WithContext(ctx).Raw(listSQL, listArgs...).Scan(&rows).Error; err != nil {
		return nil, 0, errx.GORMErr(err)
	}
	return rows, total, nil
}

type memberReportRow struct {
	ID             int64
	UserID         int64
	Username       string
	Mobile         string
	RealName       string
	Level          int32
	Amount         float64
	RechargeAmount float64
	WithdrawAmount float64
	RechargeDiff   float64
	TotalProfit    float64
	TeamCount      int32
	RegisterTime   time.Time
	LastLogin      *time.Time
	LoginIP        string
	ParentID       int64
	ParentUsername string
}

// PageReports 用户报表分页
func (d *FbMemberDal) PageReports(ctx context.Context, f MemberListFilter) (rows []memberReportRow, total int64, err error) {
	where, args := d.userWhere(f)
	if f.Level != "" {
		where += " AND u.level = ?"
		args = append(args, f.Level)
	}
	base := `
FROM fb_users u
LEFT JOIN fb_users p ON p.id = u.parent_id
LEFT JOIN (SELECT user_id, SUM(balance) AS bal FROM fb_user_wallets GROUP BY user_id) w ON w.user_id = u.id
LEFT JOIN (SELECT user_id, SUM(amount) AS amt FROM fb_deposits WHERE status = 'SUCCESS' GROUP BY user_id) dep ON dep.user_id = u.id
LEFT JOIN (SELECT user_id, SUM(amount) AS amt FROM fb_withdraws WHERE status = 'SUCCESS' GROUP BY user_id) wd ON wd.user_id = u.id
LEFT JOIN (
  SELECT user_id, login_ip FROM fb_device_login_log l1
  WHERE l1.id = (SELECT MAX(l2.id) FROM fb_device_login_log l2 WHERE l2.user_id = l1.user_id)
) ll ON ll.user_id = u.id
WHERE ` + where

	if err = d.db.WithContext(ctx).Raw("SELECT COUNT(*) "+base, args...).Scan(&total).Error; err != nil {
		return nil, 0, errx.GORMErr(err)
	}
	offset := (f.PageNum - 1) * f.PageSize
	listSQL := `SELECT u.id, u.id AS user_id, u.username, u.mobile, u.real_name, u.level,
  COALESCE(w.bal, 0) AS amount,
  COALESCE(dep.amt, 0) AS recharge_amount,
  COALESCE(wd.amt, 0) AS withdraw_amount,
  COALESCE(dep.amt, 0) - COALESCE(wd.amt, 0) AS recharge_diff,
  0 AS total_profit,
  u.team_size AS team_count,
  u.created_at AS register_time, u.last_login, COALESCE(ll.login_ip, '') AS login_ip,
  COALESCE(u.parent_id, 0) AS parent_id, COALESCE(p.username, '') AS parent_username ` +
		base + " ORDER BY u.created_at DESC LIMIT ? OFFSET ?"
	listArgs := append(append([]any{}, args...), f.PageSize, offset)
	if err = d.db.WithContext(ctx).Raw(listSQL, listArgs...).Scan(&rows).Error; err != nil {
		return nil, 0, errx.GORMErr(err)
	}
	return rows, total, nil
}

type memberFlowRow struct {
	ID           int64
	UserID       int64
	Username     string
	Mobile       string
	RealName     string
	AccountType  string
	FlowType     string
	BeforeAmount float64
	FlowAmount   float64
	AfterAmount  float64
	BusinessNo   string
	Remark       string
	CreatedAt    time.Time
	WalletID     int64
	Currency     string
	Description  string
	Status       string
	UpdatedAt    *time.Time
}

// PageReportFlow 用户流水明细分页
func (d *FbMemberDal) PageReportFlow(ctx context.Context, userID string, f MemberListFilter) (rows []memberFlowRow, total int64, err error) {
	where := []string{"f.user_id = ?"}
	args := []any{userID}
	if f.BeginTime != "" {
		where = append(where, "f.created_at >= ?")
		args = append(args, f.BeginTime)
	}
	if f.EndTime != "" {
		where = append(where, "f.created_at <= ?")
		args = append(args, f.EndTime)
	}
	w := strings.Join(where, " AND ")
	base := `FROM fb_account_flow_records f LEFT JOIN fb_users u ON u.id = f.user_id WHERE ` + w
	if err = d.db.WithContext(ctx).Raw("SELECT COUNT(*) "+base, args...).Scan(&total).Error; err != nil {
		return nil, 0, errx.GORMErr(err)
	}
	offset := (f.PageNum - 1) * f.PageSize
	listSQL := `SELECT f.id, f.user_id, u.username, u.mobile, u.real_name, f.account_type, f.flow_type,
  f.before_amount, f.flow_amount, f.after_amount, f.business_no, f.remark, f.created_at,
  COALESCE(f.wallet_id, 0) AS wallet_id, f.currency, f.description, f.status, f.updated_at ` +
		base + " ORDER BY f.created_at DESC LIMIT ? OFFSET ?"
	listArgs := append(append([]any{}, args...), f.PageSize, offset)
	if err = d.db.WithContext(ctx).Raw(listSQL, listArgs...).Scan(&rows).Error; err != nil {
		return nil, 0, errx.GORMErr(err)
	}
	return rows, total, nil
}

type memberTeamRow struct {
	ID               int64
	Username         string
	RealName         string
	Level            int32
	Status           string
	TeamSize         int32
	AgentLevel       int32
	Balance          float64
	TotalTeamBalance float64
	CreatedAt        time.Time
	ParentID         int64
	ParentUsername   string
	ParentRealName   string
	Level1Members    int64
	WalletCount      int64
}

// PageTeams 团队列表分页
func (d *FbMemberDal) PageTeams(ctx context.Context, f MemberListFilter) (rows []memberTeamRow, total int64, err error) {
	where, args := d.userWhere(f)
	if f.Status != "" {
		where += " AND u.status = ?"
		args = append(args, f.Status)
	}
	base := `
FROM fb_users u
LEFT JOIN fb_users p ON p.id = u.parent_id
LEFT JOIN (SELECT user_id, SUM(balance) AS bal FROM fb_user_wallets GROUP BY user_id) w ON w.user_id = u.id
LEFT JOIN (SELECT parent_id, COUNT(*) AS cnt FROM fb_users WHERE flag = 0 AND parent_id IS NOT NULL GROUP BY parent_id) c1 ON c1.parent_id = u.id
LEFT JOIN (SELECT user_id, COUNT(*) AS cnt FROM fb_user_wallets GROUP BY user_id) wc ON wc.user_id = u.id
WHERE ` + where

	if err = d.db.WithContext(ctx).Raw("SELECT COUNT(*) "+base, args...).Scan(&total).Error; err != nil {
		return nil, 0, errx.GORMErr(err)
	}
	offset := (f.PageNum - 1) * f.PageSize
	listSQL := `SELECT u.id, u.username, u.real_name, u.level, u.status, u.team_size, u.agent_level,
  COALESCE(w.bal, 0) AS balance, 0 AS total_team_balance, u.created_at,
  COALESCE(u.parent_id, 0) AS parent_id, COALESCE(p.username, '') AS parent_username, COALESCE(p.real_name, '') AS parent_real_name,
  COALESCE(c1.cnt, 0) AS level1_members, COALESCE(wc.cnt, 0) AS wallet_count ` +
		base + " ORDER BY u.created_at DESC LIMIT ? OFFSET ?"
	listArgs := append(append([]any{}, args...), f.PageSize, offset)
	if err = d.db.WithContext(ctx).Raw(listSQL, listArgs...).Scan(&rows).Error; err != nil {
		return nil, 0, errx.GORMErr(err)
	}
	return rows, total, nil
}

// TeamStats 团队代理统计（按 fb_users.level 1～5 分层计数）
func (d *FbMemberDal) TeamStats(ctx context.Context, f MemberListFilter) (levels [6]int64, totalMembers int64, totalBalance float64, err error) {
	where, args := d.userWhere(f)
	if f.Status != "" {
		where += " AND u.status = ?"
		args = append(args, f.Status)
	}
	type levelRow struct {
		Level int32
		Cnt   int64
	}
	var lr []levelRow
	if err = d.db.WithContext(ctx).Raw(
		"SELECT u.level, COUNT(*) AS cnt FROM fb_users u WHERE "+where+" GROUP BY u.level", args...,
	).Scan(&lr).Error; err != nil {
		return levels, 0, 0, errx.GORMErr(err)
	}
	for _, r := range lr {
		totalMembers += r.Cnt
		if r.Level >= 1 && r.Level <= 5 {
			levels[r.Level] = r.Cnt
		}
	}
	if err = d.db.WithContext(ctx).Raw(`
SELECT COALESCE(SUM(w.balance), 0) FROM fb_user_wallets w
INNER JOIN fb_users u ON u.id = w.user_id WHERE `+where, args...,
	).Scan(&totalBalance).Error; err != nil {
		return levels, 0, 0, errx.GORMErr(err)
	}
	return levels, totalMembers, totalBalance, nil
}

type memberLoginLogRow struct {
	ID            int64
	UserID        int64
	Username      string
	RealName      string
	DeviceID      string
	LoginTime     time.Time
	LoginIP       string
	LoginLocation string
	LoginType     string
	LoginResult   string
	FailReason    string
	RiskLevel     string
	RiskDetail    string
}

// PageLoginLogs 登录记录分页
func (d *FbMemberDal) PageLoginLogs(ctx context.Context, f MemberListFilter, loginResult, loginMethod, riskLevel string) (rows []memberLoginLogRow, total int64, err error) {
	where := []string{"1=1"}
	var args []any
	if f.Keyword != "" {
		where = append(where, "(u.username LIKE ? OR u.real_name LIKE ? OR l.device_id LIKE ? OR l.login_ip LIKE ?)")
		kw := "%" + f.Keyword + "%"
		args = append(args, kw, kw, kw, kw)
	}
	if loginResult != "" {
		where = append(where, "l.login_result = ?")
		args = append(args, loginResult)
	}
	if loginMethod != "" {
		where = append(where, "l.login_type = ?")
		args = append(args, loginMethod)
	}
	if riskLevel != "" {
		where = append(where, "l.risk_level = ?")
		args = append(args, riskLevel)
	}
	if f.BeginTime != "" {
		where = append(where, "l.login_time >= ?")
		args = append(args, f.BeginTime)
	}
	if f.EndTime != "" {
		where = append(where, "l.login_time <= ?")
		args = append(args, f.EndTime)
	}
	w := strings.Join(where, " AND ")
	base := `FROM fb_device_login_log l LEFT JOIN fb_users u ON u.id = l.user_id WHERE ` + w
	if err = d.db.WithContext(ctx).Raw("SELECT COUNT(*) "+base, args...).Scan(&total).Error; err != nil {
		return nil, 0, errx.GORMErr(err)
	}
	offset := (f.PageNum - 1) * f.PageSize
	listSQL := `SELECT l.id, l.user_id, u.username, u.real_name, l.device_id, l.login_time,
  l.login_ip, l.login_location, l.login_type, l.login_result, l.fail_reason, l.risk_level, l.risk_detail ` +
		base + " ORDER BY l.login_time DESC LIMIT ? OFFSET ?"
	listArgs := append(append([]any{}, args...), f.PageSize, offset)
	if err = d.db.WithContext(ctx).Raw(listSQL, listArgs...).Scan(&rows).Error; err != nil {
		return nil, 0, errx.GORMErr(err)
	}
	return rows, total, nil
}

func (d *FbMemberDal) userWhere(f MemberListFilter) (string, []any) {
	where := []string{"1=1"}
	var args []any
	if f.Deleted == "1" {
		where = append(where, "u.flag = 1")
	} else {
		where = append(where, "u.flag = 0")
	}
	if f.Keyword != "" {
		where = append(where, "(u.username LIKE ? OR u.mobile LIKE ? OR u.real_name LIKE ? OR CAST(u.id AS CHAR) LIKE ?)")
		kw := "%" + f.Keyword + "%"
		args = append(args, kw, kw, kw, kw)
	}
	if f.AuthStatus != "" {
		where = append(where, "u.verification_status = ?")
		args = append(args, f.AuthStatus)
	}
	if f.BeginTime != "" {
		where = append(where, "u.created_at >= ?")
		args = append(args, f.BeginTime)
	}
	if f.EndTime != "" {
		where = append(where, "u.created_at <= ?")
		args = append(args, f.EndTime)
	}
	return strings.Join(where, " AND "), args
}

// FormatFbTime 格式化可空时间供 API 返回
func FormatFbTime(t *time.Time) string {
	if t == nil || t.IsZero() {
		return ""
	}
	return t.Format("2006-01-02 15:04:05")
}

func FormatFbTimeVal(t time.Time) string {
	if t.IsZero() {
		return ""
	}
	return t.Format("2006-01-02 15:04:05")
}

func ParseOnlineStatus(s string) int64 {
	n, _ := strconv.ParseInt(s, 10, 64)
	return n
}

func IDStr(id int64) string {
	return fmt.Sprintf("%d", id)
}

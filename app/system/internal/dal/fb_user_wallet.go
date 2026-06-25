package dal

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"math"
	"strconv"
	"time"

	"ovra/app/system/internal/dal/model"
	"ovra/toolkit/errx"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// FbUserWalletDal 用户钱包写操作（加减款、流水），对齐 Java FbUserWalletsServiceImpl
type FbUserWalletDal struct {
	db *gorm.DB
}

func NewFbUserWalletDal(db *gorm.DB) *FbUserWalletDal {
	return &FbUserWalletDal{db: db}
}

// 管理端加款可选 flowType
var adminAddFlowTypes = map[string]struct{}{
	"ADD_AMOUNT":   {},
	"ADD_BONUS":    {},
	"ADD_DIVIDEND": {},
	"ADD_TRANSFER": {},
}

// 管理端减款可选 flowType
var adminSubtractFlowTypes = map[string]struct{}{
	"SUBTRACT_AMOUNT": {},
	"ADD_TRANSFER":    {},
}

// ParseWalletID 解析钱包主键 id
func ParseWalletID(id string) (int64, error) {
	if id == "" {
		return 0, errx.BizErr("钱包ID不能为空")
	}
	walletID, err := strconv.ParseInt(id, 10, 64)
	if err != nil || walletID <= 0 {
		return 0, errx.BizErr("钱包ID无效")
	}
	return walletID, nil
}

// NormalizeAddOrSubtractFlowType 校验并回填 flowType（加款默认 ADD_AMOUNT，减款默认 SUBTRACT_AMOUNT）
func NormalizeAddOrSubtractFlowType(isAdd bool, flowType string) (string, error) {
	if flowType == "" {
		if isAdd {
			return "ADD_AMOUNT", nil
		}
		return "SUBTRACT_AMOUNT", nil
	}
	if isAdd {
		if _, ok := adminAddFlowTypes[flowType]; !ok {
			return "", errx.BizErr("加款流水类型无效")
		}
		return flowType, nil
	}
	if _, ok := adminSubtractFlowTypes[flowType]; !ok {
		return "", errx.BizErr("减款流水类型无效")
	}
	return flowType, nil
}

// AddBalance 按钱包 id 加款并写 fb_account_flow_records（对齐 addBalance(walletId, ...)）
func (d *FbUserWalletDal) AddBalance(ctx context.Context, walletID int64, amount float64, remark, flowType string) error {
	amount = math.Abs(amount)
	return d.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		wallet, err := d.lockWallet(tx, walletID)
		if err != nil {
			return err
		}
		before := wallet.Balance
		after := before + amount
		now := time.Now()
		if err := tx.Model(&model.FbUserWallet{}).Where("id = ?", walletID).Updates(map[string]any{
			"balance":    after,
			"updated_at": now,
		}).Error; err != nil {
			return errx.GORMErr(err)
		}
		flow := &model.FbAccountFlowRecord{
			UserID:       wallet.UserID,
			AccountType:  wallet.AccountType,
			FlowType:     flowType,
			BeforeAmount: before,
			FlowAmount:   amount,
			AfterAmount:  after,
			BusinessNo:   genWalletBusinessNo("ADD_AMOUNT"),
			Remark:       remark,
			CreatedAt:    now,
			WalletID:     wallet.ID,
			Currency:     wallet.Currency,
			Description:  "加款",
			Status:       "SUCCESS",
		}
		if err := tx.Create(flow).Error; err != nil {
			return errx.BizErr("资金流水写入失败")
		}
		return nil
	})
}

// ReduceBalance 按钱包 id 减款并写流水（对齐 reduceBalance(walletId, ...)）
func (d *FbUserWalletDal) ReduceBalance(ctx context.Context, walletID int64, amount float64, remark, flowType string) error {
	amount = math.Abs(amount)
	return d.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		wallet, err := d.lockWallet(tx, walletID)
		if err != nil {
			return err
		}
		before := wallet.Balance
		if before < amount {
			return errx.BizErr("余额不足")
		}
		after := before - amount
		now := time.Now()
		if err := tx.Model(&model.FbUserWallet{}).Where("id = ?", walletID).Updates(map[string]any{
			"balance":    after,
			"updated_at": now,
		}).Error; err != nil {
			return errx.GORMErr(err)
		}
		flow := &model.FbAccountFlowRecord{
			UserID:       wallet.UserID,
			AccountType:  wallet.AccountType,
			FlowType:     flowType,
			BeforeAmount: before,
			FlowAmount:   -amount,
			AfterAmount:  after,
			BusinessNo:   genWalletBusinessNo("SUBTRACT_AMOUNT"),
			Remark:       remark,
			CreatedAt:    now,
			WalletID:     wallet.ID,
			Currency:     wallet.Currency,
			Description:  "减款",
			Status:       "SUCCESS",
		}
		if err := tx.Create(flow).Error; err != nil {
			return errx.BizErr("资金流水写入失败")
		}
		return nil
	})
}

// lockWallet 行锁读取钱包，避免并发加减款覆盖
func (d *FbUserWalletDal) lockWallet(tx *gorm.DB, walletID int64) (*model.FbUserWallet, error) {
	var wallet model.FbUserWallet
	err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).
		Where("id = ?", walletID).
		First(&wallet).Error
	if err != nil {
		return nil, errx.GORMErrMsg(err, "用户钱包不存在")
	}
	return &wallet, nil
}

func genWalletBusinessNo(prefix string) string {
	buf := make([]byte, 4)
	_, _ = rand.Read(buf)
	return fmt.Sprintf("%s_%d%s", prefix, time.Now().UnixMilli(), hex.EncodeToString(buf))
}

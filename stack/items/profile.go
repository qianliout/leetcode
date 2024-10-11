package items

import "fmt"

// 利润表

type Profile struct {
	ID       int64  `gorm:"column:id"`
	UniqueID string `gorm:"column:unique_id"`

	Code           string  `gorm:"column:code"`             // 代码
	Name           string  `gorm:"column:name"`             // 名字
	ReportPeriod   string  `gorm:"column:report_period"`    // 报告期
	OperateIn      int64   `gorm:"column:operate_in"`       // 营业收入
	OperateAllCost int64   `gorm:"column:operate_all_cost"` // 营业总成本
	OperateCost    int64   `gorm:"column:operate_cost"`     // 营业成本
	Tax            int64   `gorm:"column:tax"`              // 营业税金及附加
	SalesCost      int64   `gorm:"column:sales_cost"`       // 销售费用用
	ManageCost     int64   `gorm:"column:manage_cost"`      // 管理费用
	RDCost         int64   `gorm:"column:rd_cost"`          // 研发费用
	FinancialCost  int64   `gorm:"column:financial_cost"`   // 财务费用
	Invest         int64   `gorm:"column:invest"`           // 投资收益
	FairIn         int64   `gorm:"column:fair_in"`          // 公允价值变动收益
	NetProfit      int64   `gorm:"column:net_profit"`       // 净利润
	EarnPerShare   float64 `gorm:"column:earn_per_share"`   // 稀释每股收益

	CreatedAt int64 `gorm:"autoCreateTime:milli;column:created_at"` // milliseconds
	UpdatedAt int64 `gorm:"autoUpdateTime:milli;column:updated_at"` // milliseconds
}

func (vi *Profile) TableName() string {
	return "profile"
}

func (vi *Profile) Serialize() {
	vi.UniqueID = fmt.Sprintf("%s-%s", vi.Code, vi.ReportPeriod)
}

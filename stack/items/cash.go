package items

import "fmt"

// 现金流量表

type CashFlow struct {
	ID       int64  `gorm:"column:id"`
	UniqueID string `gorm:"column:unique_id"`

	Name         string `gorm:"column:name"`          // 名字
	Code         string `gorm:"column:code"`          // 代码
	ReportPeriod string `gorm:"column:report_period"` // 报告期
	SaleIn       int64  `gorm:"column:sale_in"`       // 销售商口流入
	TaxIn        int64  `gorm:"column:tax_in"`        // 税费返还
	SumIn        int64  `gorm:"column:sum_in"`        // 经营活动流入小计
	SaleOut      int64  `gorm:"column:sale_out"`      // 购买商口的流出
	EmpOut       int64  `gorm:"column:emp_out"`       // 支付给员工的流出
	SumOut       int64  `gorm:"column:sum_out"`       // 流出小计
	Netflow      int64  `gorm:"column:netflow"`       // 经营活动现金净额

	CreatedAt int64 `gorm:"autoCreateTime:milli;column:created_at"` // milliseconds
	UpdatedAt int64 `gorm:"autoUpdateTime:milli;column:updated_at"` // milliseconds
}

func (vi *CashFlow) TableName() string {
	return "cash_flow"
}

func (vi *CashFlow) Serialize() {
	vi.UniqueID = fmt.Sprintf("%s-%s", vi.Code, vi.ReportPeriod)
}

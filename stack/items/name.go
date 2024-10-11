package items

// 股票列表

type NameCode struct {
	ID   int64  `gorm:"column:id"`
	Name string `gorm:"column:name"` // 名字
	Code string `gorm:"column:code"` // 代码

	Profile    int64   `gorm:"profile"`
	CashFlow   int64   `gorm:"cashFlow"`
	Balance    int64   `gorm:"balance"`
	StockPrice float64 `gorm:"stockPrice"`
	CrawlDate  int64   `gorm:"crawlDate"`
	SHSZ       string  `gorm:"column:shsz" gorm:"SHSZ"`

	CreatedAt int64 `gorm:"autoCreateTime:milli;column:created_at"` // milliseconds
	UpdatedAt int64 `gorm:"autoUpdateTime:milli;column:updated_at"` // milliseconds
}

func (vi *NameCode) TableName() string {
	return "names"
}

func (vi *NameCode) Serialize() {
}

type Industry struct {
	Name string
	Code string
}

type NubSh struct {
	Result []DataSH `json:"result"`
}

type DataSH struct {
	Code string `json:"SECURITY_CODE_A"`
	Name string `json:"COMPANY_ABBR"`
}

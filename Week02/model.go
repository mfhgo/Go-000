package httpserver

type PayOrder struct {
	ID          uint64 `gorm:"column:id; primary_key; auto_increment; type:bigint(18); not null"`
	Platform    uint8  `gorm:"column:platform; index:idx_platform_orderid,idx_platform_userid; type:int(8); not null"`
	OrderId     string `gorm:"column:orderid; index:idx_platform_orderid; type:varchar(32); not null"`
	UserId      string `gorm:"column:userid; index:idx_platform_userid;type:varchar(32); not null"`
	RoleId      uint64 `gorm:"column:roleid; type:bigint(17); not null"`
	Channel     uint8  `gorm:"column:channel; type:int(8); not null; default:0"`
	ProductId   string `gorm:"column:productid; type:varchar(32); not null; default:''"`
	Amount      int32  `gorm:"column:amount; type:int(11); not null; default:0"`
	Timestamp   uint64 `gorm:"column:timestamp; type:bigint(17); not null; default:0"`
	Status      uint8  `gorm:"column:status; type:int(8); not null; default:0"`
	TimeReceive uint64 `gorm:"column:time_receive; type:datetime"`
	TimeFinish  uint64 `gorm:"column:time_finish; type:datetime"`
}

func CreateTablePayOrder() {
	t := DB.HasTable(PayOrder{})
	if !t {
		DB.CreateTable(PayOrder{})
	}
}

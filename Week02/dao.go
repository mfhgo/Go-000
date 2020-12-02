package httpserver

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
)

func (p *PayOrder) MysqlFirst() (bool, error) {
	if p.Channel <= 0 || p.OrderId == "" {
		return false, errors.Wrap(fmt.Errorf("channel <=0 or orderid empty"), "")
	}

	ret := DB.Where(&PayOrder{Channel: p.Channel, OrderId: p.OrderId}).First(&p)
	if errors.Is(ret.Error, gorm.ErrRecordNotFound) {
		return false, nil
	}
	if ret.Error != nil {
		return false, errors.Wrap(ret.Error, "")
	}
	return true, nil
}

//
func (p *PayOrder) MysqlCreate() error {
	ret := DB.Create(&p)
	if ret.Error != nil {
		return errors.Wrap(ret.Error, "")
	}
	return nil
}

//
func (p *PayOrder) MysqlUpdateStatus(status int32) error {
	ret := DB.Model(&p).Select("status").Updates(map[string]interface{}{"status": status})
	if ret.Error != nil {
		return errors.Wrap(ret.Error, "")
	}
	return nil
}

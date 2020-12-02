package httpserver

import (
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
)

func (p *PayOrder) IsOrderExist() (bool, error) {
	exist, err := p.MysqlFirst()
	if err != nil {
		by, _ := json.Marshal(p)
		return false, errors.WithMessagef(err, "order:%s", string(by))
	}
	return exist, err
}

func (p *PayOrder) AddOrder() error {
	err := p.MysqlCreate()
	if err != nil {
		by, _ := json.Marshal(p)
		return errors.WithMessagef(err, "order:%s", string(by))
	}
	return err
}

func (p *PayOrder) UpdateStatus() error {
	err := p.MysqlUpdateStatus(3)
	if err != nil {
		by, _ := json.Marshal(p)
		return errors.WithMessagef(err, "order:%s", string(by))
	}
	return err
}

func (p *PayOrder) Add() error {
	exist, err := p.IsOrderExist()
	if err != nil {
		return errors.WithMessage(err, "")
	}
	if exist {
		return errors.Wrapf(fmt.Errorf(""), "orderid:%s", p.OrderId)
	}

	err = p.AddOrder()
	if err != nil {
		return errors.WithMessage(err, "")
	}
	return nil
}

package service

import (
	"fmt"
	"github.com/pkg/errors"
	"learn/week2/dao"
)

var InfoPrice = errors.New("info price")

// CheckInfo 此处模拟service层
func CheckInfo() error {

	info, err := dao.GetInfo(100)
	if err != nil {
		return errors.Wrap(err, "service.get GetInfo error")
	}
	if info.Price < 200 {
		return fmt.Errorf("info.price < %v ，%w", info.Price, errors.New("info price"))
	}

	return nil
}

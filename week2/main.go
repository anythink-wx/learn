package main

import (
	"fmt"
	"github.com/pkg/errors"
	"learn/week2/g"
	"learn/week2/service"
)

// 问题
// 我们在数据库操作的时候，比如 dao 层中当遇到一个 sql.ErrNoRows 的时候，是否应该 Wrap 这个 error，抛给上层。
// 为什么，应该怎么做请写出代码？

// 我之前是做 php 的，好像不叫dao层，我们叫model层。我理解Dao是数据访问层，service 是业务层。
// Dao 的作用是封装对数据库的访问：增删改查，不涉及业务逻辑，只是达到按某个条件获得指定数据的要求。
// 既然不涉及业务逻辑，那么就说明这块的数据是由更上层进行组合使用的，因此，不应该wrap这个error，而是抛给上层service层处理完业务逻辑后
// 再由service层去wrap。

func init() {
	g.Init()
}
func main() {

	err := service.CheckInfo()

	if err != nil {
		fmt.Println("service error:", errors.Cause(err))
		fmt.Printf("---------\r\ntrace:::%+v", err)
	}

	if errors.As(err, &service.InfoPrice) {
		fmt.Println("错误类型是 service.InfoPrice:::", err)
	}

}

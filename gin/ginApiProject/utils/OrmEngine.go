package utils

import (
	"github.com/go-xorm/xorm"
)

var DbEngine *Orm

type Orm struct {
	*xorm.Engine
}

// OrmEngine 封装ORM
func OrmEngine(cfg *Config) (*Orm, error) {
	database := cfg.Database

	// 拼装数据库字符串
	conn := database.User + ":" + database.Password + "@tcp(" + database.Host + ":" + database.Port + ")/" + database.DbName + "?charset=" + database.Charset

	// 创建数据库操作引擎，连接数据库
	engine, err := xorm.NewEngine("mysql", conn)

	if err != nil {
		return nil, err
	}

	orm := new(Orm)
	orm.Engine = engine

	return orm, err
}

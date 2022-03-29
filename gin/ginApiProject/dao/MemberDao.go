package dao

import (
	"github.com/prometheus/common/log"
	"gogo_learn/gin/ginApiProject/model"
	"gogo_learn/gin/ginApiProject/utils"
)

type MemberDao struct {
	*utils.Orm
}

func (md *MemberDao) InsertCode(sms model.SmsCode) int64 {
	result, err := md.InsertOne(&sms)
	if err != nil {
		log.Error(err.Error())
	}
	return result
}

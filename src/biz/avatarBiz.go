package biz

import (
	"errors"
	"model"
)
//帖子管理中心
type AvatarBiz struct{

}

//获取所有图像
func(this *AvatarBiz) GetAllAvatar(status string) []model.AvatarModel{
	var avatarModel []model.AvatarModel
	err := GetDbInstance().Where("status = ?", status).Scan(&avatarModel).Error
	if err != nil {
		Debug("get avatar fail:", err.Error())
	}
	return avatarModel
}

//添加头像资源
func(this *AvatarBiz) AddAvatar(avatarUrl string) (string, error){
	var avatar model.AvatarModel
	err := GetDbInstance().Where("url = ?", avatarUrl).First(&avatar).Error
	if avatar.Id > 0 {
		return "图像资源已存在", err
	}
	avatar = model.AvatarModel{ Url: avatarUrl, Status: "A" }
	err = GetDbInstance().Create(&avatar).Error
	if err == nil{
		return "", nil
	}
	Debug("user Register: err: %v\n", err.Error())
	return "添加图像资源失败",errors.New("添加图像资源失败")
}

//删除头像资源
func(this *AvatarBiz) DeleteAvatar(avatar model.AvatarModel) bool{
	updateValue := map[string]interface{}{"status": "C"}
	err := GetDbInstance().Model(&avatar).Updates(updateValue).Error
	if err != nil {
		Debug("delete avatar failed:", err.Error())
		return false
	}
	return true
}
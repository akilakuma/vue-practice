package controller

import "hotel-server/model"

// GetRoomList 取得房間資料
func GetRoomList() []model.RoomStruct {
	return model.GetRoomList()
}
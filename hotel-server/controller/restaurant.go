package controller

import "hotel-server/model"

// GetRoomList 取得房間資料
func GetRestaurantList() []model.RestaurantStruct {
	return model.GetRestaurantList()
}
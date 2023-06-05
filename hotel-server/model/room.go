package model

type RoomStruct struct {
	ID int
	RoomName string
	Price int
	Size string
	Floor int 
	HasView bool
}


func GetRoomList() []RoomStruct{
	data :=[]RoomStruct{
		{
			ID: 1,
			RoomName : "標準舒適2人客房",
			Price : 3000,
			Size : "雙人房",
			Floor : 10,
		},
		{
			ID: 2,
			RoomName : "標準舒適4人客房",
			Price : 4200,
			Size : "四人房",
			Floor : 11,
		},
		{
			ID: 3,
			RoomName : "景致舒適2人客房",
			Price : 5000,
			Size : "雙人房",
			Floor : 10,
			HasView : true,
		},
		{
			ID: 4,
			RoomName : "景致舒適4人客房",
			Price : 6000,
			Size : "四人房",
			Floor : 11,
			HasView : true,
		},
		{
			ID: 5,
			RoomName : "豪華VIP客房",
			Price : 6600,
			Size : "雙人房",
			Floor : 12,
			HasView : true,
		},
		{
			ID: 6,
			RoomName : "頂級客房",
			Price : 10000,
			Size : "雙人房",
			Floor : 12,
			HasView : true,
		},
	}


	return data
}
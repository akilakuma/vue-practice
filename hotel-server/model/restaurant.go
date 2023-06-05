package model

type RestaurantStruct struct {
	ID int
	Name string
	Price int

}


func GetRestaurantList() []RestaurantStruct{
	data :=[]RestaurantStruct{
		{
			ID: 1,
			Name : "美味樓",
			Price : 3000,
		},
		{
			ID: 2,
			Name : "豪饗廳",
			Price : 2000,
		},
		{
			ID: 3,
			Name : "行政酒吧",
			Price : 800,
		},
		{
			ID: 4,
			Name : "親子餐廳",
			Price : 1000,
		},
	}


	return data
}
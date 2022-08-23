package home

type DisplayMenuStruct struct {
	ItemName  string  `json:"itemName"`
	ItemImage string  `json:"itemImage"`
	ItemPrice float32 `json:"itemPrice"`
}

func DisplayMenu(name string) *DisplayMenuStruct {

	displayMenu := &DisplayMenuStruct{
		ItemImage: "https://static.toiimg.com/photo/76425511.cms",
		ItemName:  "Samosa",
		ItemPrice: 30,
	}

	return displayMenu
}

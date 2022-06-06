package model

type Mobile struct {
	Manufacturer string `json: "manufacturer`
	Model        string `json: "model"`
	Form         string `json: "form"`
	Smartphone   string `json: "smartphone"`
	Year_        int16  `json: "year"`
	Units_sold_m int16  `json: "unit_sold"`
	Ids          int16  `json: "ids"`
}
type Mobiles struct {
	Mobiles []Mobile `json: "mobiles"`
}

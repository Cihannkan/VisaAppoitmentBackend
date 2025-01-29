package models

type Appoitment struct {
	Source_country   string `json:"source_country"`
	Mission_country  string `json:"mission_country"`
	Visa_type_id     int    `json:"visa_type_id"`
	Visa_category    string `json:"visa_category"`
	Visa_subcategory string `json:"visa_subcategory"`
	People_looking   int    `json:"people_looking"`
	Center_name      string `json:"center_name"`
	Appointment_date string `json:"appointment_date"`
	Book_now_link    string `json:"book_now_link"`
	Last_checked     string `json:"last_checked"`
}
type Source_country struct {
	Source_country string `json:"source_country"`
}

package model

// var WorkflowList = []Workflow{
// 	{ID: "1", Company: "Teknosa", Sector: "technology",
// 	 Project: "DSP", Technology: "adtech", Year: time.Now().Year()},
// 	{ID: "2", Company: "Trendyol", Sector: "ecommorce",
// 	Project: "SSP", Technology: "adtech", Year: time.Now().Year()},
// 	{ID: "3", Company: "Turkcell", Sector: "telecommunication",
// 	 Project: "DMP", Technology: "adtech", Year: time.Now().Year()},
// 	{ID: "4", Company: "Getir", Sector: "ecommorce",
// 	 Project: "RDD", Technology: "adtech", Year: time.Now().Year()},
// 	{ID: "5", Company: "Vodafone", Sector: "telecommunication",
// 	 Project: "SSP", Technology: "adtech", Year: time.Now().Year()},
// }

type Workflow struct{
	ID 				string		`json:"id" binding:"required"` 
	CompanyName 	string		`json:"name" binding:"required"`
	Sector			string		`json:"sector" binding:"required"`
	Project			string		`json:"projectname" binding:"required"`
	Technology		string		`json:"technology" binding:"required"`
	Year			int 		`json:"year" binding:"required"`
}

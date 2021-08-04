package models

type Vehicle struct {
	Id      uint64 `json:"id"`
	BrandId uint64 `json:"brand_id"`
	ModelId uint64 `json:"model_id"`
	Version string `gorm:"type:varchar(255)" json:"version"`
	Year    uint   `json:"year"`
	Fuel    string `gorm:"type:varchar(255)" json:"fuel"`
	Enabled bool   `json:"enabled"`
}

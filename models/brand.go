package models

type Brand struct {
	Id      uint64 `json:"id"`
	Name    string `gorm:"type:varchar(255)" json:"name"`
	Enabled bool   `json:"enabled"`
}

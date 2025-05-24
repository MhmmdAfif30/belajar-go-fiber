package entity

type About struct {
	ID          uint   `json:"id" gorm:"primaryKey"`
	UserID      uint   `json:"user_id" gorm:"unique"`
	DisplayName string `json:"displayname"`
	Gender      string `json:"gender" validate:"oneof=male female"`
	Birthday    string `json:"birthday"`
	Horoscope   string `json:"horoscope"`
	Zodiac      string `json:"zodiac"`
	Height      int    `json:"height"`
	Weight      int    `json:"weight"`

	User        User   `gorm:"foreignKey:UserID"`
}

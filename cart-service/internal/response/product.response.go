package response

type ProductResponse struct {
	Id          string  `json:"id"`
	Name        string  `json:"name"`
	Alias       string  `json:"alias"`
	Description string  `json:"description"`
	Status      bool    `json:"status"`
	Stock       int     `json:"stock"`
	Cost        float32 `json:"cost"`
	Price       float32 `json:"price"`
	Sale        float32 `json:"sale"`
	ImageID     string  `gorm:"column:image_id" json:"image_id"`
	URL         string  `json:"url"`

	ExtraImages []interface{} `json:"extra_images"`
	CategoryID  string        `json:"category_id"`
	CreatedAt   string        `json:"created_at"`
	UpdatedAt   string        `json:"updated_at"`
}

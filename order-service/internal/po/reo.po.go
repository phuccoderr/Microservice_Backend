package po

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Reo struct {
	gorm.Model
	ID uuid.UUID `gorm:"column:id; type:uuid; primaryKey; not null; autoIncrement"`
}

func (r *Reo) TableName() string {
	return "go_db_reo"
}

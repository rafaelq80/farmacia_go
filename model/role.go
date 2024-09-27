package model

type Role struct {
	ID      uint   		`gorm:"type:bigint;primaryKey;AUTO_INCREMENT" json:"id"`
	Role    string		`gorm:"type:varchar(255);not null" json:"role" validate:"required"`
	Usuario []Usuario 	`gorm:"foreignkey:RoleID;references:ID;constraint:OnDelete:CASCADE;" json:"usuario,omitempty"`
}

func (Role) TableName() string {
	return "tb_roles"
}
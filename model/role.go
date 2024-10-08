package model

type Role struct {
	ID      	uint   		`gorm:"primaryKey;autoIncrement" json:"id"`
	Role    	string		`gorm:"type:varchar(255);not null" json:"role" validate:"required"`
	Descricao   string		`gorm:"type:varchar(255);not null" json:"descricao" validate:"required"`
	Usuario 	[]Usuario 	`gorm:"foreignkey:RoleID;references:ID;constraint:OnDelete:CASCADE;" json:"usuario,omitempty"`
}

func (Role) TableName() string {
	return "tb_roles"
}
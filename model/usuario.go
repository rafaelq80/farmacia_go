package model

type Usuario struct {
	ID        uint       	`gorm:"type:bigint;primaryKey;AUTO_INCREMENT" json:"id"`
	Name      string     	`gorm:"type:varchar(255);not null" json:"name" validate:"required"`
	Usuario   string     	`gorm:"type:varchar(255);not null" json:"usuario" validate:"required,email"`
	Senha     string     	`gorm:"type:varchar(255);not null, min=8" json:"senha" validate:"required"`
	Foto      string     	`gorm:"type:varchar(5000)" json:"foto"`
	Produto	  []Produto 	`gorm:"foreignkey:UsuarioID;references:ID;constraint:OnDelete:CASCADE;" json:"produto,omitempty"`
}

func (Usuario) TableName() string {
	return "tb_usuarios"
}
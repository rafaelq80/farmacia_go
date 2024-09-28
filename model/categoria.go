package model

type Categoria struct {
	ID      	uint		`gorm:"primaryKey;autoIncrement" json:"id"`
	Grupo    	string  	`gorm:"type:varchar(255);not null" json:"grupo" validate:"required,min=3,max=255"`
	Produto		[]Produto	`gorm:"foreignkey:CategoriaID;references:ID;constraint:OnDelete:CASCADE;" json:"produto,omitempty"`
}

func (Categoria) TableName() string {
	return "tb_categorias"
}

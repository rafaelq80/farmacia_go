package model

type Produto struct {
	ID          uint       `gorm:"type:bigint;primaryKey;AUTO_INCREMENT" json:"id"`
	Nome        string     `gorm:"type:varchar(255);not null" json:"nome" validate:"required,min=3,max=255"`
	Preco       float32    `gorm:"type:decimal(8,2);not null" json:"preco" validate:"required"`
	Foto        string     `gorm:"type:varchar(5000)" json:"foto"`
	CategoriaID uint       `gorm:"column:categoria_id;not null" json:"categoria_id" validate:"required" example:"1"`
	Categoria   *Categoria `gorm:"ForeignKey:CategoriaID;association_foreignkey:ID" json:"categoria,omitempty" validate:"-"`
	UsuarioID   uint       `gorm:"column:usuario_id;not null" json:"usuario_id" validate:"required" example:"1"`
	Usuario     *Usuario   `gorm:"ForeignKey:UsuarioID;association_foreignkey:ID" json:"usuario,omitempty" validate:"-"`
}

func (Produto) TableName() string {
	return "tb_produtos"
}

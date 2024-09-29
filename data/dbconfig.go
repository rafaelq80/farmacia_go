package data

import (
	"log"
	"os"

	"github.com/rafaelq80/farmacia_go/model"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"

	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB
var err error

// Conexão com o Banco de dados local
func ConnectDB(connectionString string, database string, drop bool) {

	log.Println(connectionString)

	switch database {
	case "local":
		DB, err = gorm.Open(mysql.Open(connectionString), &gorm.Config{
			Logger: logger.Default.LogMode(logger.Info), //Exibir a Query SQL no Console
		})
	case "teste":
		DB, err = gorm.Open(sqlite.Open(connectionString), &gorm.Config{})
		case "remoto":
		DB, err = gorm.Open(postgres.Open(connectionString), &gorm.Config{})
	default:
		log.Println("Configuração inválida!")
	}

	if err != nil {
		log.Fatal("❌ Falha ao Conectar com o Banco de dados! \n", err.Error())
		os.Exit(1)
	}

	log.Println("🚀 Conexão com o Banco de dados efetuada com Sucesso!")

	if drop {
		//Deletar todas as tabelas
		DB.Migrator().DropTable(&model.Produto{}, &model.Categoria{}, &model.Usuario{})
	}

	// Lista de tabelas que deverão ser criadas
	models := []interface{}{
		&model.Produto{},
		&model.Categoria{},
		&model.Usuario{},
		&model.Role{},
	}

	// Verifica se todas as tabelas já foram criadas
	for _, model := range models {
		if !DB.Migrator().HasTable(model) {
			DB.AutoMigrate(model)
		} else {
			log.Printf("Tabela para o modelo %T já existe.\n", model)
		}
	}

	 // Adicionar dados na entidade Role caso esteja vazia
	 var count int64
	 DB.Model(&model.Role{}).Count(&count)
	 if count == 0 {
		 roles := []model.Role{
			 {Role: "admin", Descricao: "Administrador"},
			 {Role: "user", Descricao: "Usuário"},
		 }
		 for _, role := range roles {
			 DB.Create(&role)
		 }
		 log.Println("🚀 Dados iniciais adicionados à tabela Role.")
	 }

	 log.Println("🚀 Tabelas Configuradas com Sucesso!")
}

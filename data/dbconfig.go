package data

import (
	"log"
	"os"

	"github.com/rafaelq80/farmacia_go/model"
	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB
var err error

// Conexão com o Banco de dados da aplicação
func ConnectDB(connectionString string) {

	DB, err = gorm.Open(mysql.Open(connectionString), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info), //Exibir a Query SQL no Console
	})

	if err != nil {
		log.Fatal("❌ Falha ao Conectar com o Banco de dados! \n", err.Error())
		os.Exit(1)
	}

	log.Println("🚀 Conexão com o Banco de dados efetuada com Sucesso!")

	// Criar as tabelas
	DB.AutoMigrate(&model.Usuario{})
	DB.AutoMigrate(&model.Categoria{})
	DB.AutoMigrate(&model.Produto{})

	log.Println("🚀 Tabelas Configuradas com Sucesso!")

}

// Conexão com o Banco de dados de testes
func ConnectTestDB(testconnectionString string, dropTables bool) {

	DB, err = gorm.Open(sqlite.Open(testconnectionString), &gorm.Config{})

	if err != nil {
		log.Fatal("❌ Falha ao Conectar com o Banco de dados! \n", err.Error())
		os.Exit(1)
	}

	// Deletar todas as tabelas
	if dropTables {
		DB.Migrator().DropTable(&model.Produto{}, &model.Categoria{}, &model.Usuario{})
	}

	// Criar todas as tabelas
	DB.AutoMigrate(&model.Produto{}, &model.Categoria{}, &model.Usuario{})

}

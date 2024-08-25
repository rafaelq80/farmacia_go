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

// Conexão com o Banco de dados Remoto
func ConnectDBRender(connectionString string) {

	DB, err = gorm.Open(postgres.Open(connectionString), &gorm.Config{})

	if err != nil {
		log.Fatal("❌ Falha ao Conectar com o Banco de dados! \n", err.Error())
		os.Exit(1)
	}

	log.Println("🚀 Conexão com o Banco de dados efetuada com Sucesso!")

	// Criar as tabelas

	 models := []interface{}{
        &model.Produto{},
        &model.Categoria{},
        &model.Usuario{},
    }

    // Verifica se todas as tabelas existem
    for _, model := range models {
        if !DB.Migrator().HasTable(model) {
            // Se uma tabela não existir, ela será criada
            log.Printf("Tabela para o modelo %T não existe. Criando...\n", model)
            DB.AutoMigrate(model)
        } else {
            log.Printf("Tabela para o modelo %T já existe.\n", model)
        }
    }
	
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

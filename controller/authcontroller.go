package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rafaelq80/farmacia_go/data"
	"github.com/rafaelq80/farmacia_go/model"
	security "github.com/rafaelq80/farmacia_go/security/bcrypt"
	auth "github.com/rafaelq80/farmacia_go/security/service"

)

//	@Summary		Autenticar Usuario
//	@Description	Autentica um Usuario
//	@Tags			usuarios
//	@Accept			json
//	@Produce		json
//	@Param			usuario	body		model.UsuarioLogin	true	"Autenticar Usuario"
//	@Success		200		{object}	model.UsuarioLogin
//	@Failure		400		{object}	config.HTTPError
//	@Failure		401		{object}	config.HTTPError
//	@Failure		404		{object}	config.HTTPError
//	@Router			/usuarios/logar [post]
func AutenticarUsuario(c *fiber.Ctx) error {

	var usuarioLogin *model.UsuarioLogin
	var usuario *model.Usuario
	var errToken error
	var token string

	if errObjeto := c.BodyParser(&usuarioLogin); errObjeto != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "400", "message": errObjeto.Error()})
	}

	// Verifica se o usuário existe
	if !checkIfUsuarioEmailExists(usuarioLogin.Usuario) {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "404", "message": "Usuário não Encontrado!"})
	}

	data.DB.Where("usuario = ?", usuarioLogin.Usuario).First(&usuario) 

	if !security.CheckPasswordHash(usuarioLogin.Senha, usuario.Senha){
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "400", "message": "Senha Incorreta!"})
	}

	token, errToken = auth.CreateToken(usuarioLogin.Usuario)

	if errToken != nil{
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "400", "message": "Erro ao gerar o Token!"})
	}

	usuarioLogin.ID = usuario.ID
	usuarioLogin.Nome = usuario.Name
	usuarioLogin.Foto = usuario.Foto
	usuarioLogin.Senha = ""
	usuarioLogin.Token = "Bearer " + token

	return c.Status(fiber.StatusOK).JSON(&usuarioLogin)
}


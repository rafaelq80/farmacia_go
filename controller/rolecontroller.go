package controller

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/rafaelq80/farmacia_go/model"
	"github.com/rafaelq80/farmacia_go/service"
	"github.com/rafaelq80/farmacia_go/validator"
)

// Injeção de Dependências - RoleService
type RoleController struct {
	roleService *service.RoleService
}

// Método Construtor
func NewRoleController(roleService *service.RoleService) *RoleController {
	return &RoleController{roleService: roleService}
}

// @Summary		Listar Roles
// @Description	Lista todas as Roles
// @Tags		roles
// @Accept		json
// @Produce		json
// @Success		200				{array}		model.Role
// @Failure		401				{object}	config.HTTPError
// @Router		/roles 	[get]
// @Security	Bearer
func (roleController *RoleController) FindAll(c *fiber.Ctx) error {

	roles, _ := roleController.roleService.FindAll()

	return c.Status(fiber.StatusOK).JSON(roles)

}

// @Summary		Listar Role por id
// @Description	Lista uma Role por id
// @Tags			roles
// @Accept			json
// @Produce		json
// @Param			id					path		string	true	"Id da Role"
// @Success		200					{array}		model.Role
// @Failure		400					{object}	config.HTTPError
// @Failure		401					{object}	config.HTTPError
// @Failure		404					{object}	config.HTTPError
// @Router			/roles/{id} 	[get]
// @Security		Bearer
func (roleController *RoleController) FindById(context *fiber.Ctx) error {

	id := context.Params("id")

	role, err := roleController.roleService.FindById(id)

	if err != nil {
		return context.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status": "404", 
			"message": "Role não encontrada!",
		})
	}

	return context.Status(fiber.StatusOK).JSON(role)
}

// @Summary		Criar Role
// @Description	Cria uma nova Role
// @Tags			roles
// @Accept			json
// @Produce		json
// @Param			role	body		model.Role	true	"Criar Role"
// @Success		201			{object}	model.Role
// @Failure		400			{object}	config.HTTPError
// @Failure		401			{object}	config.HTTPError
// @Router			/roles [post]
// @Security		Bearer
func (roleController *RoleController) Create(context *fiber.Ctx) error {

	var role model.Role

	if err := context.BodyParser(&role); err != nil {
		return context.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status": "400", 
			"message": err.Error(),
		})
	}

	if err := validator.ValidateStruct(&role); err != nil {
		return context.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status": "400", 
			"message": err,
		})
	}

	// Verifica se o role existe (evitar duplicidade)
	if roleController.roleService.RoleExists(role.Role) {
		return context.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "400",
			"message": "Role já cadastrado!",
		})
	}

	if err := roleController.roleService.Create(&role); err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status": "500", 
			"message": err.Error(),
		})
	}

	return context.Status(fiber.StatusCreated).JSON(&role)

}

// @Summary		Atualizar Role
// @Description	Edita uma Role
// @Tags			roles
// @Accept			json
// @Produce		json
// @Param			Role	body		model.Role	true	"Atualizar Role"
// @Success		200			{object}	model.Role
// @Failure		400			{object}	config.HTTPError
// @Failure		401			{object}	config.HTTPError
// @Failure		404			{object}	config.HTTPError
// @Router			/roles [put]
// @Security		Bearer
func (roleController *RoleController) Update(context *fiber.Ctx) error {

	var role model.Role

	if err := context.BodyParser(&role); err != nil {
		return context.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status": "400", 
			"message": err.Error(),
		})
	}

	if err := validator.ValidateStruct(&role); err != nil {
		return context.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status": "400", 
			"message": err,
		})
	}

	id := strconv.FormatUint(uint64(role.ID), 10)

	roleExist, _ := roleController.roleService.ExistsById(id)

	if  !roleExist {
		return context.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status": "404", 
			"message": "Role não encontrada!",
		})
	}

	if err := roleController.roleService.Update(&role); err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status": "500", 
			"message": err.Error(),
		})
	}

	return context.Status(fiber.StatusOK).JSON(&role)

}


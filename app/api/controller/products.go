package controller

import (
	"github.com/allanfvc/inventory/api/business"
	"github.com/allanfvc/inventory/api/models"
	"github.com/gofiber/fiber"
)

/*ProductController handle the product controller
*	@createdAt 01/08/2020
* @createdBy allanfvc
 */
type ProductController struct {
}

func (c *ProductController) setupRoutes(app fiber.Router) {
	controller := "/products"
	app.Get(controller, c.listProducts)
	app.Post(controller, c.saveProduct)
	app.Put(controller+"/:id", c.updateProduct)
	app.Delete(controller+"/:id", c.deleteProduct)
}

func (c *ProductController) saveProduct(ctx *fiber.Ctx) {
	product := new(models.Product)
	if err := ctx.BodyParser(product); err != nil {
		ctx.Status(503).Send(err)
		return
	}
	business.SaveProduct(product)
	ctx.JSON(product)
}

func (c *ProductController) listProducts(ctx *fiber.Ctx) {
	name := ctx.Query("name")
	patients := business.ListProducts(name)
	ctx.JSON(patients)
}

func (c *ProductController) updateProduct(ctx *fiber.Ctx) {
	book := new(models.Product)
	if err := ctx.BodyParser(book); err != nil {
		ctx.Status(503).Send(err)
		return
	}
	ctx.JSON(book)
}

func (c *ProductController) deleteProduct(ctx *fiber.Ctx) {
	id := c.Params("id")
	if err := ctx.BodyParser(book); err != nil {
		ctx.Status(503).Send(err)
		return
	}
	ctx.JSON(book)
}

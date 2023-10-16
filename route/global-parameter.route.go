package route

import (
	"github.com/gofiber/fiber/v2"
	"github.com/tedbearr/go-learn/controller"
	"github.com/tedbearr/go-learn/database"
	"github.com/tedbearr/go-learn/service"
	"gorm.io/gorm"
)

var (
	db                        *gorm.DB                             = database.DatabaseInit()
	globalParameterService    service.GlobalParameterService       = service.NewGlobalParameterService(db)
	globalParameterController controller.GlobalParameterController = controller.NewGlobalParameterController(globalParameterService)
)

func GlobalParameterRoute(route fiber.Router) {
	group := route.Group("/global-parameter")
	group.Get("/", globalParameterController.All)
}

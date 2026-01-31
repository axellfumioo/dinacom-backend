package router

import (
	"backend-dinakom/app/controllers"
	"backend-dinakom/app/middlewares"
	"backend-dinakom/app/services"
	"backend-dinakom/database"

	"github.com/gofiber/fiber/v2"
)

func MemberRoute(r fiber.Router) {
	memberService := services.NewMemberService(database.DB)
	memberController := controllers.NewMemberController(memberService)
	members := r.Group("members")
	members.Use(middlewares.AuthMiddleware())
	members.Post("/", memberController.AddFamilyMembers)
	members.Get("/family/:familyID", memberController.GetFamilyMembers)
	members.Get("/statistics/family/:familyID", memberController.GetAllMemberStatistics)
	members.Delete("/:ID/family/:familyID", memberController.DeleteFamilyMember)
}

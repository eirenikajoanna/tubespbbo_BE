package router

import (
	"tubespbbo/modules/controller"

	"github.com/gofiber/fiber/v2"
)

func RouteExpenses(e *fiber.App) {
	e.Get("/expenses", controller.FindExpense)
	e.Get("/expenses/:id", controller.FindOneExpense)
	e.Post("/expenses", controller.CreateExpense)
	e.Put("/expenses/:id", controller.UpdateExpense)
	e.Delete("/expenses/:id", controller.DeleteExpense)
}

func RoutePayments(e *fiber.App) {
	e.Get("/payments", controller.FindPayment)
	e.Get("/payments/:id", controller.FindOnePayment)
	e.Post("/payments", controller.CreatePayment)
	e.Put("/payments/:id", controller.UpdatePayment)
	e.Delete("/payments/:id", controller.DeletePayment)
}

func RoutePaymentMethods(e *fiber.App) {
	e.Get("/payment-methods", controller.FindPaymentMethod)
	e.Get("/payment-methods/:id", controller.FindOnePaymentMethod)
	e.Post("/payment-methods", controller.CreatePaymentMethod)
	e.Put("/payment-methods/:id", controller.UpdatePaymentMethod)
	e.Delete("/payment-methods/:id", controller.DeletePaymentMethod)
}

func RouteProducts(e *fiber.App) {
	e.Get("/products", controller.FindProduct)
	e.Get("/products/:id", controller.FindOneProduct)
	e.Post("/products", controller.CreateProduct)
	e.Put("/products/:id", controller.UpdateProduct)
	e.Delete("/products/:id", controller.DeleteProduct)
}

func RouteTransactions(e *fiber.App) {
	e.Get("/transactions", controller.FindTransaction)
	e.Get("/transactions/:id", controller.FindOneTransaction)
	e.Post("/transactions", controller.CreateTransaction)
	e.Put("/transactions/:id", controller.UpdateTransaction)
	e.Delete("/transactions/:id", controller.DeleteTransaction)
}

func RouteTransactionMethods(e *fiber.App) {
	e.Get("/transaction-details", controller.FindTransactionDetail)
	e.Get("/transaction-details/:id", controller.FindOneTransactionDetail)
	e.Post("/transaction-details", controller.CreateTransactionDetail)
	e.Put("/transaction-details/:id", controller.UpdateTransactionDetail)
	e.Delete("/transaction-details/:id", controller.DeleteTransactionDetail)
}

func RouteUsers(e *fiber.App) {
	e.Get("/users", controller.FindUser)
	e.Get("/users/:id", controller.FindOneUser)
	e.Post("/users", controller.CreateUser)
	e.Put("/users/:id", controller.UpdateUser)
	e.Delete("/users/:id", controller.DeleteUser)

	e.Post("/login", controller.Login)
}

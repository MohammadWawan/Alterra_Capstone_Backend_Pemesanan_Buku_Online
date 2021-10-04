package routes

import (
	"alterra/controllers/books"
	"alterra/controllers/categories"
	"alterra/controllers/descriptions"
	"alterra/controllers/karyawans"
	"alterra/controllers/payment_methods"
	transactiondetails "alterra/controllers/transaction_details"
	"alterra/controllers/transactions"
	"alterra/controllers/users"
	"alterra/controllers/wishlists"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type ControllerList struct {
	JWTMiddleware                middleware.JWTConfig
	UserController               users.UserController
	KaryawanController           karyawans.KaryawanController
	BookController               books.BookController
	CategoryController           categories.CategoryController
	DescriptionController        descriptions.DescriptionController
	Payment_MethodController     payment_methods.Payment_MethodController
	WishlistController           wishlists.WishlistController
	Transaction_DetailController transactiondetails.Transaction_DetailController
	TransactionController        transactions.TransactionController
}

func (cl *ControllerList) RouteRegister(e *echo.Echo) {
	//user
	e.POST("users/login", cl.UserController.Login)
	e.POST("users/register", cl.UserController.Register, middleware.JWTWithConfig(cl.JWTMiddleware))
	e.GET("users", cl.UserController.GetUsers)
	e.GET("users/:id", cl.UserController.GetDetailUser)
	e.DELETE("users/:id", cl.UserController.DeleteUser)
	e.PUT("users/:id", cl.UserController.UpdateUser)

	//karyawan
	e.POST("karyawans/login", cl.KaryawanController.Login)
	e.POST("karyawans/register", cl.KaryawanController.Register, middleware.JWTWithConfig(cl.JWTMiddleware))
	e.GET("karyawans", cl.KaryawanController.GetKaryawans)
	e.GET("karyawans/:id", cl.KaryawanController.GetDetailKaryawan)
	e.DELETE("karyawans/:id", cl.KaryawanController.DeleteKaryawan)
	e.PUT("karyawans/:id", cl.KaryawanController.UpdateKaryawan)

	//book
	e.GET("books", cl.BookController.GetBooks)
	e.GET("books/:id", cl.BookController.GetBookById)
	e.POST("books/insertbook", cl.BookController.InsertBook)
	e.PUT("books/:id", cl.BookController.UpdateBook)
	e.DELETE("books/:id", cl.BookController.DeleteBook)

	//categories
	e.GET("categories", cl.CategoryController.GetCategories)
	e.GET("categories/:id", cl.CategoryController.GetCategoryById)
	e.POST("categories/insertcategory", cl.CategoryController.InsertCategory)
	e.PUT("categories/:id", cl.CategoryController.UpdateCategory)
	e.DELETE("categories/:id", cl.CategoryController.DeleteCategory)

	//descriptions
	e.GET("descriptions", cl.DescriptionController.GetDescriptions)
	e.GET("descriptions/:id", cl.DescriptionController.GetDescriptionById)
	e.POST("descriptions/insertdescription", cl.DescriptionController.InsertDescription)
	e.PUT("descriptions/:id", cl.DescriptionController.UpdateDescription)
	e.DELETE("descriptions/:id", cl.DescriptionController.DeleteDescription)

	//payment methods
	e.GET("payment_methods", cl.Payment_MethodController.GetPayment_Methods)
	e.GET("payment_methods/:id", cl.Payment_MethodController.GetPayment_MethodById)
	e.POST("payment_methods/insertpayment_method", cl.Payment_MethodController.InsertPayment_Method)
	e.PUT("payment_methods/:id", cl.Payment_MethodController.UpdatePayment_Method)
	e.DELETE("payment_methods/:id", cl.Payment_MethodController.DeletePayment_Method)

	//wishlists
	e.GET("wishlists", cl.WishlistController.GetWishlists)
	e.GET("wishlists/:id", cl.WishlistController.GetWishlistById)
	e.POST("wishlists/insertwishlist", cl.WishlistController.InsertWishlist)
	e.PUT("wishlists/:id", cl.WishlistController.UpdateWishlist)
	e.DELETE("wishlists/:id", cl.WishlistController.DeleteWishlist)

	//transaction_details
	e.GET("transaction_details", cl.Transaction_DetailController.GetTransaction_Details)
	e.GET("transaction_details/:id", cl.Transaction_DetailController.GetTransaction_DetailsById)
	e.POST("transaction_details/inserttransaction_detail", cl.Transaction_DetailController.InsertTransaction_Detail)
	e.PUT("transaction_details/:id", cl.Transaction_DetailController.UpdateTransaction_Detail)
	e.DELETE("transaction_details/:id", cl.Transaction_DetailController.DeleteTransaction_Details)

	//transactions
	e.GET("transaction", cl.TransactionController.GetTransactions)
	e.GET("transaction/:id", cl.TransactionController.GetTransactionById)
	e.POST("transaction/inserttransaction", cl.TransactionController.InsertTransaction)
	e.PUT("transaction/:id", cl.TransactionController.UpdateTransaction)
	e.DELETE("transaction/:id", cl.TransactionController.DeleteTransactions)
}

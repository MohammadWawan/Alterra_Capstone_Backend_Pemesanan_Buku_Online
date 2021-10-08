package main

import (
	"alterra/app/middlewares"
	"alterra/app/routes"

	_karyawanUsecase "alterra/business/karyawans"
	_karyawanController "alterra/controllers/karyawans"
	"alterra/drivers/databases/karyawans"

	_bookUsecase "alterra/business/books"
	_bookController "alterra/controllers/books"
	"alterra/drivers/databases/books"

	_userUsecase "alterra/business/users"
	_userController "alterra/controllers/users"
	"alterra/drivers/databases/users"

	_categoryUsecase "alterra/business/categories"
	_categoryController "alterra/controllers/categories"
	"alterra/drivers/databases/categories"

	_descriptionUsecase "alterra/business/descriptions"
	_descriptionController "alterra/controllers/descriptions"
	"alterra/drivers/databases/descriptions"

	_payment_methodUsecase "alterra/business/payment_methods"
	_payment_methodController "alterra/controllers/payment_methods"
	"alterra/drivers/databases/payment_methods"

	_wishlistUsecase "alterra/business/wishlists"
	_wishlistController "alterra/controllers/wishlists"
	"alterra/drivers/databases/wishlists"

	_transaction_detailUsecase "alterra/business/transaction_details"
	_transaction_detailController "alterra/controllers/transaction_details"
	"alterra/drivers/databases/transaction_details"

	_transactionUsecase "alterra/business/transactions"
	_transactionController "alterra/controllers/transactions"
	"alterra/drivers/databases/transactions"

	_mysqlDriver "alterra/drivers/mysql"
	"log"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

func init() {
	viper.SetConfigFile(`config.json`)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	if viper.GetBool(`debug`) {
		log.Println("Service RUN on DEBUG mode")
	}
}

func DbMigrate(db *gorm.DB) {
	db.AutoMigrate(&users.Users{})
	db.AutoMigrate(&karyawans.Karyawans{})
	db.AutoMigrate(&books.Books{})
	db.AutoMigrate(&categories.Categories{})
	db.AutoMigrate(&descriptions.Descriptions{})
	db.AutoMigrate(&payment_methods.Payment_Method{})
	db.AutoMigrate(&wishlists.Wishlist{})
	db.AutoMigrate(&transaction_details.Transaction_Detail{})
	db.AutoMigrate(&transactions.Transaction{})
}

func main() {
	// init koneksi database
	configDB := _mysqlDriver.ConfigDB{
		DB_Username: viper.GetString(`database.user`),
		DB_Password: viper.GetString(`database.pass`),
		DB_Host:     viper.GetString(`database.host`),
		DB_Port:     viper.GetString(`database.port`),
		DB_Database: viper.GetString(`database.name`),
	}
	configJWT := middlewares.ConfigJWT{
		SecretJWT:       viper.GetString(`jwt.secret`),
		ExpiresDuration: viper.GetInt(`jwt.expired`),
	}

	Conn := configDB.InitialDB()
	DbMigrate(Conn)

	e := echo.New()
	timeoutContext := time.Duration(viper.GetInt("context.timeout")) * time.Second

	userRepository := users.NewMysqlUserRepository(Conn)
	userUseCase := _userUsecase.NewUserUsecase(userRepository, timeoutContext, &configJWT)
	userController := _userController.NewUserController(userUseCase)

	karyawanRepository := karyawans.NewMysqlKaryawanRepository(Conn)
	karyawanUseCase := _karyawanUsecase.NewKaryawanUsecase(karyawanRepository, timeoutContext, &configJWT)
	karyawanController := _karyawanController.NewKaryawanController(karyawanUseCase)

	bookRepository := books.NewMysqlBookRepository(Conn)
	bookUseCase := _bookUsecase.NewBookUsecase(bookRepository, timeoutContext)
	bookcontroller := _bookController.NewBookController(bookUseCase)

	categoryRepository := categories.NewMysqlCategoryRepository(Conn)
	categoryUseCase := _categoryUsecase.NewCategoryUsecase(categoryRepository, timeoutContext)
	categorycontroller := _categoryController.NewCategoryController(categoryUseCase)

	descriptionRepository := descriptions.NewMysqlDescriptionRepository(Conn)
	descriptionUseCase := _descriptionUsecase.NewDescriptionUsecase(descriptionRepository, timeoutContext)
	descriptioncontroller := _descriptionController.NewDescriptionController(descriptionUseCase)

	payment_methodRepository := payment_methods.NewMysqlPayment_MethodRepository(Conn)
	payment_methodUseCase := _payment_methodUsecase.NewPayment_MethodUsecase(payment_methodRepository, timeoutContext)
	payment_methodcontroller := _payment_methodController.NewPayment_MethodController(payment_methodUseCase)

	wishlistRepository := wishlists.NewMysqlWishlistRepository(Conn)
	wishlistUseCase := _wishlistUsecase.NewWishlistUsecase(wishlistRepository, timeoutContext)
	wishlistcontroller := _wishlistController.NewWishlistController(wishlistUseCase)

	transaction_detailRepository := transaction_details.NewMysqlTransaction_DetailRepository(Conn)
	transaction_detailUseCase := _transaction_detailUsecase.NewTransaction_DetailUsecase(transaction_detailRepository, timeoutContext)
	transaction_detailcontroller := _transaction_detailController.NewTransaction_DetailController(transaction_detailUseCase)

	transactionRepository := transactions.NewMysqlTransactionRepository(Conn)
	transactionUseCase := _transactionUsecase.NewTransactionUsecase(transactionRepository, timeoutContext)
	transactioncontroller := _transactionController.NewTransactionController(transactionUseCase.Repo)

	routesInit := routes.ControllerList{
		UserController:               *userController,
		WishlistController:           *wishlistcontroller,
		Transaction_DetailController: *transaction_detailcontroller,
		TransactionController:        *transactioncontroller,
		KaryawanController:           *karyawanController,
		BookController:               *bookcontroller,
		CategoryController:           *categorycontroller,
		DescriptionController:        *descriptioncontroller,
		Payment_MethodController:     *payment_methodcontroller,
		JWTMiddleware:                configJWT.Init(),
	}

	routesInit.RouteRegister(e)
	log.Fatal(e.Start(viper.GetString("server.address")))
}

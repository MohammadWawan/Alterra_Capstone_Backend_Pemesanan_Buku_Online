package main

import (
	"alterra/app/middlewares"
	"alterra/app/routes"

	_karyawanUsecase "alterra/business/karyawans"
	_karyawanController "alterra/controllers/karyawans"
	_karyawanRepository "alterra/drivers/databases/karyawans"
	_karyawandb "alterra/drivers/databases/karyawans"

	_bookUsecase "alterra/business/books"
	_bookController "alterra/controllers/books"
	_bookRepository "alterra/drivers/databases/books"
	_bookdb "alterra/drivers/databases/books"

	_userUsecase "alterra/business/users"
	_userController "alterra/controllers/users"
	_userRepository "alterra/drivers/databases/users"
	_userdb "alterra/drivers/databases/users"

	_categoryUsecase "alterra/business/categories"
	_categoryController "alterra/controllers/categories"
	_categoryRepository "alterra/drivers/databases/categories"
	_categorydb "alterra/drivers/databases/categories"

	_descriptionUsecase "alterra/business/descriptions"
	_descriptionController "alterra/controllers/descriptions"
	_descriptionRepository "alterra/drivers/databases/descriptions"
	_descriptiondb "alterra/drivers/databases/descriptions"

	_payment_methodUsecase "alterra/business/payment_methods"
	_payment_methodController "alterra/controllers/payment_methods"
	_payment_methodRepository "alterra/drivers/databases/payment_methods"
	_payment_methoddb "alterra/drivers/databases/payment_methods"

	_wishlistUsecase "alterra/business/wishlists"
	_wishlistController "alterra/controllers/wishlists"
	_wishlistRepository "alterra/drivers/databases/wishlists"
	_wishlistdb "alterra/drivers/databases/wishlists"

	_transaction_detailUsecase "alterra/business/transaction_details"
	_transaction_detailController "alterra/controllers/transaction_details"
	_transaction_detailRepository "alterra/drivers/databases/transaction_details"
	_transaction_detaildb "alterra/drivers/databases/transaction_details"

	_transactionUsecase "alterra/business/transactions"
	_transactionController "alterra/controllers/transactions"
	_transactionRepository "alterra/drivers/databases/transactions"
	_transactiondb "alterra/drivers/databases/transactions"

	_mysqlDriver "alterra/drivers/mysql"
	"log"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

func init() {
	viper.SetConfigFile(`app/config.json`)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	if viper.GetBool(`debug`) {
		log.Println("Service RUN on DEBUG mode")
	}
}

func DbMigrate(db *gorm.DB) {
	db.AutoMigrate(&_userdb.Users{})
	db.AutoMigrate(&_karyawandb.Karyawans{})
	db.AutoMigrate(&_bookdb.Books{})
	db.AutoMigrate(&_categorydb.Categories{})
	db.AutoMigrate(&_descriptiondb.Descriptions{})
	db.AutoMigrate(&_payment_methoddb.Payment_Method{})
	db.AutoMigrate(&_wishlistdb.Wishlist{})
	db.AutoMigrate(&_transaction_detaildb.Transaction_Detail{})
	db.AutoMigrate(&_transactiondb.Transaction{})
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

	userRepository := _userRepository.NewMysqlUserRepository(Conn)
	userUseCase := _userUsecase.NewUserUsecase(userRepository, timeoutContext, &configJWT)
	userController := _userController.NewUserController(userUseCase)

	karyawanRepository := _karyawanRepository.NewMysqlKaryawanRepository(Conn)
	karyawanUseCase := _karyawanUsecase.NewKaryawanUsecase(karyawanRepository, timeoutContext, &configJWT)
	karyawanController := _karyawanController.NewKaryawanController(karyawanUseCase)

	bookRepository := _bookRepository.NewMysqlBookRepository(Conn)
	bookUseCase := _bookUsecase.NewBookUsecase(bookRepository, timeoutContext)
	bookcontroller := _bookController.NewBookController(bookUseCase)

	categoryRepository := _categoryRepository.NewMysqlCategoryRepository(Conn)
	categoryUseCase := _categoryUsecase.NewCategoryUsecase(categoryRepository, timeoutContext)
	categorycontroller := _categoryController.NewCategoryController(categoryUseCase)

	descriptionRepository := _descriptionRepository.NewMysqlDescriptionRepository(Conn)
	descriptionUseCase := _descriptionUsecase.NewDescriptionUsecase(descriptionRepository, timeoutContext)
	descriptioncontroller := _descriptionController.NewDescriptionController(descriptionUseCase)

	payment_methodRepository := _payment_methodRepository.NewMysqlPayment_MethodRepository(Conn)
	payment_methodUseCase := _payment_methodUsecase.NewPayment_MethodUsecase(payment_methodRepository, timeoutContext)
	payment_methodcontroller := _payment_methodController.NewPayment_MethodController(payment_methodUseCase)

	wishlistRepository := _wishlistRepository.NewMysqlWishlistRepository(Conn)
	wishlistUseCase := _wishlistUsecase.NewWishlistUsecase(wishlistRepository, timeoutContext)
	wishlistcontroller := _wishlistController.NewWishlistController(wishlistUseCase)

	transaction_detailRepository := _transaction_detailRepository.NewMysqlTransaction_DetailRepository(Conn)
	transaction_detailUseCase := _transaction_detailUsecase.NewTransaction_DetailUsecase(transaction_detailRepository, timeoutContext)
	transaction_detailcontroller := _transaction_detailController.NewTransaction_DetailController(transaction_detailUseCase)

	transactionRepository := _transactionRepository.NewMysqlTransactionRepository(Conn)
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

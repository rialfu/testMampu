package providers

import (
	"rialfu/wallet/config"
	authController "rialfu/wallet/modules/auth/controller"
	authService "rialfu/wallet/modules/auth/service"
	userController "rialfu/wallet/modules/user/controller"
	userRepo "rialfu/wallet/modules/user/repository"
	userService "rialfu/wallet/modules/user/service"
	walletController "rialfu/wallet/modules/wallet/controller"
	walletRepo "rialfu/wallet/modules/wallet/repository"
	walletService "rialfu/wallet/modules/wallet/service"
	"rialfu/wallet/pkg/constants"

	"github.com/samber/do"
	"gorm.io/gorm"
)

func InitDatabase(injector *do.Injector) {
	do.ProvideNamed(injector, constants.DB, func(i *do.Injector) (*gorm.DB, error) {
		return config.SetUpTestDatabaseConnection(), nil
	})
}

func RegisterDependencies(injector *do.Injector) {
	InitDatabase(injector)

	do.ProvideNamed(injector, constants.JWTService, func(i *do.Injector) (authService.JWTService, error) {
		return authService.NewJWTService(), nil
	})

	db := do.MustInvokeNamed[*gorm.DB](injector, constants.DB)
	jwtService := do.MustInvokeNamed[authService.JWTService](injector, constants.JWTService)

	userRepository := userRepo.NewUserRepository(db)
	walletRepository := walletRepo.NewWalletRepository(db)
	transactionRepository := walletRepo.NewTransactionRepository(db)
	masterBankRepository := walletRepo.NewMasterBankRepositry(db)
	walletLedgerRepository := walletRepo.NewWalletLedgerRepository(db)
	withdrawRepository := walletRepo.NewWithdrawRepository(db)
	depositRepository := walletRepo.NewDepositRepository(db)
	// refreshTokenRepository := authRepo.NewRefreshTokenRepository(db)

	userService := userService.NewUserService(userRepository, db)
	authService := authService.NewAuthService(userRepository, walletRepository, jwtService, db)
	walletService := walletService.NewUserService(
		walletRepository, transactionRepository, masterBankRepository,
		walletLedgerRepository, withdrawRepository, depositRepository, db)

	do.Provide(injector, func(i *do.Injector) (userController.UserController, error) {
		return userController.NewUserController(i, userService), nil
	})

	do.Provide(injector, func(i *do.Injector) (authController.AuthController, error) {
		return authController.NewAuthController(i, authService), nil
	})

	do.Provide(injector, func(i *do.Injector) (walletController.WalletController, error) {
		return walletController.NewUserController(i, walletService), nil
	})
}

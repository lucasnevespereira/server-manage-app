package services

import (
	"log"
	"server-manage-api/configs"
	"server-manage-api/internal/repository"
)

type Services struct {
	ServerManagerService ServerManagerService
}

func InitServices(config configs.Config) *Services {
	serverRepository, err := repository.NewServerRepository(repository.Config{
		DbHost:     config.DbHost,
		DbPort:     config.DbPort,
		DbUser:     config.DbUser,
		DbPassword: config.DbPassword,
		DbName:     config.DbName,
		DbSsl:      config.DbSsl,
	})
	if err != nil {
		log.Printf("could not init serverRepository: %v \n", err)
	}
	err = serverRepository.AutoMigrate()
	if err != nil {
		log.Printf("could not auto migrate serverRepository: %v \n", err)
	}

	return &Services{
		ServerManagerService: NewServerManager(serverRepository),
	}

}

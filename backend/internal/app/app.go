package app

import (
	"fmt"
	"neurotech-assignment/backend/internal/config"
	"neurotech-assignment/backend/internal/repo"
	"neurotech-assignment/backend/internal/server"
	"neurotech-assignment/backend/internal/storage"
	"neurotech-assignment/backend/internal/utils"
	"os"
)

func Run() {
	cfg, err := config.Init()
	if err != nil {
		fmt.Printf("app.Run failed to init config: %s\n", err.Error())
		os.Exit(1)
	}

	logger, err := utils.NewLogger()
	if err != nil {
		fmt.Printf("app.Run failed to init logger: %s\n", err.Error())
		os.Exit(1)
	}

	logger.Info("initializing storages...")
	var fileStorage repo.Storage
	if cfg.FileStorage.IsEncrypted {
		fileStorage, err = storage.NewEncryptedFileStorage(cfg.FileStorage.Path, cfg.FileStorage.Secret)
		if err != nil {
			fmt.Printf("app.Run failed to init file storage: %s\n", err.Error())
			os.Exit(1)
		}
	} else {
		fileStorage = storage.NewFileStorage(cfg.FileStorage.Path)
	}
	guidGenerator := utils.NewGUIDGenerator()

	logger.Info("initializing repo...")
	repo := repo.NewPatientRepo(fileStorage, guidGenerator)

	logger.Info("initializing server...")
	server := server.NewServer(cfg.HTTPServer, repo)

	if err := server.Serve(); err != nil {
		fmt.Printf("app.Run failed to run server: %s\n", err.Error())
		os.Exit(1)
	}

	fmt.Println("app.Run server stopped successfully")
}

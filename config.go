package main

import (
	"log"
	"os"

	"github.com/caarlos0/env"
	"github.com/joho/godotenv"
)

// UploadServerConfig contains settings that control the behavior of the UploadServer
type UploadServerConfig struct {
	ListenAddr         string `env:"LISTEN_ADDR" envDefault:"127.0.0.1:8088"`
	StoragePath        string `env:"STORAGE_PATH" envDefault:"./uploads"`
	StorageShardLayers int    `env:"STORAGE_SHARD_LAYERS" envDefault:"6"`
	DBPath             string `env:"DATABASE_PATH" envDefault:"./uploads.db"`
}

// LoadFromEnv populates the config from the process environment and .env file
func (cfg *UploadServerConfig) LoadFromEnv() {
	// load values from .env file
	err := godotenv.Load()
	if err != nil {
		if _, ok := err.(*os.PathError); ok {
			log.Println("no .env file loaded")
		} else {
			log.Fatalf("%#v", err)
		}
	}

	// populate config struct
	err = env.Parse(cfg)
	if err != nil {
		log.Fatal(err)
	}
}
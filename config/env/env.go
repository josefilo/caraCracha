package config

import "github.com/joho/godotenv"

err := godotenv.Load()
if err != nil {
	logger.Error("Error loading .env file",
		zap.Any("error", err))
}

type MongodbConnection struct {
	HOST string
	PORT string
	USER string
	PASSWORD string
	DB   string
}

func getMongodbConnection() MongodbConnection {
	return MongodbConnection{
		HOST: os.Getenv("MONGODB_HOST"),
		PORT: os.Getenv("MONGODB_PORT"),
		USER: os.Getenv("MONGODB_USER"),
		PASSWORD: os.Getenv("MONGODB_PASS"),
		DB:   os.Getenv("MONGODB_DB"),
	}
}

func getMongodbConnectionString() string {
	connection := getMongodbConnection()
	return fmt.Sprintf("mongodb://%s:%s@%s:%s/%s",
		connection.USER,
		connection.PASSWORD,
		connection.HOST,
		connection.PORT,
		connection.DB)
}





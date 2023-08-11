package config

type ConfigStruct struct {
	Production          bool   `env:"PRODUCTION" env-description:"dev or prod" env-default:"true"`
	LogDebug            bool   `env:"LOG_DEBUG" env-description:"should log at debug level" env-default:"false"`
	Port                int32  `env:"PORT" env-description:"server port" env-default:"50051"`
	Firebase_Project_Id string `env:"FIREBASE_PROJECT_ID" env-description:"firebase project id" env-default:"NO_PROJECT"`
}

var Env ConfigStruct

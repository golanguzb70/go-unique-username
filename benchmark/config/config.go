package config

import (
	"os"
)

type Config struct {
	CharList []byte
	CharMp   map[byte]byte
	GrpcPort string
}

var GlobalConfig Config

func Load() {
	charList, ok := GetEnvOrDefault("SUPPORTED_CHARS", "abcdefghijklmnopqrstuvwxyz0123456789").(string)
	if ok {
		GlobalConfig.CharList = []byte(charList)
	}

	GlobalConfig.CharMp = make(map[byte]byte)
	for i := range len(GlobalConfig.CharList) {
		GlobalConfig.CharMp[GlobalConfig.CharList[i]] = byte(i)
	}

	GlobalConfig.GrpcPort, _ = GetEnvOrDefault("GRPC_PORT", "9000").(string)
}

func GetEnvOrDefault(key string, defaultValue any) any {
	v, exists := os.LookupEnv(key)
	if !exists {
		return defaultValue
	}
	return v
}

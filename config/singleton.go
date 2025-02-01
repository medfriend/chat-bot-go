package config

import "sync"

type singleton struct {
	OpenAIKey string
}

var instance *singleton
var once sync.Once

// Init inicializa el singleton con valores predeterminados o de configuración
func Init(apiKey string) {
	once.Do(func() {
		instance = &singleton{
			OpenAIKey: apiKey,
		}
	})
}

// GetInstance devuelve una instancia del singleton configurado
func GetInstance() string {
	if instance == nil {
		panic("Config singleton is not initialized")
	}

	return instance.OpenAIKey
}

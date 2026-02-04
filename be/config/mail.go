package config

import "strconv"

type EmailConfig struct {
	Host         string `mapstructure:"SMTP_HOST"`
	Port         int    `mapstructure:"SMTP_PORT"`
	SenderName   string `mapstructure:"SMTP_SENDER_NAME"`
	AuthEmail    string `mapstructure:"SMTP_AUTH_EMAIL"`
	AuthPassword string `mapstructure:"SMTP_AUTH_PASSWORD"`
}

func NewEmailConfig() (*EmailConfig, error) {
	port, err := strconv.Atoi(getEnvOrDefault("SMTP_PORT", "1234"))
	if err != nil {
		port = 1234
	}
	config := EmailConfig{
		Host:         getEnvOrDefault("SMTP_HOST", ""),
		Port:         port,
		SenderName:   getEnvOrDefault("SMTP_SENDER_NAME", ""),
		AuthEmail:    getEnvOrDefault("SMTP_AUTH_EMAIL", ""),
		AuthPassword: getEnvOrDefault("SMTP_AUTH_PASSWORD", ""),
	}
	return &config, nil
}

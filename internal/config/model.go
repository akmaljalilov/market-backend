package config

import "fmt"

type Postgres struct {
	Username string  `yaml:"username"      env:"PG_USERNAME"`
	Host     string  `yaml:"host"          env:"PG_HOST"`
	Password string  `yaml:"password"      env:"PG_PASSWORD"`
	DbName   string  `yaml:"dbname"        env:"PG_DBNAME"`
	Port     *string `yaml:"port"          env:"PG_PORT"`
}

func (p Postgres) URL() string {
	host := p.Host
	if p.Port != nil {
		host += fmt.Sprintf(":%s", *p.Port)
	}
	return "postgresql://" + p.Username + ":" + p.Password + "@" + host + "/" + p.DbName + "?sslmode=disable"
}
func (p Postgres) JDBC() string {
	host := p.Host
	if p.Port != nil {
		host += fmt.Sprintf(":%s", *p.Port)
	}
	return "jdbc:postgresql://" + host + "/" + p.DbName + "?sslmode=disable&extra_float_digits=0"
}

type Config struct {
	Postgres Postgres `yaml:"postgres"`
}

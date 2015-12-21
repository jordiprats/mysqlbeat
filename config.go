package main

type MySQLConfig struct {
	URLs []string

	Period *int64
}

type ConfigSettings struct {
	Input MySQLConfig
}

package internal

type SchemaPortMapping struct {
	schema            string
	transportProtocol string
	port              int
}

var SchemaPortTable = map[string]*SchemaPortMapping{
	"http":     {"http", "tcp", 80},
	"https":    {"https", "tcp", 443},
	"postgres": {"postgres", "tcp", 5432},
	"mysql":    {"mysql", "tcp", 3306},
	"mariadb":  {"mariadb", "tcp", 3306},
	"redis":    {"redis", "tcp", 6379},
	"mssql":    {"mssql", "tcp", 1433},
	"ftp":      {"ftp", "tcp", 21},
	"":         {"tcp", "tcp", 80},
	"tcp":      {"tcp", "tcp", 80},
	"udp":      {"tcp", "tcp", 80},
}

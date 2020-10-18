package setting

//Config is the service configuration struct
type Config struct {
	RunMode string `yaml:"run_mod"`
	App App `yaml:"app"`
	Server Server `yaml:"server"`
	Database Database `yaml:"database"`
}

//App is the service application related configuration struct
type App struct {
	ServiceName string `yaml:"name"`
	PageSize int `yaml:"page_size"`
	JwtSecret string `yaml:"jwt_secret"`
}


//Server is the service server related configuration struct
type Server struct {
	HTTPProto string `yaml:"http_proto"`
	HTTPPort int `yaml:"http_port"`
	ReadTimeout int `yaml:"read_timeout"`
	WriteTimeout int `yaml:"write_timeout"`
}

//Database is the service db related configuration struct
type Database struct {
	Type string `yaml:"type"`
	User string `yaml:"user"`
	Password string `yaml:"password"`
	Host string `yaml:"host"`
	Name string `yaml:"name"`
	TablePrefix string `yaml:"table_prefix"`
}
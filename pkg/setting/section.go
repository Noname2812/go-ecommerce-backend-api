package setting

type Config struct {
	Server ServerSetting `mapstructure:"server"`
	Mysql  MySQLSetting  `mapstructure:"mysql"`
	Logger LoggerSetting `mapstructure:"logger"`
	Redis  RedisSetting  `mapstructure:"redis"`
	JWT    JWTSetting    `mapstructure:"jwt"`
	GRPC   GRPCSettings  `mapstructure:"GRPC"`
	Kafka  KafkaSetting  `mapstructure:"kafka"`
	Email  EmailSetting  `mapstructure:"email"`
}

type EmailSetting struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
}

type GRPCSettings struct {
	AuthServicePort         int `mapstructure:"AUTH_SERVICE_PORT"`
	UserServicePort         int `mapstructure:"USER_SERVICE_PORT"`
	OrderServicePort        int `mapstructure:"ORDER_SERVICE_PORT"`
	ProductServicePort      int `mapstructure:"PRODUCT_SERVICE_PORT"`
	NotificationServicePort int `mapstructure:"NOTIFICATION_SERVICE_PORT"`
}

type ServerSetting struct {
	Port int    `mapstructure:"port"`
	Mode string `mapstructure:"mode"`
	Host string `mapstructure:"host"`
}

type RedisSetting struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Password string `mapstructure:"password"`
	Database int    `mapstructure:"database"`
}

type MySQLSetting struct {
	Host            string `mapstructure:"host"`
	Port            int    `mapstructure:"port"`
	Username        string `mapstructure:"username"`
	Password        string `mapstructure:"password"`
	Dbname          string `mapstructure:"dbname"`
	MaxIdleConns    int    `mapstructure:"maxIdleConns"`
	MaxOpenConns    int    `mapstructure:"maxOpenConns"`
	ConnMaxLifetime int    `mapstructure:"connMaxLifetime"`
}

type LoggerSetting struct {
	Log_level     string `mapstructure:"log_level"`
	File_log_name string `mapstructure:"file_log_name"`
	Max_backups   int    `mapstructure:"max_backups"`
	Max_age       int    `mapstructure:"max_age"`
	Max_size      int    `mapstructure:"max_size"`
	Compress      bool   `mapstructure:"compress"`
}

// JWT Settings
type JWTSetting struct {
	TOKEN_HOUR_LIFESPAN uint   `mapstructure:"TOKEN_HOUR_LIFESPAN"`
	API_SECRET_KEY      string `mapstructure:"API_SECRET_KEY"`
	JWT_EXPIRATION      string `mapstructure:"JWT_EXPIRATION"`
}

// Kafka Settings
type KafkaSetting struct {
	Brokers []string `mapstructure:"brokers"`
}

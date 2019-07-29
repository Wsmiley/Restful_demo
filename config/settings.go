package config

//"github.com/spf13/viper"

var (
	
	PlayerDataURL = "https://nba.hupu.com/stats/players/pts/"

	TeamURL       = "https://nba.hupu.com/standings"
)

// func init() {
// 	viper.SetConfigName("settings")
// 	viper.AddConfigPath("$GOPATH/src/NBA-master/config")
// 	viper.SetConfigType("yml")
// 	if err != nil {
// 		panic(fmt.Errorf("Fatal error config file: %s \n", err))
// 	}
// }

// func GetPostGreConifg() string {

// 	var (
// 		dbName   string
// 		port     string
// 		user     string
// 		sslMode  string
// 		password string
// 		host     string
// 	)
// 	dbName = viper.GetString("db.dbname")
// 	port = viper.GetString("db.port")
// 	user = viper.GetString("db.user")
// 	sslMode = viper.GetString("db.sslmode")
// 	password = viper.GetString("db.password")
// 	host = viper.GetString("db.host")
// 	return fmt.Sprintf("host:%s user=%s dbname=%s,sslmode=%s,password=%s", host, user, sslMode, password)

// }

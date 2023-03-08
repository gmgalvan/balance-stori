package schema

// Secret data base configurations secrets
type Secret struct {
	DatabasePassworld string
	DatabaseUser      string
	DatabaseName      string
	AWSAccessToken    string
	AWSSecret         string
	AWSRegion         string
}

// Config app configurations for server and db
type Config struct {
	ServerPort         string `default:":8080"`
	DatabasePort       int
	DataBaseHost       string `default:"localhost"`
	SSLDatabaseMode    string `default:"disable"`
	MigrationsPath     string `default:"./internal/store/migrations"`
	S3Bucket           string
	TransactionsFile   string `default:"transactions.csv"`
	EmailAddressFrom   string
	EmailToSendBalance string
}

// PopulatedConfigs configurations and secret combined
type PopulatedConfigs struct {
	Secret
	Config
}

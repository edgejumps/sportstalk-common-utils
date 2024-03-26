package redisfactory

type Config struct {
	// Connection string sample: "redis://<user>:<pass>@localhost:6379/<db>"
	// once connection string is provided, other connection parameters will be ignored
	ConnectionString string `toml:"redis_connection_string" json:"redis_connection_string"`

	// server address, e.g., "localhost:6379", or ":6379" (default)
	Addr     string `toml:"redis_address" json:"redis_address"`
	Username string `toml:"redis_username" json:"redis_username"`
	Password string `toml:"redis_password" json:"redis_password"`

	// sentinel related parameters
	SentinelMasterName string `toml:"redis_sentinel_master_name" json:"redis_sentinel_master_name"`
	SentinelAddrs      string `toml:"redis_sentinel_addrs" json:"redis_sentinel_addrs"`
	SentinelUsername   string `toml:"redis_sentinel_username" json:"redis_sentinel_username"`
	SentinelPassword   string `toml:"redis_sentinel_password" json:"redis_sentinel_password"`

	// generic parameters
	Timeout int `toml:"redis_timeout" json:"redis_timeout"`

	// TLS related parameters
	TlsX509CertFile string `toml:"redis_tls_x509_cert_file" json:"redis_tls_x509_cert_file"`
	TlsX509KeyFile  string `toml:"redis_tls_x509_key_file" json:"redis_tls_x509_key_file"`
	TlsCACertFile   string `toml:"redis_tls_ca_cert_file" json:"redis_tls_ca_cert_file"`
}

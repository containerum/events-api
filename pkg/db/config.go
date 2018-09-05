package db

import (
	"time"

	"encoding/json"

	"github.com/globalsign/mgo"
	"github.com/sirupsen/logrus"
)

type MongoConfig struct {
	Logger logrus.FieldLogger

	// Debug enable the delivery of debug messages to the logger.  Only meaningful
	// if a logger is also set.
	Debug bool
	mgo.DialInfo
}

type StorableMongoConfig struct {
	// Debug enable the delivery of debug messages to the logger.  Only meaningful
	// if a logger is also set.
	Debug bool `json:"debug"`

	// Addrs holds the addresses for the seed servers.
	Addrs []string `json:"addrs"`

	// Timeout is the amount of time to wait for a server to respond when
	// first connecting and on follow up operations in the session. If
	// timeout is zero, the call may block forever waiting for a connection
	// to be established. Timeout does not affect logic in DialServer.
	Timeout time.Duration `json:"timeout"`

	// Database is the default database name used when the Session.DB method
	// is called with an empty name, and is also used during the initial
	// authentication if Source is unset.
	Database string `json:"database"`

	// ReplicaSetName, if specified, will prevent the obtained session from
	// communicating with any server which is not part of a replica set
	// with the given name. The default is to communicate with any server
	// specified or discovered via the servers contacted.
	ReplicaSetName string `json:"replica_set_name"`

	// Source is the database used to establish credentials and privileges
	// with a MongoDB server. Defaults to the value of Database, if that is
	// set, or "admin" otherwise.
	Source string `json:"source"`

	// Service defines the service name to use when authenticating with the GSSAPI
	// mechanism. Defaults to "mongodb".
	Service string `json:"service"`

	// ServiceHost defines which hostname to use when authenticating
	// with the GSSAPI mechanism. If not specified, defaults to the MongoDB
	// server's address.
	ServiceHost string `json:"service_host"`

	// Mechanism defines the protocol for credential negotiation.
	// Defaults to "MONGODB-CR".
	Mechanism string `json:"mechanism"`

	// Username and Password inform the credentials for the initial authentication
	// done on the database defined by the Source field. See Session.Login.
	Username string `json:"username"`
	Password string `json:"password"`

	// PoolLimit defines the per-server socket pool limit. Defaults to 4096.
	// See Session.SetPoolLimit for details.
	PoolLimit int `json:"pool_limit"`

	// PoolTimeout defines max time to wait for a connection to become available
	// if the pool limit is reaqched. Defaults to zero, which means forever.
	// See Session.SetPoolTimeout for details
	PoolTimeout time.Duration `json:"pool_timeout"`

	// The identifier of the client application which ran the operation.
	AppName string `json:"app_name"`

	// ReadPreference defines the manner in which servers are chosen. See
	// Session.SetMode and Session.SelectServers.
	ReadPreference *mgo.ReadPreference `json:"read_preference,omitempty"`

	// FailFast will cause connection and query attempts to fail faster when
	// the server is unavailable, instead of retrying until the configured
	// timeout period. Note that an unavailable server may silently drop
	// packets instead of rejecting them, in which case it's impossible to
	// distinguish it from a slow server, so the timeout stays relevant.
	FailFast bool `json:"fail_fast"`

	// Direct informs whether to establish connections only with the
	// specified seed servers, or to obtain information for the whole
	// cluster and establish connections with further servers too.
	Direct bool `json:"direct"`

	// MinPoolSize defines The minimum number of connections in the connection pool.
	// Defaults to 0.
	MinPoolSize int `json:"min_pool_size"`

	//The maximum number of milliseconds that a connection can remain idle in the pool
	// before being removed and closed.
	MaxIdleTimeMS int `json:"max_idle_time_ms"`
}

func (config MongoConfig) Storable() StorableMongoConfig {
	return StorableMongoConfig{
		Debug:          config.Debug,
		Addrs:          append(make([]string, 0, len(config.Addrs)), config.Addrs...),
		Timeout:        config.Timeout,
		Database:       config.Database,
		ReplicaSetName: config.ReplicaSetName,
		Source:         config.Source,
		Service:        config.Service,
		ServiceHost:    config.ServiceHost,
		Mechanism:      config.Mechanism,
		Username:       config.Username,
		Password:       config.Password,
		PoolLimit:      config.PoolLimit,
		PoolTimeout:    config.PoolTimeout,
		ReadPreference: config.ReadPreference,
		FailFast:       config.FailFast,
		Direct:         config.Direct,
		MinPoolSize:    config.MinPoolSize,
		MaxIdleTimeMS:  config.MaxIdleTimeMS,
	}
}

func (config MongoConfig) MarshalJSON() ([]byte, error) {
	return json.MarshalIndent(config.Storable(), "", "  ")
}

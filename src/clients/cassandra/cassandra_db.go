package cassandra

import (
	"fmt"
	"github.com/gocql/gocql"
	"time"
)

var (
	// cluster *gocql.ClusterConfig
	session *gocql.Session
)

func init() {
	// Connect to cassandra cluster
	cluster := gocql.NewCluster("127.0.0.1")
	cluster.Keyspace = "oauth"
	cluster.Consistency = gocql.Quorum
	cluster.ConnectTimeout = time.Second * 10
	cluster.DisableInitialHostLookup = true

	var err error
	session, err = cluster.CreateSession()
	if err != nil {
		panic(err)
	}
	fmt.Println("cassandra connection successfully created")
}

func GetSession() *gocql.Session {
	return session
}

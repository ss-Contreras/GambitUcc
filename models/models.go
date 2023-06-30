package models

type SecretRDSJon struct {
	Username            string `json:"username"`
	Password            string `json:"passowrd"`
	Engine              string `json:"engine"`
	Host                string `json:"host"`
	Port                int    `json:"port"`
	DbClusterIdentifier string `json:"dbCluterIdentifier"`
}

type SignUp struct {
	UserEmail string `json:"UserEmail"`
	UserUUID  string `json:"v"`
}

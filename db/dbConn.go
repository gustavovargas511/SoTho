package db

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//MongoConn es el objeto de conexion a la db
var MongoConn = ConectarBD()
var clientOptions = options.Client().ApplyURI("mongodb+srv://user01:picadorcriminalmutilador@sotho-57omd.mongodb.net/test?retryWrites=true&w=majority")

//ConectarBD es la funcion que me permite conetar a la base de datos
func ConectarBD() *mongo.Client {
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}

	log.Println("Connection SUCCESSFUL")
	return client
}

//ConnCheck es el ping a la database
func ConnCheck() int {
	err := MongoConn.Ping(context.TODO(), nil)
	if err != nil {
		return 0
	}

	return 1
}

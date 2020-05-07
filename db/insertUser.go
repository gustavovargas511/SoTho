package db

import (
	"context"
	"time"
	"github.com/gustavovargas511/sotho/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//InsertUser es la parada final con la DB para insertar datos de usuario
func InsertUser(u models.User) (string, bool, error){
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoConn.Database("sotho_db")
	col := db.Collection("users")

	u.Password, _ = EncryptPass(u.Password)

	result, err := col.InsertOne(ctx, u)
	
	if err != nil{
		return "",false, err
	}

	ObjID, _ := result.InsertedID.(primitive.ObjectID)
	return ObjID.String(), true, nil
}
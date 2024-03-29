package db

import (
	"context"
	"time"

	"github.com/mnarvaezm96/doom_go/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/* InsertoRegistro es para la parte final con la db para insertar los datos del usuario*/
func InsertoRegistro(u models.Usuario) (string, bool, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("twittor")
	col := db.Collection("user")

	u.Password, _ = EncriptarPassword(u.Password)

	result, err := col.InsertOne(ctx, u)
	if err != nil {
		return "", false, err

	}
	objID, _ := result.InsertedID.(primitive.ObjectID)
	return objID.String(), true, nil

}

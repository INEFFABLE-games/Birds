package service

import (
	"context"
	"github.com/INEFFABLE-games/Birds/models"
	"github.com/INEFFABLE-games/Birds/repository"
	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"testing"
	"time"
)

func InitDataBase() *mongo.Client {
	clientOptions := options.Client().ApplyURI("mongodb+srv://dbuser1:dbuser1@petscluster.82emx.mongodb.net/Birds?retryWrites=true&w=majority")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()

	mongoConn, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.WithFields(log.Fields{
			"handler": "birds",
			"action":  "connect",
		}).Errorf("unable connect to MonogDB %v", err)
	}
	return mongoConn
}

func TestBirdService_Create(t *testing.T) {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()

	//connect to mongo

	mongoConn := InitDataBase()

	r := repository.NewBirdRepository(mongoConn)
	serv := NewBirdService(r)

	err := serv.Create(ctx, models.Bird{
		Owner: "Me",
		Name:  "Phoenix",
		Type:  "Firebird",
	})
	if err != nil {
		log.WithFields(log.Fields{
			"handler": "birds",
			"action":  "create",
		}).Errorf("unable to insert bird %v", err)
	}

	require.Nil(t, err)
}

func TestBirdService_Update(t *testing.T) {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()

	//connect to mongo

	mongoConn := InitDataBase()

	r := repository.NewBirdRepository(mongoConn)
	serv := NewBirdService(r)

	err := serv.Update(ctx, models.Bird{
		Owner: "Me",
		Name:  "Daeg",
		Type:  "Firesrnabird",
	})
	if err != nil {
		log.WithFields(log.Fields{
			"handler": "birds",
			"action":  "update",
		}).Errorf("unable to insert bird %v", err)
	}

	require.Nil(t, err)
}

func TestBirdService_Get(t *testing.T) {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()

	//connect to mongo

	mongoConn := InitDataBase()

	r := repository.NewBirdRepository(mongoConn)
	serv := NewBirdService(r)

	result, err := serv.Get(ctx, "Me")
	if err != nil {
		log.WithFields(log.Fields{
			"handler": "birds",
			"action":  "get",
		}).Errorf("unable to insert bird %v", err)
	}
	require.Equal(t, models.Bird{
		Owner: "Me",
		Name:  "Daeg",
		Type:  "Firesrnabird",
	}, result)
}

func TestBirdService_Delete(t *testing.T) {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()

	//connect to mongo

	mongoConn := InitDataBase()

	r := repository.NewBirdRepository(mongoConn)
	serv := NewBirdService(r)

	err := serv.Delete(ctx, "Me")
	if err != nil {
		log.WithFields(log.Fields{
			"handler": "birds",
			"action":  "delete",
		}).Errorf("unable to insert bird %v", err)
	}
	require.Nil(t, err)
}

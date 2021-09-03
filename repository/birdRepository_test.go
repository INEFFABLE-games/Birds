package repository

import (
	"context"
	"github.com/INEFFABLE-games/Birds/models"
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

func TestBirdRepository_Insert(t *testing.T) {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()

	//connect to mongo

	mongoConn := InitDataBase()

	r := NewBirdRepository(mongoConn)

	err := r.Insert(ctx, models.Bird{
		Owner: "Me",
		Name:  "Phoenix",
		Type:  "Firebird",
	})
	if err != nil {
		log.WithFields(log.Fields{
			"handler": "birds",
			"action":  "insert",
		}).Errorf("unable to insert bird %v", err)
	}

	require.Nil(t, err)
}

func TestBirdRepository_Get(t *testing.T) {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()

	//connect to mongo

	mongoConn := InitDataBase()

	r := NewBirdRepository(mongoConn)

	bird, err := r.FindByOwner(ctx, "Me")
	if err != nil {
		log.WithFields(log.Fields{
			"handler": "birds",
			"action":  "find",
		}).Errorf("unable to find bird %v", err)
	}

	require.Equal(t, models.Bird{
		Owner: "Me",
		Name:  "Phoenix",
		Type:  "Firebird",
	}, bird)
}

func TestBirdRepository_Update(t *testing.T) {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()

	//connect to mongo

	mongoConn := InitDataBase()

	r := NewBirdRepository(mongoConn)

	err := r.Update(ctx, models.Bird{
		Owner: "Me",
		Name:  "Dragon",
		Type:  "sejgnose",
	})

	require.Nil(t, err)
}

func TestBirdRepository_DeleteByOwner(t *testing.T) {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()

	//connect to mongo

	mongoConn := InitDataBase()

	r := NewBirdRepository(mongoConn)

	err := r.DeleteByOwner(ctx, "Me")

	require.Nil(t, err)
}

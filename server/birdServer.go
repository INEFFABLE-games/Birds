package server

import (
	"context"
	"github.com/INEFFABLE-games/Birds/models"
	"github.com/INEFFABLE-games/Birds/protocol"
	"github.com/INEFFABLE-games/Birds/service"
	log "github.com/sirupsen/logrus"
)

type BirdHandler struct {
	birdService *service.BirdService

	protocol.UnimplementedBirdServiceServer
}

func (b BirdHandler) Create(ctx context.Context, request *protocol.CreateRequest) (*protocol.CreateReply, error) {

	bird := models.Bird{}
	bird.Owner = request.GetOwner()
	bird.Name = request.GetName()
	bird.Type = request.GetType()

	err := b.birdService.Create(ctx, bird)
	if err != nil {
		log.WithFields(log.Fields{
			"handler": "birdHandler",
			"action":  "create",
		}).Errorf("unable to create bird %v", err.Error())

		message := "unable to create bird"
		return &protocol.CreateReply{Message: &message}, err
	}

	message := "bird created"
	return &protocol.CreateReply{Message: &message}, err
}

func (b BirdHandler) Get(ctx context.Context, request *protocol.GetRequest) (*protocol.GetReply, error) {

	owner := request.GetOwner()

	result, err := b.birdService.Get(ctx, owner)
	if err != nil {
		log.WithFields(log.Fields{
			"handler": "birdHandler",
			"action":  "get",
		}).Errorf("unable to get bird %v", err.Error())

		return &protocol.GetReply{}, err
	}

	return &protocol.GetReply{Owner: &result.Owner, Name: &result.Name, Type: &result.Type}, err
}

func (b BirdHandler) Change(ctx context.Context, request *protocol.ChangeRequest) (*protocol.ChangeReply, error) {

	bird := models.Bird{}
	bird.Owner = request.GetOwner()
	bird.Name = request.GetName()
	bird.Type = request.GetType()

	err := b.birdService.Update(ctx, bird)
	if err != nil {
		log.WithFields(log.Fields{
			"handler": "birdHandler",
			"action":  "get",
		}).Errorf("unable to change bird %v", err.Error())

		return &protocol.ChangeReply{}, err
	}

	message := "bird cahnged"

	return &protocol.ChangeReply{Message: &message}, err
}

func (b BirdHandler) Delete(ctx context.Context, request *protocol.DeleteRequest) (*protocol.DeleteReply, error) {

	owner := request.GetOwner()

	err := b.birdService.Delete(ctx, owner)
	if err != nil {
		log.WithFields(log.Fields{
			"handler": "birdHandler",
			"action":  "get",
		}).Errorf("unable to delete bird %v", err.Error())

		return &protocol.DeleteReply{}, err
	}

	message := "bird deleted"

	return &protocol.DeleteReply{Message: &message}, err
}

func NewBirdHandler(birdService *service.BirdService) protocol.BirdServiceServer {
	return BirdHandler{birdService: birdService}
}

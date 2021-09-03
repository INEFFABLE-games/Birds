package server

import (
	"Birds/internal/models"
	proto "Birds/internal/protocol"
	"Birds/internal/service"
	"context"
	log "github.com/sirupsen/logrus"
)

type BirdHandler struct {
	birdService *service.BirdService

	proto.UnimplementedBirdServiceServer
}

func (b BirdHandler) CreateBird(ctx context.Context, request *proto.CreateRequest) (*proto.CreateReply, error) {

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
		return &proto.CreateReply{Message: &message}, err
	}

	message := "bird created"
	return &proto.CreateReply{Message: &message}, err
}

func (b BirdHandler) GetBird(ctx context.Context, request *proto.GetRequest) (*proto.GetReply, error) {

	owner := request.GetOwner()

	result, err := b.birdService.Get(ctx, owner)
	if err != nil {
		log.WithFields(log.Fields{
			"handler": "birdHandler",
			"action":  "get",
		}).Errorf("unable to get bird %v", err.Error())

		return &proto.GetReply{}, err
	}

	return &proto.GetReply{Owner: &result.Owner, Name: &result.Name, Type: &result.Type}, err
}

func (b BirdHandler) ChangeBird(ctx context.Context, request *proto.ChangeRequest) (*proto.ChangeReply, error) {

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

		return &proto.ChangeReply{}, err
	}

	message := "bird cahnged"

	return &proto.ChangeReply{Message: &message}, err
}

func (b BirdHandler) DeleteBird(ctx context.Context, request *proto.DeleteRequest) (*proto.DeleteReply, error) {

	owner := request.GetOwner()

	err := b.birdService.Delete(ctx, owner)
	if err != nil {
		log.WithFields(log.Fields{
			"handler": "birdHandler",
			"action":  "get",
		}).Errorf("unable to delete bird %v", err.Error())

		return &proto.DeleteReply{}, err
	}

	message := "bird deleted"

	return &proto.DeleteReply{Message: &message}, err
}

func NewBirdHandler(birdService *service.BirdService) proto.BirdServiceServer {
	return BirdHandler{birdService: birdService}
}

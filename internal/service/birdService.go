package service

import (
	"Birds/internal/models"
	"Birds/internal/repository"
	"context"
)

type BirdService struct {
	birdRepo *repository.BirdRepository
}

func (b *BirdService) Create(ctx context.Context, bird models.Bird) error {
	return b.birdRepo.Insert(ctx, bird)
}

func (b *BirdService) Update(ctx context.Context, bird models.Bird) error {
	return b.birdRepo.Update(ctx, bird)
}

func (b *BirdService) Get(ctx context.Context, owner string) (models.Bird, error) {
	return b.birdRepo.FindByOwner(ctx, owner)
}

func (b *BirdService) Delete(ctx context.Context, owner string) error {
	return b.birdRepo.DeleteByOwner(ctx, owner)
}

func NewBirdService(birdRepo *repository.BirdRepository) *BirdService {
	return &BirdService{birdRepo: birdRepo}
}

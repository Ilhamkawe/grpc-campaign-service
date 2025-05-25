package campaign

import (
	campaignPort "github.com/Ilhamkawe/grpc-campaign-service/internal/port/campign"
	"gorm.io/gorm"
)

type CampaignRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) campaignPort.CampaignRepositoryPort {
	return &CampaignRepository{db: db}
}

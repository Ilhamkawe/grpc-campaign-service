package port

import (
	db "github.com/Ilhamkawe/grpc-campaign-service/internal/adapter/database"
	"github.com/Ilhamkawe/grpc-campaign-service/internal/application/domain/campaign"
)

type CampaignServicePort interface {
	GetCampaigns(UserID int) ([]db.Campaign, error)
	GetCampaignByID(input campaign.GetCampaignDetailInput) (db.Campaign, error)
	GetUserCampaignByID(inputID campaign.GetCampaignDetailInput, inputUser campaign.GetUserCampaign) (db.Campaign, error)
	GetAllCampaign() ([]db.Campaign, error)
	GetCampaignByStatus(status string) ([]db.Campaign, error)
	GetCampaignByIDWoStatus(id int) (db.Campaign, error)
	CreateCampaign(input campaign.CreateCampaignInput) (db.Campaign, error)
	UpdateCampaign(inputID campaign.GetCampaignDetailInput, inputData campaign.UpdateCampaignInput) (db.Campaign, error)
	UpdateAttachment(input campaign.UpdateAttachmentInput) (db.Campaign, error)
	SaveCampaignImage(input campaign.CreateCampaignImageInput, fileLocation string) (db.CampaignImage, error)
	SaveCampaignReward(input campaign.CreateCampaignRewardInput) (db.CampaignReward, error)
	Limit(num int) ([]db.Campaign, error)
	GetRewards(input campaign.GetCampaignDetailInput) ([]db.CampaignReward, error)
	SearchCampaign(input campaign.SearchCampaignInput) ([]db.Campaign, error)
	DeleteReward(input campaign.DeleteCampaignRewardInput) (bool, error)
	DeleteImage(input campaign.DeleteCampaignImageInput) (bool, error)
	ChangeStatus(Status string, ID int) (db.Campaign, error)
	CreateActivity(input campaign.CreateCampaignActivityInput) (db.CampaignActivity, error)
	UpdateActivity(input campaign.UpdateCampaignActivityInput) (db.CampaignActivity, error)
	DeleteActivity(input campaign.DeleteCampaignActivityInput) (bool, error)
	FindActivity(activityID int) (db.CampaignActivity, error)
	FindActivityByUser(input campaign.GetCampaignActivityInput, campaignUser campaign.GetUserCampaign) (db.CampaignActivity, error)
	FindAllActivityByCampaignID(campaignID int) ([]db.CampaignActivity, error)
	Paginate(input campaign.PaginateCampaignInput) (db.PaginateCampaigns, error)
	IsCollectAbleByDate()
	CreateCattegory(input campaign.CattegoryInput) (db.Cattegory, error)
	DeleteCattegory(id int) (bool, error)
	FindAllCattegory() ([]db.Cattegory, error)
	SearchCampaignPaginate(input campaign.SearchCampaignPaginate) (db.campaign.PaginateCampaigns, error)
}

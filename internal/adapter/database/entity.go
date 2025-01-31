package database

import (
	"github.com/leekchan/accounting"
	"os/user"
	"time"
)

type Campaign struct {
	ID               int
	UserID           int
	Name             string
	ShortDescription string
	Description      string
	BackerCount      int
	GoalAmount       int
	Cattegory        string
	CurrentAmount    int
	Slug             string
	Status           string
	Attachment       string
	Collectable      bool
	FinishAt         string
	CreatedAt        time.Time
	UpdatedAt        time.Time
	User             user.User
	CampaignImages   []CampaignImage
	CampaignRewards  []CampaignReward
}

func (Campaign) TableName() string {
	return "campaigns"
}

func (c Campaign) GoalAmountFormatIDR() string {
	ac := accounting.Accounting{Symbol: "Rp", Precision: 2, Decimal: ",", Thousand: "."}
	return ac.FormatMoney(c.GoalAmount)
}
func (c Campaign) CurrentAmountFormatIDR() string {
	ac := accounting.Accounting{Symbol: "Rp", Precision: 2, Decimal: ",", Thousand: "."}
	return ac.FormatMoney(c.CurrentAmount)
}

type CampaignImage struct {
	ID         int
	CampaignID int
	FileName   string
	IsPrimary  int
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

func (CampaignImage) TableName() string {
	return "campaign_images"
}

type CampaignReward struct {
	ID          int
	CampaignID  int
	Description string
	Perks       string
	MinDonate   int
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (CampaignReward) TableName() string {
	return "campaign_rewards"
}

type CampaignActivity struct {
	ID               int
	CampaignID       int
	Name             string
	ShortDescription string
	Description      string
	ImageUrl         string
	Slug             string
	Campaign         Campaign
	CreatedAt        time.Time
	UpdatedAt        time.Time
}

func (CampaignActivity) TableName() string {
	return "campaign_activities"
}

type Cattegory struct {
	ID   int
	Name string
}

func (Cattegory) TableName() string {
	return "cattegories"
}

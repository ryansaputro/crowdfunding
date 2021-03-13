package campaign

import (
	"errors"
	"fmt"

	"github.com/gosimple/slug"
)

type Service interface {
	GetCampaigns(UserID int) ([]Campaign, error)
	GetCampaignByID(input GetCampaignDetailInput) (Campaign, error)
	CreateCampaign(input CreateCampaignInput) (Campaign, error)
	Update(inputID GetCampaignDetailInput, inputData CreateCampaignInput) (Campaign, error)
	SaveCampaignImage(input CreateCampaignImageInput, fileLocation string) (CampaignImage, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) GetCampaigns(UserID int) ([]Campaign, error) {
	if UserID != 0 {
		campaigns, err := s.repository.FindByUserID(UserID)

		if err != nil {
			return campaigns, err
		}

		return campaigns, nil
	}

	campaigns, err := s.repository.FindAll()

	if err != nil {
		return campaigns, err
	}

	return campaigns, nil

}

func (s *service) GetCampaignByID(input GetCampaignDetailInput) (Campaign, error) {

	campaign, err := s.repository.FindByID(input.ID)

	if err != nil {
		return campaign, err
	}

	return campaign, nil
}

func (s *service) CreateCampaign(input CreateCampaignInput) (Campaign, error) {
	campaign := Campaign{}
	campaign.Name = input.Name
	campaign.ShortDescription = input.ShortDescription
	campaign.Description = input.Description
	campaign.GoalAmount = input.GoalAmount
	campaign.Perks = input.Perks
	campaign.UserID = input.User.ID

	slugCandidate := fmt.Sprintf("%s %d", input.Name, input.User.ID)

	campaign.Slug = slug.Make(slugCandidate)

	newCampaign, err := s.repository.Save(campaign)

	if err != nil {
		return newCampaign, err
	}

	return newCampaign, nil

}

func (s *service) Update(inputID GetCampaignDetailInput, inputData CreateCampaignInput) (Campaign, error) {
	campaign, err := s.GetCampaignByID(inputID)

	if err != nil {
		return campaign, nil
	}

	if campaign.UserID != inputData.User.ID {
		return campaign, errors.New("Maaf anda bukan pemilik campaign")
	}

	campaign.Name = inputData.Name
	campaign.ShortDescription = inputData.ShortDescription
	campaign.Description = inputData.Description
	campaign.GoalAmount = inputData.GoalAmount
	campaign.Perks = inputData.Perks

	// slug tidak diupdate karena kalo di crawling biar masih ada
	updatedCampaign, err := s.repository.Update(campaign)

	if err != nil {
		return updatedCampaign, err
	}

	return updatedCampaign, nil
}

func (s *service) SaveCampaignImage(input CreateCampaignImageInput, fileLocation string) (CampaignImage, error) {
	isPrimary := 0
	GetCampaignImage, err := s.repository.FindByID(input.CampaignID)

	if err != nil {
		return CampaignImage{}, nil
	}

	if GetCampaignImage.UserID != input.User.ID {
		return CampaignImage{}, errors.New("Maaf anda bukan pemilik campaign")
	}

	if input.IsPrimary {
		isPrimary = 1
		_, err := s.repository.MarkAllImagesNonPrimary(input.CampaignID)

		if err != nil {
			return CampaignImage{}, err
		}
	}

	campaignImage := CampaignImage{}
	campaignImage.CampaignID = input.CampaignID
	campaignImage.IsPrimary = isPrimary
	campaignImage.FileName = fileLocation

	newCampaignImage, err := s.repository.CreateImage(campaignImage)

	if err != nil {
		return newCampaignImage, err
	}

	return newCampaignImage, nil
}

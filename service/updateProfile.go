package service

import (
	"context"

	"github.com/teohrt/cruddyAPI/entity"
	"github.com/teohrt/cruddyAPI/utils"
)

func (svc ServiceImpl) UpdateProfile(ctx context.Context, profileData entity.ProfileData, profileID string) error {
	emailHash := utils.Hash(profileData.Email)
	if profileID != emailHash {
		return EmailIncsonsistentWithProfileIDError{"Email inconsistent with ProfileID"}
	}

	_, err := svc.GetProfile(ctx, profileID)
	if err != nil {
		svc.Logger.Error().Err(err).Msg("UpdateProfile: GetProfile failed")
		return err
	}

	updateProfile := entity.Profile{
		ID:          profileID,
		ProfileData: profileData,
	}

	_, err = svc.Client.UpsertItem(ctx, updateProfile)
	if err != nil {
		svc.Logger.Printf("%v", profileID)
		svc.Logger.Error().Err(err).Msg("UpdateProfile: UpsertItem failed")
		return err
	}

	return nil
}

package repositories

import "tacoProject/internal/models"

type PfcInfoRepository interface {
	CreatePfcInfo(info models.PfcInfo) error
}

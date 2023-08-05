package collector

import (
	"time"

	"github.com/aarontanjaya/finance-data-collector/dto"
	"github.com/aarontanjaya/finance-data-collector/entity"
)

func MapPasardanaFundSnp(r []dto.PasardanaSnapshotResponse) ([]entity.FundSnapshot, error) {
	result := make([]entity.FundSnapshot, len(r))

	for idx, item := range r {
		tAUM, err := time.Parse("2006-01-02T15:04:05", item.AUMLastUpdate)
		if err != nil {
			return nil, err
		}

		tUpdate, err := time.Parse("2006-01-02T15:04:05", item.LastUpdate)
		if err != nil {
			return nil, err
		}

		result[idx] = entity.FundSnapshot{
			Name:          item.Name,
			IMName:        item.IMName,
			Type:          item.ConservativeCategory,
			TypeID:        -1,
			IsIndex:       item.IsIndex,
			IsETF:         item.IsETF,
			IsSharia:      item.IsSharia,
			NAV:           item.NAV,
			DailyRR:       item.DailyReturn,
			MonthlyRR:     item.MonthlyReturn,
			YTDRR:         item.YTDReturn,
			YearlyRR:      item.YearlyReturn,
			AUM:           item.AUM,
			AUMLastUpdate: uint64(tAUM.Unix()),
			Units:         item.TotalUnit,
			LastUpdate:    uint64(tUpdate.Unix()),
		}
	}

	return result, nil
}

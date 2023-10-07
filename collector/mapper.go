package collector

import (
	"github.com/aarontanjaya/finance-data-collector/dto"
	"github.com/aarontanjaya/finance-data-collector/entity"
	util "github.com/aarontanjaya/finance-data-collector/utils"
)

func MapPasardanaFundDetail(r dto.PasardanaFundDetailResponse) (*entity.FundDetail, error) {
	var res *entity.FundDetail

	category := GetPasardanaMappedCategory(r.Type)
	NAV, err := MapPasardanaDataPoint(r.NAVTimeSeries)
	if err != nil {
		return res, err
	}

	AUM, err := MapPasardanaDataPoint(r.AUM)
	if err != nil {
		return res, err
	}

	totalUnits, err := MapPasardanaDataPoint(r.TotalUnits)
	if err != nil {
		return res, err
	}

	benchmarks, err := MapPasardanaBenchmarks(r.Benchmarks)
	if err != nil {
		return res, err
	}

	res = &entity.FundDetail{
		Source:            SourceIdPasardana,
		Id:                r.Id,
		Name:              r.Name,
		Type:              category,
		Currency:          r.Currency,
		IsSharia:          r.IsSharia,
		IsETF:             r.IsETF,
		IsIndex:           r.IsIndex,
		HasDividend:       r.HasDividend,
		IMId:              r.IMId,
		CustodianBankId:   r.CustodianBankId,
		CustodianBankName: r.CustodianBankName,
		Benchmarks:        benchmarks,
		NAVTimeSeries:     NAV,
		AUM:               AUM,
		TotalUnits:        totalUnits,
	}

	return res, nil
}

func MapPasardanaDataPoint(r []dto.PasardanaDataPoint) ([]entity.DataPoint, error) {
	result := make([]entity.DataPoint, len(r))
	for idx, item := range r {
		date, err := util.ParseTimeInLocationOrDefault("2006-01-02T15:04:05", item.Date, "Asia/Jakarta")
		if err != nil {
			return nil, err
		}

		result[idx].Date = uint64(date.Unix())
		result[idx].Rate = item.Rate
		result[idx].Value = item.Value
	}

	return result, nil
}

func MapPasardanaBenchmarks(r []dto.PasardanaFundBenchmark) ([]entity.FundBenchmark, error) {
	result := make([]entity.FundBenchmark, len(r))
	for idx, item := range r {
		dp, err := MapPasardanaDataPoint(item.Data)
		if err != nil {
			return result, err
		}
		result[idx].Name = item.Name
		result[idx].Id = item.Id
		result[idx].Data = dp
	}

	return result, nil
}

func MapPasardanaFundSnp(r []dto.PasardanaSnapshotResponse) ([]entity.FundSnapshot, error) {
	result := make([]entity.FundSnapshot, len(r))

	for idx, item := range r {
		tAUM, err := util.ParseTimeInLocationOrDefault("2006-01-02T15:04:05", item.AUMLastUpdate, "Asia/Jakarta")
		if err != nil {
			return nil, err
		}

		tUpdate, err := util.ParseTimeInLocationOrDefault("2006-01-02T15:04:05", item.LastUpdate, "Asia/Jakarta")
		if err != nil {
			return nil, err
		}

		result[idx] = entity.FundSnapshot{
			Source:        SourceIdPasardana,
			Id:            item.Id,
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

func MapPasardanaPagination(p entity.Pagination) *entity.PasardanaPagination {
	return &entity.PasardanaPagination{
		Start:     p.Start,
		Length:    p.Length,
		SortField: p.SortField,
		SortOrder: p.SortOrder,
		KeyWord:   p.KeyWord,
	}
}

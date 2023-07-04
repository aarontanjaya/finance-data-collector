package collector

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/aarontanjaya/finance-data-collector/dto"
	"github.com/aarontanjaya/finance-data-collector/entity"
)

type FundSnapshotCollector interface {
	GetAll() ([]entity.FundSnapshot, error)
}

type pasardanaSnapshotImpl struct {
	url string
}

func NewPasardanaSnaphotCollector(url string) FundSnapshotCollector {
	return &pasardanaSnapshotImpl{
		url: url,
	}
}

func (snp *pasardanaSnapshotImpl) GetAll() ([]entity.FundSnapshot, error) {
	var res []dto.PasardanaSnapshotResponse
	resp, err := http.Get(snp.url + "FundSearchResult/GetAll?pageBegin=1&pageLength=100&sortField=Name&sortOrder=ASC")
	if err != nil {
		log.Fatalln(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	json.Unmarshal(body, &res)

	return []entity.FundSnapshot{}, nil
}

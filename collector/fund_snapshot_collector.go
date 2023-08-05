package collector

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/aarontanjaya/finance-data-collector/dto"
	"github.com/aarontanjaya/finance-data-collector/entity"
	util "github.com/aarontanjaya/finance-data-collector/utils"
)

type FundSnapshotCollector interface {
	GetAll(entity.PasardanaPagination) ([]entity.FundSnapshot, error)
}

type pasardanaSnapshotImpl struct {
	url *url.URL
}

func NewPasardanaSnaphotCollector(address string) (FundSnapshotCollector, error) {
	url, err := url.Parse(address)
	if err != nil {
		return nil, err
	}
	return &pasardanaSnapshotImpl{
		url: url,
	}, nil
}

func (snp *pasardanaSnapshotImpl) GetAll(q entity.PasardanaPagination) ([]entity.FundSnapshot, error) {
	var res []dto.PasardanaSnapshotResponse

	u := snp.url.JoinPath("FundSearchResult/GetAll")
	util.BuildURLQuery(u, q)

	resp, err := http.Get(u.String())
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(body, &res)
	if err != nil {
		return nil, err
	}

	return MapPasardanaFundSnp(res)
}

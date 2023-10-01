package collector

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/aarontanjaya/finance-data-collector/dto"
	"github.com/aarontanjaya/finance-data-collector/entity"
	util "github.com/aarontanjaya/finance-data-collector/utils"
)

type FundSnapshotCollector interface {
	GetAll(entity.Pagination) ([]entity.FundSnapshot, error)
	GetById(id int) (*entity.FundDetail, error)
}

type pasardanaSnapshotImpl struct {
	url *url.URL
}

func NewPasardanaSnaphotCollector() (FundSnapshotCollector, error) {
	url, err := url.Parse(PasardanaBaseURL)
	if err != nil {
		return nil, err
	}
	return &pasardanaSnapshotImpl{
		url: url,
	}, nil
}

func (snp *pasardanaSnapshotImpl) GetAll(q entity.Pagination) ([]entity.FundSnapshot, error) {
	var res []dto.PasardanaSnapshotResponse

	u := snp.url.JoinPath("FundSearchResult/GetAll")
	pdnPagination := MapPasardanaPagination(q)
	util.BuildURLQuery(u, *pdnPagination)

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

func (snp *pasardanaSnapshotImpl) GetById(id int) (*entity.FundDetail, error) {
	var res dto.PasardanaFundDetailResponse

	u := snp.url.JoinPath("FundService/GetSnapshot")
	q := url.Values{}
	q.Add("fundid", fmt.Sprintf("%d", id))
	u.RawQuery = q.Encode()

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

	return MapPasardanaFundDetail(res)
}

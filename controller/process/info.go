package process

import (
	"M-Gate/response"

	"net/http"
)

type GetBillerMsg struct {
	ResponseCode string     `json:"responseCode"`
	Description  string     `json:"description"`
	MyData       DataDetail `json:"dataDetail"`
}

type DataDetail struct {
	TheData []DataListVar `json:"billers"`
}

type DataListVar struct {
	BillerId   string `json:"billerId"`
	Name       string `json:"name"`
	BillerType string `json:"type"`
}

func GetBiller(w http.ResponseWriter, r *http.Request) {
	dataList := []DataListVar{
		DataListVar{
			BillerId:   "21120",
			Name:       "PLN - POSTPAID",
			BillerType: "Electricity",
		}, DataListVar{
			BillerId:   "21140",
			Name:       "PLN - NONTAGLIS",
			BillerType: "Electricity",
		},
	}
	BillerData := DataDetail{
		TheData: dataList,
	}
	GetBillers := GetBillerMsg{
		ResponseCode: "00",
		Description:  "Success",
		MyData:       BillerData,
	}
	response.RespondJSON(w, http.StatusOK, GetBillers)
}

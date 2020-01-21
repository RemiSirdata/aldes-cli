package api

type Product struct {
	Indicator    *ProductIndicator `json:"indicator"`
	Modem        string            `json:"modem"`
	Reference    string            `json:"reference"`
	SerialNumber string            `json:"serial_number"`
	Type         string            `json:"type"`
	Name         string            `json:"name"`
	Address      string            `json:"address"`
	PostalCode   string            `json:"postal_code"`
	City         string            `json:"city"`
	Country      string            `json:"country"`
	UsureFiltre  float64           `json:"usure_filtre"`
	HasFilter    *bool             `json:"has_filter"`
	GpsLatitude  *float64          `json:"gps_latitude"`
	GpsLongitude *float64          `json:"gps_longitude"`
}

type ProductIndicator map[string]interface{}

func (p ProductIndicator) GetType() string {
	if t, found := p["indicatorType"]; found {
		return t.(string)
	}
	return ""
}

/**
Product type OQAI_Indicator
*/

type ProductIndicatorOQai struct {
	IndicatorType string `json:"indicatorType"`
}

type ProductIndicatorQaiData struct {
	Id           string                     `json:"id"`
	Temperature  float64                    `json:"temperature"`
	SerialNumber string                     `json:"serial_number"`
	AlertMode    bool                       `json:"alert_mode"`
	QaiIndex     *ProductIndicatorQaiMetric `json:"qaiIndex"`
	CO2          *ProductIndicatorQaiMetric `json:"co2"`
	H20          *ProductIndicatorQaiMetric `json:"h2o"`
	COV          *ProductIndicatorQaiMetric `json:"cov"`
	PM25         *ProductIndicatorQaiMetric `json:"pm25"`
}

type ProductIndicatorQaiMetric struct {
	CurrentValue float64   `json:"currentValue"`
	L1           []float64 `json:"l1"`
	L2           []float64 `json:"l2"`
	L3           []float64 `json:"l3"`
	L4           []float64 `json:"l4"`
}

/**
Product type AIR_Indicator
*/

type ProductIndicatorAir struct {
	IndicatorType   string      `json:"indicatorType"`
	AIR_OUTSIDE_TPT *float64    `json:"AIR_OUTSIDE_TPT"`
	AIR_REJECT_TPT  *float64    `json:"AIR_REJECT_TPT"`
	AIR_TII_EST     *float64    `json:"AIR_TII_EST"`
	AIR_EXTF_SPD    *float64    `json:"AIR_EXTF_SPD"`
	AIR_EXTF_TSN    *float64    `json:"AIR_EXTF_TSN"`
	AIR_VI_SPD      *float64    `json:"AIR_VI_SPD"`
	AIR_VV_TSN      *float64    `json:"AIR_VV_TSN"`
	AIR_FF_CPT      *float64    `json:"AIR_FF_CPT"`
	AIR_FFE_FLW     *int        `json:"AIR_FFE_FLW"`
	AIR_EXTF_FLW    *int        `json:"AIR_EXTF_FLW"`
	AIR_DEP_IND     *int        `json:"AIR_DEP_IND"`
	AIR_CVE_CSN     *int        `json:"AIR_CVE_CSN"`
	AIR_VI_CSN      *int        `json:"AIR_VI_CSN"`
	AIR_TTE_CSN     *int        `json:"AIR_TTE_CSN"`
	AIR_DTB_IND     *int        `json:"AIR_DTB_IND"`
	AIR_ECHANGE_PWR *int        `json:"AIR_ECHANGE_PWR"`
	AIR_SET_SPD     *int        `json:"AIR_SET_SPD"`
	AIR_EXCH_ENG    interface{} `json:"AIR_EXCH_ENG"`
	AIR_CO2_VMC     struct {
		ActualValue   interface{} `json:"actualValue"`
		PreviousValue interface{} `json:"previousValue"`
	} `json:"AIR_CO2_VMC"`
	AIR_TEMP_CAPT    *float64    `json:"AIR_TEMP_CAPT"`
	AIR_HYGR_CAPT    *float64    `json:"AIR_HYGR_CAPT"`
	AIR_CO2_CAPT     *float64    `json:"AIR_CO2_CAPT"`
	AIR_COV1_CAPT    *float64    `json:"AIR_COV1_CAPT"`
	AIR_COV2_CAPT    *float64    `json:"AIR_COV2_CAPT"`
	AIR_PM2_5_CAPT   *float64    `json:"AIR_PM2_5_CAPT"`
	AIR_PM10_CAPT    *float64    `json:"AIR_PM10_CAPT"`
	AIR_CURRENT_MODE string      `json:"AIR_CURRENT_MODE"`
	AIR_START_MODE   interface{} `json:"AIR_START_MODE"`
	AIR_END_MODE     interface{} `json:"air_end_mode"`
}

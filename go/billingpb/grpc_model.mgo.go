package billingpb

import (
	tools "github.com/paysuper/paysuper-tools/number"
	"go.mongodb.org/mongo-driver/bson"
)

type MgoDashboardAmountItemWithChart struct {
	Amount   float64                    `bson:"amount"`
	Currency string                     `bson:"currency"`
	Chart    []*DashboardChartItemFloat `bson:"chart"`
}

type MgoDashboardRevenueDynamicReportItem struct {
	Label    int64   `bson:"label"`
	Amount   float64 `bson:"amount"`
	Currency string  `bson:"currency"`
	Count    int64   `bson:"count"`
}

type MgoDashboardRevenueByCountryReportTop struct {
	Country string  `bson:"_id"`
	Amount  float64 `bson:"amount"`
}

type MgoDashboardRevenueByCountryReportChartItem struct {
	Label  int64   `bson:"label"`
	Amount float64 `bson:"amount"`
}

type MgoDashboardRevenueByCountryReport struct {
	Currency      string                                      `bson:"currency"`
	Top           []*DashboardRevenueByCountryReportTop       `bson:"top"`
	TotalCurrent  float64                                     `bson:"total"`
	TotalPrevious float64                                     `bson:"total_previous"`
	Chart         []*DashboardRevenueByCountryReportChartItem `bson:"chart"`
}

func (m *DashboardAmountItemWithChart) UnmarshalBSON(raw []byte) error {
	decoded := new(MgoDashboardAmountItemWithChart)
	err := bson.Unmarshal(raw, decoded)

	if err != nil {
		return err
	}

	m.AmountCurrent = tools.FormatAmount(decoded.Amount)
	m.Currency = decoded.Currency

	for _, v := range decoded.Chart {
		item := &DashboardChartItemFloat{
			Label: v.Label,
			Value: tools.FormatAmount(v.Value),
		}
		m.Chart = append(m.Chart, item)
	}

	return nil
}

func (m *DashboardRevenueDynamicReportItem) UnmarshalBSON(raw []byte) error {
	decoded := new(MgoDashboardRevenueDynamicReportItem)
	err := bson.Unmarshal(raw, decoded)

	if err != nil {
		return err
	}

	m.Label = decoded.Label
	m.Amount = tools.FormatAmount(decoded.Amount)
	m.Currency = decoded.Currency
	m.Count = decoded.Count

	return nil
}

func (m *DashboardRevenueByCountryReportTop) UnmarshalBSON(raw []byte) error {
	decoded := new(MgoDashboardRevenueByCountryReportTop)
	err := bson.Unmarshal(raw, decoded)

	if err != nil {
		return err
	}

	m.Amount = tools.FormatAmount(decoded.Amount)
	m.Country = decoded.Country

	return nil
}

func (m *DashboardRevenueByCountryReportChartItem) UnmarshalBSON(raw []byte) error {
	decoded := new(MgoDashboardRevenueByCountryReportChartItem)
	err := bson.Unmarshal(raw, decoded)

	if err != nil {
		return err
	}

	m.Amount = tools.FormatAmount(decoded.Amount)
	m.Label = decoded.Label

	return nil
}

func (m *DashboardRevenueByCountryReport) UnmarshalBSON(raw []byte) error {
	decoded := new(MgoDashboardRevenueByCountryReport)
	err := bson.Unmarshal(raw, decoded)

	if err != nil {
		return err
	}

	m.Currency = decoded.Currency
	m.Top = decoded.Top
	m.TotalCurrent = tools.FormatAmount(decoded.TotalCurrent)
	m.TotalPrevious = tools.FormatAmount(decoded.TotalPrevious)
	m.Chart = decoded.Chart

	return nil
}

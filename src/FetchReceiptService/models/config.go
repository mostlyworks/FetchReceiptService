package models

type PointConfig struct {
	TotalRoundedPoints            int     `json:"totalRoundedPoints"`
	TotalMutiplePoints            int     `json:"totalMutiplePoints"`
	TotalMutiple                  float64 `json:"totalMutiple"`
	TotalRoundMod                 float64 `json:"totalRoundMod"`
	ItemCountPoints               int     `json:"itemCountPoints"`
	ItemCountDivsor               int     `json:"itemCountDivsor"`
	ItemDescriptionMutiple        int     `json:"itemDescriptionMutiple"`
	ItemDescriptionPriceMutiplier float64 `json:"itemDescriptionPriceMutiplier"`
	PriceMutiplierRoundingPoints  int32   `json:"priceMutiplierRoundingPoints"`
	PurchaseDatePoints            int     `json:"purchaseDatePoints"`
	PurchaseDateCheckMod          int     `json:"purchaseDateCheckMod"`
	PurchaseTimeLowerBound        int     `json:"purchaseTimeLowerBound"`
	PurchaseTimeUpperBound        int     `json:"purchaseTimeUpperBound"`
	PurchaseTimePoints            int     `json:"purchaseTimePoints"`
	RetailerNamePointMutiplier    int     `json:"retailerNamePointMutiplier"`
	DefaultPointReturn            int     `json:"defaultPointReturn"`
}

// Probably pulled from SSM or something in the real world.
type DatabaseConfig struct {
	Host     string `json:"host"`
	Password string `json:"password"`
	Port     string `json:"port"`
}

type HttpServerConfig struct {
	UseHttps bool   `json:"useHttps"`
	Port     int    `json:"port"`
	Cert     string `json:"certFile"`
	Key      string `json:"keyFile"`
}

package model

type FeeTemplate struct {
	Max    string `json:"max"`
	Actual string `json:"actual"`
}

type MutualFundFee struct {
	Management FeeTemplate `json:"management"`
	Total      FeeTemplate `json:"total"`
}

type UnitHolderFee struct {
	Sale      FeeTemplate `json:"sale"`
	BuyBack   FeeTemplate `json:"buy_back"`
	SwitchIn  FeeTemplate `json:"switch_in"`
	SwitchOut FeeTemplate `json:"switch_out"`
}

type ReturnTemplate struct {
	FundCode      string        `json:"fund_code"`
	MutualFundFee MutualFundFee `json:"mutual_fund_fee"`
	UnitHolderFee UnitHolderFee `json:"unit_holder_fee"`
}

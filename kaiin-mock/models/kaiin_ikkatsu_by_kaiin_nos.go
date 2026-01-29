package models

type KaiinIkkatsuByKaiinNosResponse struct {
	Model    []KaiinIkkatsuItem `json:"model"`
	Errors   []ErrorItem        `json:"errors,omitempty"`
	Warnings []WarningItem      `json:"warnings,omitempty"`
}

type KaiinIkkatsuItem struct {
	KaiinNo             string  `json:"kaiinNo"`
	IkkatsuKaiinKbn     *string `json:"ikkatsuKaiinKbn,omitempty"`
	IkkatsuKaiinKbnName *string `json:"ikkatsuKaiinKbnName,omitempty"`
	Shumoku1            *string `json:"shumoku1,omitempty"`
	Shumoku2            *string `json:"shumoku2,omitempty"`
	Shumoku3            *string `json:"shumoku3,omitempty"`
	Shumoku4            *string `json:"shumoku4,omitempty"`
	BuCd                *string `json:"buCd,omitempty"`
	KaCd                *string `json:"kaCd,omitempty"`
	HanCd               *string `json:"hanCd,omitempty"`
	TantoCd             *string `json:"tantoCd,omitempty"`
	EigyoshoName        *string `json:"eigyoshoName,omitempty"`
	TantoName           *string `json:"tantoName,omitempty"`
	CreateDate          *string `json:"createDate,omitempty"`
	CreateTime          *string `json:"createTime,omitempty"`
	Creator             *string `json:"creator,omitempty"`
	TanmatsuId          *string `json:"tanmatsuId,omitempty"`
	UpdateDate          *string `json:"updateDate,omitempty"`
	UpdateTime          *string `json:"updateTime,omitempty"`
	VersionNo           *int    `json:"versionNo,omitempty"`
	TantoVersionNo      *int    `json:"tantoVersionNo,omitempty"`
}

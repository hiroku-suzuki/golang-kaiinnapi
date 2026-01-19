package models

type BukkakuResponse struct {
	Model    *BukkakuModel `json:"model"`
	Errors   []ErrorItem   `json:"errors,omitempty"`
	Warnings []WarningItem `json:"warnings,omitempty"`
}

type BukkakuModel struct {
	KaiinNo string `json:"kaiinNo"`

	Account       *string `json:"account,omitempty"`
	UserAccessKey *string `json:"userAccessKey,omitempty"`

	RiyoStartShinseiDate *string `json:"riyoStartShinseiDate,omitempty"`
	RiyoStartDate        *string `json:"riyoStartDate,omitempty"`
	RiyoStopShinseiDate  *string `json:"riyoStopShinseiDate,omitempty"`
	RiyoStopDate         *string `json:"riyoStopDate,omitempty"`

	BukkakuTelNo *string `json:"bukkakuTelNo,omitempty"`

	AtbbDisplayStartDate   *string `json:"atbbDisplayStartDate,omitempty"`
	AtbbDisplayStopDate    *string `json:"atbbDisplayStopDate,omitempty"`
	AtbbRelationStartDate  *string `json:"atbbRelationStartDate,omitempty"`
	AtbbRelationStopDate   *string `json:"atbbRelationStopDate,omitempty"`

	KeiyakuTantoshaCd *string `json:"keiyakuTantoshaCd,omitempty"`
	KaiyakuTantoshaCd *string `json:"kaiyakuTantoshaCd,omitempty"`
	BknApiGetTime     *string `json:"bknApiGetTime,omitempty"`

	VersionNo  *int    `json:"versionNo,omitempty"`
	CreateDate *string `json:"createDate,omitempty"`
	CreateTime *string `json:"createTime,omitempty"`
	Creator    *string `json:"creator,omitempty"`
	UpdateDate *string `json:"updateDate,omitempty"`
	UpdateTime *string `json:"updateTime,omitempty"`
	Updater    *string `json:"updater,omitempty"`

	HonKaiinNo *string `json:"honKaiinNo,omitempty"`

	Biko1 *string `json:"biko1,omitempty"`
	Biko2 *string `json:"biko2,omitempty"`
	Biko3 *string `json:"biko3,omitempty"`
}
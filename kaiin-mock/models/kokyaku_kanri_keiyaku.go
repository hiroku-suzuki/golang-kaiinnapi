package models

type KokyakuKanriKeiyakuResponse struct {
	Model    *KokyakuKanriKeiyakuModel `json:"model"`
	Errors   []ErrorItem               `json:"errors,omitempty"`
	Warnings []WarningItem             `json:"warnings,omitempty"`
}

type KokyakuKanriKeiyakuModel struct {
	KaiinNo string `json:"kaiinNo"`

	RiyoKbn        *string `json:"riyoKbn,omitempty"`
	RiyoKbnName    *string `json:"riyoKbnName,omitempty"`
	MailBoxAddress *string `json:"mailBoxAddress,omitempty"`

	RiyoStartShinseiDate *string `json:"riyoStartShinseiDate,omitempty"`
	RiyoStartDate        *string `json:"riyoStartDate,omitempty"`

	KeiyakuTantoCd   *string `json:"keiyakuTantoCd,omitempty"`
	KeiyakuTantoName *string `json:"keiyakuTantoName,omitempty"`

	RiyoStopShinseiDate *string `json:"riyoStopShinseiDate,omitempty"`
	RiyoStopDate        *string `json:"riyoStopDate,omitempty"`

	KaiyakuTantoCd   *string `json:"kaiyakuTantoCd,omitempty"`
	KaiyakuTantoName *string `json:"kaiyakuTantoName,omitempty"`
	KaiyakuRiyu      *string `json:"kaiyakuRiyu,omitempty"`

	Biko5 *string `json:"biko5,omitempty"`
	Biko6 *string `json:"biko6,omitempty"`

	CreateDate *string `json:"createDate,omitempty"`
	CreateTime *string `json:"createTime,omitempty"`
	Creator    *string `json:"creator,omitempty"`
	UpdateDate *string `json:"updateDate,omitempty"`
	UpdateTime *string `json:"updateTime,omitempty"`
	Updater    *string `json:"updater,omitempty"`

	ThankyouRiyoStartDate   *string `json:"thankyouRiyoStartDate,omitempty"`
	ThankyouRiyoShinseiDate *string `json:"thankyouRiyoShinseiDate,omitempty"`

	VersionNo          *int    `json:"versionNo,omitempty"`
	FirstRiyoStartDate *string `json:"firstRiyoStartDate,omitempty"`
}

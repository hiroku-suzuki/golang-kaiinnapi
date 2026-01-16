package models

type SeikyuMasterResponse struct {
	Model    *SeikyuMasterModel `json:"model"`
	Errors   []ErrorItem        `json:"errors,omitempty"`
	Warnings []WarningItem      `json:"warnings,omitempty"`
}

type SeikyuMasterModel struct {
	KaiinNo      string  `json:"kaiinNo"`
	SeikyuKaiinNo *string `json:"seikyuKaiinNo,omitempty"`
	NyukinKaiinNo *string `json:"nyukinKaiinNo,omitempty"`

	SeikyuShimeDate   *string `json:"seikyuShimeDate,omitempty"`
	ShiharaiJyokenMm  *string `json:"shiharaiJyokenMm,omitempty"`
	ShiharaiJyokenDd  *string `json:"shiharaiJyokenDd,omitempty"`
	ShiharaiKbn       *string `json:"shiharaiKbn,omitempty"`
	ShiharaiKbnName   *string `json:"shiharaiKbnName,omitempty"`
	SeikyuKbn         *string `json:"seikyuKbn,omitempty"`
	SeikyuKbnName     *string `json:"seikyuKbnName,omitempty"`
	SeikyuStopMm      *string `json:"seikyuStopMm,omitempty"`
	YusoRyoshuKbn     *string `json:"yusoRyoshuKbn,omitempty"`
	YusoRyoshuKbnName *string `json:"yusoRyoshuKbnName,omitempty"`

	JifuriBankCd      *string `json:"jifuriBankCd,omitempty"`
	JifuriBankNm      *string `json:"jifuriBankNm,omitempty"`
	JifuriBankKana    *string `json:"jifuriBankKana,omitempty"`
	JifuriShitenNm    *string `json:"jifuriShitenNm,omitempty"`
	JifuriShitenKana  *string `json:"jifuriShitenKana,omitempty"`
	JifuriKozaSbt     *string `json:"jifuriKozaSbt,omitempty"`
	JifuriKozaNo      *string `json:"jifuriKozaNo,omitempty"`
	JifuriKozaNm      *string `json:"jifuriKozaNm,omitempty"`
	JifuriTorokuDate  *string `json:"jifuriTorokuDate,omitempty"`

	DaikoGyoshaCd1 *string `json:"daikoGyoshaCd1,omitempty"`
	DaikoGyoshaCd2 *int    `json:"daikoGyoshaCd2,omitempty"`
	JifuriTeishiMm *string `json:"jifuriTeishiMm,omitempty"`

	IsTokushuNyukin  *bool   `json:"isTokushuNyukin,omitempty"`
	TokushuNyukinMemo *string `json:"tokushuNyukinMemo,omitempty"`

	VersionNo  *int    `json:"versionNo,omitempty"`
	CreateDate *string `json:"createDate,omitempty"`
	CreateTime *string `json:"createTime,omitempty"`
	Creator    *string `json:"creator,omitempty"`
	UpdateDate *string `json:"updateDate,omitempty"`
	UpdateTime *string `json:"updateTime,omitempty"`
	Updater    *string `json:"updater,omitempty"`
}
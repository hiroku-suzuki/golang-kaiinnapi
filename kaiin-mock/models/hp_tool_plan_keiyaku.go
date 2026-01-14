package models

type HpToolPlanKeiyakuResponse struct {
	Model    *HpToolPlanKeiyakuModel `json:"model"`
	Errors   []ErrorItem             `json:"errors,omitempty"`
	Warnings []WarningItem           `json:"warnings,omitempty"`
}

type HpToolPlanKeiyakuModel struct {
	KaiinNo string `json:"kaiinNo"`

	KameitenId  *int    `json:"kameitenId,omitempty"`
	KeiyakuType *string `json:"keiyakuType,omitempty"`
	KaiinName   *string `json:"kaiinName,omitempty"`
	Shogo       *string `json:"shogo,omitempty"`
	Domain      *string `json:"domain,omitempty"`

	RiyoStartShinseiDate *string `json:"riyoStartShinseiDate,omitempty"`
	RiyoStartDate        *string `json:"riyoStartDate,omitempty"`
	KeiyakuTantoCd       *string `json:"keiyakuTantoCd,omitempty"`
	KeiyakuTantoName     *string `json:"keiyakuTantoName,omitempty"`
	KeiyakuEigyoshoName  *string `json:"keiyakuEigyoshoName,omitempty"`

	RiyoEndShinseiDate  *string `json:"riyoEndShinseiDate,omitempty"`
	RiyoEndDate         *string `json:"riyoEndDate,omitempty"`
	KaiyakuTantoCd      *string `json:"kaiyakuTantoCd,omitempty"`
	KaiyakuTantoName    *string `json:"kaiyakuTantoName,omitempty"`
	KaiyakuEigyoshoName *string `json:"kaiyakuEigyoshoName,omitempty"`

	FtpServer    *string `json:"ftpServer,omitempty"`
	FtpPort      *string `json:"ftpPort,omitempty"`
	FtpUserId    *string `json:"ftpUserId,omitempty"`
	FtpDirectory *string `json:"ftpDirectory,omitempty"`
	IsFtpPasv    *bool   `json:"isFtpPasv,omitempty"`

	ConpaneUrl    *string `json:"conpaneUrl,omitempty"`
	ConpaneUserId *string `json:"conpaneUserId,omitempty"`

	IsGroup *bool `json:"isGroup,omitempty"`
	IsOya   *bool `json:"isOya,omitempty"`

	NijiKokokuRiyoStartShinseiDate *string `json:"nijiKokokuRiyoStartShinseiDate,omitempty"`
	NijiKokokuRiyoStartDate        *string `json:"nijiKokokuRiyoStartDate,omitempty"`
	NijiKokokuKeiyakuTantoCd       *string `json:"nijiKokokuKeiyakuTantoCd,omitempty"`
	NijiKokokuKeiyakuTantoName     *string `json:"nijiKokokuKeiyakuTantoName,omitempty"`
	NijiKokokuKeiyakuEigyoshoName  *string `json:"nijiKokokuKeiyakuEigyoshoName,omitempty"`
	NijiKokokuRiyoEndShinseiDate   *string `json:"nijiKokokuRiyoEndShinseiDate,omitempty"`
	NijiKokokuRiyoEndDate          *string `json:"nijiKokokuRiyoEndDate,omitempty"`
	NijiKokokuKaiyakuTantoCd       *string `json:"nijiKokokuKaiyakuTantoCd,omitempty"`
	NijiKokokuKaiyakuTantoName     *string `json:"nijiKokokuKaiyakuTantoName,omitempty"`
	NijiKokokuKaiyakuEigyoshoName  *string `json:"nijiKokokuKaiyakuEigyoshoName,omitempty"`

	IsBukkenGroup *bool `json:"isBukkenGroup,omitempty"`
	IsBukkenOya   *bool `json:"isBukkenOya,omitempty"`

	KoushinDate      *string `json:"koushinDate,omitempty"`
	KoukaiTeishiDate *string `json:"koukaiTeishiDate,omitempty"`
	ShokaiKoukaiDate *string `json:"shokaiKoukaiDate,omitempty"`

	CreateDate *string `json:"createDate,omitempty"`
	CreateTime *string `json:"createTime,omitempty"`
	UpdateDate *string `json:"updateDate,omitempty"`
	UpdateTime *string `json:"updateTime,omitempty"`

	KeiyakuStatus              *string `json:"keiyakuStatus,omitempty"`
	KeiyakuPlan                *string `json:"keiyakuPlan,omitempty"`
	YoyakuKeiyakuPlan          *string `json:"yoyakuKeiyakuPlan,omitempty"`
	YoyakuRiyoStartShinseiDate *string `json:"yoyakuRiyoStartShinseiDate,omitempty"`
	YoyakuRiyoStartDate        *string `json:"yoyakuRiyoStartDate,omitempty"`
	YoyakuKeiyakuTantoCd       *string `json:"yoyakuKeiyakuTantoCd,omitempty"`
	YoyakuKeiyakuTantoName     *string `json:"yoyakuKeiyakuTantoName,omitempty"`
	YoyakuKeiyakuEigyoshoName  *string `json:"yoyakuKeiyakuEigyoshoName,omitempty"`

	ShokaiRiyoStartDate *string `json:"shokaiRiyoStartDate,omitempty"`

	ChizuRiyoStartShinseiDate *string `json:"chizuRiyoStartShinseiDate,omitempty"`
	ChizuRiyoStartDate        *string `json:"chizuRiyoStartDate,omitempty"`
	ChizuKeiyakuTantoCd       *string `json:"chizuKeiyakuTantoCd,omitempty"`
	ChizuKeiyakuTantoName     *string `json:"chizuKeiyakuTantoName,omitempty"`
	ChizuKeiyakuEigyoshoName  *string `json:"chizuKeiyakuEigyoshoName,omitempty"`
	ChizuRiyoEndShinseiDate   *string `json:"chizuRiyoEndShinseiDate,omitempty"`
	ChizuRiyoEndDate          *string `json:"chizuRiyoEndDate,omitempty"`
	ChizuKaiyakuTantoCd       *string `json:"chizuKaiyakuTantoCd,omitempty"`
	ChizuKaiyakuTantoName     *string `json:"chizuKaiyakuTantoName,omitempty"`
	ChizuKaiyakuEigyoshoName  *string `json:"chizuKaiyakuEigyoshoName,omitempty"`

	TopOriginalRiyoStartShinseiDate *string `json:"topOriginalRiyoStartShinseiDate,omitempty"`
	TopOriginalRiyoStartDate        *string `json:"topOriginalRiyoStartDate,omitempty"`
	TopOriginalKeiyakuTantoCd       *string `json:"topOriginalKeiyakuTantoCd,omitempty"`
	TopOriginalKeiyakuTantoName     *string `json:"topOriginalKeiyakuTantoName,omitempty"`
	TopOriginalKeiyakuEigyoshoName  *string `json:"topOriginalKeiyakuEigyoshoName,omitempty"`
	TopOriginalRiyoEndShinseiDate   *string `json:"topOriginalRiyoEndShinseiDate,omitempty"`
	TopOriginalRiyoEndDate          *string `json:"topOriginalRiyoEndDate,omitempty"`
	TopOriginalKaiyakuTantoCd       *string `json:"topOriginalKaiyakuTantoCd,omitempty"`
	TopOriginalKaiyakuTantoName     *string `json:"topOriginalKaiyakuTantoName,omitempty"`
	TopOriginalKaiyakuEigyoshoName  *string `json:"topOriginalKaiyakuEigyoshoName,omitempty"`

	FdpRiyoStartDate *string `json:"fdpRiyoStartDate,omitempty"`
	FdpRiyoEndDate   *string `json:"fdpRiyoEndDate,omitempty"`
}

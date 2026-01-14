package models

// レスポンスは一旦「空の model {}」を返す
// errors / warnings は必要になったら追加で返す

type EngineRentalResponse struct {
	Model    *EngineRentalModel `json:"model"`
	Errors   []ErrorItem        `json:"errors,omitempty"`
	Warnings []WarningItem      `json:"warnings,omitempty"`
}

type EngineRentalModel struct {
	EngineRentalId      *string `json:"engineRentalId,omitempty"`
	SiteName            *string `json:"siteName,omitempty"`
	TownInfoInterlockId *string `json:"townInfoInterlockId,omitempty"`
	Memo1               *string `json:"memo1,omitempty"`
	Memo2               *string `json:"memo2,omitempty"`
	VersionNo           *int    `json:"versionNo,omitempty"`
	CreateDate          *string `json:"createDate,omitempty"`
	CreateTime          *string `json:"createTime,omitempty"`
	Creator             *string `json:"creator,omitempty"`
	UpdateDate          *string `json:"updateDate,omitempty"`
	UpdateTime          *string `json:"updateTime,omitempty"`
	Updater             *string `json:"updater,omitempty"`

	ControlPanelKaiins []EngineRentalControlPanelKaiin `json:"controlPanelKaiins,omitempty"`
	MemberKaiins       []EngineRentalMemberKaiin       `json:"memberKaiins,omitempty"`
	BasicServices      []EngineRentalBasicService      `json:"basicServices,omitempty"`
	OptionServices     []EngineRentalOptionService     `json:"optionServices,omitempty"`
	NijikokokuAreas    []EngineRentalNijikokokuArea    `json:"nijikokokuAreas,omitempty"`
}

type EngineRentalControlPanelKaiin struct {
	RecordNo      *string `json:"recordNo,omitempty"`
	KaiinNo       *string `json:"kaiinNo,omitempty"`
	RiyoStartDate *string `json:"riyoStartDate,omitempty"`
	RiyoStopDate  *string `json:"riyoStopDate,omitempty"`
	VersionNo     *int    `json:"versionNo,omitempty"`
	CreateDate    *string `json:"createDate,omitempty"`
	CreateTime    *string `json:"createTime,omitempty"`
	Creator       *string `json:"creator,omitempty"`
	UpdateDate    *string `json:"updateDate,omitempty"`
	UpdateTime    *string `json:"updateTime,omitempty"`
	Updater       *string `json:"updater,omitempty"`
}

type EngineRentalMemberKaiin struct {
	RecordNo      *string `json:"recordNo,omitempty"`
	KaiinNo       *string `json:"kaiinNo,omitempty"`
	KaiinLinkNo   *string `json:"kaiinLinkNo,omitempty"`
	RiyoStartDate *string `json:"riyoStartDate,omitempty"`
	RiyoStopDate  *string `json:"riyoStopDate,omitempty"`
	VersionNo     *int    `json:"versionNo,omitempty"`
	CreateDate    *string `json:"createDate,omitempty"`
	CreateTime    *string `json:"createTime,omitempty"`
	Creator       *string `json:"creator,omitempty"`
	UpdateDate    *string `json:"updateDate,omitempty"`
	UpdateTime    *string `json:"updateTime,omitempty"`
	Updater       *string `json:"updater,omitempty"`
}

type EngineRentalBasicService struct {
	RecordNo            *string `json:"recordNo,omitempty"`
	SyohinCd            *string `json:"syohinCd,omitempty"`
	SyohinName          *string `json:"syohinName,omitempty"`
	RiyoStartSinseiDate *string `json:"riyoStartSinseiDate,omitempty"`
	RiyoStopSinseiDate  *string `json:"riyoStopSinseiDate,omitempty"`
	RiyoStartDate       *string `json:"riyoStartDate,omitempty"`
	RiyoStopDate        *string `json:"riyoStopDate,omitempty"`
	KeiyakuTantoCd      *string `json:"keiyakuTantoCd,omitempty"`
	KeiyakuTantoName    *string `json:"keiyakuTantoName,omitempty"`
	KaiyakuTantoCd      *string `json:"kaiyakuTantoCd,omitempty"`
	KaiyakuTantoName    *string `json:"kaiyakuTantoName,omitempty"`
	VersionNo           *int    `json:"versionNo,omitempty"`
	CreateDate          *string `json:"createDate,omitempty"`
	CreateTime          *string `json:"createTime,omitempty"`
	Creator             *string `json:"creator,omitempty"`
	UpdateDate          *string `json:"updateDate,omitempty"`
	UpdateTime          *string `json:"updateTime,omitempty"`
	Updater             *string `json:"updater,omitempty"`
}

type EngineRentalOptionService struct {
	RecordNo            *string `json:"recordNo,omitempty"`
	ServiceType         *string `json:"serviceType,omitempty"`
	ServiceName         *string `json:"serviceName,omitempty"`
	RiyoStartSinseiDate *string `json:"riyoStartSinseiDate,omitempty"`
	RiyoStopSinseiDate  *string `json:"riyoStopSinseiDate,omitempty"`
	RiyoStartDate       *string `json:"riyoStartDate,omitempty"`
	RiyoStopDate        *string `json:"riyoStopDate,omitempty"`
	KeiyakuTantoCd      *string `json:"keiyakuTantoCd,omitempty"`
	KeiyakuTantoName    *string `json:"keiyakuTantoName,omitempty"`
	KaiyakuTantoCd      *string `json:"kaiyakuTantoCd,omitempty"`
	KaiyakuTantoName    *string `json:"kaiyakuTantoName,omitempty"`
	VersionNo           *int    `json:"versionNo,omitempty"`
	CreateDate          *string `json:"createDate,omitempty"`
	CreateTime          *string `json:"createTime,omitempty"`
	Creator             *string `json:"creator,omitempty"`
	UpdateDate          *string `json:"updateDate,omitempty"`
	UpdateTime          *string `json:"updateTime,omitempty"`
	Updater             *string `json:"updater,omitempty"`
}

type EngineRentalNijikokokuArea struct {
	RecordNo      *string `json:"recordNo,omitempty"`
	KenCd         *string `json:"kenCd,omitempty"`
	KenName       *string `json:"kenName,omitempty"`
	RiyoStartDate *string `json:"riyoStartDate,omitempty"`
	RiyoStopDate  *string `json:"riyoStopDate,omitempty"`
	VersionNo     *int    `json:"versionNo,omitempty"`
	CreateDate    *string `json:"createDate,omitempty"`
	CreateTime    *string `json:"createTime,omitempty"`
	Creator       *string `json:"creator,omitempty"`
	UpdateDate    *string `json:"updateDate,omitempty"`
	UpdateTime    *string `json:"updateTime,omitempty"`
	Updater       *string `json:"updater,omitempty"`
}

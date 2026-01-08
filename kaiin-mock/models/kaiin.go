package models

type KaiinResponse struct {
	Model    *KaiinModel   `json:"model"`
	Errors   []ErrorItem   `json:"errors,omitempty"`
	Warnings []WarningItem `json:"warnings,omitempty"`
}

type KaiinModel struct {
	VersionNo                  *int            `json:"versionNo,omitempty"`
	KaiinNo                    string          `json:"kaiinNo"`
	ShozaichiCd1               *string         `json:"shozaichiCd1,omitempty"`
	ShozaichiCd2               *string         `json:"shozaichiCd2,omitempty"`
	PostNo                     *string         `json:"postNo,omitempty"`
	TodofukenName              *string         `json:"todofukenName,omitempty"`
	CityName                   *string         `json:"cityName,omitempty"`
	TownName                   *string         `json:"townName,omitempty"`
	Banchi                     *string         `json:"banchi,omitempty"`
	BuildingName               *string         `json:"buildingName,omitempty"`
	RailLineCd                 *string         `json:"railLineCd,omitempty"`
	RailLineName               *string         `json:"railLineName,omitempty"`
	StationCd                  *string         `json:"stationCd,omitempty"`
	StationName                *string         `json:"stationName,omitempty"`
	SearchForRailLineStationCd *string         `json:"searchForRailLineStationCd,omitempty"`
	TohoJikan                  *int            `json:"tohoJikan,omitempty"`
	TohoKyori                  *int            `json:"tohoKyori,omitempty"`
	BusJikan                   *int            `json:"busJikan,omitempty"`
	BusteiJikan                *int            `json:"busteiJikan,omitempty"`
	BusteiKyori                *int            `json:"busteiKyori,omitempty"`
	BusteiName                 *string         `json:"busteiName,omitempty"`
	Seikyusakis                []Seikyusaki    `json:"seikyusakis,omitempty"`
	DaihyoshaInfo              *DaihyoshaInfo  `json:"daihyoshaInfo,omitempty"`
	Shihonkin                  *int            `json:"shihonkin,omitempty"`
	SetsuritsuDate             *string         `json:"setsuritsuDate,omitempty"`
	StaffCnt                   *int            `json:"staffCnt,omitempty"`
	DaihyoTorihikiShunin       *string         `json:"daihyoTorihikiShunin,omitempty"`
	TorihikiShuninshaCnt       *int            `json:"torihikiShuninshaCnt,omitempty"`
	EigyoTimeInfo              *EigyoTimeInfo  `json:"eigyoTimeInfo,omitempty"`
	TeikyuDateInfo             *TeikyuDateInfo `json:"teikyuDateInfo,omitempty"`
	TenpoKbn                   *string         `json:"tenpoKbn,omitempty"`
	TenpoKbnName               *string         `json:"tenpoKbnName,omitempty"`
	TenpoKaisuKbn              *string         `json:"tenpoKaisuKbn,omitempty"`
	TenpoKaisuKbnName          *string         `json:"tenpoKaisuKbnName,omitempty"`
	IsGyotaiMotozuke           *bool           `json:"isGyotaiMotozuke,omitempty"`
	IsGyotaiKyakuzuke          *bool           `json:"isGyotaiKyakuzuke,omitempty"`
	TenpoSekininshaName        *string         `json:"tenpoSekininshaName,omitempty"`
	HoshoKyokaiCd              *string         `json:"hoshoKyokaiCd,omitempty"`
	HoshoKyokaiName            *string         `json:"hoshoKyokaiName,omitempty"`
	BlackKaiinKbn              *string         `json:"blackKaiinKbn,omitempty"`
	BlackKaiinKbnName          *string         `json:"blackKaiinKbnName,omitempty"`
	IsRiyoStop                 *bool           `json:"isRiyoStop,omitempty"`
	KaiinLat                   *float64        `json:"kaiinLat,omitempty"`
	KaiinLng                   *float64        `json:"kaiinLng,omitempty"`
	KaiinGmatchLv              *string         `json:"kaiinGmatchLv,omitempty"`
	KaiinLatHosei              *float64        `json:"kaiinLatHosei,omitempty"`
	KaiinLngHosei              *float64        `json:"kaiinLngHosei,omitempty"`
	Biko                       *string         `json:"biko,omitempty"`
	IsMaisokuHaifu             *bool           `json:"isMaisokuHaifu,omitempty"`
	IsMenkyo                   *bool           `json:"isMenkyo,omitempty"`
	KiyakuShubetsu             *string         `json:"kiyakuShubetsu,omitempty"`
	CreateDate                 *string         `json:"createDate,omitempty"`
	CreateTime                 *string         `json:"createTime,omitempty"`
	Creator                    *string         `json:"creator,omitempty"`
	UpdateDate                 *string         `json:"updateDate,omitempty"`
	UpdateTime                 *string         `json:"updateTime,omitempty"`
	Updater                    *string         `json:"updater,omitempty"`
	KotoriCd                   *string         `json:"kotoriCd,omitempty"`
	KotoriName                 *string         `json:"kotoriName,omitempty"`
	Shozokus                   []Shozoku       `json:"shozokus,omitempty"`
	GyoseiShobuns              []GyoseiShobun  `json:"gyoseiShobuns,omitempty"`
	MenkyoKenCd                *string         `json:"menkyoKenCd,omitempty"`
	MenkyoKenName              *string         `json:"menkyoKenName,omitempty"`
	MenkyoKbn                  *string         `json:"menkyoKbn,omitempty"`
	MenkyoKbnName              *string         `json:"menkyoKbnName,omitempty"`
	MenkyoUpdateCnt            *int            `json:"menkyoUpdateCnt,omitempty"`
	MenkyoNo                   *string         `json:"menkyoNo,omitempty"`
	MenkyoName                 *string         `json:"menkyoName,omitempty"`
	MenkyoYukoDate             *string         `json:"menkyoYukoDate,omitempty"`
	MenkyoUmuKbn               *string         `json:"menkyoUmuKbn,omitempty"`
	MenkyoUmuKbnName           *string         `json:"menkyoUmuKbnName,omitempty"`
	BuCd                       *string         `json:"buCd,omitempty"`
	KaCd                       *string         `json:"kaCd,omitempty"`
	HanCd                      *string         `json:"hanCd,omitempty"`
	KaiinLinkNo                *string         `json:"kaiinLinkNo,omitempty"`
	SeikiShogo                 *Shogo          `json:"seikiShogo,omitempty"`
	IppanShogo                 *Shogo          `json:"ippanShogo,omitempty"`
	HanbaiShogo                *Shogo          `json:"hanbaiShogo,omitempty"`
	BukkenShogo                *Shogo          `json:"bukkenShogo,omitempty"`
	Contact                    *Contact        `json:"contact,omitempty"`
	IsNotUsedMapSearch         *bool           `json:"isNotUsedMapSearch,omitempty"`
	IsNotUsedMapDisplay        *bool           `json:"isNotUsedMapDisplay,omitempty"`
	StatusCd                   *string         `json:"statusCd,omitempty"`
}

type Seikyusaki struct {
	KaiEdaNo                *string `json:"kaiEdaNo,omitempty"`
	SeikyusakiName          *string `json:"seikyusakiName,omitempty"`
	SeikyusakiPostNo        *string `json:"seikyusakiPostNo,omitempty"`
	SeikyusakiShozaichiCd1  *string `json:"seikyusakiShozaichiCd1,omitempty"`
	SeikyusakiShozaichiCd2  *string `json:"seikyusakiShozaichiCd2,omitempty"`
	SeikyusakiTodofukenName *string `json:"seikyusakiTodofukenName,omitempty"`
	SeikyusakiCityName      *string `json:"seikyusakiCityName,omitempty"`
	SeikyusakiTownName      *string `json:"seikyusakiTownName,omitempty"`
	SeikyusakiBanchi        *string `json:"seikyusakiBanchi,omitempty"`
	SeikyusakiBuildingName  *string `json:"seikyusakiBuildingName,omitempty"`
	SeikyusakiSofusakiName  *string `json:"seikyusakiSofusakiName,omitempty"`
	CreateDate              *string `json:"createDate,omitempty"`
	CreateTime              *string `json:"createTime,omitempty"`
	Creator                 *string `json:"creator,omitempty"`
	UpdateDate              *string `json:"updateDate,omitempty"`
	UpdateTime              *string `json:"updateTime,omitempty"`
	Updater                 *string `json:"updater,omitempty"`
}

type DaihyoshaInfo struct {
	DaihyoshaName          *string `json:"daihyoshaName,omitempty"`
	DaihyoshaKana          *string `json:"daihyoshaKana,omitempty"`
	DaihyoshaBirthDate     *string `json:"daihyoshaBirthDate,omitempty"`
	DaihyoshaShozaichiCd1  *string `json:"daihyoshaShozaichiCd1,omitempty"`
	DaihyoshaShozaichiCd2  *string `json:"daihyoshaShozaichiCd2,omitempty"`
	DaihyoshaPostNo        *string `json:"daihyoshaPostNo,omitempty"`
	DaihyoshaTodofukenName *string `json:"daihyoshaTodofukenName,omitempty"`
	DaihyoshaCityName      *string `json:"daihyoshaCityName,omitempty"`
	DaihyoshaTownName      *string `json:"daihyoshaTownName,omitempty"`
	DaihyoshaBanchi        *string `json:"daihyoshaBanchi,omitempty"`
	DaihyoshaBuildingName  *string `json:"daihyoshaBuildingName,omitempty"`
	DaihyoshaTel           *string `json:"daihyoshaTel,omitempty"`
}

type EigyoTimeInfo struct {
	EigyoStartTime *string `json:"eigyoStartTime,omitempty"`
	EigyoEndTime   *string `json:"eigyoEndTime,omitempty"`
	OtherEigyoTime *string `json:"otherEigyoTime,omitempty"`
}

type TeikyuDateInfo struct {
	TeikyuDay      *string `json:"teikyuDay,omitempty"`
	OtherTeikyuDay *string `json:"otherTeikyuDay,omitempty"`
}

type Shozoku struct {
	KaiEdaNo                *string `json:"kaiEdaNo,omitempty"`
	ShozokuDantaiCd         *string `json:"shozokuDantaiCd,omitempty"`
	ShozokuDantaiName       *string `json:"shozokuDantaiName,omitempty"`
	ShozokuDantaiNyukaiDate *string `json:"shozokuDantaiNyukaiDate,omitempty"`
	ShozokuDantaiTaikaiDate *string `json:"shozokuDantaiTaikaiDate,omitempty"`
	CreateDate              *string `json:"createDate,omitempty"`
	CreateTime              *string `json:"createTime,omitempty"`
	Creator                 *string `json:"creator,omitempty"`
	UpdateDate              *string `json:"updateDate,omitempty"`
	UpdateTime              *string `json:"updateTime,omitempty"`
	Updater                 *string `json:"updater,omitempty"`
}

type GyoseiShobun struct {
	SeqNo                 *int    `json:"seqNo,omitempty"`
	IsKokaiStop           *bool   `json:"isKokaiStop,omitempty"`
	GyoseiShobunNaiyoKbn  *string `json:"gyoseiShobunNaiyoKbn,omitempty"`
	GyoseiShobunStartDate *string `json:"gyoseiShobunStartDate,omitempty"`
	GyoseiShobunEndDate   *string `json:"gyoseiShobunEndDate,omitempty"`
	TorihikiStopKikanFrom *string `json:"torihikiStopKikanFrom,omitempty"`
	TorihikiStopKikanTo   *string `json:"torihikiStopKikanTo,omitempty"`
	CreateDate            *string `json:"createDate,omitempty"`
	CreateTime            *string `json:"createTime,omitempty"`
	Creator               *string `json:"creator,omitempty"`
	UpdateDate            *string `json:"updateDate,omitempty"`
	UpdateTime            *string `json:"updateTime,omitempty"`
	Updater               *string `json:"updater,omitempty"`
}

type Shogo struct {
	ShogoName     *string `json:"shogoName,omitempty"`
	ShogoKana     *string `json:"shogoKana,omitempty"`
	ShogoRyakusho *string `json:"shogoRyakusho,omitempty"`
}

type Contact struct {
	DaihyoTel         *string `json:"daihyoTel,omitempty"`
	SeiyakuCheckTel   *string `json:"seiyakuCheckTel,omitempty"`
	IppanTel          *string `json:"ippanTel,omitempty"`
	DaihyoFax         *string `json:"daihyoFax,omitempty"`
	MailToFax         *string `json:"mailToFax,omitempty"`
	ZumenDeliveryFax  *string `json:"zumenDeliveryFax,omitempty"`
	DaihyoMail        *string `json:"daihyoMail,omitempty"`
	HankyoReceiveMail *string `json:"hankyoReceiveMail,omitempty"`
	SeiyakuReportMail *string `json:"seiyakuReportMail,omitempty"`
}

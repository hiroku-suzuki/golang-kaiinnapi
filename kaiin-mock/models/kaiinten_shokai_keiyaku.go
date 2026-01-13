package models

type KaiintenShokaiKeiyakuByKaiinNoResponse struct {
	Model    *KaiintenShokaiKeiyakuByKaiinNoModel `json:"model"`
	Errors   []ErrorItem                          `json:"errors,omitempty"`
	Warnings []WarningItem                        `json:"warnings,omitempty"`
}

type KaiintenShokaiKeiyakuByKaiinNoModel struct {
	KaiinLinkNo string `json:"kaiinLinkNo"`

	KaiintenShokaiKeiyakuKaiin []KaiintenShokaiKeiyakuKaiinItem `json:"kaiintenShokaiKeiyakuKaiin,omitempty"`

	IchiranPr *string `json:"ichiranPr,omitempty"`
	ShosaiPr  *string `json:"shosaiPr,omitempty"`

	DispNaiyoKbn     *string `json:"dispNaiyoKbn,omitempty"`
	DispNaiyoKbnName *string `json:"dispNaiyoKbnName,omitempty"`
	KokaiKbn         *string `json:"kokaiKbn,omitempty"`
	KokaiKbnName     *string `json:"kokaiKbnName,omitempty"`

	IsKokaiFs            *bool `json:"isKokaiFs,omitempty"`
	IsKokaiOnline        *bool `json:"isKokaiOnline,omitempty"`
	IsKokaiChintai       *bool `json:"isKokaiChintai,omitempty"`
	IsKokaiMyRoom        *bool `json:"isKokaiMyRoom,omitempty"`
	IsKokaiNavi          *bool `json:"isKokaiNavi,omitempty"`
	IsKokaiMaisoku       *bool `json:"isKokaiMaisoku,omitempty"`
	IsKokaiReshizu       *bool `json:"isKokaiReshizu,omitempty"`
	IsKokaiBaiyazu       *bool `json:"isKokaiBaiyazu,omitempty"`
	IsKokaiSelazu        *bool `json:"isKokaiSelazu,omitempty"`
	IsKokaiNagano        *bool `json:"isKokaiNagano,omitempty"`
	IsKokaiMaisokuOnline *bool `json:"isKokaiMaisokuOnline,omitempty"`
	IsKokaiVisa          *bool `json:"isKokaiVisa,omitempty"`
	IsKokaiYobi1         *bool `json:"isKokaiYobi1,omitempty"`
	IsKokaiYobi2         *bool `json:"isKokaiYobi2,omitempty"`
	IsKokaiYobi3         *bool `json:"isKokaiYobi3,omitempty"`

	BukkenLinkKbn   *string `json:"bukkenLinkKbn,omitempty"`
	IsBukkenShokai  *bool   `json:"isBukkenShokai,omitempty"`
	IsMailToFaxFuka *bool   `json:"isMailToFaxFuka,omitempty"`

	PatternKbn     *string `json:"patternKbn,omitempty"`
	PatternKbnName *string `json:"patternKbnName,omitempty"`
	IsFlashRiyo    *bool   `json:"isFlashRiyo,omitempty"`

	// isTokucho1..80
	IsTokucho1  *bool `json:"isTokucho1,omitempty"`
	IsTokucho2  *bool `json:"isTokucho2,omitempty"`
	IsTokucho3  *bool `json:"isTokucho3,omitempty"`
	IsTokucho4  *bool `json:"isTokucho4,omitempty"`
	IsTokucho5  *bool `json:"isTokucho5,omitempty"`
	IsTokucho6  *bool `json:"isTokucho6,omitempty"`
	IsTokucho7  *bool `json:"isTokucho7,omitempty"`
	IsTokucho8  *bool `json:"isTokucho8,omitempty"`
	IsTokucho9  *bool `json:"isTokucho9,omitempty"`
	IsTokucho10 *bool `json:"isTokucho10,omitempty"`
	IsTokucho11 *bool `json:"isTokucho11,omitempty"`
	IsTokucho12 *bool `json:"isTokucho12,omitempty"`
	IsTokucho13 *bool `json:"isTokucho13,omitempty"`
	IsTokucho14 *bool `json:"isTokucho14,omitempty"`
	IsTokucho15 *bool `json:"isTokucho15,omitempty"`
	IsTokucho16 *bool `json:"isTokucho16,omitempty"`
	IsTokucho17 *bool `json:"isTokucho17,omitempty"`
	IsTokucho18 *bool `json:"isTokucho18,omitempty"`
	IsTokucho19 *bool `json:"isTokucho19,omitempty"`
	IsTokucho20 *bool `json:"isTokucho20,omitempty"`
	IsTokucho21 *bool `json:"isTokucho21,omitempty"`
	IsTokucho22 *bool `json:"isTokucho22,omitempty"`
	IsTokucho23 *bool `json:"isTokucho23,omitempty"`
	IsTokucho24 *bool `json:"isTokucho24,omitempty"`
	IsTokucho25 *bool `json:"isTokucho25,omitempty"`
	IsTokucho26 *bool `json:"isTokucho26,omitempty"`
	IsTokucho27 *bool `json:"isTokucho27,omitempty"`
	IsTokucho28 *bool `json:"isTokucho28,omitempty"`
	IsTokucho29 *bool `json:"isTokucho29,omitempty"`
	IsTokucho30 *bool `json:"isTokucho30,omitempty"`
	IsTokucho31 *bool `json:"isTokucho31,omitempty"`
	IsTokucho32 *bool `json:"isTokucho32,omitempty"`
	IsTokucho33 *bool `json:"isTokucho33,omitempty"`
	IsTokucho34 *bool `json:"isTokucho34,omitempty"`
	IsTokucho35 *bool `json:"isTokucho35,omitempty"`
	IsTokucho36 *bool `json:"isTokucho36,omitempty"`
	IsTokucho37 *bool `json:"isTokucho37,omitempty"`
	IsTokucho38 *bool `json:"isTokucho38,omitempty"`
	IsTokucho39 *bool `json:"isTokucho39,omitempty"`
	IsTokucho40 *bool `json:"isTokucho40,omitempty"`
	IsTokucho41 *bool `json:"isTokucho41,omitempty"`
	IsTokucho42 *bool `json:"isTokucho42,omitempty"`
	IsTokucho43 *bool `json:"isTokucho43,omitempty"`
	IsTokucho44 *bool `json:"isTokucho44,omitempty"`
	IsTokucho45 *bool `json:"isTokucho45,omitempty"`
	IsTokucho46 *bool `json:"isTokucho46,omitempty"`
	IsTokucho47 *bool `json:"isTokucho47,omitempty"`
	IsTokucho48 *bool `json:"isTokucho48,omitempty"`
	IsTokucho49 *bool `json:"isTokucho49,omitempty"`
	IsTokucho50 *bool `json:"isTokucho50,omitempty"`
	IsTokucho51 *bool `json:"isTokucho51,omitempty"`
	IsTokucho52 *bool `json:"isTokucho52,omitempty"`
	IsTokucho53 *bool `json:"isTokucho53,omitempty"`
	IsTokucho54 *bool `json:"isTokucho54,omitempty"`
	IsTokucho55 *bool `json:"isTokucho55,omitempty"`
	IsTokucho56 *bool `json:"isTokucho56,omitempty"`
	IsTokucho57 *bool `json:"isTokucho57,omitempty"`
	IsTokucho58 *bool `json:"isTokucho58,omitempty"`
	IsTokucho59 *bool `json:"isTokucho59,omitempty"`
	IsTokucho60 *bool `json:"isTokucho60,omitempty"`
	IsTokucho61 *bool `json:"isTokucho61,omitempty"`
	IsTokucho62 *bool `json:"isTokucho62,omitempty"`
	IsTokucho63 *bool `json:"isTokucho63,omitempty"`
	IsTokucho64 *bool `json:"isTokucho64,omitempty"`
	IsTokucho65 *bool `json:"isTokucho65,omitempty"`
	IsTokucho66 *bool `json:"isTokucho66,omitempty"`
	IsTokucho67 *bool `json:"isTokucho67,omitempty"`
	IsTokucho68 *bool `json:"isTokucho68,omitempty"`
	IsTokucho69 *bool `json:"isTokucho69,omitempty"`
	IsTokucho70 *bool `json:"isTokucho70,omitempty"`
	IsTokucho71 *bool `json:"isTokucho71,omitempty"`
	IsTokucho72 *bool `json:"isTokucho72,omitempty"`
	IsTokucho73 *bool `json:"isTokucho73,omitempty"`
	IsTokucho74 *bool `json:"isTokucho74,omitempty"`
	IsTokucho75 *bool `json:"isTokucho75,omitempty"`
	IsTokucho76 *bool `json:"isTokucho76,omitempty"`
	IsTokucho77 *bool `json:"isTokucho77,omitempty"`
	IsTokucho78 *bool `json:"isTokucho78,omitempty"`
	IsTokucho79 *bool `json:"isTokucho79,omitempty"`
	IsTokucho80 *bool `json:"isTokucho80,omitempty"`

	IsMiniKonyu           *bool `json:"isMiniKonyu,omitempty"`
	IsNaviKensakuChinKyo  *bool `json:"isNaviKensakuChinKyo,omitempty"`
	IsNaviKensakuChinJi   *bool `json:"isNaviKensakuChinJi,omitempty"`
	IsNaviKensakuUriKyo   *bool `json:"isNaviKensakuUriKyo,omitempty"`
	IsNaviKensakuUriJi    *bool `json:"isNaviKensakuUriJi,omitempty"`
	IsLastYearMemberJuchu *bool `json:"isLastYearMemberJuchu,omitempty"`
	IsThisYearMemberJuchu *bool `json:"isThisYearMemberJuchu,omitempty"`

	// isToriatsukaiBukken1..25
	IsToriatsukaiBukken1  *bool `json:"isToriatsukaiBukken1,omitempty"`
	IsToriatsukaiBukken2  *bool `json:"isToriatsukaiBukken2,omitempty"`
	IsToriatsukaiBukken3  *bool `json:"isToriatsukaiBukken3,omitempty"`
	IsToriatsukaiBukken4  *bool `json:"isToriatsukaiBukken4,omitempty"`
	IsToriatsukaiBukken5  *bool `json:"isToriatsukaiBukken5,omitempty"`
	IsToriatsukaiBukken6  *bool `json:"isToriatsukaiBukken6,omitempty"`
	IsToriatsukaiBukken7  *bool `json:"isToriatsukaiBukken7,omitempty"`
	IsToriatsukaiBukken8  *bool `json:"isToriatsukaiBukken8,omitempty"`
	IsToriatsukaiBukken9  *bool `json:"isToriatsukaiBukken9,omitempty"`
	IsToriatsukaiBukken10 *bool `json:"isToriatsukaiBukken10,omitempty"`
	IsToriatsukaiBukken11 *bool `json:"isToriatsukaiBukken11,omitempty"`
	IsToriatsukaiBukken12 *bool `json:"isToriatsukaiBukken12,omitempty"`
	IsToriatsukaiBukken13 *bool `json:"isToriatsukaiBukken13,omitempty"`
	IsToriatsukaiBukken14 *bool `json:"isToriatsukaiBukken14,omitempty"`
	IsToriatsukaiBukken15 *bool `json:"isToriatsukaiBukken15,omitempty"`
	IsToriatsukaiBukken16 *bool `json:"isToriatsukaiBukken16,omitempty"`
	IsToriatsukaiBukken17 *bool `json:"isToriatsukaiBukken17,omitempty"`
	IsToriatsukaiBukken18 *bool `json:"isToriatsukaiBukken18,omitempty"`
	IsToriatsukaiBukken19 *bool `json:"isToriatsukaiBukken19,omitempty"`
	IsToriatsukaiBukken20 *bool `json:"isToriatsukaiBukken20,omitempty"`
	IsToriatsukaiBukken21 *bool `json:"isToriatsukaiBukken21,omitempty"`
	IsToriatsukaiBukken22 *bool `json:"isToriatsukaiBukken22,omitempty"`
	IsToriatsukaiBukken23 *bool `json:"isToriatsukaiBukken23,omitempty"`
	IsToriatsukaiBukken24 *bool `json:"isToriatsukaiBukken24,omitempty"`
	IsToriatsukaiBukken25 *bool `json:"isToriatsukaiBukken25,omitempty"`

	HttpLink *string `json:"httpLink,omitempty"`

	ChinKyoBukkenHyojiKbn *string `json:"chinKyoBukkenHyojiKbn,omitempty"`
	ChinJiBukkenHyojiKbn  *string `json:"chinJiBukkenHyojiKbn,omitempty"`
	UriBukkenHyojiKbn     *string `json:"uriBukkenHyojiKbn,omitempty"`

	// model直下の updater/updateDate/updateTime
	ModelUpdater    *string `json:"updater,omitempty"`
	ModelUpdateDate *string `json:"updateDate,omitempty"`
	ModelUpdateTime *string `json:"updateTime,omitempty"`

	IsOsusumeKokaiChinKyo *bool `json:"isOsusumeKokaiChinKyo,omitempty"`
	IsOsusumeKokaiChinJi  *bool `json:"isOsusumeKokaiChinJi,omitempty"`
	IsOsusumeKokaiUriKyo  *bool `json:"isOsusumeKokaiUriKyo,omitempty"`
	IsOsusumeKokaiUriJi   *bool `json:"isOsusumeKokaiUriJi,omitempty"`
	IsOsusumeKokaiResort  *bool `json:"isOsusumeKokaiResort,omitempty"`
	IsStaffPrelusion      *bool `json:"isStaffPrelusion,omitempty"`

	UriJiBukkenHyojiKbn  *string `json:"uriJiBukkenHyojiKbn,omitempty"`
	ResortBukkenHyojiKbn *string `json:"resortBukkenHyojiKbn,omitempty"`
	ColorKbn             *string `json:"colorKbn,omitempty"`
	ShogoDispPosition    *string `json:"shogoDispPosition,omitempty"`

	IsConpaneRiyo        *bool   `json:"isConpaneRiyo,omitempty"`
	ConpaneRiyoStartDate *string `json:"conpaneRiyoStartDate,omitempty"`
	ConpaneRiyoEndDate   *string `json:"conpaneRiyoEndDate,omitempty"`

	RiyoStartShinseiDate *string `json:"riyoStartShinseiDate,omitempty"`
	RiyoStartDate        *string `json:"riyoStartDate,omitempty"`
	RiyoStopShinseiDate  *string `json:"riyoStopShinseiDate,omitempty"`
	RiyoStopDate         *string `json:"riyoStopDate,omitempty"`

	LastUpdateUserId *string `json:"lastUpdateUserId,omitempty"`
	UrlShubetsu      *string `json:"urlShubetsu,omitempty"`
	UrlShort         *string `json:"urlShort,omitempty"`
	UrlLong          *string `json:"urlLong,omitempty"`

	VersionNo  *int    `json:"versionNo,omitempty"`
	CreateDate *string `json:"createDate,omitempty"`
	CreateTime *string `json:"createTime,omitempty"`
	Creator    *string `json:"creator,omitempty"`

	IsShowLogo *bool `json:"isShowLogo,omitempty"`
}

type KaiintenShokaiKeiyakuKaiinItem struct {
	SeqNo   int    `json:"seqNo,omitempty"`
	KaiinNo string `json:"kaiinNo"`

	IsUsed                *bool   `json:"isUsed,omitempty"`
	KameitenSearchHikokai *string `json:"kameitenSearchHikokai,omitempty"`
	Memo                  *string `json:"memo,omitempty"`
	VersionNo             *int    `json:"versionNo,omitempty"`

	CreateDate *string `json:"createDate,omitempty"`
	CreateTime *string `json:"createTime,omitempty"`
	Creator    *string `json:"creator,omitempty"`

	UpdateDate *string `json:"updateDate,omitempty"`
	UpdateTime *string `json:"updateTime,omitempty"`
	Updater    *string `json:"updater,omitempty"`
}

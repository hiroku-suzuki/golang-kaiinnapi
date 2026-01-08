package models

// API共通レスポンス

type ErrorItem struct {
	Message string   `json:"message"`
	Fields  []string `json:"fields,omitempty"`
}

type WarningItem struct {
	Message string   `json:"message"`
	Fields  []string `json:"fields,omitempty"`
}

type AtbbKeiyakuResponse struct {
	Model    *AtbbKeiyakuModel `json:"model"`
	Errors   []ErrorItem       `json:"errors,omitempty"`
	Warnings []WarningItem     `json:"warnings,omitempty"`
}

// メインモデル（提示された型に対応）
// NOTE: モックで“例”を返すため、値が入っていないフィールドは省略（omitempty）

type AtbbKeiyakuModel struct {
	KaiinNo                  string  `json:"kaiinNo"`
	HokatsuKaiinNo           *string `json:"hokatsuKaiinNo,omitempty"`
	IsKaiyaku                *bool   `json:"isKaiyaku,omitempty"`
	IsRiyochu                *bool   `json:"isRiyochu,omitempty"`
	DispKaiinKbn             *string `json:"dispKaiinKbn,omitempty"`
	DispKaiinKbnName         *string `json:"dispKaiinKbnName,omitempty"`
	AtbbKeiyakuKbn           *string `json:"atbbKeiyakuKbn,omitempty"`
	AtbbKeiyakuKbnName       *string `json:"atbbKeiyakuKbnName,omitempty"`
	IsMuryoKeiyaku           *bool   `json:"isMuryoKeiyaku,omitempty"`
	IsAthomeEmployee         *bool   `json:"isAthomeEmployee,omitempty"`
	CanMailSend              *bool   `json:"canMailSend,omitempty"`
	YuryoKeiyakuShinseiDate  *string `json:"yuryoKeiyakuShinseiDate,omitempty"`
	YuryoKaiyakuShinseiDate  *string `json:"yuryoKaiyakuShinseiDate,omitempty"`
	MuryoRiyoStartDate       *string `json:"muryoRiyoStartDate,omitempty"`
	MuryoRiyoStopDate        *string `json:"muryoRiyoStopDate,omitempty"`
	YuryoRiyoStartDate       *string `json:"yuryoRiyoStartDate,omitempty"`
	YuryoRiyoStopDate        *string `json:"yuryoRiyoStopDate,omitempty"`
	MuryoLoginIdCnt          *int    `json:"muryoLoginIdCnt,omitempty"`
	Biko1                    *string `json:"biko1,omitempty"`
	Biko2                    *string `json:"biko2,omitempty"`
	Biko3                    *string `json:"biko3,omitempty"`
	Biko4                    *string `json:"biko4,omitempty"`
	KaiyakuRiyu              *string `json:"kaiyakuRiyu,omitempty"`
	ShubetsuBiko             *string `json:"shubetsuBiko,omitempty"`
	Biko5                    *string `json:"biko5,omitempty"`
	Biko6                    *string `json:"biko6,omitempty"`
	Biko7                    *string `json:"biko7,omitempty"`
	Biko8                    *string `json:"biko8,omitempty"`
	PcRentalNo               *string `json:"pcRentalNo,omitempty"`
	SenpoTantoName           *string `json:"senpoTantoName,omitempty"`
	RiyoKanoMenuRiyoArea     *string `json:"riyoKanoMenuRiyoArea,omitempty"`
	RiyoKanoMenuRiyoAreaName *string `json:"riyoKanoMenuRiyoAreaName,omitempty"`
	Tantos                   []Tanto `json:"tantos,omitempty"`
	IsDemo                   *bool   `json:"isDemo,omitempty"`
	CanMultiLogin            *bool   `json:"canMultiLogin,omitempty"`
	PwdLimitKbn              *string `json:"pwdLimitKbn,omitempty"`
	PwdLimitKbnName          *string `json:"pwdLimitKbnName,omitempty"`
	PwdLimit                 *int    `json:"pwdLimit,omitempty"`
	CreateProgram            *string `json:"createProgram,omitempty"`
	Creator                  *string `json:"creator,omitempty"`
	CreateTimestamp          *string `json:"createTimestamp,omitempty"`
	UpdateProgram            *string `json:"updateProgram,omitempty"`
	UpdateTimestamp          *string `json:"updateTimestamp,omitempty"`
	Updater                  *string `json:"updater,omitempty"`
	VersionNo                *int    `json:"versionNo,omitempty"`
}

type Tanto struct {
	TantoShubetsu        *string `json:"tantoShubetsu,omitempty"`
	TantoBunrui          *string `json:"tantoBunrui,omitempty"`
	AtbbKeiyakuCreateKbn *string `json:"atbbKeiyakuCreateKbn,omitempty"`
	TantoCd              *string `json:"tantoCd,omitempty"`
	TantoName            *string `json:"tantoName,omitempty"`
	IsTaishoku           *bool   `json:"isTaishoku,omitempty"`
	CreateProgram        *string `json:"createProgram,omitempty"`
	Creator              *string `json:"creator,omitempty"`
	CreateTimestamp      *string `json:"createTimestamp,omitempty"`
	UpdateProgram        *string `json:"updateProgram,omitempty"`
	UpdateTimestamp      *string `json:"updateTimestamp,omitempty"`
	Updater              *string `json:"updater,omitempty"`
}

// ポインタ生成ヘルパ
func Bool(v bool) *bool    { return &v }
func Int(v int) *int       { return &v }
func Str(v string) *string { return &v }

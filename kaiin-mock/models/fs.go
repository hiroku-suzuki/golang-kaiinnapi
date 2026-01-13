package models

type FSResponse struct {
	Model    *FSModel      `json:"model"`
	Errors   []ErrorItem   `json:"errors,omitempty"`
	Warnings []WarningItem `json:"warnings,omitempty"`
}

type FSModel struct {
	KaiinNo string `json:"kaiinNo"`

	InternetTokubetsuKaiinKbn     string `json:"internetTokubetsuKaiinKbn,omitempty"`
	InternetTokubetsuKaiinKbnName string `json:"internetTokubetsuKaiinKbnName,omitempty"`
	FsKeiyakuKbn                  string `json:"fsKeiyakuKbn,omitempty"`
	FsKeiyakuKbnName              string `json:"fsKeiyakuKbnName,omitempty"`
	FsKeiyakuDate                 string `json:"fsKeiyakuDate,omitempty"`
	FsKaiyakuDate                 string `json:"fsKaiyakuDate,omitempty"`
	InsatsuTankaKbn               string `json:"insatsuTankaKbn,omitempty"`
	InsatsuTankaKbnName           string `json:"insatsuTankaKbnName,omitempty"`
	ZumenKbn                      string `json:"zumenKbn,omitempty"`
	ZumenKbnName                  string `json:"zumenKbnName,omitempty"`
	FarboRiyoKbn                  string `json:"farboRiyoKbn,omitempty"`
	FarboRiyoKbnName              string `json:"farboRiyoKbnName,omitempty"`
	Course                        string `json:"course,omitempty"`
	Cycle                         string `json:"cycle,omitempty"`
	CreateDate                    string `json:"createDate,omitempty"`
	CreateTime                    string `json:"createTime,omitempty"`
	Creator                       string `json:"creator,omitempty"`
	UpdateDate                    string `json:"updateDate,omitempty"`
	UpdateTime                    string `json:"updateTime,omitempty"`
	Updater                       string `json:"updater,omitempty"`
	IsKodokuStop                  bool   `json:"isKodokuStop,omitempty"`
	HaifuEndDate                  string `json:"haifuEndDate,omitempty"`
	VersionNo                     int    `json:"versionNo,omitempty"`
}

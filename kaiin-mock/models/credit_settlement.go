package models

type CreditSettlementResponse struct {
	Model    *CreditSettlementModel `json:"model"`
	Errors   []ErrorItem            `json:"errors,omitempty"`
	Warnings []WarningItem          `json:"warnings,omitempty"`
}

type CreditSettlementModel struct {
	KaiinNo string `json:"kaiinNo"`

	IsUsingJcb             *bool   `json:"isUsingJcb,omitempty"`
	UsingJcbName           *string `json:"usingJcbName,omitempty"`
	IsUsingEarlyPayment    *bool   `json:"isUsingEarlyPayment,omitempty"`
	UsingEarlyPaymentName  *string `json:"usingEarlyPaymentName,omitempty"`
	IsUsingBonusPayment    *bool   `json:"isUsingBonusPayment,omitempty"`
	UsingBonusPaymentName  *string `json:"usingBonusPaymentName,omitempty"`

	Memo1 *string `json:"memo1,omitempty"`
	Memo2 *string `json:"memo2,omitempty"`
	Memo3 *string `json:"memo3,omitempty"`

	ContractStatusCd   *string `json:"contractStatusCd,omitempty"`
	ContractStatusName *string `json:"contractStatusName,omitempty"`

	CreateDate *string `json:"createDate,omitempty"`
	CreateTime *string `json:"createTime,omitempty"`
	Creator    *string `json:"creator,omitempty"`
	UpdateDate *string `json:"updateDate,omitempty"`
	UpdateTime *string `json:"updateTime,omitempty"`
	Updater    *string `json:"updater,omitempty"`

	CreditSettlementService []CreditSettlementServiceItem `json:"creditSettlementService,omitempty"`
}

type CreditSettlementServiceItem struct {
	KaiinNo string `json:"kaiinNo"`

	ServiceType *string `json:"serviceType,omitempty"`
	OrderDate   *string `json:"orderDate,omitempty"`
	ContractDate *string `json:"contractDate,omitempty"`
	ContractKaCd *string `json:"contractKaCd,omitempty"`

	ContractTantoshaCd   *string `json:"contractTantoshaCd,omitempty"`
	ContractTantoshaName *string `json:"contractTantoshaName,omitempty"`

	ContractStatusCd   *string `json:"contractStatusCd,omitempty"`
	ContractStatusName *string `json:"contractStatusName,omitempty"`

	DeliveryScheduledDate *string `json:"deliveryScheduledDate,omitempty"`
	NumberUnitUsing       *int    `json:"numberUnitUsing,omitempty"`
	UseStartDate          *string `json:"useStartDate,omitempty"`

	CancelOrderDate *string `json:"cancelOrderDate,omitempty"`
	CancelDate      *string `json:"cancelDate,omitempty"`
	CancelKaCd      *string `json:"cancelKaCd,omitempty"`
	CancelTantoCd   *string `json:"cancelTantoCd,omitempty"`
	CancelTantoName *string `json:"cancelTantoName,omitempty"`

	CreateDate *string `json:"createDate,omitempty"`
	CreateTime *string `json:"createTime,omitempty"`
	Creator    *string `json:"creator,omitempty"`
	UpdateDate *string `json:"updateDate,omitempty"`
	UpdateTime *string `json:"updateTime,omitempty"`
	Updater    *string `json:"updater,omitempty"`
}
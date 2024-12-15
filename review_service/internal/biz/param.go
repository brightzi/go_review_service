package biz

type ReplyParam struct {
	ReviewID  int64
	StoreID   int64
	Content   string
	PicInfo   string
	VideoInfo string
}

type ReviewAppealParam struct {
	ReviewID  int64
	StoreID   int64
	Resaon    string
	Content   string
	PicInfo   string
	VideoInfo string
}

type AuditParam struct {
	ReviewID  int64
	Status    int32
	OpUser    string
	OpReason  string
	OpRemarks string
}

type AuditAppealParam struct {
	ReviewID int64
	AppealID int64
	Status   int32
	OpUser   string
	OpReason string
}

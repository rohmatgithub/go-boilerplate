package common

type ContextModel struct {
	UserID    string
	CompanyID string
	BranchID  string
	KassaID   string
}

type LoggerModel struct {
	Pid         string
	RequestID   string
	Application string
	Version     string
	ByteIn      int
	Status      int
	Path        string
}

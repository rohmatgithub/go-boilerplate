package common

type ContextModel struct {
	UserID      string
	CompanyID   string
	BranchID    string
	KassaID     string
	LoggerModel LoggerModel
}

type LoggerModel struct {
	Pid         string
	RequestID   string
	Application string
	Version     string
	ByteIn      int
	Status      int
	Path        string
	Message     string
	Code        string
}

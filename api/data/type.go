package data

type BackendRequest struct {
	ID string `json:"id"`
}

type FunctionGetResponse struct {
	Data        string `json:"data"`
	DownloadURL string `json:"url"`
}

type FormatRequest struct {
	ID   string `json:"id"`
	Data string `json:"data"`
}

type FormatResponse struct {
	Data string `json:"data"`
}

type CompileRequest struct {
	ID   string `json:"id"`
	Data string `json:"data"`
	OS   string `json:"os"`
}

type ErrResponse struct {
	Err string `json:"err"`
}

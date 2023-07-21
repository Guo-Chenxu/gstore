package model

type R struct {
	StatusCode int    `json:"statusCode"`
	Msg        string `json:"msg"`
	Data       Data   `json:"data"`
	Success    bool   `json:"success"`
}

type Head struct {
	Link []interface{} `json:"link"`
	Vars []string      `json:"vars"`
}

type Results struct {
	Bindings []map[string]map[string]string `json:"bindings"`
}

type Data struct {
	Head        Head    `json:"head"`
	Results     Results `json:"results"`
	ThreadID    string  `json:"threadId"`
	StatusCode  int     `json:"statusCode"`
	AnsNum      int     `json:"ansNum"`
	StatusMsg   string  `json:"statusMsg"`
	OutputLimit int     `json:"outputLimit"`
	QueryTime   string  `json:"queryTime"`
}

type GstoreRequest struct {
	Database string `json:"database"`
	Sparql   string `json:"sparql"`
}

type TestJson struct {
	Label   string `json:"label"`
	Predict string `json:"predict"`
}

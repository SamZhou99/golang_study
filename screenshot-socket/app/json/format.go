package json

// JSON格式
type Screenshot struct {
	Top    int `json:top`
	Left   int `json:left`
	Width  int `json:width`
	Height int `json:height`
}
type Socket struct {
	Ip   string `json:ip`
	Port int    `json:port`
}

type Config struct {
	Screenshot Screenshot `json:screenshot`
	Socket     Socket     `json:socket`
}

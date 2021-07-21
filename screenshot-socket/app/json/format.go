package json

// JSON格式
type Screenshot struct {
	Top    int32 `json:top`
	Left   int32 `json:left`
	Width  int32 `json:width`
	Height int32 `json:height`
}
type Socket struct {
	Ip   string `json:ip`
	Port int32  `json:port`
}

type Config struct {
	Screenshot Screenshot `json:screenshot`
	Socket     Socket     `json:socket`
}

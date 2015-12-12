package jsonstruct

type Photo struct {
	ID    string `json:"id"`
	Owner string `json:"owner"`
	Title string `json:"title"`
}

type Photos struct {
	Page    int64   `json:"page"`
	Pages   int64   `json:"pages"`
	Perpage int64   `json:"perpage"`
	Total   int64   `json:"total"`
	Photo   []Photo `json:"photo"`
}

type PhotosSearch struct {
	Photos Photos `json:"photos"`
	Stat   string `json:"stat"`
}

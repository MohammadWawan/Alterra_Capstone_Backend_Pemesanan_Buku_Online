package books

type Getbook struct {
	VolumeInfo struct {
		Title      string   `json:"title"`
		Categories []string `json:"categories"`
		Publisher  string   `json:"publisher"`
		Price      struct {
			Harga int `json:"amount"`
		}
		Description string   `json:"description"`
		Author      []string `json:"authors"`
	} `json:"volumeInfo"`
}

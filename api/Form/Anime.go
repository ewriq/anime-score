package Form


type Anime struct {
	Name string `json:"name"`
	Token string `json:"token"`
	Description    string `json:"description"`
	Thumbnail    string `json:"thumbnail"`
	Five int `json:"five"`
	Four int `json:"four"`
	Three int `json:"three"`
	Two int `json:"two"`
	One int `json:"one"`
	Average float64 `json:"rating"`
}

type AddAnime struct {
	Name string `json:"name"`
	Description string `json:"description"`
	Thumbnail string `json:"thumbnail"`
}

type DelAnime struct {
	Name string `json:"name"`
}

type Search struct {
	Query string `json:"query"`
}
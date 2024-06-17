package Form


type Vote struct {
	Token string `json:"token"`
	Voter string `json:"voter"`
	Score    int `json:"score"`
	Anime    string `json:"anime"`
}

type ListVote struct {
	Token string `json:"token"`
}

type AddVote struct {
	Voter string `json:"voter"`
	Score    int `json:"score"`
	Anime    string `json:"anime"`
}

type DelVote struct {
	Voter string `json:"voter"`
	Anime    string `json:"anime"`
	Score    int `json:"score"`
}
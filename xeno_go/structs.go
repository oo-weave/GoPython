package xeno_go


type Response struct {
	NumRecordings string      `json:"numRecordings"`
	NumSpecies    string      `json:"numSpecies"`
	Page          int         `json:"page"`
	NumPages      int         `json:"numPages"`
	Recordings    []Recording `json:"recordings"`
}

type Recording struct {
	Id           string   `json:"id"`
	Gen          string   `json:"gen"`
	Sp           string   `json:"sp"`
	Ssp          string   `json:"ssp"`
	En           string   `json:"en"`
	Rec          string   `json:"rec"`
	Cnt          string   `json:"cnt"`
	Loc          string   `json:"loc"`
	Lat          string   `json:"lat"`
	Lng          string   `json:"lng"`
	Alt          string   `json:"alt"`
	Type         string   `json:"type"`
	Url          string   `json:"url"`
	File         string   `json:"file"`
	FileName     string   `json:"file-name"`
	Sono         Sonogram `json:"sono"`
	Lic          string   `json:"lic"`
	Q            string   `json:"q"`
	Length       string   `json:"length"`
	Time         string   `json:"time"`
	Date         string   `json:"date"`
	Uploaded     string   `json:"uploaded"`
	Also         []string `json:"also"`
	Rmk          string   `json:"rmk"`
	BirdSeen     string   `json:"bird-seen"`
	PlaybackUsed string   `json:"playback-used"`
}

type Sonogram struct {
	Small string `json:"small"`
	Med   string `json:"med"`
	Large string `json:"large"`
	Full  string `json:"full"`
}

package types

type Media struct {
	Bin       string `json:"bin"`
	Text      string `json:"text"`
	Icon      string `json:"icon"`
	ImagePath string `json:"image_path,omitempty"`
	AudioPath string `json:"audio_path,omitempty"`
	VideoPath string `json:"video_path,omitempty"`
	TimeStamp string `json:"time_stamp,omitempty"`
}

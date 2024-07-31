package domain

// 変数の頭文字を大文字にしないとScanできなくなる
type Report struct {
	Id               string `json:"id"`
	User_id          string `json:"user_id"`
	Title            string `json:"title"`
	Character_counts int    `json:"character_counts"`
	Style            int    `json:"style"`
	Language         int    `json:"language"`
}

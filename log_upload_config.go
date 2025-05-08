package webaccel

// LogUploadConfig ログアップロード設定
type LogUploadConfig struct {
	Bucket          string `json:","`
	Endpoint        string `json:"," validate:"omitempty,equal=https://s3.isk01.sakurastorage.jp"`
	Region          string `json:"," validate:"omitempty,equal=jp-north-1"`
	AccessKeyID     string `json:","`
	SecretAccessKey string `json:","`
	Status          string `json:"," validate:"oneof=enabled disabled"`
}

package webaccel

// LogUploadConfig ログアップロード設定
type LogUploadConfig struct {
	Bucket          string `json:",omitempty"`
	Endpoint        string `json:",omitempty" validate:"omitempty,equal=https://s3.isk01.sakurastorage.jp"`
	Region          string `json:",omitempty" validate:"omitempty,equal=jp-north-1"`
	AccessKeyID     string `json:",omitempty"`
	SecretAccessKey string `json:",omitempty"`
	Status          string `json:",omitempty" validate:"oneof=enabled disabled"`
}

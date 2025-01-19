package config

type (
	FaceConfig struct {
		RunAddress      string `json:"run_address"`       // Адрес запуска web-сервера
		StaticFilesPath string `json:"static_files_path"` // Расположение статических файлов
		StaticURLPrefix string `json:"static_url_prefix"` // url-префикс статических файлов
	}
)

package content

type ExtractContentDTO struct {
	UserID uint64
	email  string
	Url    string
}

type ExtractContentPreviewDTO struct {
	email string
	Url   string
}

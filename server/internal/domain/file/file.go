package file

const (
	TypeImageMain = "image_main"
	TypeImage     = "image"
)

type File struct {
	ID   int64  `json:"id"`
	Path string `json:"path"`
	Type string `json:"type"`
}

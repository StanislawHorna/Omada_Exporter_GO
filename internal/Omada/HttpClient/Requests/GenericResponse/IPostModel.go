package GenericResponse

type PayloadModel interface {
	Path() string
	Call() any
}

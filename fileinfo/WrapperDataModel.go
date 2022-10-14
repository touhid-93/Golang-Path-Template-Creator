package fileinfo

type WrapperDataModel struct {
	RawPath     string
	IsDirectory bool
	IsFile      bool
	IsEmptyPath bool
	Separator   string
}

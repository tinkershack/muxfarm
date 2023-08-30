package plumber

func NewMediaIn() *MediaIn {
	min := new(MediaIn)
	return min
}

func (min *MediaIn) Add(st StorageType, objURI isURI_Uri) {
	min.Ins = append(min.Ins, &Media{Type: *st.Enum(), Path: &URI{Uri: objURI}})
}

func NewMuxfarmID(id string) *MuxfarmID {
	mid := new(MuxfarmID)
	mid.Mid = id
	return mid
}

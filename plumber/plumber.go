package plumber

func NewMediaIn() *MediaIn {
	min := new(MediaIn)
	min.StrictValidate = false
	return min
}

func (min *MediaIn) Add(st StorageType, objURI isURI_Uri) {
	min.Ins = append(min.Ins, &Media{Type: *st.Enum(), Path: &URI{Uri: objURI}})
}

func (min *MediaIn) Revelio() {}

func (min *MediaIn) Acid() {}

func (min *MediaIn) Doobie() {}

func (min *MediaIn) Pulse() {}

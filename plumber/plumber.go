package plumber

import (
	"github.com/google/uuid"
)

func NewMediaIn() *MediaIn {
	min := new(MediaIn)
	return min
}

func (min *MediaIn) Add(st StorageType, objURI isURI_Uri) {
	min.Input = append(min.Input, &Media{Storagetype: *st.Enum(), Uri: &URI{Uri: objURI}})
}

func (mid *MuxfarmID) ID() {
	mid.Mid = uuid.NewString()
}

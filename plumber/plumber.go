package plumber

import (
	"fmt"

	"github.com/google/uuid"
)

func NewMediaIn() *MediaIn {
	min := new(MediaIn)
	return min
}

func (min *MediaIn) Add(st StorageType, uri string) {
	min.Input = append(min.Input, &Media{Storagetype: *st.Enum(), Uri: uri})
}

func (mid *MuxfarmID) ID() {
	mid.Mid = uuid.NewString()
}

func (m *Media) FormGetterURI() string {
	sourceURI := m.GetUri()
	switch st := m.GetStoragetype(); st {
	case StorageType_STORAGE_UNSPECIFIED:
	case StorageType_STORAGE_HTTP:
	case StorageType_STORAGE_S3:
		sourceURI = fmt.Sprintf("s3::%s", sourceURI)
	case StorageType_STORAGE_GCS:
		sourceURI = fmt.Sprintf("gcs::%s", sourceURI)
	}
	return sourceURI
}

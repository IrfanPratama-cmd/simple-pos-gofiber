package model

import "github.com/google/uuid"

type ProductAsset struct {
	Base
	ProductAssetAPI
}

type ProductAssetAPI struct {
	ProductID *uuid.UUID `json:"product_id,omitempty" gorm:"type:varchar(36)" format:"uuid"`
	FileName  string     `json:"file_name,omitempty" `
	FilePath  string     `json:"file_path,omitempty" `
	IsPrimary bool       `json:"is_primary,omitempty" `
}

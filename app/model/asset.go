package model

// Asset struct
type Asset struct {
	Base
	DataOwner
	AssetAPI
}

// AssetAPI Asset API
type AssetAPI struct {
	Title            *string  `json:"title,omitempty" example:"Image" gorm:"type:varchar(256)"`                 // Title
	Filename         *string  `json:"filename,omitempty" example:"image.png" gorm:"type:varchar(256)"`          // Filename
	FileSize         *float64 `json:"file_size,omitempty" format:"float" swaggertype:"number"`                  // File Size
	OriginalFilename *string  `json:"original_filename,omitempty" example:"image.png" gorm:"type:varchar(256)"` // OriginalFilename
	FilePath         *string  `json:"file_path,omitempty" gorm:"type:varchar(256)"`                             // File Path
	AbsolutePath     *string  `json:"absolute_path,omitempty" example:"image.png" gorm:"type:varchar(256)"`     // AbsolutePath
	RelativePath     *string  `json:"relative_path,omitempty" example:"image.png" gorm:"type:varchar(256)"`     // RelativePath
	Description      *string  `json:"description,omitempty" example:"string" gorm:"type:text"`                  // Description
}

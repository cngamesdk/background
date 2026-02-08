package response

import "github.com/flipped-aurora/gin-vue-admin/server/model/example"

type ExaFileResponse struct {
	File     example.ExaFileUploadAndDownload `json:"file"`
	Size     int64                            `json:"size"`
	Width    int                              `json:"width"`
	Height   int                              `json:"height"`
	Duration int                              `json:"duration"`
	Bitrate  int                              `json:"bitrate"`
	Fps      int                              `json:"fps"`
}

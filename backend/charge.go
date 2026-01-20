package backend

import (
	"ahut-tool/backend/models"
)

// GetIMS 获取IMS信息
func (a *App) GetIMS(xiaoqu, ld_Name, ld_Id, Room_No, etype string) (*models.IMSResponse, error) {
	formData := PayInstance.GetIMSFormData(xiaoqu, ld_Name, ld_Id, Room_No, etype)
	imsResponse, err := PayInstance.GetIMS(formData)
	if err != nil {
		return nil, err
	}

	return imsResponse, nil
}

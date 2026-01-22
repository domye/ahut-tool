package pay

const (
	// BaseURL 支付系统基础URL
	BaseURL = "https://pay.ahut.edu.cn"

	// LoginURL 登录URL
	LoginURL = BaseURL + "/Account/Login"

	// LoginServiceURL 登录服务URL
	LoginServiceURL = BaseURL + "/Account/LoginService"

	// IMSURL IMS页面URL
	IMSURL = BaseURL + "/Charge/IMS?state=WXSTATEFLAG"

	// IMSServiceURL IMS服务URL
	IMSServiceURL = BaseURL + "/Charge/GetIMS_AHUTService"
)

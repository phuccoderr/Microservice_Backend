package constants

const (
	API_ENPOINT       = "/api/v1"
	EMAIL_FROM        = "phuctapcode@gmail.com"
	BODY              = "Chẳng có gì cả"
	CHANGE_PASSWORD   = "Thay đổi mật khẩu"
	VERIFY_ACCOUNT    = "Xác thực tài khoản"
	VERIFY_PLACEORDER = "Xác thực đơn hàng"
	FEEDBACK_CUSTOMER = "Khách hàng báo cáo"

	TITLE_VERIFY_ACCOUNT = `Vui lòng nhấn đường link dưới đây để xác thực tài khoản của bạn!`
	TITLE_RESET_PASSWORD = `Vui lòng nhấn đường link dưới đây để thay đổi mật khẩu tài khoản của bạn!`
	TITLE_PLACE_ORDER    = `Cảm ơn bạn đã mua hàng từ cửa hàng chúng tôi!`

	KAFKA_TOPIC_VERIFY_ACCOUNT  = "customer-verify-events-topics"
	KAFKA_TOPIC_VERIFY_PASSWORD = "customer-password-events-topics"
	KAFKA_TOPIC_PLACE_ORDER     = "place-order-events-topics"
)

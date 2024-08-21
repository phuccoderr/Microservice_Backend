package constants

const (
	API_ENPOINT = "/api/v1"

	JWT_NOT_FOUND     = "Authorization header is empty!"
	JWT_TOKEN_INVALID = "Token is invalid!"
	JWT_PARSE_CLAIMS  = "Cannot parse claims!"
	JWT_SIGNING_KEY   = "Unexpected signing method!"

	GET_SUCCESS         = "Get cart successfully!"
	CREATE_SUCCESS      = "Create cart successfully!"
	DELETE_SUCCESS      = "Delete cart successfully!"
	CHECKOUT_SUCCESS    = "Check out successfully!"
	PLACE_ORDER_SUCCESS = "Place order successfully!"

	MICROSERVICE_FAIL      = "Failed to call request!"
	MICROSERVICE_READ_RESP = "Failed to read microservice response!"

	STATUS_BAD_REQUEST                 = "Bad request!"
	STATUS_UNAUTHORIZED                = "Unauthorized!"
	STATUS_NOT_FOUND                   = "Not found!"
	STATUS_INTERNAL_ERROR              = "Internal server error!"
	STATUS_STATUS_UNPROCESSABLEENTITTY = "Status unprocessable entity!"

	DB_NOT_FOUND = "Data not found!"

	KAFKA_TOPIC_PLACE_ORDER = "place-order-events-topics"
)

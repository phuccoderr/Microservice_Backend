package constants

const (
	API_CART = "/api/v1/cart"

	JWT_NOT_FOUND     = "Authorization header is empty!"
	JWT_TOKEN_INVALID = "Token is invalid!"
	JWT_PARSE_CLAIMS  = "Cannot parse claims!"
	JWT_SIGNING_KEY   = "Unexpected signing method!"

	GET_SUCCESS      = "Get cart successfully!"
	CREATE_SUCCESS   = "Create cart successfully!"
	DELETE_SUCCESS   = "Delete cart successfully!"
	CHECKOUT_SUCCESS = "Check out successfully!"

	MICROSERVICE_FAIL      = "Failed to call request!"
	MICROSERVICE_READ_RESP = "Failed to read microservice response!"

	STATUS_BADREQUEST          = "Bad request!"
	STATUS_UNAUTHORIZED        = "Unauthorized!"
	STATUS_UNPROCESSABLEENTITY = "Unprocessable entity!"
	STATUS_NOT_FOUND           = "Not found!"
	STATUS_INTERNAL_ERROR      = "Internal server error!"

	DB_NOT_FOUND = "Data not found!"
)

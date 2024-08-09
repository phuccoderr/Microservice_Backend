package constants

const (
	API_CART = "/api/v1/cart"

	JWT_NOT_FOUND     = "Authorization header is empty!"
	JWT_TOKEN_INVALID = "Token is invalid!"
	JWT_PARSE_CLAIMS  = "Cannot parse claims!"
	JWT_SIGNING_KEY   = "Unexpected signing method!"

	MICROSERVICE_FAIL      = "Failed to call request!"
	MICROSERVICE_READ_RESP = "Failed to read microservice response!"

	STATUS_BADREQUEST          = "Bad request!"
	STATUS_UNAUTHORIZED        = "Unauthorized!"
	STATUS_UNPROCESSABLEENTITY = "Unprocessable entity!"
)

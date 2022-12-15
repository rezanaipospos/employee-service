package Constant

import "net/http"

type InternalError int

const (
	UnavailablePicPhoneAndCompanyPhone                        = "pic phone and company phone cannot be the same"
	EmptyForm                                                 = "Form Cannot Be Empty"
	UnsupportedMediaFile                                      = "Form File Unsupported File Type"
	UnsupportedConstantType                                   = "Form File Unsupported Constant Type"
	UnsupportedSizeFile                                       = "Form File Unsupported Size File"
	NoImageFile                                               = "No Image File"
	InvalidToken                                              = "Invalid Token"
	TokenExpired                                              = "Token Expired"
	EmailNotExistOrAlreadyVerified                            = "Email Not Exists Or Email Already Verified"
	DataNotExists                                             = "Data Not Exists"
	KtpFileNotExists                                          = "Ktp File Not Exists"
	NpwpFileNotExists                                         = "Npwp File Not Exists"
	UserNotExists                                             = "User Not Exists"
	ErrorRemovePassword                                       = "Error Remove Password"
	UnavailableToken                                          = "Token No Longer Exists"
	UnavailablePIN                                            = "PIN Must be 16 Digits"
	InvalidParameterLimit                                     = "Invalid Parameter Limit"
	InvalidParameterOffset                                    = "Invalid Parameter Offset"
	DuplicateData                                             = "Duplicate Data"
	ErrorNoRows                                               = "no rows in result set"
	StatusBadRequestJson                        InternalError = 4001
	StatusBadRequestNotExist                    InternalError = 4002
	StatusBadRequestAlreadyExists               InternalError = 4003
	StatusBadRequestInvalidData                 InternalError = 4004
	StatusBadRequestInvalidType                 InternalError = 4005
	StatusBadRequestInvalidPhoneNumber          InternalError = 4006
	StatusBadRequestInvalidPin                  InternalError = 4007
	StatusBadRequestInvalidLicenseNumber        InternalError = 4008
	StatusBadRequestInvalidDriverId             InternalError = 4009
	StatusBadRequestInvalidPoliceNumber         InternalError = 40010
	StatusBadRequestInvalidParameter            InternalError = 40011
	StatusBadRequestRefNoAlreadyExists          InternalError = 40012
	StatusBadRequestMobilePhoneAlreadyExists    InternalError = 40013
	StatusBadRequestEmailAlreadyExists          InternalError = 40014
	StatusBadRequestCompanyAlreadyExists        InternalError = 40015
	StatusBadRequestContractStartNull           InternalError = 40016
	StatusBadRequestContractEndNull             InternalError = 40017
	StatusBadRequestEmployeeAlreadyResign       InternalError = 40018
	StatusUnauthorizedInvalidToken              InternalError = 4011
	StatusUnauthorizedTokenNotExists            InternalError = 4012
	StatusUnauthorizedApiKeyNotExists           InternalError = 4013
	StatusUnauthorizedInvalidApiKey             InternalError = 4014
	StatusUnauthorizedInvalidUsernamePassword   InternalError = 4015
	StatusUnauthorizedBearerNotFound            InternalError = 4016
	StatusUnauthorizedErrorVerifying            InternalError = 4017
	StatusPreconditionFailedTokenExpired        InternalError = 4121
	StatusPreconditionFailedTokenNoLongerExists InternalError = 4122
	StatusPreconditionFailedInvalidParameter    InternalError = 4123
	StatusInternalServerErrorDB                 InternalError = 5001
	StatusInternalServerErrorRabbit             InternalError = 5002
	StatusInternalServerErrorRedis              InternalError = 5003
	InternalServerErrorMessageBroker            InternalError = 5004
	StatusInternalServerError                   InternalError = 5005
	StatusInternalServerErrorPubSub             InternalError = 5006
	StatusInternalServerErrorDapr               InternalError = 5007
)

type ErrorInfo struct {
	HttpCode    int
	Description string
	Title       string
}

var errorInfo = map[InternalError]ErrorInfo{
	StatusBadRequestJson: {
		HttpCode:    http.StatusBadRequest,
		Description: "Json Error",
		Title:       http.StatusText(http.StatusBadRequest),
	},
	StatusBadRequestNotExist: {
		HttpCode:    http.StatusBadRequest,
		Description: "Data Not Exists",
		Title:       http.StatusText(http.StatusBadRequest),
	},
	StatusBadRequestAlreadyExists: {
		HttpCode:    http.StatusBadRequest,
		Description: "Data Already Exists",
		Title:       http.StatusText(http.StatusBadRequest),
	},
	StatusBadRequestInvalidData: {
		HttpCode:    http.StatusBadRequest,
		Description: "Invalid Data",
		Title:       http.StatusText(http.StatusBadRequest),
	},
	StatusBadRequestInvalidType: {
		HttpCode:    http.StatusBadRequest,
		Description: "Type can only increase or decrease",
		Title:       http.StatusText(http.StatusBadRequest),
	},
	StatusBadRequestInvalidPhoneNumber: {
		HttpCode:    http.StatusBadRequest,
		Description: "Invalid Phone Number",
		Title:       http.StatusText(http.StatusBadRequest),
	},
	StatusBadRequestInvalidPin: {
		HttpCode:    http.StatusBadRequest,
		Description: "invalid pin",
		Title:       http.StatusText(http.StatusBadRequest),
	},
	StatusBadRequestInvalidLicenseNumber: {
		HttpCode:    http.StatusBadRequest,
		Description: "invalid license number",
		Title:       http.StatusText(http.StatusBadRequest),
	},
	StatusBadRequestInvalidDriverId: {
		HttpCode:    http.StatusBadRequest,
		Description: "invalid driver id",
		Title:       http.StatusText(http.StatusBadRequest),
	},
	StatusBadRequestInvalidPoliceNumber: {
		HttpCode:    http.StatusBadRequest,
		Description: "invalid police number",
		Title:       http.StatusText(http.StatusBadRequest),
	},
	StatusBadRequestInvalidParameter: {
		HttpCode:    http.StatusBadRequest,
		Description: "invalid parameter",
		Title:       http.StatusText(http.StatusBadRequest),
	},
	StatusBadRequestRefNoAlreadyExists: {
		HttpCode:    http.StatusBadRequest,
		Description: "the refs no have already existed",
		Title:       http.StatusText(http.StatusBadRequest),
	},
	StatusBadRequestMobilePhoneAlreadyExists: {
		HttpCode:    http.StatusBadRequest,
		Description: "Mobile Phone already exists",
		Title:       http.StatusText(http.StatusBadRequest),
	},
	StatusBadRequestEmailAlreadyExists: {
		HttpCode:    http.StatusBadRequest,
		Description: "Email already exists",
		Title:       http.StatusText(http.StatusBadRequest),
	},
	StatusBadRequestCompanyAlreadyExists: {
		HttpCode:    http.StatusBadRequest,
		Description: "Company already exists",
		Title:       http.StatusText(http.StatusBadRequest),
	},
	StatusBadRequestContractStartNull: {
		HttpCode:    http.StatusBadRequest,
		Description: "Contract Start cannot be null",
		Title:       http.StatusText(http.StatusBadRequest),
	},
	StatusBadRequestContractEndNull: {
		HttpCode:    http.StatusBadRequest,
		Description: "Contract End cannot be null",
		Title:       http.StatusText(http.StatusBadRequest),
	},
	StatusBadRequestEmployeeAlreadyResign: {
		HttpCode:    http.StatusBadRequest,
		Description: "Work Status Cannot be Changed, Because This Employee Has Resigned",
		Title:       http.StatusText(http.StatusBadRequest),
	},
	StatusUnauthorizedInvalidToken: {
		HttpCode:    http.StatusUnauthorized,
		Description: "invalid token",
		Title:       http.StatusText(http.StatusUnauthorized),
	},
	StatusUnauthorizedTokenNotExists: {
		HttpCode:    http.StatusUnauthorized,
		Description: "token not exists",
		Title:       http.StatusText(http.StatusUnauthorized),
	},
	StatusUnauthorizedApiKeyNotExists: {
		HttpCode:    http.StatusUnauthorized,
		Description: "api key not exists",
		Title:       http.StatusText(http.StatusUnauthorized),
	},
	StatusUnauthorizedInvalidApiKey: {
		HttpCode:    http.StatusUnauthorized,
		Description: "invalid api key",
		Title:       http.StatusText(http.StatusUnauthorized),
	},
	StatusUnauthorizedInvalidUsernamePassword: {
		HttpCode:    http.StatusUnauthorized,
		Description: "invalid username or password",
		Title:       http.StatusText(http.StatusUnauthorized),
	},
	StatusUnauthorizedBearerNotFound: {
		HttpCode:    http.StatusUnauthorized,
		Description: "Error verifying JWT token: 'Bearer ' Not Found",
		Title:       http.StatusText(http.StatusUnauthorized),
	},
	StatusUnauthorizedErrorVerifying: {
		HttpCode:    http.StatusUnauthorized,
		Description: "Error verifying JWT token",
		Title:       http.StatusText(http.StatusUnauthorized),
	},
	StatusPreconditionFailedTokenExpired: {
		HttpCode:    http.StatusPreconditionFailed,
		Description: "token expired",
		Title:       http.StatusText(http.StatusPreconditionFailed),
	},
	StatusPreconditionFailedTokenNoLongerExists: {
		HttpCode:    http.StatusPreconditionFailed,
		Description: "token no longer exists",
		Title:       http.StatusText(http.StatusPreconditionFailed),
	},
	StatusPreconditionFailedInvalidParameter: {
		HttpCode:    http.StatusPreconditionFailed,
		Description: "invalid parameter",
		Title:       http.StatusText(http.StatusPreconditionFailed),
	},
	StatusInternalServerErrorDB: {
		HttpCode:    http.StatusInternalServerError,
		Description: "internal server error",
		Title:       http.StatusText(http.StatusInternalServerError),
	},
	InternalServerErrorMessageBroker: {
		HttpCode:    http.StatusInternalServerError,
		Description: "Internal Server Error: Message Broker Service Not Found",
		Title:       http.StatusText(http.StatusInternalServerError),
	},
	StatusInternalServerError: {
		HttpCode:    http.StatusInternalServerError,
		Description: "internal server error",
		Title:       http.StatusText(http.StatusInternalServerError),
	},
	StatusInternalServerErrorPubSub: {
		HttpCode:    http.StatusInternalServerError,
		Description: "Internal Server Error: PubSub Err",
		Title:       http.StatusText(http.StatusInternalServerError),
	},
	StatusInternalServerErrorDapr: {
		HttpCode:    http.StatusInternalServerError,
		Description: "Internal Server Error: Dapr Err",
		Title:       http.StatusText(http.StatusInternalServerError),
	},
}

func (i InternalError) Info() ErrorInfo {
	return errorInfo[i]
}

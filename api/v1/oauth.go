package v1

import (
	"awesomeProject/api/common"
	"awesomeProject/config"
	v1 "awesomeProject/core/v1"
	"awesomeProject/core/v1/api"
	"awesomeProject/core/v1/service"
	"awesomeProject/enums"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
	"log"
	"strconv"
)

type oauthApi struct {
	jwt  service.Jwt
	user service.UserService
}

// Login... logged in
// @Summary login by user mail and password
// @Description api for logging in by user mail and password
// @Accept  json
// @Produce  json
// @Param loginDto body v1.LoginDto true "loginDto object that needs to be added to log in"
// @Param grant_type query string true "grant type"
// @Success 200 {object} common.ResponseDTO
// @Failure 400 {object} common.ResponseDTO
// @Router /oauth/login [post]
func (o oauthApi) Login(context echo.Context) error {
	if context.QueryParam("grant_type") == string(enums.PASSWORD) {
		return o.handlePasswordGrant(context)
	}
	return common.GenerateForbiddenResponse(context, nil, "Please provide a valid grant_type")
}

func (o oauthApi) handlePasswordGrant(context echo.Context) error {
	tokenType := context.QueryParam("token_type")
	if tokenType == "" {
		tokenType = string(enums.REGULAR_TOKEN)
	} else if tokenType != string(enums.REGULAR_TOKEN) {
		return common.GenerateErrorResponse(context, "No valid token tokenType provided!", "Please provide a valid tokenType!")
	}
	loginDto := new(v1.LoginDto)
	if err := context.Bind(&loginDto); err != nil {
		log.Println("Input Error:", err.Error())
		return common.GenerateErrorResponse(context, "[ERROR]: Failed bind payload from context", err.Error())
	}
	if loginDto.Password == "" {
		return common.GenerateForbiddenResponse(context, "[ERROR]: Passoword is required!", "Please provide valid password!")
	}
	existingUser := o.user.GetByEmail(loginDto.Email)
	if existingUser.Id == "" || existingUser.Status != enums.ACTIVE {
		return common.GenerateForbiddenResponse(context, "[ERROR]: No User found!", "Please login with actual user email!")
	}
	err := bcrypt.CompareHashAndPassword([]byte(existingUser.Password), []byte(loginDto.Password))
	if err != nil {
		return common.GenerateForbiddenResponse(context, "[ERROR]: Password not matched!", "Please login with due credential!"+err.Error())
	}
	var tokenLifeTime int64
	userPermission := v1.UserResourcePermissionDto{
		UserId: existingUser.Id,
		Email:  existingUser.Email,
		Role:   existingUser.Role,
	}
	if tokenType == string(enums.REGULAR_TOKEN) {
		i, err := strconv.ParseInt(config.RegularTokenLifetime, 10, 64)
		if err != nil {
			log.Println(err.Error())
			return common.GenerateForbiddenResponse(context, "[ERROR]: failed to read regular token lifetime from env!", err.Error())
		}
		tokenLifeTime = i
	} else {
		i, err := strconv.ParseInt(config.CTLTokenLifetime, 10, 64)
		if err != nil {
			log.Println(err.Error())
			return common.GenerateForbiddenResponse(context, "[ERROR]: failed to read ctl token lifetime from env!", err.Error())
		}
		tokenLifeTime = i
	}
	token, refreshToken, err := o.jwt.GenerateToken(userPermission.UserId, tokenLifeTime, userPermission)
	if err != nil {
		log.Println(err.Error())
		return common.GenerateForbiddenResponse(context, "[ERROR]: failed to create token!", err.Error())
	}
	return common.GenerateSuccessResponse(context, v1.JWTPayLoad{AccessToken: token, RefreshToken: refreshToken}, nil, "Operation Successful")
}

func NewOauthApi(jwt service.Jwt, user service.UserService) api.Oauth {
	return &oauthApi{jwt: jwt, user: user}
}

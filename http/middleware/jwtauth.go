package middleware

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	. "github.com/hulutech-web/frame/database"
	models "github.com/hulutech-web/frame/model/sysmodel"
	"github.com/hulutech-web/frame/request"
	"net/http"
	"reflect"
	"strconv"
	"strings"

	"time"
)


// 新建一个jwt实例
func NewJWT(signKey string) *JWT {
	return &JWT{
		[]byte(signKey),
	}
}

// JWTAuth 中间件，检查token
func JWTAuthRequired(signKey string) request.HandlerFunc {
	return func(c request.Context) {
		token := c.Request().Header.Get("Authorization")
		//提取token
		token = strings.Replace(token, "Bearer ", "", -1)
		if token == "" {
			c.JSON(http.StatusOK, gin.H{
				"code":  -1,
				"message": "请求未携带token，无权限访问",
			})
			c.Abort()
			return
		}
		j := NewJWT(signKey)
		// parseToken 解析token包含的信息
		claims, err := j.ParseToken(token)
		if err != nil {
			if err == TokenExpired {
				c.JSON(http.StatusOK, gin.H{
					"code":  -1,
					"message": "授权已过期",
				})
				c.Abort()
				return
			}
			c.JSON(http.StatusOK, gin.H{
				"code":  -1,
				"message": err.Error(),
			})
			c.Abort()
			return
		}

		// 继续交由下一个路由处理,并将解析出的信息传递下去
		c.Set("token", token)
		c.Set("claims", claims)
		c.Next()
	}
}

// JWT 签名结构
type JWT struct {
	SigningKey []byte
}

// 一些常量
var (
	TokenExpired     error = errors.New("Token is expired")
	TokenNotValidYet error = errors.New("Token not active yet")
	TokenMalformed   error = errors.New("That's not even a token")
	TokenInvalid     error = errors.New("Couldn't handle this token:")
)

// 载荷，可以加一些自己需要的信息
type CustomClaims struct {
	ID        string `json:"id"`
	UserName  string `json:"username"`
	ModelName string `json:"modelName"`
	Revoked   bool   `json:"revoked"`
	jwt.RegisteredClaims
}

// CreateToken 生成一个token
func (j *JWT) CreateToken(model interface{}) (string, error) {
	id, modelName, username := ReflectTokenModel(model)
	claims := CustomClaims{
		strconv.Itoa(id),
		username,
		modelName,
		false,
		jwt.RegisteredClaims{
			NotBefore: jwt.NewNumericDate(time.Now()),                       // 签名生效时间
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(60 * time.Minute)), // 过期时间 一小时
			Issuer:    "newtrekWang",                                        //签名的发行者
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	//这里往personal_access_token中插入一条数据
	OwnerId := id
	tokenstr, err := token.SignedString(j.SigningKey)
	//插入数据库===================
	DB().Create(&models.PersonalAccessToken{
		//string转int
		OwnerId:    uint(OwnerId),
		OwnerType:  modelName,
		Token:      tokenstr,
		LastUsedAt: time.Now(),
	})
	return tokenstr, err
}

// 解析Token
func (j *JWT) ParseToken(tokenString string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.SigningKey, nil
	})
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, TokenMalformed
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				// Token is expired
				return nil, TokenExpired
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, TokenNotValidYet
			} else {
				return nil, TokenInvalid
			}
		}
	}
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, TokenInvalid
}

// 更新token
func (j *JWT) RefreshToken(tokenString string) (string, error) {
	jwt.TimeFunc = func() time.Time {
		return time.Unix(0, 0)
	}
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.SigningKey, nil
	})
	if err != nil {
		return "", err
	}
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		jwt.TimeFunc = time.Now
		//转int64
		claims.RegisteredClaims.ExpiresAt = jwt.NewNumericDate(time.Now().Add(24 * time.Hour))
		return j.CreateToken(claims)
	}
	return "", TokenInvalid
}

// 删除token
func (j *JWT) DeleteToken(tokenString string) error {
	return nil
}

// 反射取出id和userModelName
func ReflectTokenModel(a interface{}) (int, string, string) {
	typ := reflect.TypeOf(a)
	val := reflect.ValueOf(a)
	kd := val.Kind()

	if kd != reflect.Struct {
		fmt.Println("expect struct")
		return 0, "", ""
	}
	//获取到结构体中的id字段
	idField := val.FieldByName("ID")
	usernameField := val.FieldByName("Username")

	//获取id字段的值
	id := idField.Uint()
	//获取username字段的值
	username := usernameField.String()
	//获取模型名称，比如AdminUser或者User
	modelName := typ.Name()
	return int(id), modelName, username
}

func AuthUser(token string) interface{} {
	//从PersonalAccessToken表中查询token
	var personalAccessToken models.PersonalAccessToken
	//查出最新的token按last_used_at排序，查询出owner_id和owner_type
	DB().Where("token = ?", token).Order("last_used_at desc").First(&personalAccessToken)
	//根据owner_type和owner_id查询出用户
	if personalAccessToken.OwnerType == "Admin" {
		var adminUser *models.Admin
		DB().Where("id = ?", personalAccessToken.OwnerId).First(adminUser)
		return adminUser
	} else {
		var user *models.User
		DB().Where("id = ?", personalAccessToken.OwnerId).First(user)
		return user
	}
}

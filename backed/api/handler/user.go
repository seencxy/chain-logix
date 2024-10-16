package handler

import (
	"backed/app/core/consts"
	"backed/app/core/variable"
	"backed/app/models"
	"backed/app/utils"
	"backed/internal/database"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Register 用户注册接口
func Register(ctx *gin.Context) {
	// 定义请求参数
	var requestData models.RegisterRequest
	// 绑定参数
	err := ctx.ShouldBindJSON(&requestData)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": "获取请求参数失败: " + err.Error(),
		})
		return
	}

	// 检验邮箱是否合理
	if !utils.ValidateEmail(requestData.Email) {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": "邮箱格式不正确",
		})
		return
	}

	// 判断验证码是否正确
	code, err := variable.RedisClient.Get(ctx, requestData.Email).Result()
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": "获取验证码错误: " + err.Error(),
		})
		return
	}

	// 判断是否相等
	if code != requestData.Code {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": "验证码错误",
		})
		return
	}
	userDao := database.UserInfo

	// 判断当前邮箱是否注册过账号
	count, err := userDao.WithContext(ctx).
		Where(userDao.Email.Eq(requestData.Email)).
		Count()
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": "获取邮箱信息错误: " + err.Error(),
		})
		return
	}

	if count != 0 {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": "当前邮箱已经注册过账号",
		})
		return
	}

	// 注册用户
	info := models.NewUserInfo(requestData)
	// 创建用户对象
	err = userDao.WithContext(ctx).Create(info)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": "创建用户信息失败: " + err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "注册成功",
	})
}

// SendEmailCode 发送邮箱验证码接口
func SendEmailCode(ctx *gin.Context) {
	// 请求参数
	var requestData models.SendEmailCodeRequest
	// 绑定参数
	err := ctx.ShouldBindJSON(&requestData)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": "获取请求参数失败: " + err.Error(),
		})
		return
	}

	// 生成随机验证码
	code := utils.GenerateRandomString(6)
	// 将code存入redis数据库中
	err = variable.RedisClient.Set(ctx, requestData.Email, code, consts.EmailCodeExpireDuration).Err()
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": "将验证码存储数据库失败: " + err.Error(),
		})
		return
	}

	// 发送验证码
	err = utils.SendEmail(requestData.Email,
		code,
		variable.Viper.GetString("email.username"),
		variable.Viper.GetString("email.password"),
		variable.Viper.GetString("email.host"),
		variable.Viper.GetString("email.port"))
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": "发送验证码失败: " + err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "发送验证码成功",
	})
}

// Login 用户登录接口
func Login(ctx *gin.Context) {
	// 请求参数
	var requestData models.LoginRequest
	// 绑定参数
	err := ctx.ShouldBindJSON(&requestData)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": "获取请求参数失败: " + err.Error(),
		})
		return
	}

	// 检验邮箱是否合理
	if !utils.ValidateEmail(requestData.Email) {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": "邮箱格式不正确",
		})
		return
	}

	userDao := database.UserInfo
	// 查询用户信息
	userInfo, err := userDao.WithContext(ctx).
		Where(userDao.Email.Eq(requestData.Email)).
		Where(userDao.Identity.Eq(requestData.Identity)).
		First()
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": "获取用户信息失败: " + err.Error(),
		})
		return
	}

	if userInfo.Password != requestData.Password {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": "密码错误",
		})
		return
	}

	if userInfo.Status != 1 {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": "账号未审核或账号信息已停用",
		})
		return
	}

	// 生成token
	token, err := utils.GenToken(userInfo.Email, consts.TokenExpireDuration, []byte(consts.TokenSecret))
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": "生成token失败: " + err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code":     200,
		"message":  "登录成功",
		"token":    token,
		"email":    userInfo.Email,
		"identity": userInfo.Identity,
		"pic":      userInfo.Head,
	})
}

// ResetPassword 重置密码接口
func ResetPassword(ctx *gin.Context) {
	// 请求参数
	var requestData models.RegisterRequest
	// 绑定参数
	err := ctx.ShouldBindJSON(&requestData)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": "获取请求参数失败: " + err.Error(),
		})
		return
	}

	// 检验邮箱是否合理
	if !utils.ValidateEmail(requestData.Email) {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": "邮箱格式不正确",
		})
		return
	}

	// 判断验证码是否正确
	code, err := variable.RedisClient.Get(ctx, requestData.Email).Result()
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": "获取验证码错误: " + err.Error(),
		})
		return
	}

	// 判断是否相等
	if code != requestData.Code {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": "验证码错误",
		})
		return
	}

	userDao := database.UserInfo
	// 查询用户信息
	userInfo, err := userDao.WithContext(ctx).
		Where(userDao.Email.Eq(requestData.Email)).
		Where(userDao.Identity.Eq(requestData.Identity)).
		First()
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": "获取用户信息失败: " + err.Error(),
		})
		return
	}

	// 修改密码
	userInfo.Password = requestData.Password
	err = userDao.WithContext(ctx).Save(userInfo)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": "修改密码失败: " + err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "修改密码成功",
	})
}

// ChangeUserHead 改变用户头像
func ChangeUserHead(ctx *gin.Context) {
	// 请求参数
	var requestData models.ChangeUserHeadRequest
	// 绑定参数
	err := ctx.ShouldBindJSON(&requestData)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": "获取请求参数失败: " + err.Error(),
		})
		return
	}

	// 检验邮箱是否合理
	if !utils.ValidateEmail(requestData.Email) {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": "邮箱格式不正确",
		})
		return
	}

	// 获取用户信息
	userDao := database.UserInfo
	userInfo, err := userDao.WithContext(ctx).
		Where(userDao.Email.Eq(requestData.Email)).
		First()

	if userInfo.Status == 0 || userInfo.Status == 2 {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": "账号未审核或账号信息已停用",
		})
		return
	}

	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": "获取用户信息失败: " + err.Error(),
		})
		return
	}

	// 更新用用户信息
	userInfo.Head = requestData.Head
	err = userDao.WithContext(ctx).Save(userInfo)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": "修改头像失败: " + err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "修改头像成功",
	})
}

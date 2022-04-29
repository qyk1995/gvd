package service

import (
	"fmt"

	"example.com/m/api"
	"example.com/m/model"

	"example.com/m/utils"
	"github.com/gin-gonic/gin"
)

type BaseApi struct{}

type SysUsersResponse struct {
	User model.SysUser `json:"user"`
}

func (b *BaseApi) Register(c *gin.Context) {
	var r model.SysUser
	fmt.Println("1111111111")
	_ = c.ShouldBindJSON(&r)
	if err := utils.Verify(r, utils.RegistuserVerify); err != nil { // 做校验
		utils.FailWithMessage(err.Error(), c)
		return
	}
	user := &model.SysUser{Username: r.Username, Password: r.Password}
	err, userReturn := api.Register(*user)
	if err != nil {
		utils.FailWithDetailed(SysUsersResponse{User: userReturn}, "注册失败", c)
	} else {
		utils.OkWithDetailed(SysUsersResponse{User: userReturn}, "注册成功", c)
	}

}
func (b *BaseApi) Login(c *gin.Context) {
	var l model.Login
	_ = c.ShouldBindJSON(&l)
	if err := utils.Verify(l, utils.LoginVerify); err != nil {
		utils.FailWithMessage(err.Error(), c)
		return
	}

	u := &model.SysUser{Username: l.Username, Password: l.Password}
	err, user := api.Login(u)
	if err != nil {
		utils.FailWithMessage("用户名不存在或者密码错误", c)
	} else {
		b.tokenNext(c, *user)
	}
}

func (b *BaseApi) tokenNext(c *gin.Context, user model.SysUser) {
	j := &utils.JWT{SigningKey: []byte(global.GVA_CONFIG.JWT.SigningKey)} // 唯一签名
	claims := j.CreateClaims(systemReq.BaseClaims{
		UUID:     user.UUID,
		ID:       user.ID,
		Username: user.Username,
	})
	token, err := j.CreateToken(claims)
	if err != nil {
		utils.FailWithMessage("获取token失败", c)
		return
	}
	if !global.GVA_CONFIG.System.UseMultipoint {
		utils.OkWithDetailed(systemRes.LoginResponse{
			User:      user,
			Token:     token,
			ExpiresAt: claims.StandardClaims.ExpiresAt * 1000,
		}, "登录成功", c)
		return
	}

	if err, jwtStr := jwtService.GetRedisJWT(user.Username); err == redis.Nil {
		if err := jwtService.SetRedisJWT(token, user.Username); err != nil {
			utils.FailWithMessage("设置登录状态失败", c)
			return
		}
		utils.OkWithDetailed(systemRes.LoginResponse{
			User:      user,
			Token:     token,
			ExpiresAt: claims.StandardClaims.ExpiresAt * 1000,
		}, "登录成功", c)
	} else if err != nil {
		utils.FailWithMessage("设置登录状态失败", c)
	} else {
		var blackJWT system.JwtBlacklist
		blackJWT.Jwt = jwtStr
		if err := jwtService.JsonInBlacklist(blackJWT); err != nil {
			utils.FailWithMessage("jwt作废失败", c)
			return
		}
		if err := jwtService.SetRedisJWT(token, user.Username); err != nil {
			utils.FailWithMessage("设置登录状态失败", c)
			return
		}
		utils.OkWithDetailed(systemRes.LoginResponse{
			User:      user,
			Token:     token,
			ExpiresAt: claims.StandardClaims.ExpiresAt * 1000,
		}, "登录成功", c)
	}

}

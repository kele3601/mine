package utils

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"mine/internal/db/model"
	"testing"
)

func TestGenerateJwtToken(t *testing.T) {
	// 创建一个 User 实例
	user := model.User{
		BaseModel: model.BaseModel{ID: "123456"},
	}

	// 调用 GenerateJwtToken 函数
	token, err := GenerateJwtToken(user)

	// 验证没有错误发生
	require.NoError(t, err)

	// 验证 token 不是空的
	assert.NotEmpty(t, token)

	// 可选：验证 token 是否可以被正确解析
	// 这里需要导入 "github.com/dgrijalva/jwt-go" 包来解析 token
	jwtToken, err := ParseJwtToken(token)

	// 验证 token 解析没有错误
	require.NoError(t, err)

	// 验证解析后的 UserID 是否与创建 User 实例时的 ID 相匹配
	assert.Equal(t, user.ID, jwtToken.UserID)
}

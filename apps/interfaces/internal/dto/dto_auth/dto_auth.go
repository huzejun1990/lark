// @Author huzejun 2024/3/7 1:34:00
package dto_auth

type SignUpReq struct {
	RegPlatform int32  `json:"reg_platform" binding:"required,oneof=1 2 3 4 5"` // 注册平台
	Nickname    string `json:"nickname" binding:"required"`                     // 昵称
	Password    string `json:"password" binding:"required,len=32"`              // 密码
	Firstname   string `json:"firstname,omitempty"`                             // firstname
	Lastname    string `json:"lastname,omitempty"`                              // lastname
	Gender      int32  `json:"gender,omitempty"`                                // 性别
	BirthTs     int64  `json:"birth_ts,omitempty"`                              // 生日
	Email       string `json:"email,omitempty"`                                 // Email
	Mobile      string `json:"mobile,omitempty"`                                // 手机号
	AvatarKey   string `json:"avatar_key,omitempty"`                            // 头像
	CityId      int64  `json:"city_id,omitempty"`                               // 城市ID
	Code        int64  `json:"code,omitempty"`                                  // 验证码
	Udid        string `json:"udid,omitempty"`
	ServerId    int64  `json:"server_id,omitempty"` // udid
}

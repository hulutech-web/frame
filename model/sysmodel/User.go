package sysmdel

import (
"gorm.io/gorm"
"time"
)

type User struct {
	gorm.Model
	Fid                 int                   `gorm:"column:fid;type:int(11);not null;default:0;comment:'上级id'" json:"fid"`
	Username            string                `gorm:"column:username;type:varchar(50);not null;default:'';comment:'用户名'" json:"username"`
	Password            string                `gorm:"column:password;type:varchar(100);not null;default:'';comment:'密码'" json:"password"`
	ApiToken            string                `gorm:"column:api_token;type:varchar(100);not null;unique;default:'';comment:'api_token';" json:"api"`
	Realname            string                `gorm:"column:realname;type:varchar(50);not null;default:'';comment:'真实姓名'" json:"realname"`
	Vip                 int                   `gorm:"column:vip;type:int(11);not null;default:0;comment:'vip等级'" json:"vip"`
	Level               int                   `gorm:"column:level;type:int(11);not null;default:0;comment:'等级'" json:"level"`
	Exp                 int                   `gorm:"column:exp;type:int(11);not null;default:0;comment:'经验值'" json:"exp"`
	Paypwd              string                `gorm:"column:paypwd;type:varchar(100);not null;default:'';comment:'支付密码'" json:"paypwd"`
	Isonline            int                   `gorm:"column:isonline;type:int(11);not null;default:0;comment:'是否在线'" json:"isonline"`
	Avatar              string                `gorm:"column:avatar;type:varchar(255);not null;default:'';comment:'头像'" json:"avatar"`
	Allowagent          int                   `gorm:"column:allowagent;type:int(11);not null;default:0;comment:'是否允许发展下级代理'" json:"allowagent"`
	Balance             float64               `gorm:"column:balance;type:decimal(10,2);not null;default:0.00;comment:'余额'" json:"balance"`
	Mbalance            float64               `gorm:"column:mbalance;type:decimal(10,2);not null;default:0.00;comment:'码量余额'" json:"mbalance"`
	Totalgame           float64               `gorm:"column:totalgame;type:decimal(10,2);not null;default:0.00;comment:'总共游戏'" json:"totalgame"`
	Phone               string                `gorm:"column:phone;type:varchar(20);not null;default:'';comment:'手机号'" json:"phone"`
	Mail                string                `gorm:"column:mail;type:varchar(50);not null;default:'';comment:'邮箱'" json:"mail"`
	Paysum              float64               `gorm:"column:paysum;type:decimal(10,2);not null;default:0.00;comment:'累计充值'" json:"paysum"`
	Status              int                   `gorm:"column:status;type:int(11);not null;default:0;comment:'状态'" json:"status"`
	Isdel               int                   `gorm:"column:isdel;type:int(11);not null;default:0;comment:'是否删除'" json:"isdel"`
	Isblack             int                   `gorm:"column:isblack;type:int(11);not null;default:0;comment:'是否黑名单'" json:"isblack"`
	Lastip              string                `gorm:"column:lastip;type:varchar(20);not null;default:'';comment:'最后登录ip'" json:"lastip"`
	Lastloginipaddress  string                `gorm:"column:last_login_ip_address;type:varchar(255);not null;default:'';comment:'最后登录ip地址'" json:"last_login_ip_address"`
	Logintime           time.Time             `gorm:"column:logintime;type:datetime;not null;default:CURRENT_TIMESTAMP;comment:'最后登录时间'" json:"logintime"`
	Sourceurl           string                `gorm:"column:sourceurl;type:varchar(255);not null;default:'';comment:'来源url'" json:"sourceurl"`
	Loginsum            int                   `gorm:"column:loginsum;type:int(11);not null;default:0;comment:'登录次数'" json:"loginsum"`
	Birthday            time.Time             `gorm:"column:birthday;type:datetime;not null;default:CURRENT_TIMESTAMP;comment:'生日'" json:"birthday"`
	Isagent             int                   `gorm:"column:isagent;type:int(11);not null;default:0;comment:'是否代理'" json:"isagent"`
	Pid                 int                   `gorm:"column:pid;type:int(11);not null;default:0;comment:'代理id'" json:"pid"`
	Settlementid        int                   `gorm:"column:settlement_id;type:int(11);not null;default:0;comment:'结算方案id'" json:"settlement_id"`
	Fanshuifee          float64               `gorm:"column:fanshuifee;type:decimal(10,2);not null;default:0.00;comment:'反水费率'" json:"fanshuifee"`
	Settlementday       int                   `gorm:"column:settlementday;type:int(11);not null;default:0;comment:'最后一次结算时间'" json:"settlementday"`
	Regip               string                `gorm:"column:fanshuifee;varchar(255);not null;default:'';comment:'注册ip'" json:"reg_ip"`
	Transferstatus      int8                  `gorm:"column:transferstatus;type:tinyint(1);not null;default:1;comment:'转账状态:0 转账 1免转'" json:"transferstatus"`
	Adminid             int                   `gorm:"column:admin_id;type:int(11);not null;default:0;comment:'管理员id'" json:"admin_id"`
	PersonalAccessToken []PersonalAccessToken `gorm:"foreignKey:OwnerId" json:"personal_access_token"` //定义token关联
}

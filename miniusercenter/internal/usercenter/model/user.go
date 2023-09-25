package model

// UserM mapped from table <user>
type UserM struct {
	ID       int64  `gorm:"column:id;type:bigint(20) unsigned;primaryKey;autoIncrement:true;comment:主键 ID" json:"id"`
	Username string `gorm:"column:username;type:varchar(253);not null;uniqueIndex:idx_username,priority:1;comment:用户名称" json:"username"`
	Password string `gorm:"column:password;type:varchar(64);not null;comment:用户加密后的密码" json:"password"`
}

// TableName UserM's table name
func (*UserM) TableName() string {
	return "user"
}

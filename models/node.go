package models

type Node struct {
	DefaultModel
	GroupID     int64  `gorm:"column:group_id" json:"group_id"`
	Role        int8   `gorm:"column:role" json:"role"`
	IsCandidate bool   `gorm:"column:is_candidate" json:"is_candidate"`
	Username    string `gorm:"column:username" json:"username"`
	Password    string `gorm:"column:password" json:"password"`
	Host        string `gorm:"column:host" json:"host"`
	Port        int    `gorm:"column:port" json:"port"`
	DBName      string `gorm:"column:db_name" json:"db_name"`
	ReadWeight  int    `gorm:"column:read_weight" json:"read_weight"`
	Charset     string `gorm:"column:charset" json:"charset"`
	AutoCommit  bool   `gorm:"column:auto_commit" json:"auto_commit"`
}

func (Node) TableName() string {
	return "nodes"
}
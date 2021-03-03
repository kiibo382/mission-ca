package model

// type Users struct {
//     Id    int64  `xorm:"pk autoincr int(64)" form:"id" json:"id"`
//     Name string `xorm:"varchar(40)" json:"name" form:"name"`
//     Token string `xorm:"varchar(40)" json:"token" form:"token"`
// }

type User struct {
	// User     string `form:"user" json:"user" xml:"user"  binding:"required"`
	// Password string `form:"password" json:"password" xml:"password" binding:"required"`
    // Id    int64  `gorm:"pk autoincr int(64)" json:"id"`
    // Token string `gorm:"varchar(40)" json:"token"`
    Name string `gorm:"varchar(40)" json:"name"`
    Mail string `gorm:"varchar(40)" json:"mail"`
}
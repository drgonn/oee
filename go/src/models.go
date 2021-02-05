package main
import (
	"github.com/jinzhu/gorm"
	"time"
)
var db *gorm.DB
func init() {
	var err error
	db, err = gorm.Open("mysql", "root:781117@/bridge?charset=utf8&parseTime=True&loc=Local")
	db.SingularTable(true)
	gorm.DefaultTableNameHandler = func(db * gorm.DB, defaultTableName string) string {
		return defaultTableName + "s";
	}
	if err != nil {
		panic("failed to conect database")
	}
	db.AutoMigrate(&User{})
	db.AutoMigrate(&Device{})
	db.AutoMigrate(&Worktime{})
	db.AutoMigrate(&Devicestatu{})
	db.AutoMigrate(&Valvetime{})
	db.AutoMigrate(&Valvetype{})
	db.AutoMigrate(&Valve{})
	db.AutoMigrate(&Bug{})
	db.AutoMigrate(&Bugtype{})
	db.AutoMigrate(&Alarmtype{})
	db.AutoMigrate(&Alarm{})
	db.AutoMigrate(&Project{})
	db.AutoMigrate(&Plan{})
	db.AutoMigrate(&Message{})
}
type User struct {
	ID uint
	Uid string `gorm:"size:64"`
	Name string `gorm:"size:64"`
	Createdate time.Time
	Plans []Plan `gorm:"foreignkey:UserID"`
	Messages []Message `gorm:"foreignkey:UserID"`
}
type jsonUser struct {
	ID uint  `json:"id"`
	Uid string `json:"uid"`
	Name string `json:"name"`
	Createdate time.Time `json:"createDate"`
}
type Device struct {
	ID uint
	Sn string `gorm:"size:16"`
	Name string `gorm:"size:64"`
	Ip string `gorm:"size:64"`
	Img string `gorm:"size:256"`
	Type string `gorm:"size:64"`
	Worktimes []Worktime `gorm:"foreignkey:DeviceID"`
	Devicestatus []Devicestatu `gorm:"foreignkey:DeviceID"`
	Valvetimes []Valvetime `gorm:"foreignkey:DeviceID"`
	Valves []Valve `gorm:"foreignkey:DeviceID"`
	Bugs []Bug `gorm:"foreignkey:DeviceID"`
	Alarms []Alarm `gorm:"foreignkey:DeviceID"`
}
type jsonDevice struct {
	ID uint  `json:"id"`
	Sn string `json:"sn"`
	Name string `json:"name"`
	Ip string `json:"ip"`
	Img string `json:"img"`
	Type string `json:"type"`
}
type Worktime struct {
	ID uint
	Start_Time time.Time
	End_Time time.Time
	Seconds int
	Type int
	Amount int
	Good int
	Glue float32 `gorm:"type:float"`
	DeviceID uint
}
type jsonWorktime struct {
	ID uint  `json:"id"`
	Start_Time time.Time `json:"start_time"`
	End_Time time.Time `json:"end_time"`
	Seconds int `json:"seconds"`
	Type int `json:"type"`
	Amount int `json:"amount"`
	Good int `json:"good"`
	Glue float32 `json:"glue"`
	DeviceID uint `json:"device_id"`
}
type Devicestatu struct {
	ID uint
	Start_Time time.Time
	Status int
	DeviceID uint
}
type jsonDevicestatu struct {
	ID uint  `json:"id"`
	Start_Time time.Time `json:"start_time"`
	Status int `json:"status"`
	DeviceID uint `json:"device_id"`
}
type Valvetime struct {
	ID uint
	Start_Time time.Time
	End_Time time.Time
	Seconds float32 `gorm:"type:float"`
	Volt float32 `gorm:"type:float"`
	Amount int
	Good int
	Glue float32 `gorm:"type:float"`
	DeviceID uint
	ValveID uint
}
type jsonValvetime struct {
	ID uint  `json:"id"`
	Start_Time time.Time `json:"start_time"`
	End_Time time.Time `json:"end_time"`
	Seconds float32 `json:"seconds"`
	Volt float32 `json:"volt"`
	Amount int `json:"amount"`
	Good int `json:"good"`
	Glue float32 `json:"glue"`
	DeviceID uint `json:"device_id"`
	ValveID uint `json:"valve_id"`
}
type Valvetype struct {
	ID uint
	Name string `gorm:"size:64"`
	Valves []Valve `gorm:"foreignkey:ValvetypeID"`
}
type jsonValvetype struct {
	ID uint  `json:"id"`
	Name string `json:"name"`
}
type Valve struct {
	ID uint
	Sn string `gorm:"size:16"`
	Name string `gorm:"size:64"`
	ValvetypeID uint
	DeviceID uint
	Valvetimes []Valvetime `gorm:"foreignkey:ValveID"`
	Alarms []Alarm `gorm:"foreignkey:ValveID"`
}
type jsonValve struct {
	ID uint  `json:"id"`
	Sn string `json:"sn"`
	Name string `json:"name"`
	ValvetypeID uint `json:"valvetype_id"`
	DeviceID uint `json:"device_id"`
}
type Bug struct {
	ID uint
	Reason string `gorm:"type:text"`
	Start_Time time.Time
	End_Time time.Time
	DeviceID uint
	BugtypeID uint
}
type jsonBug struct {
	ID uint  `json:"id"`
	Reason string `json:"reason"`
	Start_Time time.Time `json:"start_time"`
	End_Time time.Time `json:"end_time"`
	DeviceID uint `json:"device_id"`
	BugtypeID uint `json:"bugtype_id"`
}
type Bugtype struct {
	ID uint
	Name string `gorm:"size:64"`
	Sn string `gorm:"size:64"`
	About string `gorm:"type:text"`
	Bugs []Bug `gorm:"foreignkey:BugtypeID"`
}
type jsonBugtype struct {
	ID uint  `json:"id"`
	Name string `json:"name"`
	Sn string `json:"sn"`
	About string `json:"about"`
}
type Alarmtype struct {
	ID uint
	Code string `gorm:"size:64"`
	Mean string `gorm:"type:text"`
	Cause string `gorm:"type:text"`
	Solution string `gorm:"type:text"`
	Alarms []Alarm `gorm:"foreignkey:AlarmtypeID"`
}
type jsonAlarmtype struct {
	ID uint  `json:"id"`
	Code string `json:"code"`
	Mean string `json:"mean"`
	Cause string `json:"cause"`
	Solution string `json:"solution"`
}
type Alarm struct {
	ID uint
	Create_Time time.Time
	AlarmtypeID uint
	ValveID uint
	DeviceID uint
}
type jsonAlarm struct {
	ID uint  `json:"id"`
	Create_Time time.Time `json:"create_time"`
	AlarmtypeID uint `json:"alarmtype_id"`
	ValveID uint `json:"valve_id"`
	DeviceID uint `json:"device_id"`
}
type Project struct {
	ID uint
	Name string `gorm:"size:256"`
	Plans []Plan `gorm:"foreignkey:ProjectID"`
}
type jsonProject struct {
	ID uint  `json:"id"`
	Name string `json:"name"`
}
type Plan struct {
	ID uint
	Name string `gorm:"size:256"`
	Week int
	Current string `gorm:"type:text"`
	Follow string `gorm:"type:text"`
	Time time.Time
	ProjectID uint
	UserID uint
}
type jsonPlan struct {
	ID uint  `json:"id"`
	Name string `json:"name"`
	Week int `json:"week"`
	Current string `json:"current"`
	Follow string `json:"follow"`
	Time time.Time `json:"time"`
	ProjectID uint `json:"project_id"`
	UserID uint `json:"user_id"`
}
type Message struct {
	ID uint
	Title string `gorm:"size:256"`
	Type int
	Description string `gorm:"type:text"`
	Read bool
	Ct_Time time.Time
	UserID uint
}
type jsonMessage struct {
	ID uint  `json:"id"`
	Title string `json:"title"`
	Type int `json:"type"`
	Description string `json:"description"`
	Read bool `json:"read"`
	Ct_Time time.Time `json:"ct_time"`
	UserID uint `json:"user_id"`
}


package models

import (
	"github.com/KevinNguyen1703/FrameStoreGolangServer/pkg/config"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type Frame struct {
	gorm.Model
	Video       string `json:"Video"`
	Data        string `json:"Data"`
	Description string `json:"Description"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Frame{})
}

func (b *Frame) CreateFrame() *Frame {
	db.NewRecord(b)
	db.Create(&b)
	return b
}

func GetAllFrames() []Frame {
	var Frames []Frame
	db.Find(&Frames)
	return Frames
}

func GetFrameById(Id int64) (*Frame, *gorm.DB) {
	var getFrame Frame
	db := db.Where("ID=?", Id).Find(&getFrame)
	return &getFrame, db
}

func DeleteFrame(ID int64) Frame {
	var frame Frame
	db.Where("ID=?", ID).Delete(frame)
	return frame
}

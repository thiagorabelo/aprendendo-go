package models

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strconv"

	"gorm.io/gorm"
)

const MAX_COLLISIONS = 3

type ShortURI struct {
	gorm.Model

	FullURI     string `gorm:"full_uri;size:1024;not null"`
	URIHash     string `gorm:"uri_hash;size:32;uniqueIndex;not null"`
	Description string `gorm:"description;size:1028"`

	// Belongs to
	UserID int  `gorm:"not null"`
	User   User `gorm:"constraint:OnDelete:CASCADE"`
}

func (ShortURI) TableName() string {
	return "shorturis"
}

func (short ShortURI) String() string {
	return fmt.Sprintf("ShortURI{%s -> %s", short.URIHash, short.FullURI)
}

func (short *ShortURI) BeforeCreate(db *gorm.DB) (err error) {
	var i int
	for i = 0; i < MAX_COLLISIONS; i++ {
		hash := short.createHash(i)
		var c int64
		if db.Model(short).Where("uri_hash = ?", hash).Count(&c); c <= 0 {
			short.URIHash = hash
			return
		}
	}

	err = fmt.Errorf("max try save achieved (%d)", i)
	return
}

func (short *ShortURI) createHash(try int) string {
	// TODO: Make secret key (sk)
	sk := "123456"
	msg := sk + short.FullURI + /* username + */ strconv.Itoa(try)
	fmt.Println(msg)
	hash := md5.Sum([]byte(msg))
	slice := []rune(hex.EncodeToString(hash[:]))[:10]
	return string(slice)
}

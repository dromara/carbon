package xorm

import (
	"github.com/dromara/carbon/v2"
)

type RFC3339Layout string

func (t RFC3339Layout) Layout() string {
	return carbon.RFC3339Layout
}

type ISO8601Format string

func (t ISO8601Format) Format() string {
	return carbon.ISO8601Format
}

type MySQLModel1 struct {
	Id uint64 `json:"-" xorm:"pk autoincr"`

	Carbon1 carbon.Carbon `xorm:"varchar(50) carbon1" json:"carbon1"`
	Carbon2 carbon.Carbon `xorm:"datetime carbon2" json:"carbon2"`
	Carbon3 carbon.Carbon `xorm:"timestamp carbon3" json:"carbon3"`

	Date1 carbon.Date `xorm:"varchar(50) date1" json:"date1"`
	Date2 carbon.Date `xorm:"date date2" json:"date2"`
	Date3 carbon.Date `xorm:"timestamp date3" json:"date3"`

	Time1 carbon.Time `xorm:"varchar(50) time1" json:"time1"`
	Time2 carbon.Time `xorm:"time time2" json:"time2"`
	Time3 carbon.Time `xorm:"timestamp time3" json:"time3"`

	DateTime1 carbon.DateTime `xorm:"varchar(50) date_time1" json:"date_time1"`
	DateTime2 carbon.DateTime `xorm:"datetime date_time2" json:"date_time2"`
	DateTime3 carbon.DateTime `xorm:"timestamp date_time3" json:"date_time3"`

	RFC3339Layout1 carbon.LayoutType[RFC3339Layout] `xorm:"varchar(50) rfc3339_layout1" json:"rfc3339_layout1"`
	RFC3339Layout2 carbon.LayoutType[RFC3339Layout] `xorm:"datetime rfc3339_layout2" json:"rfc3339_layout2"`
	RFC3339Layout3 carbon.LayoutType[RFC3339Layout] `xorm:"timestamp rfc3339_layout3" json:"rfc3339_layout3"`

	ISO8601Format1 carbon.FormatType[ISO8601Format] `xorm:"varchar(50) iso8601_format1" json:"iso8601_format1"`
	ISO8601Format2 carbon.FormatType[ISO8601Format] `xorm:"datetime iso8601_format2" json:"iso8601_format2"`
	ISO8601Format3 carbon.FormatType[ISO8601Format] `xorm:"timestamp iso8601_format3" json:"iso8601_format3"`

	Timestamp1 carbon.Timestamp `xorm:"bigint timestamp1" json:"timestamp1"`

	CreatedAt carbon.DateTime `xorm:"created created_at" json:"-"`
	UpdatedAt carbon.DateTime `xorm:"updated updated_at" json:"-"`
	DeletedAt carbon.DateTime `xorm:"deleted deleted_at" json:"-"`
}

func (MySQLModel1) TableName() string {
	return "xorm_mysql1"
}

type MySQLModel2 struct {
	Id uint64 `json:"-" xorm:"pk autoincr"`

	Carbon1 *carbon.Carbon `xorm:"varchar(50) carbon1" json:"carbon1"`
	Carbon2 *carbon.Carbon `xorm:"datetime carbon2" json:"carbon2"`
	Carbon3 *carbon.Carbon `xorm:"timestamp carbon3" json:"carbon3"`

	Date1 *carbon.Date `xorm:"varchar(50) date1" json:"date1"`
	Date2 *carbon.Date `xorm:"date date2" json:"date2"`
	Date3 *carbon.Date `xorm:"timestamp date3" json:"date3"`

	Time1 *carbon.Time `xorm:"varchar(50) time1" json:"time1"`
	Time2 *carbon.Time `xorm:"time time2" json:"time2"`
	Time3 *carbon.Time `xorm:"timestamp time3" json:"time3"`

	DateTime1 *carbon.DateTime `xorm:"varchar(50) date_time1" json:"date_time1"`
	DateTime2 *carbon.DateTime `xorm:"datetime date_time2" json:"date_time2"`
	DateTime3 *carbon.DateTime `xorm:"timestamp date_time3" json:"date_time3"`

	RFC3339Layout1 *carbon.LayoutType[RFC3339Layout] `xorm:"varchar(50) rfc3339_layout1" json:"rfc3339_layout1"`
	RFC3339Layout2 *carbon.LayoutType[RFC3339Layout] `xorm:"datetime rfc3339_layout2" json:"rfc3339_layout2"`
	RFC3339Layout3 *carbon.LayoutType[RFC3339Layout] `xorm:"timestamp rfc3339_layout3" json:"rfc3339_layout3"`

	ISO8601Format1 *carbon.FormatType[ISO8601Format] `xorm:"varchar(50) iso8601_format1" json:"iso8601_format1"`
	ISO8601Format2 *carbon.FormatType[ISO8601Format] `xorm:"datetime iso8601_format2" json:"iso8601_format2"`
	ISO8601Format3 *carbon.FormatType[ISO8601Format] `xorm:"timestamp iso8601_format3" json:"iso8601_format3"`

	Timestamp1 *carbon.Timestamp `xorm:"bigint timestamp1" json:"timestamp1"`

	CreatedAt *carbon.DateTime `xorm:"created created_at" json:"-"`
	UpdatedAt *carbon.DateTime `xorm:"updated updated_at" json:"-"`
	DeletedAt *carbon.DateTime `xorm:"deleted deleted_at" json:"-"`
}

func (MySQLModel2) TableName() string {
	return "xorm_mysql2"
}

type PgSQLModel1 struct {
	Id uint64 `json:"-" xorm:"pk autoincr"`

	Carbon1 carbon.Carbon `gorm:"timestamp carbon1" json:"carbon1"`
	Carbon2 carbon.Carbon `gorm:"timestamptz carbon2" json:"carbon2"`

	Date1 carbon.Date `gorm:"date date1" json:"date1"`

	Time1 carbon.Time `gorm:"time time1" json:"time1"`
	Time2 carbon.Time `gorm:"timetz time2" json:"time2"`

	DateTime1 carbon.DateTime `gorm:"timestamp date_time1" json:"date_time1"`
	DateTime2 carbon.DateTime `gorm:"timestamptz date_time2" json:"date_time2"`

	RFC3339Layout1 carbon.LayoutType[RFC3339Layout] `gorm:"timestamp rfc3339_layout1" json:"rfc3339_layout1"`
	RFC3339Layout2 carbon.LayoutType[RFC3339Layout] `gorm:"timestamptz rfc3339_layout2" json:"rfc3339_layout2"`

	ISO8601Format1 carbon.FormatType[ISO8601Format] `gorm:"timestamp iso8601_format1" json:"iso8601_format1"`
	ISO8601Format2 carbon.FormatType[ISO8601Format] `gorm:"timestamptz iso8601_format2" json:"iso8601_format2"`

	Timestamp1 carbon.Timestamp `gorm:"int4 timestamp1" json:"timestamp1"`

	CreatedAt carbon.DateTime `xorm:"created created_at" json:"-"`
	UpdatedAt carbon.DateTime `xorm:"updated updated_at" json:"-"`
	DeletedAt carbon.DateTime `xorm:"deleted deleted_at" json:"-"`
}

func (PgSQLModel1) TableName() string {
	return "xorm_pgsql1"
}

type PgSQLModel2 struct {
	Id uint64 `json:"-" xorm:"pk autoincr"`

	Carbon1 *carbon.Carbon `gorm:"timestamp carbon1" json:"carbon1"`
	Carbon2 *carbon.Carbon `gorm:"timestamptz carbon2" json:"carbon2"`

	Date1 *carbon.Date `gorm:"date date1" json:"date1"`

	Time1 *carbon.Time `gorm:"time time1" json:"time1"`
	Time2 *carbon.Time `gorm:"timetz time2" json:"time2"`

	DateTime1 *carbon.DateTime `gorm:"timestamp date_time1" json:"date_time1"`
	DateTime2 *carbon.DateTime `gorm:"timestamptz date_time2" json:"date_time2"`

	RFC3339Layout1 *carbon.LayoutType[RFC3339Layout] `gorm:"timestamp rfc3339_layout1" json:"rfc3339_layout1"`
	RFC3339Layout2 *carbon.LayoutType[RFC3339Layout] `gorm:"timestamptz rfc3339_layout2" json:"rfc3339_layout2"`

	ISO8601Format1 *carbon.FormatType[ISO8601Format] `gorm:"timestamp iso8601_format1" json:"iso8601_format1"`
	ISO8601Format2 *carbon.FormatType[ISO8601Format] `gorm:"timestamptz iso8601_format2" json:"iso8601_format2"`

	Timestamp1 *carbon.Timestamp `gorm:"int4 timestamp1" json:"timestamp1"`

	CreatedAt *carbon.DateTime `xorm:"created created_at" json:"-"`
	UpdatedAt *carbon.DateTime `xorm:"updated updated_at" json:"-"`
	DeletedAt *carbon.DateTime `xorm:"deleted deleted_at" json:"-"`
}

func (PgSQLModel2) TableName() string {
	return "xorm_pgsql2"
}

type SQLiteModel1 struct {
	Id uint64 `json:"-" xorm:"pk autoincr"`

	Carbon carbon.Carbon `xorm:"text carbon" json:"carbon"`

	Date carbon.Date `xorm:"text date" json:"date"`

	Time carbon.Time `xorm:"text time" json:"time"`

	DateTime carbon.DateTime `xorm:"text date_time" json:"date_time"`

	RFC3339Layout carbon.LayoutType[RFC3339Layout] `xorm:"text rfc3339_layout" json:"rfc3339_layout"`
	ISO8601Format carbon.FormatType[ISO8601Format] `xorm:"text iso8601_format" json:"iso8601_format"`

	Timestamp carbon.Timestamp `xorm:"integer timestamp" json:"timestamp"`

	CreatedAt carbon.DateTime `xorm:"created created_at" json:"-"`
	UpdatedAt carbon.DateTime `xorm:"updated updated_at" json:"-"`
	DeletedAt carbon.DateTime `xorm:"deleted deleted_at" json:"-"`
}

func (SQLiteModel1) TableName() string {
	return "xorm_sqlite1"
}

type SQLiteModel2 struct {
	Id uint64 `json:"-" xorm:"pk autoincr"`

	Carbon *carbon.Carbon `xorm:"text carbon" json:"carbon"`

	Date *carbon.Date `xorm:"text date" json:"date"`

	Time *carbon.Time `xorm:"text time" json:"time"`

	DateTime *carbon.DateTime `xorm:"text date_time" json:"date_time"`

	RFC3339Layout *carbon.LayoutType[RFC3339Layout] `xorm:"text rfc3339_layout" json:"rfc3339_layout"`
	ISO8601Format *carbon.FormatType[ISO8601Format] `xorm:"text iso8601_format" json:"iso8601_format"`

	Timestamp *carbon.Timestamp `xorm:"integer timestamp" json:"timestamp"`

	CreatedAt *carbon.DateTime `xorm:"created created_at" json:"-"`
	UpdatedAt *carbon.DateTime `xorm:"updated updated_at" json:"-"`
	DeletedAt *carbon.DateTime `xorm:"deleted deleted_at" json:"-"`
}

func (SQLiteModel2) TableName() string {
	return "xorm_sqlite2"
}

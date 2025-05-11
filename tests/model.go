package tests

import (
	"github.com/dromara/carbon/v2"
	"gorm.io/gorm"
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
	ID uint64 `json:"-" gorm:"column:id;primaryKey"`

	Carbon1 carbon.Carbon `gorm:"column:carbon1;type:varchar(50);" json:"carbon1"`
	Carbon2 carbon.Carbon `gorm:"column:carbon2;type:datetime;" json:"carbon2"`
	Carbon3 carbon.Carbon `gorm:"column:carbon3;type:timestamp;" json:"carbon3"`

	Date1 carbon.Date `gorm:"column:date1;type:varchar(50);" json:"date1"`
	Date2 carbon.Date `gorm:"column:date2;type:date;" json:"date2"`
	Date3 carbon.Date `gorm:"column:date3;type:timestamp;" json:"date3"`

	Time1 carbon.Time `gorm:"column:time1;type:varchar(50);" json:"time1"`
	Time2 carbon.Time `gorm:"column:time2;type:time;" json:"time2"`
	Time3 carbon.Time `gorm:"column:time3;type:timestamp;" json:"time3"`

	DateTime1 carbon.DateTime `gorm:"column:date_time1;type:varchar(50);" json:"date_time1"`
	DateTime2 carbon.DateTime `gorm:"column:date_time2;type:datetime;" json:"date_time2"`
	DateTime3 carbon.DateTime `gorm:"column:date_time3;type:timestamp;" json:"date_time3"`

	RFC3339Layout1 carbon.LayoutType[RFC3339Layout] `gorm:"column:rfc3339_layout1;type:varchar(50);" json:"rfc3339_layout1"`
	RFC3339Layout2 carbon.LayoutType[RFC3339Layout] `gorm:"column:rfc3339_layout2;type:datetime;" json:"rfc3339_layout2"`
	RFC3339Layout3 carbon.LayoutType[RFC3339Layout] `gorm:"column:rfc3339_layout3;type:timestamp;" json:"rfc3339_layout3"`

	ISO8601Format1 carbon.FormatType[ISO8601Format] `gorm:"column:iso8601_format1;type:varchar(50);" json:"iso8601_format1"`
	ISO8601Format2 carbon.FormatType[ISO8601Format] `gorm:"column:iso8601_format2;type:datetime;" json:"iso8601_format2"`
	ISO8601Format3 carbon.FormatType[ISO8601Format] `gorm:"column:iso8601_format3;type:timestamp;" json:"iso8601_format3"`

	Timestamp1 carbon.Timestamp `gorm:"column:timestamp1;type:bigint;" json:"timestamp1"`

	CreatedAt carbon.DateTime `gorm:"column:created_at;type:timestamp;" json:"-"`
	UpdatedAt carbon.DateTime `gorm:"column:updated_at;type:timestamp;" json:"-"`
	DeletedAt gorm.DeletedAt  `gorm:"column:deleted_at;type:datetime;" json:"-"`
}

func (MySQLModel1) TableName() string {
	return "carbon_mysql1"
}

type MySQLModel2 struct {
	ID uint64 `json:"-" gorm:"column:id;primaryKey"`

	Carbon1 *carbon.Carbon `gorm:"column:carbon1;type:varchar(50);" json:"carbon1"`
	Carbon2 *carbon.Carbon `gorm:"column:carbon2;type:datetime;" json:"carbon2"`
	Carbon3 *carbon.Carbon `gorm:"column:carbon3;type:timestamp;" json:"carbon3"`

	Date1 *carbon.Date `gorm:"column:date1;type:varchar(50);" json:"date1"`
	Date2 *carbon.Date `gorm:"column:date2;type:date;" json:"date2"`
	Date3 *carbon.Date `gorm:"column:date3;type:timestamp;" json:"date3"`

	Time1 *carbon.Time `gorm:"column:time1;type:varchar(50);" json:"time1"`
	Time2 *carbon.Time `gorm:"column:time2;type:time;" json:"time2"`
	Time3 *carbon.Time `gorm:"column:time3;type:timestamp;" json:"time3"`

	DateTime1 *carbon.DateTime `gorm:"column:date_time1;type:varchar(50);" json:"date_time1"`
	DateTime2 *carbon.DateTime `gorm:"column:date_time2;type:datetime;" json:"date_time2"`
	DateTime3 *carbon.DateTime `gorm:"column:date_time3;type:timestamp;" json:"date_time3"`

	RFC3339Layout1 *carbon.LayoutType[RFC3339Layout] `gorm:"column:rfc3339_layout1;type:varchar(50);" json:"rfc3339_layout1"`
	RFC3339Layout2 *carbon.LayoutType[RFC3339Layout] `gorm:"column:rfc3339_layout2;type:datetime;" json:"rfc3339_layout2"`
	RFC3339Layout3 *carbon.LayoutType[RFC3339Layout] `gorm:"column:rfc3339_layout3;type:timestamp;" json:"rfc3339_layout3"`

	ISO8601Format1 *carbon.FormatType[ISO8601Format] `gorm:"column:iso8601_format1;type:varchar(50);" json:"iso8601_format1"`
	ISO8601Format2 *carbon.FormatType[ISO8601Format] `gorm:"column:iso8601_format2;type:datetime;" json:"iso8601_format2"`
	ISO8601Format3 *carbon.FormatType[ISO8601Format] `gorm:"column:iso8601_format3;type:timestamp;" json:"iso8601_format3"`

	Timestamp1 *carbon.Timestamp `gorm:"column:timestamp1;type:bigint;" json:"timestamp1"`

	CreatedAt *carbon.DateTime `gorm:"column:created_at;type:timestamp;" json:"-"`
	UpdatedAt *carbon.DateTime `gorm:"column:updated_at;type:timestamp;" json:"-"`
	DeletedAt *gorm.DeletedAt  `gorm:"column:deleted_at;type:datetime;" json:"-"`
}

func (MySQLModel2) TableName() string {
	return "carbon_mysql2"
}

type PgSQLModel1 struct {
	ID uint64 `json:"-" gorm:"column:id;primaryKey"`

	Carbon1 carbon.Carbon `gorm:"column:carbon1;type:timestamp without time zone;" json:"carbon1"`
	Carbon2 carbon.Carbon `gorm:"column:carbon2;type:timestamp with time zone;" json:"carbon2"`

	Date1 carbon.Date `gorm:"column:date1;type:date;" json:"date1"`

	Time1 carbon.Time `gorm:"column:time1;type:time without time zone;" json:"time1"`
	Time2 carbon.Time `gorm:"column:time2;type:time with time zone;" json:"time2"`

	DateTime1 carbon.DateTime `gorm:"column:date_time1;type:timestamp without time zone;" json:"date_time1"`
	DateTime2 carbon.DateTime `gorm:"column:date_time2;type:timestamp with time zone;" json:"date_time2"`

	RFC3339Layout1 carbon.LayoutType[RFC3339Layout] `gorm:"column:rfc3339_layout1;type:timestamp without time zone;" json:"rfc3339_layout1"`
	RFC3339Layout2 carbon.LayoutType[RFC3339Layout] `gorm:"column:rfc3339_layout2;type:timestamp with time zone;" json:"rfc3339_layout2"`

	ISO8601Format1 carbon.FormatType[ISO8601Format] `gorm:"column:iso8601_format1;type:timestamp without time zone;" json:"iso8601_format1"`
	ISO8601Format2 carbon.FormatType[ISO8601Format] `gorm:"column:iso8601_format2;type:timestamp with time zone;" json:"iso8601_format2"`

	Timestamp1 carbon.Timestamp `gorm:"column:timestamp1;type:int4;" json:"timestamp1"`

	CreatedAt carbon.DateTime `gorm:"column:created_at;type:timestamp with time zone;" json:"-"`
	UpdatedAt carbon.DateTime `gorm:"column:updated_at;type:timestamp with time zone;" json:"-"`
	DeletedAt gorm.DeletedAt  `gorm:"column:deleted_at;type:timestamp with time zone;" json:"-"`
}

func (PgSQLModel1) TableName() string {
	return "carbon_pgsql1"
}

type PgSQLModel2 struct {
	ID uint64 `json:"-" gorm:"column:id;primaryKey"`

	Carbon1 *carbon.Carbon `gorm:"column:carbon1;type:timestamp without time zone;" json:"carbon1"`
	Carbon2 *carbon.Carbon `gorm:"column:carbon2;type:timestamp with time zone;" json:"carbon2"`

	Date1 *carbon.Date `gorm:"column:date1;type:date;" json:"date1"`

	Time1 *carbon.Time `gorm:"column:time1;type:time without time zone;" json:"time1"`
	Time2 *carbon.Time `gorm:"column:time2;type:time with time zone;" json:"time2"`

	DateTime1 *carbon.DateTime `gorm:"column:date_time1;type:timestamp without time zone;" json:"date_time1"`
	DateTime2 *carbon.DateTime `gorm:"column:date_time2;type:timestamp with time zone;" json:"date_time2"`

	RFC3339Layout1 *carbon.LayoutType[RFC3339Layout] `gorm:"column:rfc3339_layout1;type:timestamp without time zone;" json:"rfc3339_layout1"`
	RFC3339Layout2 *carbon.LayoutType[RFC3339Layout] `gorm:"column:rfc3339_layout2;type:timestamp with time zone;" json:"rfc3339_layout2"`

	ISO8601Format1 *carbon.FormatType[ISO8601Format] `gorm:"column:iso8601_format1;type:timestamp without time zone;" json:"iso8601_format1"`
	ISO8601Format2 *carbon.FormatType[ISO8601Format] `gorm:"column:iso8601_format2;type:timestamp with time zone;" json:"iso8601_format2"`

	Timestamp1 *carbon.Timestamp `gorm:"column:timestamp1;type:int4;" json:"timestamp1"`

	CreatedAt *carbon.DateTime `gorm:"column:created_at;type:timestamp with time zone;" json:"-"`
	UpdatedAt *carbon.DateTime `gorm:"column:updated_at;type:timestamp with time zone;" json:"-"`
	DeletedAt *gorm.DeletedAt  `gorm:"column:deleted_at;type:timestamp with time zone;" json:"-"`
}

func (PgSQLModel2) TableName() string {
	return "carbon_pgsql2"
}

type SQLiteModel1 struct {
	ID uint64 `json:"-" gorm:"column:id;primaryKey"`

	Carbon carbon.Carbon `gorm:"column:carbon;type:text;" json:"carbon"`

	Date carbon.Date `gorm:"column:date;type:text;" json:"date"`

	Time carbon.Time `gorm:"column:time;type:text;" json:"time"`

	DateTime carbon.DateTime `gorm:"column:date_time;type:text;" json:"date_time"`

	RFC3339Layout carbon.LayoutType[RFC3339Layout] `gorm:"column:rfc3339_layout1;type:text;" json:"rfc3339_layout"`
	ISO8601Format carbon.FormatType[ISO8601Format] `gorm:"column:iso8601_format1;type:text;" json:"iso8601_format"`

	Timestamp carbon.Timestamp `gorm:"column:timestamp;type:integer;" json:"timestamp"`

	CreatedAt carbon.DateTime `gorm:"column:created_at;type:timestamp with time zone;" json:"-"`
	UpdatedAt carbon.DateTime `gorm:"column:updated_at;type:timestamp with time zone;" json:"-"`
	DeletedAt gorm.DeletedAt  `gorm:"column:deleted_at;type:timestamp with time zone;" json:"-"`
}

func (SQLiteModel1) TableName() string {
	return "carbon_sqlite1"
}

type SQLiteModel2 struct {
	ID uint64 `json:"-" gorm:"column:id;primaryKey"`

	Carbon *carbon.Carbon `gorm:"column:carbon;type:text;" json:"carbon"`

	Date *carbon.Date `gorm:"column:date;type:text;" json:"date"`

	Time *carbon.Time `gorm:"column:time;type:text;" json:"time"`

	DateTime *carbon.DateTime `gorm:"column:date_time;type:text;" json:"date_time"`

	RFC3339Layout *carbon.LayoutType[RFC3339Layout] `gorm:"column:rfc3339_layout1;type:text;" json:"rfc3339_layout"`
	ISO8601Format *carbon.FormatType[ISO8601Format] `gorm:"column:iso8601_format1;type:text;" json:"iso8601_format"`

	Timestamp *carbon.Timestamp `gorm:"column:timestamp;type:integer;" json:"timestamp"`

	CreatedAt *carbon.DateTime `gorm:"column:created_at;type:timestamp with time zone;" json:"-"`
	UpdatedAt *carbon.DateTime `gorm:"column:updated_at;type:timestamp with time zone;" json:"-"`
	DeletedAt *gorm.DeletedAt  `gorm:"column:deleted_at;type:timestamp with time zone;" json:"-"`
}

func (SQLiteModel2) TableName() string {
	return "carbon_sqlite2"
}

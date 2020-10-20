package models

import (
	"BeegoDemo/db_mysql"
	"BeegoDemo/util"
)

/**
 * 上传文件记录的   结构体定义
 */
type UploadRecord struct {
	Id        		 int
	FileName 		 string
	FileSize 		 int64
	FileCert 		 string //认证号
	FileTitle		 string
	CertTime 		 int64
	FormatCertTime   string //格式化时间，仅在前段展示页面使用
	Phone    		 string //对应的用户的phone
}

/**
 * 保存上传记录到数据库中
 */
func (u UploadRecord) SaveRecord() (int64, error) {
	rs, err := db_mysql.Db.Exec("insert into upload_record(file_name, file_size, file_cert, file_title, cert_time, phone)" +
		" values(?,?,?,?,?,?)",
		u.FileName,
		u.FileSize,
		u.FileCert,
		u.FileTitle,
		u.CertTime,
		u.Phone)

	if err != nil {
		return -1,err
	}
	id, err := rs.RowsAffected()
	if err != nil {
		return -1, err
	}
	return id,nil
}

/**
 * 读取数据库中phone用户对应的所有认证数据
 */
func QueryRecordBtPhone(phone string)([]UploadRecord, error){
	rs, err := db_mysql.Db.Query(" select id, file_name, file_size, file_cert, file_title, cert_time, phone from upload_record where phone = ?", phone)
	if err != nil{
		return nil, err
	}
	records := make([]UploadRecord,0)
	for rs.Next()  {
		var record UploadRecord
		err := rs.Scan(&record.Id,&record.FileName,&record.FileSize,&record.FileCert,&record.FileTitle,&record.CertTime,&record.Phone)
		if err != nil{
			return nil,err
		}
		//时间转换 record.CertTime
		record.FormatCertTime = util.TimeFormat(record.CertTime,0,util.TIME_FORMAT_THREE)
		records = append(records, record)
	}
	return records,nil
}
package controllers

import (
	"BeegoDemo/models"
	"BeegoDemo/util"
	"bufio"
	"fmt"
	"github.com/astaxie/beego"
	"io"
	"os"
	"time"
)

type HomeController struct {
	beego.Controller
}

func (u *HomeController) Get(){
	phone := u.GetString("phone")
	u.Data["Phone"] = phone
	u.TplName = "home.html"
}

func (r *HomeController) Post(){

	fileTitle := r.Ctx.Request.PostFormValue("upload_title")
	phone := r.Ctx.Request.PostFormValue("phone")
	f, h, err := r.GetFile("myfile")//获取上传文件
	if err != nil {
		r.Ctx.WriteString("抱歉，用户文件解析失败，请重试")
		return
	}
	//ext := path.Ext(h.Filename)
	//验证后缀名是否符合要求
	//var AllowExtMap map[string]bool = map[string]bool{
	//	".jpg":true,
	//	".jpeg":true,
	//	".png":true,
	//	".bat":true,
	//}
	//if _,ok :=AllowExtMap[ext];!ok{
	//	r.Ctx.WriteString("后缀名不符合上传要求")
	//	return
	//}
	//2、关闭文件
	defer f.Close()

	fmt.Println("自定义的文件标题：",fileTitle)
	fmt.Println("文件名称：",h.Filename)
	fmt.Println("文件的大小：",h.Size)
	//创建目录
	uploadDir := "./static/img/" + h.Filename
	//文件权限：a+b+c
	//a：文件所有者拥有的权限，读4、写2、执行1。
	//b：文件所有者所在的组的用户对文件拥有的权限，读4、写2、执行1
	//c：其他用户对文件拥有的权限，读4、写2、执行1
	//eg：某个文件m，其权限是985（错误）

	saveFile, err := os.OpenFile(uploadDir,os.O_RDWR | os.O_CREATE,777)
	dsWriter := bufio.NewWriter(saveFile)
	_, err = io.Copy(dsWriter,f)
	//err := os.MkdirAll(uploadDir, 777)
	if err != nil{
		fmt.Println(err.Error())
		r.Ctx.WriteString("抱歉，保存电子数据失败，请重试" )
		return
	}
	defer saveFile.Close()

	//计算文件的hash值
	hashFile, err := os.Open(uploadDir)
	defer hashFile.Close()
	hash, err := util.MD5HashFile(hashFile)
	if err != nil {
		return
	}

	record := models.UploadRecord{}
	record.FileName = h.Filename
	record.FileSize = h.Size
	record.FileTitle = fileTitle
	record.CertTime = time.Now().Unix()
	record.FileCert = hash
	record.Phone = phone
	_ , err = record.SaveRecord()
	if err != nil{
		r.Ctx.WriteString("抱歉，数据认证错误，请重试")
		return
	}
	//4、从数据库中读取phone用户对应的所有认证数据记录
	records, err := models.QueryRecordBtPhone(phone)

	//5、根据文件保存结果，返回相应的提示信息或者页面跳转
	if err != nil{
		r.Ctx.WriteString("抱歉，获取认证数据失败，请重试")
		return
	}
	fmt.Println(records)
	r.Data["Records"] = records
	r.Data["Phone"] = phone
	r.TplName = "list_record.html"

	//构造文件名称
	//rand.Seed(time.Now().UnixNano())
	//randNum := fmt.Sprintf("%d", rand.Intn(9999)+1000 )
	//hashName := md5.Sum( []byte( time.Now().Format("2006_01_02_15_04_05_") + randNum))
	////
	//fileName := fmt.Sprintf("%x",hashName) + ext
	//fpath := uploadDir + fileName
	//defer f.Close()//关闭上传的文件，不然的话会出现临时文件不能清除的情况
	//err = r.SaveToFile("myfile", fpath)
	//if err != nil {
	//	r.Ctx.WriteString(fmt.Sprintf("%v",err))
	//}
	//r.Ctx.WriteString("上传成功！")
}

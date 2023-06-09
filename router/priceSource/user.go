package priceSource

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"yy-price-api/errors"
	"yy-price-api/proxy"
)

type userRouter struct {
}

var User = &userRouter{}

// 查询一个用户
func (p *userRouter) Get(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusOK, errors.ApiErrGet.SetErrMsg("务必传入参数!"))
		return
	}
	data, err := proxy.User.Get(id)
	if err != nil {
		c.JSON(http.StatusOK, errors.ApiErrSql.SetErrMsg("没有数据!"))
		return
	}
	c.JSON(http.StatusOK, errors.ApiErrSuccess.SetDate(data))
	return
}
//
//func (p *AdminRouter) Update(c *gin.Context) {
//
//	var adminData modelBack.Admin
//
//	err := c.ShouldBindWith(&adminData, binding.JSON)
//	if err != nil || adminData.Uuid == "" {
//		c.JSON(http.StatusOK, errors.ApiErrPostForm)
//		return
//	}
//	//判断密码的位数
//	if l := len(adminData.Password); l < 6 || l > 20 {
//		c.JSON(http.StatusOK, errors.ApiErrPostForm.SetErrMsg("密码必须在6-20位"))
//		return
//	}
//	//判断姓名的位数
//	if l := len(adminData.Name); l < 4 || l > 30 {
//		c.JSON(http.StatusOK, errors.ApiErrPostForm.SetErrMsg("姓名必须在2-15位"))
//		return
//	}
//	adminData.Password = common.GetMD5Hash(adminData.Password)
//
//	err, adminData = proxyCenter.AdminUser.Update(&adminData)
//	if err != nil {
//		c.JSON(http.StatusOK, errors.ApiErrSql.SetDate(err.Error()))
//		return
//	}
//	adminData.Password = "为了用户安全，密码已自动隐藏"
//	c.JSON(http.StatusOK, errors.ApiErrSuccess.SetDate(adminData))
//	return
//}


//func Login(c *gin.Context) {
//	// var user models.User
//	// c.BindJSON()
//
//	var user models.User
//	err := c.BindJSON(&user)
//	if err != nil {
//		fmt.Println(err)
//		c.String(400, err.Error())
//		return
//	}
//	data, err := proxy.User.GetOne(bson.M{
//		"email":     user.Email,
//		"password": common.GetMD5Hash(user.Password),
//	})
//	if err != nil {
//		ApiStandardError := errors.ApiErrNamePwdIncorrect
//		c.JSON(400, ApiStandardError)
//		return
//	}
//	user = data.(models.User)
//
//
//	// 校验账号状态
//	if user.Status == common.Linactive || user.Status == common.Ldisable {
//		c.String(403, "账户禁止登陆")
//		return
//	}
//
//	session := sessions.Default(c)
//	// 设置一下session,然后返回给浏览器
//	session.Set("user", user)
//	// 保存这个session,登录
//	err = session.Save()
//	if err != nil {
//		c.String(400, err.Error())
//		return
//	}
//	// 返回的session
//	resSe := session.Get("user")
//
//	c.JSON(200, map[string]interface{}{
//		"session":  resSe ,
//		"errno":0,
//	})
//}
//
//func Logout(c *gin.Context) {
//	session := sessions.Default(c)
//	session.Clear()
//	err := session.Save()
//	if err != nil {
//		c.String(400, err.Error())
//		return
//	}
//}
//
//func User(c *gin.Context) {
//	user := c.MustGet("user").(models.User)
//	c.JSON(200, user)
//}
//
//func Regist(c *gin.Context) {
//
//	var user models.User
//	err := c.BindJSON(&user)
//	if err != nil {
//		c.String(400, err.Error())
//		return
//	}
//	fmt.Println(user)
//
//	// user.Status = common.Linactive
//	// 测试用的 亲
//	user.Status = common.Lnormal
//
//
//	user.Password = common.GetMD5Hash(user.Password)
//	err = proxy.User.Create(&user)
//
//	if err != nil {
//		c.String(400, err.Error())
//		return
//	}
//	fmt.Println(user)
//	go sendActiveEmail(user.ID.Hex(), user.Email)
//
//	c.JSON(http.StatusOK, errors.ApiErrSuccess)
//}
//
//func sendActiveEmail(id, to string) error {
//
//	url := os.Getenv("HOST") + ":" + os.Getenv("PORT") + "/" + "active"
//	host := os.Getenv("MAIL_HOST")
//	port, err := strconv.Atoi(os.Getenv("MAIL_PORT"))
//	if err != nil {
//		log.Println(err)
//		return err
//	}
//	name := os.Getenv("MAIL_USERNAME")
//	pasd := os.Getenv("MAIL_PASSWORD")
//
//	expire := time.Now().Add(86400 * time.Second).Unix()
//	args := models.ActiveInfo{Id: id, Expire: expire}
//	encodeByte, err := json.Marshal(args)
//	if err != nil {
//		log.Println(err)
//		return err
//	}
//	encryptByte, err := common.AesEncrypt(string(encodeByte))
//	baseEncrypt := base64.URLEncoding.EncodeToString(encryptByte)
//	tmpl, err := template.ParseFiles("templates/user/email_active.html")
//	if err != nil {
//		log.Println("Error happened..")
//		return err
//	}
//	activeUrlStruct := models.ActiveUrl{Url: url, Info: string(baseEncrypt)}
//	var b bytes.Buffer
//	tmpl.Execute(&b, activeUrlStruct)
//	m := gomail.NewMessage()
//	m.SetHeader("From", name) //发件人
//	m.SetHeader("To", to)     //收件人
//	m.SetAddressHeader("Cc", "15398381714@163.com", "goulang")
//	m.SetHeader("Subject", "够浪社区邮箱激活认证")
//	//TODO temple完成
//	m.SetBody("text/html", b.String())
//	d := gomail.NewDialer(host, port, name, pasd)
//
//	if err := d.DialAndSend(m); err != nil {
//		//TODO 完成失败提示
//		log.Println(err)
//		return err
//	}
//
//	return nil
//}
//
//// DeleteUsers delete a user
//func DeleteUser(c *gin.Context) {
//	userID := c.Param("userID")
//	err := proxy.User.Delete(userID)
//	if err != nil {
//		c.String(400, err.Error())
//		return
//	}
//}
//
//// GetUsers get all user
//func Users(c *gin.Context) {
//	data, err := proxy.User.GetMany(nil, 1, 10)
//	if err != nil {
//		c.String(400, err.Error())
//		return
//	}
//	users := data.([]models.User)
//	c.JSON(200, users)
//}
//
//// GetUser get a user
//func GetUser(c *gin.Context) {
//	userID := c.Param("userID")
//	data, err := proxy.User.Get(userID)
//	if err != nil {
//		c.String(400, err.Error())
//		return
//	}
//	user := data.(models.User)
//	c.JSON(200, user)
//}
//
//func Passwd(c *gin.Context) {
//	var user models.User
//	var password models.Password
//	userID := c.Param("userID")
//	err := c.BindJSON(&password)
//	if err != nil {
//		c.JSON(http.StatusBadRequest, errors.NewUnknownErr(err))
//		return
//	}
//	data, err := proxy.User.Get(userID)
//	if err != nil {
//		c.JSON(http.StatusBadRequest, errors.NewUnknownErr(err))
//		return
//	}
//	user = data.(models.User)
//
//	if !common.CheckPasswrd(password.Password, user.Password) {
//		c.JSON(http.StatusBadRequest, errors.ApiErrPwdIncorrect)
//		return
//	}
//
//	user.Password = common.GetMD5Hash(password.RePassword)
//	if err := proxy.User.Update(userID, &user); err != nil {
//		c.JSON(http.StatusBadRequest, errors.NewUnknownErr(err))
//		return
//	}
//
//	c.JSON(http.StatusOK, errors.ApiErrSuccess)
//}
//
//func Active(c *gin.Context) {
//	var activeInfo models.ActiveInfo
//	var active models.Active
//	base := c.Param("active")
//	encryptStr, err := base64.URLEncoding.DecodeString(base)
//	if err != nil {
//		c.JSON(http.StatusBadRequest, errors.NewUnknownErr(err))
//		return
//	}
//	info, err := common.AesDecrypt(encryptStr)
//	if err != nil {
//		c.JSON(http.StatusBadRequest, errors.NewUnknownErr(err))
//		return
//	}
//	if err := json.Unmarshal([]byte(info), &activeInfo); err != nil {
//		c.JSON(http.StatusBadRequest, errors.NewUnknownErr(err))
//		return
//	}
//	if activeInfo.Expire <= time.Now().Unix() {
//		c.JSON(http.StatusBadRequest, errors.ApiErrActiveInvalid)
//		return
//	}
//	if _, err := proxy.User.Get(activeInfo.Id); err != nil {
//		c.JSON(http.StatusBadRequest, errors.NewUnknownErr(err))
//		return
//	}
//	active.Status = common.Lnormal
//	if err := proxy.User.Update(activeInfo.Id, &active); err != nil {
//		c.JSON(http.StatusBadRequest, errors.NewUnknownErr(err))
//		return
//	}
//
//	home := os.Getenv("HOST") + ":" + os.Getenv("PORT")
//	c.Redirect(http.StatusFound, home)
//}
//
//func UpdateProfile(c *gin.Context) {
//	userID := c.Param("userID")
//	var update models.Update
//	err := c.BindJSON(&update)
//	if err != nil {
//		log.Println(err)
//		c.JSON(http.StatusBadRequest, errors.NewUnknownErr(err))
//		return
//	}
//
//	if err := proxy.User.Update(userID, &update); err != nil {
//		log.Println(err)
//		c.JSON(http.StatusBadRequest, errors.NewUnknownErr(err))
//		return
//	}
//
//	c.JSON(http.StatusOK, errors.ApiErrSuccess)
//}
//
//func Avatar(c *gin.Context) {
//	var user models.User
//	userID := c.Param("userID")
//	data, err := proxy.User.Get(userID)
//	if err != nil {
//		log.Println(err)
//		c.JSON(http.StatusBadRequest, errors.NewUnknownErr(err))
//		return
//	}
//	user = data.(models.User)
//
//	//删除原有头像
//	ok := Qiniu.Storage.DeleteFile(user.Avatar, true)
//	if ok != nil {
//		log.Println(err)
//		c.JSON(http.StatusBadRequest, errors.NewUnknownErr(err))
//		return
//	}
//
//	//上传新头像
//	file, header, err := c.Request.FormFile("file")
//	if err != nil {
//		log.Println(err)
//		c.JSON(http.StatusBadRequest, errors.NewUnknownErr(err))
//		return
//	}
//	name := common.GetFileUniqueName(header.Filename)
//	// TODO 完成从配置读取路径
//	name = time.Now().Format("avatar/2006/01/02") + "/" + name
//	byteFile, err := ioutil.ReadAll(file)
//	if err != nil {
//		log.Println(err)
//		c.JSON(http.StatusBadRequest, errors.NewUnknownErr(err))
//		return
//	}
//	_, isOk := Qiniu.Storage.PutFile(name, byteFile)
//	if isOk != nil {
//		log.Println(err)
//		c.JSON(http.StatusBadRequest, errors.NewUnknownErr(err))
//		return
//	}
//
//	//更新头像
//	user.Avatar = Qiniu.Storage.GetUrl(name)
//	if err := proxy.User.Update(userID, &user); err != nil {
//		log.Println(err)
//		c.JSON(http.StatusBadRequest, errors.NewUnknownErr(err))
//		return
//	}
//
//	c.JSON(http.StatusOK, gin.H{
//		"key":    name,
//		"access": Qiniu.Storage.GetUrl(name),
//	})
//}

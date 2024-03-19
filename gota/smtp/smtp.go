package smtp

import (
	"crypto/tls"
	"errors"
	"fmt"
	"framework-iris/gota/commons/utils"
	"framework-iris/gota/config"
	"net"
	"net/smtp"
	"strings"
)

// Sendmail 可批量发送，格式：test1;test2;test3
func Sendmail(smtpName, receive, subject string, body string) error {

	var configContent config.SmtpConfig
	for _, smtpConfig := range config.Configs.Smtp {
		if smtpConfig.Name == smtpName {
			configContent = smtpConfig
		}
	}

	if i := utils.DataCheck(configContent.Host, configContent.Port, configContent.Sender, configContent.Password); len(i) > 0 {
		return errors.New("smtp config error")
	}

	host := configContent.Host
	port := configContent.Port
	sender := configContent.Sender // 发送邮箱
	pwd := configContent.Password  // 邮箱密码

	header := make(map[string]string)
	header["From"] = "<" + sender + ">"
	header["To"] = receive
	header["Subject"] = subject
	header["Content-Type"] = "text/html;charset=UTF-8"

	message := ""
	for k, v := range header {
		message += fmt.Sprintf("%s:%s\r\n", k, v)
	}
	message += "\r\n" + body
	auth := smtp.PlainAuth(
		"",
		sender,
		pwd,
		host,
	)
	err := SendMailUsingTLS(
		fmt.Sprintf("%s:%s", host, port),
		auth,
		sender,
		receive,
		[]byte(message),
	)
	if err != nil {
		return err
	} else {
		return nil
	}
}

func SendMailUsingTLS(addr string, auth smtp.Auth, from string, to string, msg []byte) (err error) {
	c, err := Dial(addr)
	if err != nil {
		return err
	}
	defer c.Close()
	if auth != nil {
		if ok, _ := c.Extension("AUTH"); ok {
			if err = c.Auth(auth); err != nil {
				return err
			}
		}
	}
	if err = c.Mail(from); err != nil {
		return err
	}
	tos := strings.Split(to, ";")
	for _, addr := range tos {
		if err = c.Rcpt(addr); err != nil {
			fmt.Print(err)
			return err
		}
	}
	w, err := c.Data()
	if err != nil {
		return err
	}
	_, err = w.Write(msg)
	if err != nil {
		return err
	}
	err = w.Close()
	if err != nil {
		return err
	}
	return c.Quit()
}
func Dial(addr string) (*smtp.Client, error) {
	conn, err := tls.Dial("tcp", addr, nil)
	if err != nil {
		return nil, err
	}
	//分解主机端口字符串
	host, _, _ := net.SplitHostPort(addr)
	return smtp.NewClient(conn, host)
}

const example = `
<!DOCTYPE html>
<html xmlns:th="http://www.thymeleaf">
<head>
    <meta charset="utf-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="description" content="email code">
    <meta name="viewport" content="width=device-width, initial-scale=1">
</head>
<body>
<div style="background-color:#ECECEC; padding: 35px;">
    <table cellpadding="0" align="center"
           style="width: 800px;height: 100%; margin: 0px auto; text-align: left; position: relative; border-top-left-radius: 5px; border-top-right-radius: 5px; border-bottom-right-radius: 5px; border-bottom-left-radius: 5px; font-size: 14px; font-family:微软雅黑, 黑体; line-height: 1.5; box-shadow: rgb(153, 153, 153) 0px 0px 5px; border-collapse: collapse; background-position: initial initial; background-repeat: initial initial;background:#fff;">
        <tbody>
        <tr>
            <th valign="middle"
                style="height: 25px; line-height: 25px; padding: 15px 35px; border-bottom-width: 1px; border-bottom-style: solid; border-bottom-color: RGB(24, 186, 186); background-color: RGB(24, 186, 186); border-top-left-radius: 5px; border-top-right-radius: 5px; border-bottom-right-radius: 0px; border-bottom-left-radius: 0px;">
                <font face="微软雅黑" size="5" style="color: rgb(255, 255, 255); ">gota 账号验证</font>
            </th>
        </tr>
        <tr>
            <td style="word-break:break-all">
                <div style="padding:25px 35px 40px; background-color:#fff;opacity:0.8;">

                    <h2 style="margin: 5px 0px; ">
                        <font color="#333333" style="line-height: 20px; ">
                            <font style="line-height: 22px; " size="4">
                                尊敬的用户：</font>
                        </font>
                    </h2>
                    <!-- 中文 -->
                    <p>您好！感谢您使用gota，您的账号正在进行邮箱验证，验证码为：<font color="#ff8c00" size="4">` + "gota" + `</font>，有效期5分钟，请尽快填写验证码完成验证！</p><br>
  					<br>
                    <div style="width:100%;margin:0 auto;">
                        <div style="padding:10px 10px 0;border-top:1px solid #ccc;color:#747474;margin-bottom:20px;line-height:1.3em;font-size:12px;">
                            <p>gota</p>
                            <p>联系我们：gota</p>
                            <br>
                            <p>此为系统邮件，请勿回复<br>
                            </p>
                            <!--<p>©gota</p>-->
                        </div>
                    </div>
                </div>
            </td>
        </tr>
        </tbody>
    </table>
</div>
</body>
</html>
    `

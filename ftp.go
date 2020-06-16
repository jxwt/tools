package tools

import (
	"github.com/astaxie/beego/logs"
	"github.com/jlaffaye/ftp"
)

// GetFtpClient ftp连接
func GetFtpClient(ip, user, password string) *ftp.ServerConn {
	conn, err := ftp.Connect(ip)
	if err != nil {
		logs.Error(err)
		return nil
	}
	err = conn.Login(user, password)
	if err != nil {
		logs.Error(err)
		return nil
	}
	return conn
}

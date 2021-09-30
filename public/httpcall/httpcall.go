package httpcall

import (
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"strings"
	"time"
)

func init() {
	HttpCC = &ClientCallSer{}
}

type ClientCallSer struct {
	URLS string
}

var HttpCC *ClientCallSer

func (cc *ClientCallSer) ClientCall(sname, req string) string {
	return Clientcall(cc.URLS, sname, req)

}
func Clientcall(surl, sname, req string) string {
	url := surl + "/" + sname
	url = strings.ReplaceAll(url, "//", "/")
	url = strings.ReplaceAll(url, ":/", "//")
	reqstring := req
	bodys := reqstring
	client := &http.Client{
		Transport: &http.Transport{
			Dial: func(netw, addr string) (net.Conn, error) {
				conn, err := net.DialTimeout(netw, addr, time.Second*100) //设置建立连接超时
				if err != nil {
					return nil, err
				}
				conn.SetDeadline(time.Now().Add(time.Second * 100)) //设置发送接受数据超时
				return conn, nil
			},
			ResponseHeaderTimeout: time.Second * 100,
		},
	}
	reqest, err := http.NewRequest("POST", url, strings.NewReader(bodys)) //提交请求;用指定的方法，网址，可选的主体放回一个新的*Request
	reqest.Header.Set("Content-Type", "application/json;charset=UTF-8")
	if err != nil {
		// panic(err)
		return "result:false " + err.Error()
	}
	response, err := client.Do(reqest) //前面预处理一些参数，状态，Do执行发送；处理返回结果;Do:发送请求,
	if err != nil {
		fmt.Println(err)
		return "result:false " + err.Error()
	}
	defer response.Body.Close()
	// stdout := os.Stdout                     //将结果定位到标准输出，也可以直接打印出来，或定位到其他地方进行相应处理
	// _, err = io.Copy(stdout, response.Body) //将第二个参数拷贝到第一个参数，直到第二参数到达EOF或发生错误，返回拷贝的字节和遇到的第一个错误.
	// if err != nil {
	// 	fmt.Println(err)
	// 	return err.Error()
	// }
	// return status
	//获取返回状态码，正常是200
	if strings.Contains(response.Status, "200") {
		fmt.Println("response:\n", req, "http request  status:", response.Status)

	} else {
		fmt.Println("response:\n", req, "http request fail, status:", response.Status)
		return "false"
	}
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("ReadAll", err)
		return "result:false " + err.Error()
	}

	return string(body)
}

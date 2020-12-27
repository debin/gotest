package main

import (
	"fmt"
	"golang.org/x/crypto/ssh"
	"io/ioutil"
	"log"
	"net"
)

func main() {
	//正确密码
	//_, session, err := SSHConnect("ivan", "123456", "qq10086.zhidaohu.com", 10086)
	//err为空
	//fmt.Println(session, err)

	//key方式
	_, session, err := SSHConnectKey("debin","", "qq10086.zhidaohu.com", 10086)
	//err为空
	fmt.Println(session, err)

	if err != nil {
		log.Fatal("创建ssh session 失败",err)
	} else {
		defer session.Close()
		//执行远程命令
		combo,err := session.CombinedOutput("whoami; cd /; ls -al;echo hello world")
		if err != nil {
			log.Fatal("远程执行cmd 失败",err)
		}
		log.Println("命令输出:",string(combo))
	}








}


/*获取客户端连接*/
func SSHConnect(user, password, host string, port int) (*ssh.Client, *ssh.Session, error) {

	var (
		authMethods  []ssh.AuthMethod
		addr         string
		clientConfig *ssh.ClientConfig
		client       *ssh.Client
		session      *ssh.Session
		err          error
	)

	// 创建密码校验方法
	authMethods = make([]ssh.AuthMethod, 0)
	authMethods = append(authMethods, ssh.Password(password))

	// 创建一个格式合法的回调函数
	hostKeyCallbk := func(hostname string, remote net.Addr, key ssh.PublicKey) error {
		return nil
	}

	/*创建SSH客户端配置*/
	clientConfig = &ssh.ClientConfig{
		User: user,
		Auth: authMethods,
		// Timeout:             30 * time.Second,
		HostKeyCallback: hostKeyCallbk,
	}

	// 连接地址
	addr = fmt.Sprintf("%s:%d", host, port)

	// 拨号并获取SSH客户端
	if client, err = ssh.Dial("tcp", addr, clientConfig); err != nil {
		return nil, nil, err
	}

	// 创建新的会话
	if session, err = client.NewSession(); err != nil {
		return nil, nil, err
	}

	return client, session, nil
}

/*获取客户端连接 key方式*/
func SSHConnectKey(user, password, host string, port int) (*ssh.Client, *ssh.Session, error) {

	var (
		authMethods  []ssh.AuthMethod
		addr         string
		clientConfig *ssh.ClientConfig
		client       *ssh.Client
		session      *ssh.Session
		err          error
	)

	// 创建密码校验方法

	key, err := ioutil.ReadFile("C:/Users/blueblue/.ssh/id_rsa")
	if err != nil {
		log.Fatalf("unable to read private key: %v", err)
	}

	// Create the Signer for this private key.
	var signer ssh.Signer
	if password == "" {
		signer, err = ssh.ParsePrivateKey(key)
	} else {
		signer, err = ssh.ParsePrivateKeyWithPassphrase(key, []byte(password))
	}

	if err != nil {
		log.Fatalf("unable to parse private key: %v", err)
	}

	authMethods = make([]ssh.AuthMethod, 0)
	keys := ssh.PublicKeys(signer)

	authMethods = append(authMethods,keys)

	// 创建一个格式合法的回调函数
	hostKeyCallbk := func(hostname string, remote net.Addr, key ssh.PublicKey) error {
		return nil
	}

	/*创建SSH客户端配置*/
	clientConfig = &ssh.ClientConfig{
		User: user,
		Auth: authMethods,
		// Timeout:             30 * time.Second,
		HostKeyCallback: hostKeyCallbk,
	}

	// 连接地址
	addr = fmt.Sprintf("%s:%d", host, port)

	// 拨号并获取SSH客户端
	if client, err = ssh.Dial("tcp", addr, clientConfig); err != nil {
		return nil, nil, err
	}

	// 创建新的会话
	if session, err = client.NewSession(); err != nil {
		return nil, nil, err
	}

	return client, session, nil
}
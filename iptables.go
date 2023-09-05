package tcpip
import (
	"errors"
	"log"
	"os/exec"
)

type Nat struct {
	Cidr       string // 内网地址
	DeviceName string // 出口网卡
}

func CreateNat(cidr string, deviceName string) (*Nat, error) {
	if len(cidr) == 0 {
		return nil, errors.New("cidr is nil")
	}
	if len(deviceName) == 0 {
		return nil, errors.New("deviceName is nil")
	}
	re := Nat{
		Cidr:       cidr,
		DeviceName: deviceName,
	}
	// TODO 查看日志是否有脏nat 如果有先删除脏nat
	// TODO 写入日志
	cmd := exec.Command("iptables", "-t", "nat", "-A", "POSTROUTING", "-s", re.Cidr, "-o", re.DeviceName, "-j", "MASQUERADE")
	log.Println(cmd)
	err := cmd.Run()
	if err != nil {
		// TODO 如果开启失败 就panic 并且 删除日志
		return nil, err
	}
	return &re, nil
}

func (nat *Nat) Close() error {
	err := exec.Command("iptables", "-t", "nat", "-D", "POSTROUTING", "-s", nat.Cidr, "-o", nat.DeviceName, "-j", "MASQUERADE").Run()
	// TODO删除成功就删除日志
	return err
}

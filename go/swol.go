package main

import (
    "fmt"
    "os"
    "os/exec"
    "io/ioutil"
    "path/filepath"
    "encoding/json"
)

const configFileName string = ".wolhost"
const commandPath string = "/usr/bin/wakeonlan"

func Usage() {
    fmt.Println("usage : wol [ -l | -h | host ]")
    fmt.Println(" Options")
    fmt.Println("   host")
    fmt.Println("      target host name")
    fmt.Println("   -l")
    fmt.Println("      print host list in config file")
    fmt.Println("   -h")
    fmt.Println("      help information")
}

func Help() {
    fmt.Println("")
    fmt.Println("ホームディレクトリに")
    fmt.Println("\".wolhost\"という名前でホストのMACアドレスをjson")
    fmt.Println("形式で列挙したファイルを作成してください．")
    fmt.Println("")
    fmt.Println("ファイルの例: ~/.wolhost")
    fmt.Println("[")
    fmt.Println("  {\"hostname\":\"host1\" , \"mac\": \"b8:27:eb:10:f3:9c\"},")
    fmt.Println("  {\"hostname\":\"host2\" , \"mac\": \"b8:27:eb:10:f4:ab\"}")
    fmt.Println("]")
}

type Config struct {
    HostName   string `json:"hostname"`
    MacAddress string `json:"mac"`
}


func loadConfig() []Config {
    home, err := os.UserHomeDir()
    if err != nil {
	fmt.Println("Error : can not determine user home directory.")
        os.Exit(1)
    }
    configFilePath := filepath.Join(home,configFileName)
    // JSON形式configファイル読み込み
    texts, err := ioutil.ReadFile(configFilePath)
    if err != nil {
        //log.Fatal(err)
	fmt.Println("Error : can not load config file \"~/.wolhost\".")
        Usage()
	Help()
        os.Exit(1)
    }
    // configデータ(JSON)デコード
    var hostLists []Config
    //if err := json.Unmarshal(texts, &config); err != nil {
    if err := json.Unmarshal(texts, &hostLists); err != nil {
        //log.Fatal(err)
	fmt.Println("Error : can not parse config file.")
	Help()
        os.Exit(1)
    }
    //fmt.Println(texts)
    return hostLists
}

func DumpHostList(hostLists []Config) {
    for _, hostEntry := range hostLists {
        fmt.Printf("hostname = %v , mac = %v\n",hostEntry.HostName, hostEntry.MacAddress)
    }
}

func searchMacAddr(hostLists []Config, hostname string) string {
    for _, hostEntry := range hostLists {
        if hostEntry.HostName == hostname {
            return hostEntry.MacAddress
        }
    }
    return ""
}

func sendPacket(mac string){
    err := exec.Command(commandPath, mac).Run()
    if err != nil {
	fmt.Println("Error : can not exec \"wakeonlan\" command.")
        os.Exit(1)
    }
}

func main() {
    argSize := len(os.Args)
    argv := os.Args
    if argSize != 2 {
        Usage()
	os.Exit(1)
    }
    if ("-h" == argv[1]) {
        Usage()
	Help()
	os.Exit(1)
    }
    hostList := loadConfig()
    if ("-l" == argv[1]) {
        DumpHostList(hostList)
	os.Exit(1)
    }
    targetMac := searchMacAddr(hostList, argv[1])
    if targetMac == "" {
        fmt.Println("Error : can not find target host name in config file.")
	Help()
        os.Exit(1)
    }
    sendPacket(targetMac)
}




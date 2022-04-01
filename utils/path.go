package utils

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path"
	"path/filepath"
	"strings"
)

// 列出指定目录下所有文件
func ListAll(path string, curLayer int) {
	fileInfos, err := ioutil.ReadDir(path)
	if err != nil {
		fmt.Printf("ReadDir failed,error:%+v\n", err)
		return
	}
	for _, info := range fileInfos {
		if info.IsDir() {
			for tmpHier := curLayer; tmpHier > 0; tmpHier-- {
				fmt.Printf("|\t")
			}
			fmt.Println(info.Name(), "\\")
			ListAll(path+"/"+info.Name(), curLayer+1)
		} else {
			for tmpHier := curLayer; tmpHier > 0; tmpHier-- {
				fmt.Printf("|\t")
			}
			fmt.Println(info.Name())
		}
	}
}

func RenameToPNG(dirname string) {
	fileInfos, err := ioutil.ReadDir(dirname)
	if err != nil {
		log.Fatal(err)
	}
	for _, f := range fileInfos {
		itemName := path.Join(dirname, f.Name()) //记录当前文件(夹)名
		if f.IsDir() {                        //判断是否是文件夹 如果是文件夹则继续递归调用
			RenameToPNG(itemName)
		} else {
			if path.Ext(itemName) == "" { // 判断是否有扩展名
				if os.Rename(itemName, itemName+".png") == nil {
					fmt.Println("[OK]", itemName) //打印文件地址
				} else {
					fmt.Printf("[ERR] %s %s\n", itemName, err.Error())
				}
			}
		}
	}
}

// 将 / 替换成 \
func ClearPath(path string) string {
	return strings.Replace(path, "\\", "/", -1)
}

// 获取壁纸所在文件夹
func GetAssetsPath() string {
	HomeDir, _ := os.UserHomeDir()
	suffix := "AppData/Local/Packages/Microsoft.Windows.ContentDeliveryManager_cw5n1h2txyewy/LocalState/Assets"

	AssetsDir := path.Join(HomeDir, suffix)
	AssetsDir = ClearPath(AssetsDir)

	return AssetsDir
}

// 获取程序执行目录
func GetExecDir() string {
	// 返回程序执行目录绝对路径
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	ExecDir := ClearPath(dir)

	return ExecDir
}

// 获取目标文件夹
func GetDstDir() string {
	return ClearPath(path.Join(GetExecDir(), "wallpapers"))
}

func Copy(from, to string) error {
	f, err := os.Stat(from)
	if err != nil {
		return err
	}

	if f.IsDir() {
		// from 是文件夹，则使 to 也为文件夹
		if list, err := ioutil.ReadDir(from); err == nil {
			// 遍历 from 文件夹
			for _, item := range list {
				if err := Copy(filepath.Join(from, item.Name()), filepath.Join(to, item.Name())); err != nil {
					return err
				}
			}
		}
	} else {
		// from 使文件，创建 to 的文件夹
		p := filepath.Dir(to)
		if _, err := os.Stat(p); err != nil {
			if err = os.MkdirAll(p, 0777); err != nil {
				return err
			}
		}

		// 读取
		file, err := os.Open(from)
		if err != nil {
			return err
		}
		defer file.Close()

		bufReader := bufio.NewReader(file)
		out, err := os.Create(to)
		if err != nil {
			return err
		}
		defer out.Close()

		_, err = io.Copy(out, bufReader)
		if err != nil {
			return err
		}
	}
	return err
}

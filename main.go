package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math"
	"os"
)

type config struct {
	LostMoney float64 `json:"lostMoney"`
	K         float64 `json:"K"`
}

func main() {
	var (
		nowPrice  float64
		lostPrice float64
		lostMoney float64
		K         float64
	)
	// 打开json文件
	jsonFile, err := os.Open("config.json")

	// 最好要处理以下错误
	if err != nil {
		fmt.Println("config.json文件不存在，请查看该文件")
		return
	}
	byteValue, _ := ioutil.ReadAll(jsonFile)
	var con config

	err = json.Unmarshal([]byte(byteValue), &con)
	if err != nil {
		fmt.Println("config.json文件错误，请查看该文件")
	}
	defer jsonFile.Close()
	for {
		fmt.Print("当前价格：")
		_, err = fmt.Scanln(&nowPrice)
		if err != nil {
			return
		}
		fmt.Print("止损价格：")
		_, err = fmt.Scanln(&lostPrice)
		if err != nil {
			return
		}
		fmt.Print("止损金额：")
		_, err = fmt.Scanln(&lostMoney)
		if err != nil {
			lostMoney = con.LostMoney
		}
		fmt.Print("盈亏比：")
		_, err = fmt.Scanln(&K)
		if err != nil {
			K = con.K
		}
		spreads := math.Abs(nowPrice - lostPrice)
		leverage := nowPrice / spreads
		num := lostMoney / spreads
		callbackRate := spreads / nowPrice
		activatedPrice := K*spreads + nowPrice
		fmt.Println("杠杆倍数：", leverage)
		fmt.Println("应开数量：", num)
		fmt.Println("回调率：", callbackRate)
		fmt.Println("激活价格：", activatedPrice)
	}

}

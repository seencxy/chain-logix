package utils

import (
	"backed/app/models"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
)

// SelectMapInfo 查询地图信息
func SelectMapInfo(apiKey string, address string) (string, error) {
	// 创建http客户端
	client := http.Client{}

	url := fmt.Sprintf("https://restapi.amap.com/v3/geocode/geo?key=%s&address=%s", apiKey, address)

	// 创建请求
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", err
	}

	do, err := client.Do(request)
	if err != nil {
		return "", err
	}

	all, err := io.ReadAll(do.Body)
	if err != nil {
		return "", err
	}

	var response models.SelectMapInfoResponse

	err = json.Unmarshal(all, &response)
	if err != nil {
	}

	if response.Status == "1" {
		return response.Geocodes[0].Location, nil
	}

	return "", errors.New("查询地图信息失败")
}

// TransferStringToFloat 将string转化成float64
func TransferStringToFloat(str string) ([]float64, error) {
	split := strings.Split(str, ",")
	var result []float64
	for _, v := range split {
		float, err := strconv.ParseFloat(v, 64)
		if err != nil {
			return nil, err
		}
		result = append(result, float)
	}
	return result, nil
}

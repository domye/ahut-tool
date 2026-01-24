package utils

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

// SaveJSON 将结构体保存为JSON文件
func SaveJSON(data interface{}, filename string) error {
	// 获取可执行文件所在目录
	exeDir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		return fmt.Errorf("failed to get executable directory: %v", err)
	}

	// 构造数据目录路径
	dataDir := filepath.Join(exeDir, "data")
	// 如果数据目录不存在，则创建
	if _, err := os.Stat(dataDir); os.IsNotExist(err) {
		err = os.MkdirAll(dataDir, 0755)
		if err != nil {
			return fmt.Errorf("failed to create data directory: %v", err)
		}
	}

	// 构造文件路径
	filePath := filepath.Join(dataDir, filename)

	// 将结构体转换为JSON
	jsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal JSON: %v", err)
	}

	// 写入文件
	err = os.WriteFile(filePath, jsonData, 0644)
	if err != nil {
		return fmt.Errorf("failed to write file: %v", err)
	}

	return nil
}

// LoadJSON 从JSON文件读取数据到结构体
func LoadJSON(filename string, data interface{}) error {
	// 获取可执行文件所在目录
	exeDir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		return fmt.Errorf("failed to get executable directory: %v", err)
	}

	// 构造文件路径
	filePath := filepath.Join(exeDir, "data", filename)

	// 检查文件是否存在
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return fmt.Errorf("file does not exist: %v", err)
	}

	// 读取文件
	jsonData, err := os.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("failed to read file: %v", err)
	}

	// 解析JSON到结构体
	err = json.Unmarshal(jsonData, data)
	if err != nil {
		return fmt.Errorf("failed to unmarshal JSON: %v", err)
	}

	return nil
}

// FileExists 检查文件是否存在
func FileExists(filename string) bool {
	// 获取可执行文件所在目录
	exeDir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		return false
	}

	// 构造文件路径
	filePath := filepath.Join(exeDir, "data", filename)

	// 检查文件是否存在
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return false
	}

	return true
}

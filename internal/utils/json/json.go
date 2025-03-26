package json

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

func ToColorJson(obj interface{}) string {
	if obj == nil {
		return ""
	}

	str, err := json.MarshalIndent(obj, "", " ")
	if err != nil {
		fmt.Println(err)
		return ""
	}

	return string(str)
}

func WriteToFile(file *os.File, obj interface{}) error {
	jsonData, err := json.MarshalIndent(obj, "", " ")
	if err != nil {
		return err
	}

	_, err = file.Write(jsonData)
	if err != nil {
		return err
	}

	return nil
}

func ReadFromFile(file *os.File, obj interface{}) error {
	jsonData, err := io.ReadAll(file)
	if err != nil {
		return err
	}

	return json.Unmarshal(jsonData, obj)
}

package helpers

import (
	"bufio"
	"log"
	"os"
	"strings"

	"encoding/json"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"io/ioutil"
)

func RegisterEnvVars(logfile string) (map[string]string, error) {
	f, err := os.OpenFile(logfile, os.O_RDONLY, os.ModePerm)
	dict := map[string]string{}
	if err != nil {
		log.Fatalf("open file error: %v", err)
		return dict, err
	}
	defer f.Close()

	sc := bufio.NewScanner(f)

	for sc.Scan() {
		var line string = sc.Text() // GET the line string
		if len(line) > 0 {
			tempEnv := strings.Split(line, "=")
			key := tempEnv[0]
			value := tempEnv[1]
			dict[key] = value
			os.Setenv(key, value)
		}
	}

	if err := sc.Err(); err != nil {
		log.Fatalf("scan file error: %v", err)
		return dict, err
	}
	return dict, nil
}

func ReadStructFromJSON(file string) (map[string]interface{}, error) {
	// Read JSON file
	data, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatal("Error reading JSON file:", err)
		return nil, err
	}
	// Define a struct to hold JSON data
	structure := make(map[string]interface{})

	// Unmarshal JSON data into struct
	err = json.Unmarshal(data, &structure)
	if err != nil {
		log.Fatal("Error parsing JSON:", err)
		return nil, err
	}
	return structure, nil

}

func ConvertToPrimitiveE(data map[string]interface{}) (bson.D, error) {
	var d primitive.D
	for key, value := range data {
		d = append(d, primitive.E{Key: key, Value: value})
	}

	bsonD := bson.D(d)
	return bsonD, nil
}

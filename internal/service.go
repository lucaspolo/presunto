package service

import (
	"context"
	"embed"
	"encoding/csv"
	"fmt"
	"strings"
)

//go:embed resources/qcodes.csv
var csvFile embed.FS

var qCodeMap = make(map[string]string)

func GetQCode(ctx context.Context, q string) (string, error) {
	q = strings.ToUpper(q)
	value, ok := qCodeMap[q]
	if !ok {
		return "", fmt.Errorf("qcode %s not found", q)
	}

	return value, nil
}

func init() {
	data, err := csvFile.ReadFile("resources/qcodes.csv")
	if err != nil {
		panic(err)
	}

	reader := csv.NewReader(strings.NewReader(string(data)))

	records, err := reader.ReadAll()

	if err != nil {
		panic(err)
	}

	for _, record := range records {
		qCodeMap[record[0]] = record[1]
	}
}

package version

import (
	"context"
	"encoding/csv"
	"io"
	"os"
	"strconv"
	"sync"
)

type CSV struct {
	reader   io.Reader
	writer   io.Writer
	mutex    *sync.RWMutex
	versions []*Version
}

func NewCSV(file *os.File) (*CSV, error) {
	csvManager := &CSV{
		reader: file,
		writer: file,
		mutex:  &sync.RWMutex{},
	}

	return csvManager, csvManager.load()
}

func (csvManager *CSV) load() error {
	csvReader := csv.NewReader(csvManager.reader)
	csvReader.FieldsPerRecord = 2

	csvManager.mutex.Lock()
	defer csvManager.mutex.Unlock()

	for {
		record, err := csvReader.Read()

		if err == io.EOF {
			break
		}

		if err != nil {
			return err
		}

		number, err := strconv.ParseInt(record[0], 10, 32)

		if err != nil {
			return err
		}

		csvManager.versions = append(csvManager.versions, NewVersion(int32(number), record[1]))
	}

	return nil
}

func (csvManager *CSV) Exists(_ context.Context, number int32) (bool, error) {
	csvManager.mutex.RLock()
	defer csvManager.mutex.RUnlock()

	for _, version := range csvManager.versions {
		if version.Number() == number {
			return true, nil
		}
	}

	return false, nil
}

func (csvManager *CSV) Save(_ context.Context, version *Version) error {
	record := []string{
		strconv.Itoa(int(version.number)),
		version.description,
	}

	writer := csv.NewWriter(csvManager.writer)

	csvManager.mutex.Lock()
	defer csvManager.mutex.Unlock()

	if err := writer.Write(record); err != nil {
		return err
	}

	writer.Flush()

	if err := writer.Error(); err != nil {
		return err
	}

	csvManager.versions = append(csvManager.versions, version)

	return nil
}

package storage_test

import (
	"io/ioutil"
	"neurotech-assignment/backend/internal/storage"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type FileStorageSuite struct {
	suite.Suite
	tempDir      string
	tempFilePath string
	fileStorage  *storage.FileStorage
}

func (suite *FileStorageSuite) SetupTest() {
	var err error
	suite.tempDir, err = ioutil.TempDir("", "filestorage_test")
	suite.Require().NoError(err)

	suite.tempFilePath = filepath.Join(suite.tempDir, "test.txt")
	suite.fileStorage = storage.NewFileStorage(suite.tempFilePath)
}

func (suite *FileStorageSuite) TearDownTest() {
	os.RemoveAll(suite.tempDir)
}

func (suite *FileStorageSuite) TestSaveAndLoad() {
	data := []byte("test data")

	err := suite.fileStorage.Save(data)
	suite.Require().NoError(err)

	loadedData, err := suite.fileStorage.Load()
	suite.Require().NoError(err)

	assert.Equal(suite.T(), data, loadedData, "Loaded data should match saved data")

	_, err = os.Stat(suite.tempFilePath)
	assert.False(suite.T(), os.IsNotExist(err), "File should exist after saving")

	err = os.Remove(suite.tempFilePath)
	suite.Require().NoError(err)

	_, err = os.Stat(suite.tempFilePath)
	assert.True(suite.T(), os.IsNotExist(err), "File should not exist after removing")
}

func TestFileStorageSuite(t *testing.T) {
	suite.Run(t, new(FileStorageSuite))
}

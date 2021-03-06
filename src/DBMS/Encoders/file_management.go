package Encoders

import (
	"os"
	"log"
)

// Constant string for the data directory.
var dir = "./Data"

// IntializeDirectory is a function to ensure that the base directory where
// all the data files for the database are stored exists. If it does not, then
// create this base directory.
func InitializeDirectory() {

	// Get the status of the ./Data path.
	_, err := os.Stat(dir)

	// If the path does not exist, create a directory with the appropriate 
	// permissions.
	if os.IsNotExist(err){
		err := os.Mkdir("./Data", 0755)

		// If there is an error with creating the file, log that error.
		if err != nil {
			log.Println(err)
		}
	}
}

// FileExists is a function to check whether a file exists in the data directory
// given the file name. It returns a boolean representing whether the file
// exists or not.
func FileExists(name string) (bool) {

	// Get the status of the file.
	_, err := os.Stat(dir+"/"+name+".data")

	// If it does not exist return false, else return true.
	if os.IsNotExist(err){
		return false;
	}

	return true;
}

// ReadByteFile is a function to read from a file in the data directory given
// the file name. It returns the bytes of the files and the error.
func ReadByteFile(name string) ([]byte, error) {
	return os.ReadFile(dir+"/"+name+".data")
}

// WriteByteFile is a function to write to a file in the data directory given
// the file name and the bytes to write. It return any errors the file may have.
func WriteByteFile(name string, data []byte) (error) {
	err := os.WriteFile(dir+"/"+name+".data", data, 0755)
	return err
}

// DeleteFile is a function to remove a file from the data directory by the
// given name input. It returns any errors that may occur during the delete.
func DeleteFile(name string) (error) {
	err := os.Remove(dir+"/"+name+".data")
	return err
}

// SetTestingPath is a function to change the relative data file path so it can
// be accessed from the testing directory.
func SetTestingPath() {
	dir = "../Data"
}
package repo

import (
	"fmt"
	"os"
	"path/filepath"
)

func InitRepo() error {
	DirName := ".gitgo"

	// .gitgo dir creation
	if err := os.Mkdir(DirName, 0755); err != nil && !os.IsExist(err) {
		fmt.Println("Error creating .gitgo: ", err)
		return err
	}

	//.gitgo/objects dir
	ObjectsPath := filepath.Join(DirName, "objects")
	if err := os.Mkdir(ObjectsPath, 0755); err != nil && !os.IsExist(err) {
		fmt.Println("Error creating objects: ", err)
		return err
	}

	//.gitgo/refs dir
	RefsPath := filepath.Join(DirName, "refs")
	if err := os.Mkdir(RefsPath, 0755); err != nil && !os.IsExist(err) {
		fmt.Println("Error creating refs: ", err)
		return err
	}

	// .gitgo/refs/heads dir
	Refs_HeadsPath := filepath.Join(RefsPath, "heads")
	if err := os.Mkdir(Refs_HeadsPath, 0755); err != nil && !os.IsExist(err) {
		fmt.Println("Error creating refs: ", err)
		return err
	}

	// HEAD file
	HeadFile := filepath.Join(DirName, "HEAD")
	/*here firstwe are cheching that if the HeadFile exists or not, if not exists then it will trigger next if block. in next if block
	er are writing the HEAD file which is storing the current branch, for now it doesnot support multiple branch */
	if _, err := os.Stat(HeadFile); os.IsNotExist(err) {
		if err := os.WriteFile(HeadFile, []byte("ref: refs/heads/master"), 0644); err != nil {
			fmt.Println("Error creating HEAD: ", err)
			return err
		}
	}

	//index file
	IndexFile := filepath.Join(DirName, "index")
	if _, err := os.Stat(IndexFile); os.IsNotExist(err) {
		if err := os.WriteFile(IndexFile, []byte(""), 0644); err != nil {
			fmt.Println("Error creating the index: ", err)
			return err
		}
	}

	return nil
}

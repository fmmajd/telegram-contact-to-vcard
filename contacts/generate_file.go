package contacts

import "os"

func generateFile(targetPath string, content string) error {
	f, err := os.Create(targetPath)
	if err != nil {
		return err
	}
	defer f.Close()

	bytes := []byte(content)
	_, err = f.Write(bytes)
	if err != nil {
		return err
	}

	return nil
}

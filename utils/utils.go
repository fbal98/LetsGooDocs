package utils

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)


var IGNORED_DIRS = []string{"score_cal.ts", ".git", ".vscode", "node_modules", "vendor", "dist", "build", "bin", "pkg", "test", "tests", "tmp", "temp", "logs", "log", "docs", "doc", "coverage", "assets", "public", "static", "uploads", "download", "downloads", "cache", "tmp", "temp", "images", "img", "media", "fonts", "css", "scss", "js", "javascript", "html", "htm", "txt", "md", "markdown", "yml", "yaml", "json", "xml", "csv", "ts", "typescript", "php", "py", "python", "rb", "ruby", "java", "jar", "class", "war", "ear", "go", "golang", "sh", "bash", "shell", "ps1", "powershell", "bat", "cmd", "exe", "dll", "bin", "obj", "lib", "so", "a", "out", "app", "ipa", "apk", "dmg", "pkg", "deb", "rpm", "iso", "img", "tar", "gz", "zip", "7z", "rar", "bak", "backup", "sql", "db", "sqlite", "sqlite3", "db3", "log", "conf", "config", "cfg", "ini", "env", "env", "md5", "sha1", "sha256", "sha512", "key", "pem", "crt", "cert", "pub", "private", "public", "tsv", "xls", "xlsx", "ods", "doc", "docx", "odt", "pdf", "ppt", "pptx", "odp", "csv", "json", "xml", "yaml", "yml", "zip", "tar", "gz", "7z", "rar", "jpg", "jpeg", "png", "gif", "svg", "bmp", "ico", "webp", "tiff", "tif", "mp4", "webm", "mkv", "flv", "avi", "mov", "wmv", "mp3", "wav", "ogg", "flac", "wma" }

//Validation functions

func IsPathNotFound(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return true
	}
	return false	
}





//helper functions
func ListDirNames(path string) ([]string, error) {
	var files []string
	

	dir, err := os.Open(path)
	if err != nil {
		return files, err
	}

	defer dir.Close()

	fileInfo, err := dir.Readdir(-1)
	if err != nil {
		return nil, err
	}

	for _, file := range fileInfo {
		if file.IsDir() {
			//check if the directory is in the ignoredDirs
			if _contains(IGNORED_DIRS, file.Name()) {
				continue
			}
			subDirFiles, err := ListDirNames(path + "/" + file.Name())
			if err != nil {
				return files, err
			}

			files = append(files, subDirFiles...)
		} else {
			if path != "./" {
				files = append(files, path + "/" + file.Name())
			} else {
				files = append(files, file.Name())
			}
		}
	}

	return files, nil
}


func GetFilesContent(files []string) map[string]string {
	 chunkSize, sizeError := strconv.Atoi(os.Getenv("CHUNK_SIZE"))
	 chunkOverLapSize, overlapError := strconv.Atoi(os.Getenv("CHUNK_OVERLAP_SIZE"))

	 if sizeError != nil || overlapError != nil {
		return nil
	 }

	fileContents := make(map[string]string)

	for _, file := range files {

		content, err := os.ReadFile(file)

		if err != nil {
			fileContents[file] = ""
		} else {

			tmpContent := string(content)

			if len(tmpContent) > chunkSize {
				chunkNumber := 1

				for i := 0; i < len(tmpContent); i += chunkSize - chunkOverLapSize {
					end := i + chunkSize
					if end > len(tmpContent) {
						end = len(tmpContent)
					}
					fileContents[file] += "Chunk Number:" + strconv.Itoa(chunkNumber) + "\n" + tmpContent[i:end] + "$EndOFChunk$"
					chunkNumber++
				}

			} else {
				fileContents[file] = tmpContent
			}
		}

}

	return fileContents
}

func WriteToFile(fileName, document string) error {
	//if the file exist
	if _, err := os.Stat(fileName); err == nil {
		//generate timestamp 
		timestamp := strconv.FormatInt(time.Now().Unix(), 10)
		//rename the file
		err := os.Rename(fileName, timestamp + fileName)
		if err != nil {
			fmt.Println("Error renaming file: ", err)
			return err
		
		}
	}
	err := os.WriteFile(fileName, []byte(document), 0644)
	if err != nil {
		fmt.Println("Error writing to file: ", err)
		return err
	}
	return nil
}

//Write file but for map
func WriteMapToFile(fileName string, document map[string]string) error {
	if _, err := os.Stat(fileName); err == nil {
		timestamp := strconv.FormatInt(time.Now().Unix(), 10)
		err := os.Rename(fileName, timestamp + "-" +fileName)
		if err != nil {
			fmt.Println("Error renaming file: ", err)
			return err
		
		}
	}
	
	//get the content of the map
	content := ""
	for key, value := range document {
		content += key + "\n" + value + "\n"
	
	}

	err := os.WriteFile(fileName, []byte(content), 0644)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return err
	}

	return nil
	
}


//GeneratePrompt of the content
//single prompt for the entire content
//THIS FUNCTION WILL BE DEPRECATED IN THE FUTURE, 
//WHEN THE GeneratePrompts FUNCTION IS IMPLEMENTED
func GeneratePrompt(fileContents map[string] string) string {
	content := ""
	for fileName, fileContent := range fileContents {
		content += "fileName:" + fileName + "\n" + fileContent + "\n"
	}

	prompt, err := _formatPrompt("fileContent", content, "1")
	if err != nil {
		fmt.Println("Error formatting prompt: ", err)
		return ""
	
	}
	
	WriteToFile("prompt.txt", prompt)
	return prompt
}

//TODO: Complete the implementation of this function
//This function should generate prompts for the chunks of content
func GeneratePrompts(fileContents map[string]string) map[string]string {
   	prompts := make(map[string]string)

	for filename, content := range fileContents {
		prompts[filename] = ""
		chunks := strings.Split(content, "$EndOFChunk$")
		for _, chunk := range chunks {
			if len(chunk) > 0 {
				prompt, _ := _formatPrompt(filename, chunk, strconv.Itoa(len(prompts[filename])+1))
				prompts[filename] += prompt
			}
		}
	}

	return prompts
    
}

//TODO: Implement dynamic chunking and overlapping 
//chunks should be meaningful and contextually complete
func AdjustChunkBoundaries(content string) string {
	return content
}



//internal functions
func _contains(ignoredDirs []string, s string) bool {
	for _, dir := range ignoredDirs {
		if dir == s {
			return true
		}
	}
	return false
}

func _formatPrompt(filename, content, chunkNumber string) (string, error){
	templateBytes, err := os.ReadFile("templates/prompt.txt")

	if err != nil {
		return "", err
	}

	template := string(templateBytes)

	formattedPrompt := strings.ReplaceAll(template, "{filename}", filename)
	formattedPrompt = strings.ReplaceAll(formattedPrompt, "{number}", chunkNumber)
	formattedPrompt = strings.ReplaceAll(formattedPrompt, "{content}", content)

	return formattedPrompt, nil
}
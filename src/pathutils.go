package main

import (
	"bufio"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
)

type Path struct {
}

type Directory struct {
}

type File2 struct {
}

func (d *Directory) CreateDirectory(dir string) {
	if d.Exists(dir) {
		return
	}
	os.MkdirAll(dir, os.ModePerm)
}

func (d *Directory) GetExtension(file string) string {
	return filepath.Ext(file)
}

func (d *Directory) GetFilenameWithoutExtension(file string) string {
	sss := &Strings2{}
	ext := d.GetExtension(file)
	if ext == "" {
		return file
	}
	return sss.Substring(file, 0, len(file)-len(ext))
}

func (d *Directory) GetFiles(dir string) []string {
	var files []string

	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		files = append(files, path)
		return nil
	})

	if err != nil {
		return files
	}

	return files
}

func (p *Directory) Exists(dir string) bool {
	info, err := os.Stat(dir)
	if os.IsNotExist(err) {
		return false
	}
	return info.IsDir()
}

func (d *Directory) GetCurrentDirectory() string {
	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	exPath := filepath.Dir(ex)
	return exPath
}

func (d *Directory) GetFilenameOnly(str string) string {
	file := filepath.Base(str)
	return file
}

func (p *Path) Combine(dir string, file string) string {
	// if path.IsAbs(target) {
	// 	return target
	// }
	return path.Join(dir, file)
}

func (p *File2) ReadAllText(file string) string {
	content, err := ioutil.ReadFile(file)
	if err != nil {
		return ""
	}
	return string(content)
}

func (p *File2) WriteAllLines(file string, lines []string) {
	file2, err := os.OpenFile(file, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		return
	}
	defer file2.Close()

	datawriter := bufio.NewWriter(file2)

	for _, data := range lines {
		_, _ = datawriter.WriteString(data + "\r\n")
	}

	datawriter.Flush()
}

func (p *File2) ReadAllBytes(file string) []byte {
	content, err := ioutil.ReadFile(file)
	if err != nil {
		return nil
	}
	return content
}

func (p *File2) Exists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func (p *File2) Delete(filename string) {
	if p.Exists(filename) {
		os.Remove(filename)
	}
}

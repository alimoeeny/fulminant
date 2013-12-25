package fulminant

import (
	"github.com/alimoeeny/cssmin"
	"github.com/alimoeeny/htmlmin"
	"github.com/alimoeeny/jsmin"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func MinifyProject(projectPath, targetPath string) (err error) {
	err = filepath.Walk(projectPath,
		func(path string, info os.FileInfo, err error) error {
			ext := filepath.Ext(path)
			ext = strings.ToLower(ext)
			rel, err := filepath.Rel(projectPath, path)
			if err != nil {
				log.Println("something is wrong with target or maybe source path")
				return err
			}
			targetfilename := filepath.Join(targetPath, rel)
			switch ext {
			case ".css":
				err = Minifycss(path, targetfilename)
			case ".html":
				err = Minifyhtml(path, targetfilename)
			case ".js":
				err = Minifyjs(path, targetfilename)
			default:
				log.Printf("%s file not supported: %s\n", ext, path)
			}
			return err
		})
	return err
}

func Minifyhtml(sourcehtmlfilename, targethtmlfilename string) error {
	log.Printf("Reading the file:%s\n", sourcehtmlfilename)
	bb, err := ioutil.ReadFile(sourcehtmlfilename)
	if err != nil {
		return err
	}
	log.Printf("length of original bytes: %d\n", len(bb))
	log.Println("minifying")
	bmin, err := htmlmin.Minify(bb, htmlmin.DefaultOptions)
	if err != nil {
		return err
	}
	log.Printf("length of minified bytes: %d\n", len(bmin))
	log.Printf("writing the file:%s\n", targethtmlfilename)
	if _, exits := os.Stat(filepath.Dir(targethtmlfilename)); exits != nil {
		err := os.MkdirAll(filepath.Dir(targethtmlfilename), os.ModePerm)
		if err != nil {
			log.Printf("failed to create the directory: %s\n", filepath.Dir(targethtmlfilename))
			return err
		}
	}
	err = ioutil.WriteFile(targethtmlfilename, bmin, os.ModePerm)
	if err != nil {
		return err
	}
	log.Println("done")
	return nil
}

func Minifyjs(sourcejsfilename, targetjsfilename string) error {
	log.Printf("Reading the file:%s\n", sourcejsfilename)
	bb, err := ioutil.ReadFile(sourcejsfilename)
	if err != nil {
		return err
	}
	log.Printf("length of original bytes: %d\n", len(bb))
	log.Println("minifying")
	bmin, err := jsmin.Minify(bb)
	if err != nil {
		return err
	}
	log.Printf("length of minified bytes: %d\n", len(bmin))
	log.Printf("writing the file:%s\n", targetjsfilename)
	if _, exits := os.Stat(filepath.Dir(targetjsfilename)); exits != nil {
		err := os.MkdirAll(filepath.Dir(targetjsfilename), os.ModePerm)
		if err != nil {
			log.Printf("failed to create the directory: %s\n", filepath.Dir(targetjsfilename))
			return err
		}
	}
	err = ioutil.WriteFile(targetjsfilename, bmin, os.ModePerm)
	if err != nil {
		return err
	}
	log.Println("done")
	return nil
}

func Minifycss(sourcecssfilename, targetcssfilename string) error {
	log.Printf("Reading the file:%s\n", sourcecssfilename)
	bb, err := ioutil.ReadFile(sourcecssfilename)
	if err != nil {
		return err
	}
	log.Printf("length of original bytes: %d\n", len(bb))
	log.Println("minifying")
	bmin := cssmin.Minify(bb)
	if err != nil {
		return err
	}
	log.Printf("length of minified bytes: %d\n", len(bmin))
	log.Printf("writing the file:%s\n", targetcssfilename)
	if _, exits := os.Stat(filepath.Dir(targetcssfilename)); exits != nil {
		err := os.MkdirAll(filepath.Dir(targetcssfilename), os.ModePerm)
		if err != nil {
			log.Printf("failed to create the directory: %s\n", filepath.Dir(targetcssfilename))
			return err
		}
	}
	err = ioutil.WriteFile(targetcssfilename, bmin, os.ModePerm)
	if err != nil {
		return err
	}
	log.Println("done")
	return nil
}

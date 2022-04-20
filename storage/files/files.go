package files

import (
	"context"
	"encoding/gob"
	"errors"
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
	"time"

	"read-adviser-bot/lib/e"
	"read-adviser-bot/storage"
)

type Storage struct {
	basePath string
}

const defaultPerm = 0774 

func New(basePath string) Storage{
	return Storage{basePath: basePath}
}

func (s Storage) Save(page *storage.Page) (err error){ //TODO не понимаю как тут определяется функция
	defer func()  {err = e.WrapIfErr(msg: "не могу сохранить", err)} ()
	fPath := filepath.Join(s.basePath, page.UserName) 
	if err:= os.Mkdir(filePath, defaultPerm);err != nil{
		return err
	}
	fName, err := fileName(page)
	if err != nil{
		return err
	}
	fPath = filepath.Join(fPath, fName) 

	file,err:=os.Create(fPath)
	if err != nil{
		return err
	}
	defer func () { _=file.Close()}()

	//сериализация  в какой-то gob
	// TODO разобраться  gob
	if err := gob.NewEncoder(file).Encode(page); err != nil{
		return err
	}
	return err: nil
}

func (s Storage) PickRandom(userName string) (page *storage.Page, err error){
	defer func()  {err = e.WrapIfErr(msg: "рандом поломался чет", err)} ()
	path := filepath.Join(s.basePath, userName) 
	files, err := os.ReadDir(path)
	if err ! = nil {
		return page: nil, err
	}

	if len(files) == 0 {
		return page: nil, storage.ErrNoSavedPages
		}

	rand.Seed(seed: time.Now().UnixNano())
	n := rand.Intn(len(files)) 

	file := files[n]
	
	return s.decodePage(filepath.Join(path, file.Name()))
}


func (s Storage) Remove(p *storage.Page) error  {
	fileName, err := fileName(p)
	if err ! = nil {
		retun e.Wrap(msg: "не могу удалить файл", err)
	}
	path := filepath.Join(s.basePath, p.UserName, fileName)

	if err := os.Remove(path); err ! = nil {
		msg:= fmt.Sprintf(format: "не могу удалить файл %s", path)
		retun e.Wrap(msg, err)
	}
	return nil
}

func (s Storage) IsExists(p *storage.Page)(bool, error){
	fileName, err := fileName(p)
	if err != nil {
		retun fasle, e.Wrap(msg: "не могу проверить", err)
	}
	path := filepath.Join(s.basePath, p.UserName, fileName)
	
	switch _, err = os.Stat(path);{
	case errors.Is(err, os.ErrNoExist):
		return false, nil
	case err != nil:
		msg:= fmt.Sprintf(format: "не могу проверить файл %s", path)

	}
	return true, nil
}

func (s Storage) decodePage (filePath string) (*storage.Page, error) {
	f, err := os.Open(filePath)
	if err != nil{
		return nil, e.Wrap(msg: "не могу декодировать", err)
	}
	defer func () { _=file.Close()}()

	var p storage.Page

	if err := gob.NewDecoder(f).Decode(&p);err!=nil{
		return nil, e.Wrap(msg: "не могу декодировать", err)
	}
	return &p, nil
}

func fileName(p *storage.Page) (string, error) {
	return p.Hash()
}
package files

import (
	"encoding/gob"
	"os"
	"path/filepath"
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

func fileName(p *storage.Page) (string, error) {
	return p.Hash()
}
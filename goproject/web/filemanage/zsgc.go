package web

import (
	"encoding/csv"
	"io"
	"os"
)

var (
	f   *os.File
	err error
)

/*
	检查文件是否存在在磁盘上
*/
func checkFileIsExist(filename string) bool {
	var exist = true
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		exist = false
	}
	return exist
}

/*
	这个函数用来判断用户是否存在且秘密是否正确
*/
func readfile(filename, name, psw string) bool {
	defer f.Close()
	f, err = os.Open(filename)

	if err != nil {
		println("can not open the file ,err is %+v", err)
	}
	//
	r := csv.NewReader(f)
	for {
		row, err := r.Read()
		if err != nil && err != io.EOF {
			println("can not read ,the err is %+v", err)
		}
		if err == io.EOF {
			break
		}
		// println(row[0])
		if row[0] == name && row[1] == psw {
			return true

		}
	}
	return false
}

/*
	新增用户
*/

func addfile(filename, name, psw, mail string) bool {
	defer f.Close()
	f, err = os.OpenFile(filename, os.O_RDWR|os.O_APPEND, 0666)
	if err != nil {
		println("can not add ,err is %+v", err)
	}
	writerCsv := csv.NewWriter(f)
	word := []string{name, psw, mail}
	err1 := writerCsv.Write(word)
	if err1 != nil {
		println("can not add ,err is %+v", err1)
	}
	writerCsv.Flush()
	return true
}

func isexist(filename, name, mail string) (namexist, mailexist bool) {
	defer f.Close()
	f, err = os.Open(filename)
	namexist = false
	mailexist = false
	if err != nil {
		println("can not open the file ,err is %+v", err)
	}

	r := csv.NewReader(f)
	for {
		row, err := r.Read()
		if err != nil && err != io.EOF {
			println("can not read,the err is %+v", err)
		}
		if err == io.EOF {
			break
		}

		if row[0] == name {
			namexist = true
			break
		}
		if row[2] == mail {
			mailexist = true
			break
		}
	}
	return namexist, mailexist
}

func alterpsw(filename,mail string)  {
	defer f.Close()
	f, err = os.OpenFile(filename, os.O_RDWR|os.O_APPEND, 0666)

	if err != nil {
		println("can not open the file ,err is %+v", err)
	}
	//
	r := csv.NewReader(f)
	for {
		row, err := r.Read()
		if err != nil && err != io.EOF {
			println("can not read ,the err is %+v", err)
		}
		if err == io.EOF {
			break
		}
		if row[2]==mail{
			SendMailpsw(mail,row[1])
			break
		}
	}
	
}

/*
	filename:绝对地址,order:命令，add or search or check or whatpsw
*/
func FileOrder(filename, order, name, psw, mail string) (flagname, flagmail bool) {
	flagname, flagmail = false, false
	if !checkFileIsExist(filename) {
		os.Create(filename)

	}
	if order == "search" {
		flag := readfile(filename, name, psw)
		return flag, false
	} else if order == "add" {
		flag := addfile(filename, name, psw, mail)
		return flag, false
	} else if order == "check" {
		flagname, flagmail = isexist(filename, name, mail)
		return flagname, flagmail
	}else if order == "whatpsw"{
		alterpsw(filename,mail)
	}
	return false, false
}



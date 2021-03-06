package main

import (
	"bufio"
	"log"
	"os"
	"path/filepath"
	"strconv"

	gologme "github.com/erasche/gologme/util"
)

func importKeys(t *gologme.Golog, uid int, logDir string) {
	tx, err := t.Db.Begin()
	if err != nil {
		log.Fatal(err)
	}
	stmt, err := tx.Prepare("insert into keyLogs (uid, time, count) values (?, ?, ?)")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	files, err := filepath.Glob(logDir + "/keyfreq*.txt")
	for _, file := range files {
		handle, err := os.Open(file)
		if err != nil {
			log.Fatal(err)
		}
		defer handle.Close()

		scanner := bufio.NewScanner(handle)
		for scanner.Scan() {
			t := scanner.Text()
			if len(t) > 11 {
				unixtime := t[0:10]
				count := t[11:]
				i, err := strconv.ParseInt(count, 10, 64)
				if len(unixtime) > 0 && err == nil {
					_, err := stmt.Exec(uid, unixtime, i)
					if err != nil {
						log.Fatal(err)
					}
				}
			}
		}
	}
	tx.Commit()
}

func importWindows(t *gologme.Golog, uid int, logDir string) {
	tx, err := t.Db.Begin()
	if err != nil {
		log.Fatal(err)
	}
	stmt, err := tx.Prepare("insert into windowLogs (uid, time, name) values (?, ?, ?)")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	files, err := filepath.Glob(logDir + "/window*.txt")
	for _, file := range files {
		handle, err := os.Open(file)
		if err != nil {
			log.Fatal(err)
		}
		defer handle.Close()

		scanner := bufio.NewScanner(handle)
		for scanner.Scan() {
			t := scanner.Text()
			if len(t) > 10 {
				unixtime := t[0:10]
				var title string
				if len(t) > 11 {
					title = t[11:]
				} else {
					title = ""
				}
				_, err := stmt.Exec(uid, unixtime, title)
				if err != nil {
					log.Fatal(err)
				}
			}
		}
	}
	tx.Commit()
}

func importNotes(t *gologme.Golog, uid int, logDir string) {
	tx, err := t.Db.Begin()
	if err != nil {
		log.Fatal(err)
	}
	stmt, err := tx.Prepare("insert into notes (uid, time, type, contents) values (?, ?, ?, ?)")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	files, err := filepath.Glob(logDir + "/notes*.txt")
	for _, file := range files {
		handle, err := os.Open(file)
		if err != nil {
			log.Fatal(err)
		}
		defer handle.Close()

		scanner := bufio.NewScanner(handle)
		for scanner.Scan() {
			t := scanner.Text()
			if len(t) > 12 {
				unixtime := t[0:10]
				title := t[11:]
				_, err := stmt.Exec(uid, unixtime, gologme.NOTE_TYPE, title)
				if err != nil {
					log.Fatal(err)
				}
			}
		}
	}
	tx.Commit()
}

func importBlogs(t *gologme.Golog, uid int, logDir string) {
	tx, err := t.Db.Begin()
	if err != nil {
		log.Fatal(err)
	}
	stmt, err := tx.Prepare("insert into notes (uid, time, type, contents) values (?, ?, ?, ?)")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	files, err := filepath.Glob(logDir + "/blog*.txt")
	for _, file := range files {
		handle, err := os.Open(file)
		if err != nil {
			log.Fatal(err)
		}
		defer handle.Close()

		scanner := bufio.NewScanner(handle)
		for scanner.Scan() {
			t := scanner.Text()

			unixtime, err := strconv.ParseInt(
				file[len(file)-14:len(file)-4], 10, 64)
			if err != nil {
				log.Fatal(err)
			}
			_, err = stmt.Exec(uid, unixtime, gologme.BLOG_TYPE, t)
			if err != nil {
				log.Fatal(err)
			}
		}
	}
	tx.Commit()
}

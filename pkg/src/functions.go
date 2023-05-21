package src

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"math/rand"
	"net/url"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789_"

type DB struct {
	db *sqlx.DB
}

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
}

func Shorting() string {
	result := make([]byte, 10)
	for i := range result {
		result[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(result)
}

func IsValidUrl(token string) bool {
	_, err := url.ParseRequestURI(token)
	if err != nil {
		return false
	}
	u, err := url.Parse(token)
	if err != nil || u.Host == "" {
		return false
	}
	return true
}

func (d *DB) NewPostgresDB() error {
	cfg := Config{
		Host:     "172.23.0.2",
		Port:     "5432",
		Username: "postgres",
		Password: "mypassword",
		DBName:   "postgres",
		SSLMode:  "disable",
	}
	db, err := sqlx.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.Username, cfg.DBName, cfg.Password, cfg.SSLMode))
	if err != nil {
		return err
	}
	err = db.Ping()
	if err != nil {
		return err
	}
	d.db = db
	return nil
}

func (d *DB) InsertNewLink(link string, shortLink string) error {
	_, err := d.db.Exec("insert into links (original_link, short_link) values ($1, $2)", link, shortLink)
	if err != nil {
		return err
	}
	return nil
}

func (d *DB) GetShortLink(link string) (string, error) {
	res, err := d.db.Query("select short_link from links where original_link = $1", link)
	if err != nil {
		return "", err
	}
	var result string
	for res.Next() {
		if err := res.Scan(&result); err != nil {
			return "", err
		}
	}
	return result, nil
}

func (d *DB) GetOriginalLink(link string) (string, error) {
	res, err := d.db.Query("select original_link from links where short_link = $1", link)
	if err != nil {
		return "", err
	}
	var result string
	for res.Next() {
		if err := res.Scan(&result); err != nil {
			return "", err
		}
	}
	return result, nil
}

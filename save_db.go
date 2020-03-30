package main

import (
	"github.com/gomodule/redigo/redis"
	"time"
)

func saveDB(conn redis.Conn, topic string, date string) error {
	_, err := conn.Do("set", topic, date)
	return err
}

func existTopic(conn redis.Conn, topic string) (int, error) {
	return redis.Int(conn.Do("exists", topic, time.Now().Unix()))
}
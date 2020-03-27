package main

import (
	"github.com/gomodule/redigo/redis"
	"time"
)

func saveDB(topic string, conn redis.Conn) error {
	_, err := conn.Do("set", topic, time.Now().Unix())
	return err
}

func existTopic(topic string, conn redis.Conn) (int, error) {
	return redis.Int(conn.Do("exists", topic, time.Now().Unix()))
}
package logstash

import (
	"net"
	"time"
	"errors"
)

type Logstash struct {
	Host net.TCPAddr
	Connection *net.TCPConn
}

func (l *Logstash) Connect() (*net.TCPConn, error) {
	connection, err := net.DialTCP("tcp", nil, &l.Host)
	if connection != nil {
		l.Connection = connection
		l.Connection.SetKeepAlive(true)
		l.Connection.SetKeepAlivePeriod(time.Duration(5) * time.Second)
	}
	return connection, err
}

func (l *Logstash) Writeln(message []byte) (error) {
	if l.Connection != nil {
		_, err := l.Connection.Write(message)
		if err != nil {
			l.Connection.Close()
			l.Connection = nil
			return err
		} else {
			// Successful write!
			return nil
		}
	}
	return errors.New("TCP Connection is nil.")
}

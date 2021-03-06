// Package irc handles connections to irc servers

package irc

import "../config"

func Connect(conf config.Config, server string, serverToClient chan []byte) {
	conn := Connection{
		Nick: conf.Nick,
		User: conf.Nick,

		Server: server,
		Socket: nil,

		Events: nil,
	}

	conn.Connect(serverToClient)
}

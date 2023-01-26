package reader

import (
	"crypto/sha1"
	"rand"
	"github.com/jackpal/bencode-go"
)

type  torrentfile struct{
	announcer string
	infoHash [20]byte
	PiecesHash [][20]byte
	length int
	name string
}



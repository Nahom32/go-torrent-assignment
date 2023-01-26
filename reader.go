package reader

import (
	"crypto/sha1"
	"log"
	"io"
	"rand"

	bencode "github.com/jackpal/bencode-go"
)

type  torrentfile struct{
	announcer string
	infoHash [20]byte
	PiecesHash [][20]byte
	length int
	name string
}
type bencodeTorrentFile struct{
	announcer string 
	torrent torrentfile 
}
func changetoReadable(r *io.Reader)(interface{}){
	t := bencodeTorrentFile{}
	err := bencode.Unmarshal(*r,&t)
	if err != nil {
		log.Printf("Error caused as a result of unmarshalling")
	}
	return &r
}






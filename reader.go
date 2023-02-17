package reader

import (
	"crypto/sha1"
	
	"io"
	"log"
	"os"

	//"rand"

	bencode "github.com/jackpal/bencode-go"
)
type bencoded struct{
	pieces string
	plength int
	wlength int
	name string
}

type  torrentfile struct{
	announcer string
	infoHash [20]byte
	PiecesHash [][20]byte
	length int
	name string
}
type bencodeTorrentFile struct{
	announcer string 
	torrent bencoded
}


func changetoReadable(r *io.Reader)(*bencodeTorrentFile,error){
	t := bencodeTorrentFile{}
	err := bencode.Unmarshal(*r,&t)
	if err != nil {
		log.Printf("Error caused as a result of unmarshalling")
	}
	return &t,nil;
}

// func parseFile(t bencodeTorrentFile)(interface{}){
// 	var reader io.Reader
// 	val := ()

// }
func hasher(val string)([][20]byte,error){
	var fileHolder [][20]byte
	x,err:= os.ReadFile(val)
	if err != nil{
		return nil,err
	}
	for i,_ := range x{
		if i + 20 < len(fileHolder){
			fileHolder= append(fileHolder,sha1.Sum(x[i:i+20]))
		}else {
			fileHolder= append(fileHolder,sha1.Sum(x[i:i+20]))
		}
		fileHolder= append(fileHolder,sha1.Sum(x[i:len(val)-1]))
	}
	
	return fileHolder,nil
}






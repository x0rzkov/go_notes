package main
import(
	"net"
	"encoding/gob"
	"log"
)

func synch_client() {
	conn, err := net.Dial("tcp", "localhost:" + SYNCH_PORT)
	if err != nil {log.Fatal("Error connecting to server ", err)}
	defer conn.Close()
	msg := Message{} // init to empty struct
	enc := gob.NewEncoder(conn)
	dec := gob.NewDecoder(conn)
    
	// Send a message
	sendMsg(enc, Message{Type: "WhoAreYou", Param: "", NoteChg: NoteChange{}})
	rcxMsg(dec, &msg) // Decode the response

    if msg.Type == "WhoIAm" {
        peer_id := msg.Param
        println("The server's database id is", peer_id)
    } else {
        println("Peer does not respond to request for database id"); return
    }

//	sendMsg(enc, Message{Type: "ReturnChangeset", Param: "", NoteChg: NoteChange{Title: "We are really talking now", Description: "Just a test"}})
//	rcxMsg(dec, &msg)

	// Send Hangup
	sendMsg(enc, Message{Type: "Hangup", Param: "", NoteChg: NoteChange{}})
	println("Client done")
}

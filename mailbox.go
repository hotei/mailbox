// mailbox.go (c) 2010-2015 David Rook

// mbx.BulkMail can be 'freed', but won't be released to OS due to how gc works.
// 		to reduce RAM usage you could read mailbox in buffered sections
//		not important for now with 16 GB ram.

// Read mailbox in Eudora/Sendmail-BSD/Thunderbird format
package mailbox

import (
	"bytes"
	//"flag"
	"fmt"
	"io/ioutil"
	//"log"
	"os"
	//	"regexp"
	"strings"
	//"time"
)

const (
	Unknown = iota
	Eudora
	Netscape
	Thunderbird
	Sendmail
)

var MailBoxTypeStrs = []string{
	"Unknown",
	"Eudora",
	"Netscape",
	"Thunderbird",
	"Sendmail",
}

type MailMsg struct {
	MsgId   int
	Size    int
	Date    string // time.Time
	Subject string
	From    string
	To      string
	Raw     []string // mix of header/body/mime
}

type MailBox struct {
	FileName    string     // filename to load/save
	MailBoxType int8       //
	MsgSepBytes []byte     // string that indicates new msg start
	FileSize    int64      // size of entire mbx
	NumMsgs     int64      // count of individual msgs
	BulkMail    []byte     // slice of length FileSize
	MsgArray    []*MailMsg // pointers to messages
}

func NewMailBox(mbxname string) *MailBox {
	mbx := new(MailBox)
	f, err := os.Open(mbxname)
	if err != nil {
		fmt.Printf("can't open file %s err=%v\n", mbxname, err)
		os.Exit(1)
	}
	Verbose.Printf("Opened %s without error\n", mbxname)
	mbx.FileName = mbxname

	Verbose.Printf("Reading all of %s at once ...\n", mbxname)
	mbx.BulkMail, err = ioutil.ReadAll(f)
	if err != nil {
		fmt.Printf("can't read file %s err=%v\n", mbxname, err)
		os.Exit(1)
	}
	mbx.FileSize = int64(len(mbx.BulkMail)) // safe ? BUG ?
	Verbose.Printf("Read %s  %d bytes\n", mbx.FileName, mbx.FileSize)
	rv := mbx.determineType()
	if rv == Unknown {
		fmt.Printf("Cant determine mbox type of %s", mbx.FileName) // BUG to panic ?
		os.Exit(-1)
	}
	mbx.MsgArray = make([]*MailMsg, 0, 100)
	mbx.splitIntoMsgs()
	return mbx
}

func (mbx *MailBox) determineType() int {
	msgSepStr := ""
	msgSepBytes := []byte{}

	ctNewLines := bytes.Count(mbx.BulkMail, []byte{'\n'})
	fmt.Printf("found %d newlines\n", ctNewLines)
	ctCarriageRtns := bytes.Count(mbx.BulkMail, []byte{'\r'})
	fmt.Printf("found %d carriage returns\n", ctCarriageRtns)
	if ctCarriageRtns > 0 {
		fmt.Printf("Found %d NL and %d CR\n", ctNewLines, ctCarriageRtns)
		fmt.Printf("You need to use 'tr' to fix this first\n")
		fmt.Printf("Maybe tr -d \"\\r\" < mail > mail.mbx\n\n")
		fmt.Printf("or tr \"\\r\" \"\\n\" < mail > mail.mbx\n\n")
		fmt.Printf("if NL > 0 and CR > 0  --> MSDOS\n")
		fmt.Printf("if only CR            --> Apple\n")
		fmt.Printf("if only NL            --> Unix\n")
		os.Exit(-1)
	}

	// try Eudora
	msgSepStr = "From ???@??? "
	msgSepBytes = []byte(msgSepStr)
	mbx.NumMsgs = int64(bytes.Count(mbx.BulkMail, msgSepBytes))
	fmt.Printf("If Eudora then we have %d msgs\n", mbx.NumMsgs)
	if mbx.NumMsgs > 0 {
		fmt.Printf("\tAvg msg size = %d\n", mbx.FileSize/mbx.NumMsgs)
		Verbose.Printf("Eudora format mailbox found\n")
		mbx.MailBoxType = Eudora
		mbx.MsgSepBytes = msgSepBytes
		return Eudora
	}

	// try tbird
	msgSepStr = "From - "
	msgSepBytes = []byte(msgSepStr)
	mbx.NumMsgs = int64(bytes.Count(mbx.BulkMail, msgSepBytes))
	fmt.Printf("If thunderbird then we have %d msgs\n", mbx.NumMsgs)
	if mbx.NumMsgs > 0 {
		fmt.Printf("\tAvg msg size = %d bytes\n", mbx.FileSize/mbx.NumMsgs)
		Verbose.Printf("Thunderbird format mailbox found\n")
		mbx.MailBoxType = Thunderbird
		mbx.MsgSepBytes = msgSepBytes
		return Thunderbird
	}

	// try BSD (sendmail)
	msgSepStr = "From "
	msgSepBytes = []byte(msgSepStr)
	mbx.NumMsgs = int64(bytes.Count(mbx.BulkMail, msgSepBytes))
	fmt.Printf("If BSD (sendmail) then we have %d msgs\n", mbx.NumMsgs)
	if mbx.NumMsgs > 0 {
		fmt.Printf("\tAvg msg size = %d\n", mbx.FileSize/mbx.NumMsgs)
		Verbose.Printf("Sendmail format mailbox found\n")
		mbx.MailBoxType = Sendmail
		mbx.MsgSepBytes = msgSepBytes
		return Sendmail
	}

	// out of options0
	return Unknown
}

// returns number of Msg structs created
func (mbx *MailBox) splitIntoMsgs() int {
	// msgs[][]byte
	msgs := bytes.Split(mbx.BulkMail, mbx.MsgSepBytes)
	MsgCt := len(msgs)
	Verbose.Printf("MsgCt = %d\n", MsgCt)
	for i := 0; i < MsgCt; i++ {
		if len(msgs[i]) <= 0 {
			continue
		}
		Verbose.Printf("length(msg[%d] bytes) = %d\n", i, len(msgs[i]))
	}
	n := 0 // number of message created so far
	for i := 0; i < MsgCt; i++ {
		newMsg := new(MailMsg)
		s := string(msgs[i])
		if len(s) <= 0 {
			continue
		}
		newMsg.Raw = strings.Split(s, "\n")
		n++
		fmt.Printf("Raw portion of msg[%d] has %d lines and %d bytes\n",
			n, len(newMsg.Raw), len(s))
		newMsg.parseSubject()
		newMsg.parseDate()
		newMsg.parseFrom()
		newMsg.parseTo()
		mbx.MsgArray = append(mbx.MsgArray, newMsg)
	}
	return len(mbx.MsgArray)
}

func (h *MailMsg) parseSubject() {
	for _, line := range h.Raw {
		if strings.Contains(line, "Subject: ") {
			h.Subject = line
			break // only first one is useful
		}
	}
}

// date probably best taken from separator line as "Date:" is often
// (but not always) missing from headers in Eudora.

func (h *MailMsg) parseDate() {
	for _, line := range h.Raw {
		if strings.Contains(line, "Date: ") {
			h.Date = line
			break // only first one is useful
		}
	}
}

func (h *MailMsg) parseFrom() {
	for _, line := range h.Raw {
		if strings.Contains(line, "From: ") {
			h.From = line
			break // only first one is useful
		}
	}
}

func (h *MailMsg) parseTo() {
	for _, line := range h.Raw {
		if strings.Contains(line, "To: ") {
			h.To = line
			break // only first one is useful
		}
	}
}

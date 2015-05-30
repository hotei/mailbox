// mailbox_test.go (c) 2010-2015 David Rook

package mailbox

import (
	"fmt"
	"log"
	"testing"
)

func (mbx *MailBox) DumpHeaders() {
	Verbose.Printf("MsgArray len = %d\n", len(mbx.MsgArray))
	if len(mbx.MsgArray) > 0 {
		for i := 0; i < int(mbx.NumMsgs); i++ {
			fmt.Printf("%d %s\n", i, mbx.MsgArray[i].Date)
			fmt.Printf("%d %s\n", i, string(mbx.MsgArray[i].From))
			fmt.Printf("%d %s\n", i, string(mbx.MsgArray[i].To))
			fmt.Printf("%d %s\n", i, string(mbx.MsgArray[i].Subject))
			fmt.Printf("\n")
		}
	}
}

func (mbx *MailBox) DumpSummary() {
	fmt.Printf("MailBox Name = %s\n", mbx.FileName)
	fmt.Printf("        Size = %d\n", mbx.FileSize)
	fmt.Printf("        Type = %s\n", MailBoxTypeStrs[mbx.MailBoxType])
	fmt.Printf("    Num Msgs = %d\n", mbx.NumMsgs)
	fmt.Printf("Avg Msg Size = %d\n", mbx.FileSize/mbx.NumMsgs)
	fmt.Printf("Msg Sep Str  = %q\n", string(mbx.MsgSepBytes))
	fmt.Printf("Msg Array Ct = %d\n", len(mbx.MsgArray))
}

func (mbx *MailBox) DisplayMsg(n int) {
	if mbx.NumMsgs < 1 {
		log.Printf("no messages to displayy\n")
		return
	}
	if n < 0 {
		log.Printf("index of requested message [%d] out of range\n", n)
		return
	}
	if n >= int(mbx.NumMsgs) {
		log.Printf("index of requested message [%d] out of range\n", n)
		return
	}

	for _, line := range mbx.MsgArray[n].Raw {
		fmt.Printf("%s", line)
	}
}

func Test_001(t *testing.T) {

	mbx := NewMailBox("bsd.mbx")
	mbx.DumpSummary()
	mbx.DumpHeaders()

	mbx = NewMailBox("eud.mbx")
	mbx.DumpSummary()
	mbx.DumpHeaders()

	mbx = NewMailBox("tbird.mbx")
	mbx.DumpSummary()
	mbx.DumpHeaders()

	mbx.DisplayMsg(56)
}

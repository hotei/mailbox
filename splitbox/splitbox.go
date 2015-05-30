// splitbox.go (c) 2015 David Rook

// working

package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
	//
	"github.com/hotei/mailbox"
	"github.com/hotei/mdr"
	"github.com/hotei/tempfile"
)

var (
	fl_MailboxName string
	fl_Verbose     bool
)

func init() {
	flag.StringVar(&fl_MailboxName, "mbox", "", "mailbox name")
	flag.BoolVar(&fl_Verbose, "v", false, "verbose mode")
}

func usage() {
	fmt.Println(`usage: splitbox -mbox="MailboxName.mbx"`)
	os.Exit(1)
}

func DumpSummary(mbx *mailbox.MailBox) {
	fmt.Printf("MailBox Name = %s\n", mbx.FileName)
	fmt.Printf("        Size = %d\n", mbx.FileSize)
	fmt.Printf("        Type = %s\n", mailbox.MailBoxTypeStrs[mbx.MailBoxType])
	fmt.Printf("    Num Msgs = %d\n", mbx.NumMsgs)
	fmt.Printf("Avg Msg Size = %d\n", mbx.FileSize/mbx.NumMsgs)
	fmt.Printf("Msg Sep Str  = %q\n", string(mbx.MsgSepBytes))
	fmt.Printf("Msg Array Ct = %d\n", len(mbx.MsgArray))
}

// m is the mailbox filename
func splitOne(m string) {
	mbx := mailbox.NewMailBox(m)
	DumpSummary(mbx)
	if mbx.NumMsgs <= 0 {
		fmt.Printf("Mailbox %s appears to be empty\n", fl_MailboxName)
		return
	}
	for _, msg := range mbx.MsgArray {
		f, err := tempfile.New("/home/mdr/tmp", "msg-", ".txt")
		if err != nil {
			fmt.Printf("tempfile create failed %v\n", err)
			os.Exit(-1)
		}
		tmpstr := strings.Join(msg.Raw, "\n")
		r := strings.NewReader(tmpstr)
		_, err = io.Copy(f, r)
		if err != nil {
			fmt.Printf("tempfile %s write failed\n", f.Name())
			os.Exit(-1)
		}
		fmt.Printf("Created file %s\n", f.Name())
		f.Close()
	}
}

func main() {
	flag.Parse()
	fmt.Printf("start splitbox.go on %s\n", fl_MailboxName)
	if fl_Verbose {
		Verbose = true
	} else {
		Verbose = false
	}
	argList := mdr.GetAllArgs()
	fmt.Printf("argList = %v\n", argList)
	if len(argList) <= 0 {
		usage()
	}
	for _, mboxName := range argList {
		splitOne(mboxName)
		fmt.Printf("Done with mailbox named %s\n", mboxName)
	}
}

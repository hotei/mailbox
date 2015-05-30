<center>
# mailbox
</center>

## What Is It?

mailbox is a package that attempts to read several flavors of mailboxes and
deliver them to the user with a consistent format.  The API is intentionally
very sparse. 

### Installation

If you have a working go installation on a Unix-like OS:

> ```go get github.com/hotei/mailbox```

Will copy github.com/hotei/program to the first entry of your $GOPATH

or if go is not installed yet :

> ```cd DestinationDirectory```

> ```git clone https://github.com/hotei/mailbox.git```

### Features

* reads Eudora format mailboxes - retested 2015
* reads BSD format mailboxes - retested 2015
* reads Thunderbird - retested 2015

### Limitations

* <font color="red">Is array of strings the best way to deliver message?</font>
Perhaps an io.Reader is a better choice (consistent with collections).
* <font color="red">RAM usage is 2x mailbox size.  That's avoidable with io.Reader.</font>

### Usage

Typical usage is demonstrated in the test code and examples.

One example is a utility called
splitbox will decompose a mailbox into individual message files:

```
find $HOME -type f -name "*.mbx" | splitbox
```






### BUGS

### To-Do

* Essential:
 * TBD
* Nice:
 * TBD
* Nice but no immediate need:
 * TBD

### Change Log

* 2015-05-29 added _splitbox_ program to demonstrate package usage
	* NOTE: tempfile package is not efficient at creating large number of temp files
* 2015-05-27 validated with go 1.4.2
* 2010-xx-xx Started

 
### Resources

* [go language reference] [1] 
* [go standard library package docs] [2]
* [Source for mailbox] [3]

[1]: http://golang.org/ref/spec/ "go reference spec"
[2]: http://golang.org/pkg/ "go package docs"
[3]: http://github.com/hotei/mailbox "github.com/hotei/mailbox"

Comments can be sent to <hotei1352@gmail.com> or to user "hotei" at github.com.

License
-------
The 'mailbox' go package/program is distributed under the Simplified BSD License:

> Copyright (c) 2010-2015 David Rook. All rights reserved.
> 
> Redistribution and use in source and binary forms, with or without modification, are
> permitted provided that the following conditions are met:
> 
>    1. Redistributions of source code must retain the above copyright notice, this list of
>       conditions and the following disclaimer.
> 
>    2. Redistributions in binary form must reproduce the above copyright notice, this list
>       of conditions and the following disclaimer in the documentation and/or other materials
>       provided with the distribution.
> 
> THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDER ``AS IS'' AND ANY EXPRESS OR IMPLIED
> WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND
> FITNESS FOR A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL <COPYRIGHT HOLDER> OR
> CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR
> CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR
> SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON
> ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT (INCLUDING
> NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF THIS SOFTWARE, EVEN IF
> ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.

Documentation (c) 2015 David Rook 

// EOF README-mailreader.md  (this is a markdown document and tested OK with blackfriday)
# gotftp

Go TFTP client

## TFTP Background

This is a very simple implementation of a TFTP server.
TFTP is a very simple protocol to transfer files.
It is from this that the name comes, Trivial File Transfer Protocol or TFTP.

TFTP is defined in [RFC 1350](https://tools.ietf.org/html/rfc1350)

## Usage

```
  -addr string
    	Server address (default "localhost:69")
  -file string
    	Name of the file on server (default "<filename>")
  -mode string
    	Transfer mode: 'octet' or 'netascii' (default "octet")
  -op string
    	What to do: download or upload file (default "<get|put>")
  -path string
    	Local file path (default ".")
```

`-addr` is the address/port that the system will connect to.

`-file` is the filename to be transferred

`-mode` is the mode using either ascii or octer (formerly binary) reference in Section 1 Paragraph 2 of RFC 1350 linked above.

`-op` is the transfer mode `get` downloads, `put` uploads

`-path` is the directory where the files that need to be transferred are stored.

### Example

We want to download a firmware image for a router. The file is called `newfirmware.bin` and we are going to store it in our `downloads` folder. Our computer and user has the permissions to run on privileged ports (<1024).

Change directory to where the file is to be stored.

```
cd downloads
```

List the files to make sure the file is stored.

```
ls -l

total 0
```

Download the file from the tftp server.

```
gotftp -addr tftpserver:1069 -file newfirmware.bin -op get
1678336 bytes recieved
```

When the client downloads the file we can that the transfer happened.



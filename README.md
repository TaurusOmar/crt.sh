# crt.go
Subdomain Enumeration with crt.go | This Go script simplifies the process of efficiently saving and analyzing subdomain output from the crt.sh website.

## Automated Subdomain Enumeration site => crt.sh

![image](https://github.com/TaurusOmar/crt.sh/blob/main/crt.gif?raw=true)


This script automates the process of enumerating and scanning subdomains based on data extraction from crt.sh, providing detailed results and saving them in a directory. The main objective of this tool is to simplify and expedite the process of discovering subdomains and their related information for a given domain.

## Installation

If you have Go installed and configured (i.e. with $GOPATH/bin in your $PATH):

```
go install github.com/TaurusOmar/crt.sh@latest
```

## Usage

```
crt hackerone.com
```

## Output

```
/UserHomeDir/result_directory/domain.com.crt.txt
```

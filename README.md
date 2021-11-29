# netuniq

Netuniq accepts a list of hosts, CIDR ranges, and IPs on stdin and deduplicates them. It will also filter out any domains that do not have an A record associated with them.

## Usage

```sh
$ cat test.txt                                                                                                                                                master
93.184.216.34
93.184.216.32/27
example.com
dsafasdfdsafsdafsadfsadfsadfasdfasdfasdfasdasdf.com
4.2.2.2
example.com
93.184.216.34

$ cat test.txt | ./netuniq                                                                                                                                    master
example.com
93.184.216.32/27
4.2.2.2

$ cat test.txt | ./netuniq -json
[
    {
        "ips": [
            "93.184.216.34"
        ],
        "hostname": "example.com"
    },
    {
        "subnet": "93.184.216.32/27"
    },
    {
        "ips": [
            "4.2.2.2"
        ]
    }
]
```

## Install

```sh
go install github.com/leesoh/netuniq/cmd/netuniq@latest
```


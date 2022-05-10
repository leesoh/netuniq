# netuniq

Netuniq accepts a list of hosts, CIDR ranges, and IPs on stdin and deduplicates them, preferring hostnames. It will also filter out any domains that do not have an A record associated with them.

## Usage

```sh
$ cat testdata/targets.txt
45.33.32.156
scanme.nmap.org
45.33.32.0/24
93.184.216.34
93.184.216.32/27
example.com
dsafasdfdsafsdafsadfsadfsadfasdfasdfasdfasdasdf.com
4.2.2.2
example.com
93.184.216.34

$ cat testdata/targets.txt| netuniq
scanme.nmap.org
example.com
45.33.32.0/24
93.184.216.32/27
4.2.2.2

$ cat testdata/targets.txt| netuniq -json
[
    {
        "ips": [
            "93.184.216.34"
        ],
        "hostname": "example.com"
    },
    {
        "ips": [
            "45.33.32.156"
        ],
        "hostname": "scanme.nmap.org"
    },
    {
        "subnet": "45.33.32.0/24"
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


# sumcheck

Quick & dirty checksum validator for files, written in Go.

## Installation

### Using `curl`

```bash
curl -s https://raw.githubusercontent.com/sby051/sumcheck/main/sumcheck -o /usr/local/bin/sumcheck && chmod +x /usr/local/bin/sumcheck
```

### Using `wget`

```bash
wget https://raw.githubusercontent.com/sby051/sumcheck/main/sumcheck -O /usr/local/bin/sumcheck && chmod +x /usr/local/bin/sumcheck
```

### Building from source

```bash
git clone https://github.com/sby051/sumcheck.git
cd sumcheck
go build
mv ./sumcheck /usr/local/bin
```

## Usage

```bash
sumcheck -f <file> -s <checksum> -a <algorithm_name>
```

...for example:

```bash
sumcheck -f file.txt -s 1234567890 -a sha256
```

## License

MIT

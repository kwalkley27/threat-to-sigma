# Threat to Sigma

Threat to Sigma is a tool that scrapes open-source threat intelligence feeds and converts them into Sigma rules written to stdout.

## Installation

To install Threat to Sigma, you will need to have Go installed on your system. You can then clone the repository and build the project:

```bash
git clone https://github.com/user/threat-to-sigma.git
cd threat-to-sigma
go build ./cmd/app
```

## Usage

To use Threat to Sigma, you can run the following command:

```bash
./app
```

This will scrape the supported threat intelligence feeds and generate Sigma rules in stdout.

## Supported Feeds

The following threat intelligence feeds are currently supported:

*   Spamhaus

## License

Threat to Sigma is licensed under the Apache License.
# Threat to Sigma

Threat to Sigma is a tool that scrapes open-source threat intelligence feeds and converts them into Sigma rules using Google's Gemini generative AI.

## Features

*   Retrieves threat intelligence from open-source feeds.
*   Converts indicators of compromise (IOCs) into Sigma rules.
*   Concurrent processing of IOCs for faster results.
*   Flexible configuration via file and environment variables.
*   Unit tests for core functionality.

## Installation

1.  **Clone the repository:**
    ```bash
    git clone https://github.com/user/threat-to-sigma.git
    cd threat-to-sigma
    ```

2.  **Install dependencies:**
    ```bash
    go mod tidy
    ```

3.  **Build the application:**
    ```bash
    go build -o threat-to-sigma ./cmd/app
    ```

## Configuration

Configuration is handled by `viper` and can be set in `config.yaml` or via environment variables.

**1. Configuration File:**

Create a `config.yaml` file in the root of the project with the following structure:

```yaml
gemini_api_key: "your-api-key-here"
spamhaus_feed_url: "https://www.spamhaus.org/drop/drop.txt"
feed_limit: 10
model_name: "gemini-1.5-pro"
max_concurrency: 5
```

**2. Environment Variables:**

You can override any of the values from the configuration file by setting an environment variable. The environment variable name is the uppercase version of the config key.

For example, to set the Gemini API key, you would use:

```bash
export GEMINI_API_KEY="your-api-key-here"
```

**Note on Precedence:** Environment variables will always take precedence over values in the `config.yaml` file. This is useful for providing sensitive credentials like API keys.

## Usage

Ensure your configuration is set up correctly, then run the application:

```bash
./threat-to-sigma
```

The application will then scrape the configured threat intelligence feeds, generate Sigma rules, and print them to standard output.

### Example Output

```yaml
title: Suspicious Network Connection to C2 Server
id: 5b2e7f7a-3b5d-4d7e-8f3a-2b1c0d1e4f9a
status: experimental
description: Detects network connections to a known command and control (C2) server.
author: Gemini
date: 2023/10/27
logsource:
  category: network_connection
  product: zeek
detection:
  selection:
    - dst_ip: "1.0.0.0/24"
  condition: selection
falsepositives:
  - Legitimate traffic to the IP address.
level: high
```

## Testing

Unit tests have been added for the `feeds` and `genai` packages. To run the tests, use the following command:

```bash
go test ./...
```

## Supported Feeds

*   Spamhaus DROP list

## Future Improvements

*   Add support for more threat intelligence feeds.
*   Implement command-line flags for configuration.
*   Add integration tests.
*   Allow users to specify a custom output file for the generated Sigma rules.

## License

Threat to Sigma is licensed under the Apache License.

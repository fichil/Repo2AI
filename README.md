# Repo2AI

> Convert repositories into AI-ready context packs for ChatGPT, Claude, Cursor and coding assistants.

Repo2AI helps developers transform source code repositories into clean, structured, size-limited files that AI tools can understand efficiently.

## Why Repo2AI?

Large repositories are difficult to paste into AI tools.

Problems:

- Context too large
- Token limits exceeded
- Files messy and noisy
- AI cannot understand project structure clearly

Repo2AI solves this by converting repositories into organized AI-ready packs.

## Features

- Scan local repositories
- Split output into 10MB / custom size chunks
- Generate Markdown and TXT context packs
- Preserve project structure
- Ignore useless files automatically
- Optimized for Java projects (initial support)
- CLI mode
- GUI mode (planned)

## Example Output

```bash
repo2ai scan ./my-project
```

Output:

```
output/
├── project-summary.md
├── controllers_01.md
├── services_01.md
├── entities_01.md
├── sql_01.md
└── manifest.json
```

## Why not Repomix?

Repo2AI focuses on:

- Better chunk control
- Cleaner enterprise project parsing
- AI feeding workflow
- Java-first optimization
- GUI for non-terminal users

## Installation

### Windows

Download from Releases.

### Go Install

```
go install github.com/fichil/Repo2AI@latest
```

## Usage

### CLI

```
repo2ai scan ./demo
repo2ai scan ./demo --max-size=10mb
repo2ai scan ./demo --format=txt
```

### GUI

```
repo2ai gui
```

## Roadmap

### v0.1

#### Features

- Scan Java projects

- Generate AI-ready context packs

- Generate project-summary.md

- Generate manifest.json

- Split context packs by max size

- Support Markdown and TXT output

- Clean output before generation

#### Usage

```bash
repo2ai scan ./examples/demo-java
repo2ai scan ./examples/demo-java --format=txt
repo2ai scan ./examples/demo-java --max-size=5mb
repo2ai scan ./examples/demo-java --format=txt --max-size=5mb
```

## Supported Parameters

| Parameter  | Description              |
| ---------- | ------------------------ |
| --max-size | Max size per output file |
| --format   | Output format (md / txt) |

### v0.2

- GUI desktop app
- ZIP export
- Better summaries

### v0.3

- Spring Boot deep parsing
- Multi-language support

## Tech Stack

- Go
- JavaParser
- Fyne GUI

## Contributing

PRs are welcome.

## Star History

If this project helps you, give it a star.

## License

MIT

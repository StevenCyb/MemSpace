# MemSpace
This a simple fun CLI project that provides a tool to get the memory usage of file and directories as a tree.

```bash
$ MemSpace -r -m
Size:      926.35GB
Free:      208.63GB - 22.52%
Available: 208.63GB - 22.52%
Used:      717.72GB - 77.48%
📁. [153.61KB]
│-📄.DS_Store [8.00KB]
│-📄.editorconfig [355.00B]
│-📁.git [92.40KB]
│ │-📄FETCH_HEAD [89.00B]
│ ...
│ │-📄description [73.00B]
│ │-📁hooks [25.22KB]
│ │ │-📄applypatch-msg.sample [478.00B]
│ ...
│-📄.gitignore [325.00B]
│-📄.golangci.yml [16.30KB]
│-📁.vscode [135.00B]
│ └-📄settings.json [135.00B]
│-📄LICENSE [1.46KB]
│-📄README.md [590.00B]
│-📄go.mod [450.00B]
│-📄go.sum [2.01KB]
│-📁internal [30.77KB]
│ ...
└-📄main.go [876.00B]
```

## Installation
```bash
go install github.com/StevenCyb/MemSpace@latest
```

## Usage
```bash
Usage:
  main [OPTIONS]

Application Options:
  -p, --path=      The base path to start scanning from (default: .)
  -d, --dir        Only show directories
  -r, --recursive  Show files (and directories) Recursively
  -e, --depth=     The depth of recursion (default: -1)
  -t, --threshold= Show only files or directories larger than the threshold
  -m, --memory     Show driver memory

Help Options:
  -h, --help       Show this help message
```

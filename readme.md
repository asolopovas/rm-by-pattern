## Overview

**rm-by-pattern** is a powerful tool designed for managing files by identifying and organizing them based on filename patterns. It scans directories recursively, extracts patterns from filenames using regular expressions, and allows users to selectively remove files matching specific criteria. This document aims to guide users through the process of utilizing **rm-by-pattern** effectively.

To begin, **rm-by-pattern** can analyze all files within a directory (and its subdirectories) to identify and group filename patterns. This is particularly useful for handling files in a structured manner, such as images or documents with size dimensions or date stamps in their names.

For instance, running the following command:

```
rm-by-pattern . -r '\d+x\d+'
```

will examine each file in the current directory (and all subdirectories) to find matches for the regular expression `\d+x\d+`, which targets patterns resembling dimensions (e.g., "600x400").

### Example

Given a set of files like:

- "0001-name-1-608x434-addition-text.jpg"
- "0001-name-2-400x434-addition-text.jpg"
- "0001-name-3-200x434-addition-text.jpg"
- "0001-name-4-608x434-addition-text.jpg"
- "0001-name-5-608x434-addition-text.png"

**rm-by-pattern** generates a YAML configuration file reflecting the identified patterns and file extensions, as shown below:

```yaml
patterns:
    - 608x434
    - 400x434
    - 200x434
extentions:
    - jpg
    - png
rmPatterns:
    # User specifies patterns to remove here
rmExtentions:
    # User specifies extensions to remove here
```

## Configuration for File Removal

After generating the initial configuration, users can specify which patterns and file extensions to target for removal. This step involves manually editing the YAML file to add desired patterns and extensions under `rmPatterns` and `rmExtentions`.

### Customizing Removal Criteria

For example, to remove all ".jpg" files with the "200x434" pattern, update the YAML configuration as follows:

```yaml
patterns:
    - 608x434
    - 400x434
    - 200x434
extentions:
    - jpg
    - png
rmPatterns:
    - 200x434
rmExtentions:
    - jpg
```

### Executing File Removal
Finally, to execute the file removal process based on the specified criteria in the YAML configuration file (e.g., `pattern.yml`), run:

```
rm-by-pattern . -c pattern.yml
```

This command will remove all files matching the user-defined patterns and extensions from the directory recursively, effectively cleaning up the file system as per user preferences.

## Prerequisites
- Go 1.20 or later.
- Access to both local and remote servers.
- SSH access to the remote server.

## Installation
Using `go install`:
```
go install github.com/asolopovas/rm-by-pattern@latest
```

Clone the repo and build the app:
```
git clone https://github.com/asolopovas/rm-by-pattern.git go build -o $ABSOLUTE_PATH_TO_RMBYPATTERN main.go
```
Add it to your system path or copy executable `sudo cp rm-by-pattern /usr/local/bin/`.

## Usage
**rm-by-pattern** operates on a config file (`pattern.yaml`). Generate a default config file using `-g` flag:
```
rm-by-pattern -g
```

### Get Version:
```
rm-by-pattern -v
```

## Contributing
Open issues, submit pull requests, and share feedback.

## License
MIT License.

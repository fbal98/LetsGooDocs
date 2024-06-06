## LetsGooDocs

<p align="center"><img src="https://github.com/fbal98/LetsGooDocs/assets/55301529/c7e883c3-101f-492b-b0cd-528458fc31b4" width="300px"/></p>

<p align="center">
  <h1>LetsGooDocs</h1>
  This repository contains the source code for LetsGooDocs, a Go-based project that automates the generation of documentation. LetsGooDocs is designed to help developers create comprehensive, well-structured documentation for their projects with minimal effort.
  <br />
</p>

## Table of Contents
- [Implemented Commands](#implemented-commands)
- [Ideas for More Commands](#ideas-for-more-commands)
- [Installation](#installation)
- [Usage](#usage)
- [Contributing](#contributing)
- [TODOs](#todos)

## Implemented Commands
- **start**
  - **Description**: This is the primary command to initiate the documentation generation process.
  - **Usage**: `go run main.go start`
  - **Subcommands**:
    - **api**
      - **Description**: Generates API documentation for your project.
      - **Usage**: `go run main.go start api --path <path>`
      - **Flags**:
        - `--path` or `-p` (Optional): Specify the root directory of the API. Defaults to the current directory (`./`).
      - **Example**:
        ```bash
        go run main.go start api --path ./my-api
        ```
      - **Output**: Creates `Documentation.md` in the specified path with the generated API documentation.

## Ideas for More Commands
- **readme**: Generate README file for a project
- **explain**: Explains the main purposes of a project to make onboarding easier
- **functions**: Lists all functions or classes with their inputs and outputs
- **testing-docs**: Generates documentation for testing procedures and test cases
- **erd**: Creates an Entity-Relationship Diagram (ERD) for the database schema
- **deployment-guide**: Generates a guide for deploying the project to various environments
- **troubleshooting**: Provides common troubleshooting steps and solutions for known issues

## Installation
To get started with LetsGooDocs, clone the repository and install the dependencies:

```bash
git clone https://github.com/fbal98/LetsGooDocs.git
cd LetsGooDocs
go mod tidy
```

## Usage
To run LetsGooDocs, use the following command:

```bash
go run main.go <command> <subcommand> <flags>
```

### Example
Generate API documentation:

```bash
go run main.go start api --path ./path/to/api
```

## TODOs
- Solidify and utilize data chunking
- Implement dynamic chunking and overlapping
- Complete the `GeneratePrompts` function to generate prompts for chunks of content
- Improve error handling and validation

## Contributing
We welcome contributions! Follow these steps to contribute:

1. Fork the repository
2. Create a new branch (`git checkout -b feature-branch`)
3. Make your changes
4. Commit your changes (`git commit -m 'Add new feature'`)
5. Push to the branch (`git push origin feature-branch`)
6. Open a Pull Request



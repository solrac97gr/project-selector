# Project Selector

Project Selector is a command-line tool that allows you to quickly select and open a project from a predefined set of directories. It uses an interactive prompt to let you choose a project and then opens it with a specified command.

## Features

- Scans a predefined set of directories for projects.
- Ignores hidden directories and `.git` folders.
- Provides an interactive prompt to select a project.
- Opens the selected project with a specified command.

## Installation

1. Clone the repository:

    ```sh
    git clone https://github.com/solrac97gr/project-selector.git
    cd project-selector
    ```

2. Install the project:

    ```sh
    make install
    ```

## Configuration

The configuration is managed through a configuration file. If the configuration file is not found, default settings will be used.

`$HOME/.config/project-selector/config.json`

```json
// This is the default config
{
  "cmd": "code",
  "project_dirs": ["Development/work"], // From $HOME directory
  "dirs_to_ignore": ["secret-projects"], // If the path contains this strings, it will be ignored
  "number_of_projects": 10,
  "style": {
    "title": {
      "template": "{{ . | blue | bold }}",
      "icon": "🔎"
    },
    "active": {
      "template": "{{ . | blue | underline | bold}}",
      "icon": "🚀"
    },
    "inactive": {
      "template": "{{ . | cyan }}",
      "icon": "📁"
    }
  }
}
```

## Usage

1. Run the project selector:

    ```sh
    project-selector
    ```

2. Follow the interactive prompt to select a project or use the Search input.

3. The selected project will be opened with the specified command.

## Example

```sh
$ project-selector
Search: my-
Select a Project 🔎 (ctrl+c to cancel)
  🚀 📁 my-project
   📁 my-website
   📁 my-blog
   📁 my-app
   📁 my-API
↓  📁 my-AI
```
## Use your own configuration

1. Create a configuration file in the default path:

    ```sh
    mkdir -p $HOME/.config/project-selector
    touch $HOME/.config/project-selector/config.json
    ```
2. Edit the configuration file with your settings.

      ```json
      {
        "cmd": "zed", // Open the project with the zed editor instead of code
        "project_dirs": ["Development/work"],
        "dirs_to_ignore": ["secret-projects"],
        "number_of_projects": 10,
        "style": {
          "title": {
            "template": "{{ . | blue | bold }}",
            "icon": "🔎"
          },
          "active": {
            "template": "{{ . | blue | underline | bold}}",
            "icon": "🚀"
          },
          "inactive": {
            "template": "{{ . | cyan }}",
            "icon": "📁"
          }
        }
      }
      ```

3. Run the project selector for loading the new configuration:

    ```sh
    project-selector
    ```

## Contributing

Contributions are welcome! Please open an issue or submit a pull request for any improvements or bug fixes.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

## Acknowledgements

- [Promptui](https://github.com/manifoldco/promptui) for the interactive prompt.
- [Go](https://golang.org) for the programming language.

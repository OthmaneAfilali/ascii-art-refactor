# ASCII Art Generator

The ASCII Art Generator is a Go program that converts text into an ASCII art representation. It provides users the ability to transform a string into artistic text using predefined banner styles. This tool can be useful for creating stylized text for terminal outputs, websites, or other creative applications.

## Features

- Converts any string of text into ASCII art
- Supports multiple banner styles (such as standard, shadow, and thinkertoy)
- Handles spaces, newlines, and special characters in input strings
- Extensible to add more features in the future, such as color formatting, alignment and file output

## Installation

To run the program, clone the repository and use Go to build and execute the project.


```
git clone https://github.com/yourusername/ascii-art-generator.git
cd ascii-art-generator
go run .
```


## Usage

You can use the program by passing a string as an argument. The program will output the string in ASCII art using the default banner style.

Example:

```
go run . "Hello, World!"
```

This will print the ASCII art representation of the string "Hello, World!" using the default banner (standard).

## Banners

The program comes with several banner files which contain different styles of ASCII art. These banners are located in the `banners` folder and include:

- Standard: The default banner style
- Shadow: A shadowed version of ASCII characters
- Thinkertoy: A playful, toy-like font

To specify a different banner, pass the banner name as an additional argument.

Example:

```
go run . "ASCII Art" shadow
```

This will print the string "ASCII Art" in the shadow banner style.

## Input Handling

The program can handle various types of input including:

- Letters: Any alphabetical characters (A-Z, a-z)
- Numbers: Digits (0-9)
- Spaces: Spaces between words are preserved in the ASCII output
- Special Characters: Symbols like !@#$%^&*() are supported
- Newlines: The program can process multiple lines of input, maintaining formatting as needed

Example:

```
go run . "Hello\nWorld"
```

This will print "Hello" and "World" on separate lines in ASCII art.

## Error Handling

If no arguments are provided or more than one string argument is given, the program will display usage instructions.
If the input contains characters outside the supported ASCII range, the program will also show an error message and example usage.

Usage Instructions Example:

```
go run .
# Outputs:
# Usage: go run . [STRING] [BANNER]
# Example: go run . "Hello" standard
```


## Roadmap

Future versions of the ASCII Art Generator may include:

- Color Options: Add the ability to colorize specific parts of the ASCII output
- File Output: Allow users to write ASCII art output to a file
- More Banner Styles: Expand the selection of banner templates
- Alignment: Additional support for advanced text alignment (left, right, center, justify)

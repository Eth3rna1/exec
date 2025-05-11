## Exec Python CLI Tool

## About
`exec` is a CLI tool that is made to improve the CL experience by offering a placeholder that in turn, the user can use to inject an entry name, dealing with the tedious task of copying and pasting long file paths manually.

**FYI, the token is: `{.}`**

> Before continuing with the installation, make sure you have the `PyInstaller` builder tool to compile the python spec file

## Getting Started
* Clone the repo
```console
git clone https://github.com/Eth3rna1/exec.git
```
* Enter the project
```console
cd exec
```
* Locate the `exec.spec` file and compile it
```console
pyinstaller exec.spec
```
* Binary will be located in `./dist/exec.exe`

## Functionality
* `exec` is flexible with different scenarios, one being with the sole call of exec.
```console
exec
 ```
Subsequently, a UI will appear waiting for user selection of the entry they want to work with; this is especially helpful when the entry has a long name.

```text
1.) a
2.) b
3.) c
4.) file_with_a_long_name.txt

Entry: 4
```

Then a command line will be introduced, letting the user inject their command and use a placeholder token to indicate the entry selected

```text
Command>> move {.} ..
```

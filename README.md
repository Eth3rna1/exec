## Exec Python CLI Tool

## About
`exec` is a CLI tool that is made to improve the CL experience by offering a placeholder that in turn, the user can use to inject an entry name, dealing with the tedious task of copying and pasting long file paths manually.

**FYI, the token is: `{.}`**

> Before continuing with the installation, make sure you have the `PyInstaller` builder tool to compile the python spec file
>
> If you don't have it, run pip.
> ```console
> pip install PyInstaller
> ``` 
> Or
> ```console
> python.exe -m pip install PyInstaller
> ```

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
`exec` is flexible with multiple scenarios, them being:

---
### 1. A sole call of `exec`.
```console
exec
 ```
Subsequently, a UI will appear waiting for user selection of the entry they want to work with; this is especially helpful when the entry has a long name.

```text
1.) a
2.) b
3.) c
4.) file_with_a_long_name.txt

Select an entry by its index or value.
Entry: 4
```

Then a command line will be introduced, letting the user inject their command and use a placeholder token to indicate the entry selected

```text
Place `{.}` as a placeholder to replace with the entry
Command>> move {.} ..
```

---
### 2. Calling `exec` along with its placeholder
```console
exec nvim {.}
```
The UI drop box then appears, waiting for user selection on entry

```text
1.) a
2.) b
3.) c
4.) file_with_a_long_name.txt

Select an entry by its index or value.
Entry: 4
```
There is **NO** prompt for command in this scenario since its already implied in the CL where the user already used a placeholder

---
### 3. Injecting the command without placeholder
```console
exec seek -r -l -u
```
Although a placeholder isn't explicitly used, the program is always going to assume that there should be an entry in the command. A UI drop box will appear, waiting for user selection, and will append the entry to the end of the command.

```text
1.) a
2.) b
3.) c
4.) file_with_a_long_name.txt

Select an entry by its index or value.
Entry: 4
```
**NO** command prompt follows in this instance since its already defined
### 4. Scenario: No entries in the current working directory
In this scenario, `exec` cannot work with any entry because they're non-existant, in this case, a new prompt will appear; a prompt that asks to specify a directory to work with.
```text
Dir:
```
Then depending, on what followed when the user called `exec`, it will run any previous defined scenarios.

package functions

import (
	_ "exec/utils"
)

// When the token is already predefined in the command line
//     Example: exec example_command.exe {TOKEN} {TOKEN}
func exec_token_in_cl(dir string, token string, args []string) {
//     selection = entries_drop_box(directory)
//     if selection is None:
//         return 1
//     if sys.platform == "win32":
//         command = ["cmd", "/C"]
//     else:
//         command = []
//     selection = os.path.join(directory, selection)
//     command.extend(args[1:]) # skipping the first exec argument
//     swap_token(command, selection, token)
//     return subprocess.run(command).returncode
	panic("Not implemented yet")
}

// When only `exec` is called in the command line
//     Example: exec
func exec_only_exec_in_cl(dir string, token string) int {
    // selection = entries_drop_box(directory)
    // if selection is None:
    //     return 1
    // selection = os.path.join(directory, selection)
    // print("\nPlace `{.}` as a placeholder to replace with the entry")
    // command = parse(input("Command>> "))
    // swap_token(command, selection, token)
    // return subprocess.run(command).returncode
	panic("Not implemented yet")
}

// Regardless of the cl, there are no entries in the
// current working directory to work with
//
// Asks for directory to work with
func exec_no_entries_in_cwd(args []string, token string) {
//     directory = input("Dir: ")
//     assert os.path.exists(directory), f"`{directory}` does not exist"
//     assert os.path.isdir(directory), f"`{directory}` must be a directory"
//     assert len(
//         os.listdir(directory)
//     ), f"`{directory}` is empty, please select a directory with entries"
//     if len(args) == 1:
//         return exec_only_exec_in_cl(directory, token)
//     if any(i == token for i in args):
//         return exec_token_in_cl(directory, token, args)
//     return exec_append_entry_to_args(directory, token, args)
	panic("Not implemented yet")
}

// When a command given along with calling `exec`, although the token isn't present,
// the program will always assume a token is present, thus, it will open a UI and have
// the client to select and append such entry to the command
//
// Example: exec seek -e -r --use-cache /* Here */     <----
//															 \
//															  |
// The program will have the user select an entry to append  /
func exec_append_entry_to_args(dir string, token string, args []string) {
//     if sys.platform == "win32":
//         command = ["cmd", "/C"]
//     else:
//         command = []
//     command.extend(args[1:]) # skipping the first exec argument
//     selection = entries_drop_box(directory)
//     if selection is None:
//         return 1
//     selection = os.path.join(directory, selection)
//     command.append(selection)
//     return subprocess.run(command).returncode
	panic("Not implemented yet")
}

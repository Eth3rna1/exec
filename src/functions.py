"""
Functions that cover every possible case when dealing with the command
"""
import sys
import subprocess
try:
    from src.utils import parse
    from src.utils import swap_token
    from src.utils import entries_drop_box
except ImportError:
    from .utils import parse
    from .utils import swap_token
    from .utils import entries_drop_box


def exec_only_exec_in_cl(directory: str, token: str) -> int:
    """
    When only `exec` is called in the command line
        Example: exec
    """
    selection = entries_drop_box(directory)
    if selection is None:
        return 1
    command = parse(input("Command>> "))
    swap_token(command, selection, token)
    return subprocess.run(command).returncode

def exec_token_in_cl(directory: str, token: str, args: list[str]) -> int:
    """
    When the token is already predefined in the command line
        Example: exec example_command.exe {TOKEN} {TOKEN}
    """
    selection = entries_drop_box(directory)
    if selection is None:
        return 1
    if sys.platform == "win32":
        command = ["cmd", "/C"]
    else:
        command = []
    command.extend(args)
    swap_token(command, selection, token)
    return subprocess.run(command).returncode

def exec_append_entry_to_args(directory: str, token: str, args: list[str]) -> int:
    if sys.platform == "win32":
        command = ["cmd", "/C"]
    else:
        command = []
    command.extend(args)
    selection = entries_drop_box(directory)
    if selection is None:
        return 1
    command.append(selection)
    return subprocess.run(command).returncode


def exec_no_entries_in_cwd(args: list[str], token: str) -> int:
    """
    Regardless of the cl, there are no entries in the
    current working directory to work with

    Asks for directory to work with
    """
    directory = input("Dir: ")
    assert os.path.exists(directory), f"`{directory}` does not exist"
    assert os.path.isdir(directory), f"`{directory}` must be a directory"
    assert len(os.listdir(directory)), f"`{directory}` is empty, please select a directory with entries"
    if len(args) == 1:
        return exec_only_exec_in_cl(directory, token)
    if any(i == token for i in args):
        return exec_token_in_cl(directory, token, args)
    else:
        return exec_append_entry_to_args(directory, token, args[1:])

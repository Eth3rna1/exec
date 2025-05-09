"""
Functions that cover every possible case when dealing with the command
"""
import subprocess
from .__init__ import TOKEN

def exec_only_exec_in_cl(cmd: list[str]):
    """
    When only `exec` is called in the command line
        Example: exec
    """
    ...

def exec_token_in_cl(cmd: list[str]):
    """
    When the token is already predefined in the command line
        Example: exec example_command.exe {TOKEN} {TOKEN}
    """
    ...

def exec_no_entries_in_cwd(cmd: list[str]):
    """
    Regardless of the cl, there are no entries in the
    current working directory to work with

    Asks for directory to work with
    """
    ...

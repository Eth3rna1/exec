from .utils import execute
from .utils import get_file
from .utils import parse_as_cl

def no_context_given(dir: str, token: str) -> int:
    """
    for when `exec` is called alone
    """
    obj = get_file(dir)
    if obj is None:
        print("No sub entries were found")
        return 1
    print(f"Provide your command and replace the object with `{token}`")
    cmd = input(">> ").replace(token, obj)
    return execute(parse_as_cl(cmd))


def append_obj_to_cmd(dir: str, token: str, cmd: list[str]) -> int:
    """
    For when a command line command has been provided
    """
    obj = get_file(dir)
    if obj is None:
        print("No sub entries were found")
        return 1
    cmd.append(obj)
    return execute(cmd)

def replace_token_in_cmd(dir: str, token: str, cmd: list[str]) -> int:
    """
    For when arguments along with the token have been provided
    in the command line
    """
    obj = get_file(dir)
    if obj is None:
        print("No sub entries were found")
        return 1
    command = []
    for tok in cmd:
        if tok == token:
            command.append(obj)
            continue
        command.append(obj)
    return execute(command)

import os
import sys

try:
    from src.options import Options
except ImportError:
    from .options import Options


def parse(query: str) -> list[str]:
    """
    Parses a string like in a command line
    """
    if sys.platform == "win32":
        cmd = ["cmd", "/C"]
    else:
        cmd = []
    buffer = []
    quoted = False
    for c in query:
        match c:
            case " ":
                if quoted:
                    buffer.append(c)
                    continue
                else:
                    cmd.append("".join(buffer))
                    buffer.clear()
            case '"':
                quoted = not quoted  # flipping the boolean
            case _:
                buffer.append(c)
    if len(buffer) != 0:
        cmd.append("".join(buffer))
    return cmd


def entries_drop_box(directory: str) -> str | None:
    """
    A UI for the user to pick an entry from the current working directory
    """
    entries = os.listdir(directory)
    if len(entries) == 0:
        return
    o = Options(entries)
    print(o.display())
    print("\nSelect an entry by its index or value.")
    while True:
        selection = input("Entry: ")
        result = o.evaluate(selection)
        if result is None and len(selection) == 0:
            return None
        if result is not None:
            return result


def swap_token(args: str, entry: str, token: str):
    for i in range(len(args)):
        if token in args[i]:
            args[i] = args[i].replace(token, entry)

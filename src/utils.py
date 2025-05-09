import os
import sys
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
            case "\"":
                quoted = not quoted # flipping the boolean
            case _:
                buffer.append(c)
    if len(buffer) != 0:
        cmd.append("".join(buffer))
    return cmd


def entries_drop_box() -> str | None:
    """
    A UI for the user to pick an entry from the current working directory
    """
    entries = os.listdir(os.getcwd())
    if len(entries) == 0:
        return
    o = Options(entries)
    print(o.display())
    selection = input("\nSelection: ")
    return o.evaluate(selection)

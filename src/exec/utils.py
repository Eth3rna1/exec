import os
import platform
import subprocess
from .options import Options

def get_file(dir) -> str | None:
    opts = Options(os.listdir(dir))
    if len(opts) == 0:
        return None
    elif len(opts) == 1:
        return opts.hashmap["1"] # returns the only object
    print(opts.display())
    while True:
        entry = opts.eval(input("Object: "))
        if entry is not None:
            if " " in entry:
                # if entry has spaces; add quotations
                entry = '"' + entry + '"'
            return entry

def execute(cmd: list[str]) -> int:
    if platform.system() == "Windows":
        command = ["cmd", "/C"]
    else:
        command = ["bash", "-c"]
    command.extend(cmd)
    return subprocess.run(command).returncode

def parse_as_cl(c: str) -> list[str]:
    cmd = []
    buffer = ""
    quoted = False
    for char in c:
        if char == '"':
            quoted = not quoted # toggles the boolean
            continue
        if char == " ":
            if quoted:
                buffer += char
            else:
                cmd.append(buffer)
                buffer = ""
            continue
        buffer += char
    if len(buffer) != 0:
        cmd.append(buffer)
    return cmd


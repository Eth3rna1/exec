import os
import sys
from .functionalities import no_context_given
from .functionalities import replace_token_in_cmd
from .functionalities import append_obj_to_cmd

class Process:
    def __init__(self, dir: str) -> None:
        self.dir = dir
        self.args = sys.argv

    def dir_has_subobjects(self) -> bool:
        return os.listdir(self.dir) != 0

    def build_and_run(self, token: str) -> int:
        match len(sys.argv):
            case 0:
                return 1
            case 1:
                return no_context_given(self.dir, token)
            case _:
                if token in self.args:
                    return replace_token_in_cmd(self.dir, token, self.args[1:])
                return append_obj_to_cmd(self.dir, token, self.args[1:])

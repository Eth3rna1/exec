import os
import sys
try:
    from src.__init__ import exec_token_in_cl
    from src.__init__ import exec_only_exec_in_cl
    from src.__init__ import exec_no_entries_in_cwd
    from src.__init__ import exec_append_entry_to_args
except ImportError:
    from .src.__init__ import exec_token_in_cl
    from .src.__init__ import exec_only_exec_in_cl
    from .src.__init__ import exec_no_entries_in_cwd
    from .src.__init__ import exec_append_entry_to_args

DIR = os.getcwd()
TOKEN = "{.}"

def main():
    global DIR, TOKEN
    if len(os.listdir(DIR)) == 0:
        return exec_no_entries_in_cwd(sys.argv)
    if len(sys.argv) == 1:
        # No args were provided
        return exec_only_exec_in_cl(DIR, TOKEN)
    elif any(i == TOKEN for i in sys.argv):
        return exec_token_in_cl(DIR, TOKEN, sys.argv[1:])
    else:
        # User provided extra arguments
        return exec_append_entry_to_args(DIR, TOKEN, sys.argv[1:])

if __name__ == "__main__":
    main()

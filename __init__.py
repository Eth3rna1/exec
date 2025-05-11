try:
    from .src.utils import parse
    from .src.functions import exec_token_in_cl
    from .src.functions import exec_only_exec_in_cl
    from .src.functions import exec_no_entries_in_cwd
    from .src.functions import exec_append_entry_to_args
except ImportError:
    from src.utils import parse
    from src.functions import exec_token_in_cl
    from src.functions import exec_only_exec_in_cl
    from src.functions import exec_no_entries_in_cwd
    from src.functions import exec_append_entry_to_args

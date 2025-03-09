import sys
from .__init__ import Process

DIR: str = "."
TOKEN: str = "{.}"

if __name__ == "__main__":
    try:
        pr = Process(DIR)
        if not pr.dir_has_subobjects():
            print("No subobjects were found")
            sys.exit(1)
        sys.exit(pr.build_and_run(TOKEN)) # exits with the sub status code
    except KeyboardInterrupt:
        sys.exit(1)
    except Exception as e:
        print(e)
        sys.exit(1)

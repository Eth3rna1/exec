"""
A class for index selection in the UI
"""

class Options:
    def __init__(self, values: list[str]) -> None:
        assert len(values) != 0, "Options: Value length cannot be 0"
        self.hashmap = {i: v for i, v in enumerate(values, 1)}
    
    def evaluate(self, selection: str) -> str | None:
        for k, v in self.hashmap.items():
            if str(k) == selection:
                return self.hashmap[k]
            if v == selection:
                return v
        return None

    def display(self) -> str:
        buffer = []
        for i in range(1, len(self.hashmap) + 1):
            buffer.append(f"{i}.) {self.hashmap[i]}")
        return "\n".join(buffer)
            

if __name__ == "__main__":
    o = Options(["a", "b"])
    print(o.display())
    selection = input("\nSelection: ")
    result = o.evaluate(selection)
    print(result)

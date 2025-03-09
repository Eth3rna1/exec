
class Options:
    def __init__(self, values: list[str]) -> None:
        self.hashmap = {}
        for i, value in enumerate(values):
            self.hashmap[str(i + 1)] = value

    def __len__(self) -> int:
        return len(self.hashmap)

    def display(self):
        buffer = []
        for i in range(1, len(self.hashmap) + 1):
            buffer.append(str(i) + ".) " + self.hashmap[str(i)])
        return "\n".join(buffer)

    def eval(self, v: str) -> str | None:
        for key, val in self.hashmap.items():
            if v == key or v == val:
                return val
        return None

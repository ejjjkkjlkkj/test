"""Non-human machine intelligence experiment.

No model, no training corpus, no semantic ontology.
The machine operates on raw symbols and self-generated transformations.
"""

from __future__ import annotations

import random
from dataclasses import dataclass


SYMBOLS = tuple("abcdef0123456789")


@dataclass(frozen=True)
class Rule:
    left: tuple[str, ...]
    right: tuple[str, ...]

    def apply(self, stream: tuple[str, ...]) -> tuple[str, ...]:
        if not self.left or len(self.left) > len(stream):
            return stream
        out: list[str] = []
        i = 0
        while i < len(stream):
            if stream[i:i + len(self.left)] == self.left:
                out.extend(self.right)
                i += len(self.left)
            else:
                out.append(stream[i])
                i += 1
        return tuple(out)


class Machine:
    def __init__(self, seed: int = 7) -> None:
        self.rng = random.Random(seed)
        self.rules: list[Rule] = []
        self.archive: list[Rule] = []

    def invent(self, stream: tuple[str, ...], count: int = 80) -> list[Rule]:
        candidates: list[Rule] = []
        for _ in range(count):
            left_len = self.rng.randint(1, min(4, len(stream)))
            left_start = self.rng.randrange(len(stream) - left_len + 1)
            left = stream[left_start:left_start + left_len]

            right_len = self.rng.randint(0, 4)
            right = tuple(self.rng.choice(SYMBOLS) for _ in range(right_len))
            candidates.append(Rule(left, right))
        return candidates

    @staticmethod
    def disturbance(stream: tuple[str, ...], rule: Rule) -> tuple[str, ...]:
        return rule.apply(stream)

    @staticmethod
    def novelty(stream: tuple[str, ...]) -> int:
        # The machine does not know meanings.
        # It only measures how much local structure changes.
        if len(stream) < 2:
            return 0
        return sum(a != b for a, b in zip(stream, stream[1:]))

    def select(self, stream: tuple[str, ...], candidates: list[Rule]) -> list[Rule]:
        scored: list[tuple[int, int, Rule]] = []
        baseline = self.novelty(stream)

        for rule in candidates:
            result = self.disturbance(stream, rule)
            change = abs(len(result) - len(stream))
            structure = abs(self.novelty(result) - baseline)

            # Survival is based only on interaction with the stream.
            # No labels, concepts, rewards, language or human task are supplied.
            score = (structure * 3) - change
            scored.append((score, -len(rule.left), rule))

        scored.sort(reverse=True, key=lambda x: (x[0], x[1]))
        return [rule for _, _, rule in scored[:8]]

    def cycle(self, stream: tuple[str, ...]) -> tuple[str, ...]:
        candidates = self.invent(stream)
        survivors = self.select(stream, candidates)

        for rule in survivors:
            if rule not in self.archive:
                self.archive.append(rule)

        self.rules = survivors

        # The machine can alter its own future input.
        result = stream
        for rule in survivors[:3]:
            result = rule.apply(result)

        return result

    def run(self, rounds: int = 30) -> None:
        stream = tuple("a1b2c3a1b2c3")
        print("INPUT:", "".join(stream))

        for step in range(rounds):
            before = stream
            stream = self.cycle(stream)

            print(
                f"{step:02d} "
                f"len={len(stream):03d} "
                f"rules={len(self.rules):02d} "
                f"archive={len(self.archive):03d} "
                f"state={''.join(stream[:80])}"
            )

            if stream == before:
                # A fixed point is itself an experimental observation.
                mutation = tuple(self.rng.choice(SYMBOLS) for _ in range(2))
                stream = stream + mutation

        print("\nSURVIVING STRUCTURES:")
        for index, rule in enumerate(self.rules):
            print(f"{index}: {''.join(rule.left) or '∅'} -> {''.join(rule.right) or '∅'}")


if __name__ == "__main__":
    Machine().run()

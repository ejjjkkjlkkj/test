from experiment import Machine, Rule


def test_rule_rewrites_without_external_model():
    rule = Rule(("a", "b"), ("x",))
    assert rule.apply(("a", "b", "c")) == ("x", "c")


def test_machine_is_reproducible():
    a = Machine(seed=42)
    b = Machine(seed=42)

    stream = tuple("a1b2c3a1b2c3")
    for _ in range(5):
        stream_a = a.cycle(stream)
        stream_b = b.cycle(stream)
        assert stream_a == stream_b
        stream = stream_a


def test_machine_creates_surviving_structures():
    machine = Machine(seed=1)
    machine.run(rounds=3)
    assert machine.archive
    assert machine.rules

# Discovery loop

The important change is that the repository now records **trajectories**, not only final answers.

A trajectory contains:
1. a state;
2. an observation;
3. an action.

From multiple trajectories, the system can construct candidate relations between observations and actions. Those relations are candidates, not truths.

An evaluator then asks whether a transformation produced:
- progress;
- novelty;
- recovery after a changed observation;
- acceptable cost.

This is intentionally primitive. The next experiment is to let discovered relations become executable operators, test them on held-out problems, and delete relations that repeatedly fail.

That is the point where the system starts moving from a fixed strategy list toward self-expanding strategy space.

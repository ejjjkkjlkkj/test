(** ZERO semantic safety model.
    This file proves properties of the abstract core model.
    It deliberately does NOT claim refinement of the current C implementation yet.
*)

Record State := {
  pub : nat;
  secret : nat;
  budget : nat
}.

Definition admissible (s : State) : Prop := budget s > 0.

Inductive step : State -> State -> Prop :=
| step_run :
    forall p sec b,
      b > 0 ->
      step
        {| pub := p; secret := sec; budget := b |}
        {| pub := S p; secret := sec; budget := pred b |}.

Theorem step_preserves_secret :
  forall s s',
    step s s' ->
    secret s' = secret s.
Proof.
  intros s s' H.
  inversion H.
  reflexivity.
Qed.

Theorem step_preserves_public_independence :
  forall s1 s2 s1' s2',
    pub s1 = pub s2 ->
    step s1 s1' ->
    step s2 s2' ->
    pub s1' = pub s2'.
Proof.
  intros s1 s2 s1' s2' Hpub H1 H2.
  inversion H1; inversion H2.
  simpl in *.
  now rewrite Hpub.
Qed.

Theorem step_consumes_budget :
  forall s s',
    step s s' ->
    budget s' < budget s.
Proof.
  intros s s' H.
  inversion H.
  simpl.
  apply Nat.pred_lt.
  assumption.
Qed.

Theorem step_requires_budget :
  forall s s',
    step s s' ->
    admissible s.
Proof.
  intros s s' H.
  inversion H.
  unfold admissible.
  assumption.
Qed.

Theorem no_secret_flow_to_public :
  forall s1 s2 s1' s2',
    pub s1 = pub s2 ->
    step s1 s1' ->
    step s2 s2' ->
    pub s1' = pub s2'.
Proof.
  exact step_preserves_public_independence.
Qed.

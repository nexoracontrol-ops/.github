# NEXORA

**Deterministic Economic Control Plane**

NEXORA is an infrastructure project for proving, authorizing, executing, verifying, and reconciling economic state transitions across existing systems.

## Identity

- Organization: [nexoracontrol-ops](https://github.com/nexoracontrol-ops)
- Identity: [identity/identity.json](https://github.com/nexoracontrol-ops/.github/blob/main/identity/identity.json)
- Evidence index: [evidence/index.json](https://github.com/nexoracontrol-ops/.github/blob/main/evidence/index.json)
- Evidence schema: [schemas/evidence-v0.1.schema.json](https://github.com/nexoracontrol-ops/.github/blob/main/schemas/evidence-v0.1.schema.json)
- Verifier: [tools/nxverify](https://github.com/nexoracontrol-ops/.github/tree/main/tools/nxverify)

## Trust model

NEXORA separates:

- **Identity** — who publishes an artifact
- **Evidence** — what was observed or produced
- **Authority** — who is allowed to authorize a transition
- **State** — what the Kernel has committed
- **Provenance** — how an artifact can be reproduced and attributed

Cryptographic proof does not by itself establish legal ownership or contractual authority. Those claims are represented separately.

## Architecture boundary

```
Sources
  -> Evidence
  -> Canonical Event
  -> Deterministic Rules
  -> Authority
  -> Kernel Decision
  -> State Transition
  -> Receipt
  -> Verification
  -> Settlement
```

AI may propose. The Kernel is the authority for state transition.

## Evidence policy

Public evidence MUST be:

1. addressable,
2. versioned,
3. hashable,
4. attributable,
5. independently verifiable where applicable.

Customer data, credentials, secrets, and private evidence are never published merely to increase transparency.

## Status

Identity & Evidence Layer v0.1 is the initial public provenance baseline. It is an engineering specification, not a certification of legal status, security, or economic performance.


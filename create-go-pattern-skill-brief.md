---
title: Create Go Pattern Skill Brief
---

# Create Go Pattern Skill Brief

Use this brief to create a Go design-pattern skill from a working pattern in
this repository. The output is a reusable skill in
`/Users/kevinronu/.agents/skills`, not a repository-specific implementation
guide.

## Inputs

Provide these concrete inputs:

- `Guideline file`: `<path-to-guideline-file>`
- `Repository root`: `<path-to-this-repo>`
- `Pattern family`: `<creational|structural|behavioral>`
- `Pattern name`: `<pattern-name>`
- `Reference implementation path`: `<path-to-pattern-implementation-in-this-repo>`
- `Output draft path`: `<path-where-the-new-draft-should-be-created>`

Provide one or two reference pattern skills only when their structure clarifies
the target. They are examples, not a requirement to duplicate their wording or
load every available skill.

## Goal and Boundaries

Create a lightweight, file-by-file Go pattern skeleton that:

- is primarily loaded by `go` after `design-pattern-decision` selects it;
- is also useful for a direct request to implement, scaffold, or inspect that
  pattern in Go;
- captures the important non-obvious Go adaptation present in the reference
  implementation;
- generalizes the structural lesson without carrying over demo business names.

Do not turn the skill into a router, decision tree, validation or testing
guide, repository manual, or duplicate of `go` or `design-pattern-decision`.
Those skills retain orchestration, validation, and pattern-classification
responsibility.

## Read and Extract

Read the guideline file and the target implementation before drafting. Read a
reference skill only when it helps resolve a genuine structural question.

Extract the decisions that affect how the pattern is expressed in Go:

- package layout and file naming when they make roles or dependency direction
  clear;
- shared types and contracts placed together to prevent duplication or import
  cycles;
- selectors, factory lookup points, and interface boundaries;
- embedding or composition that carries shared state or behavior;
- a helper that substitutes for behavior inherited from an abstract base class;
- zero-value concrete types that need no construction or dependency wiring;
- import aliases that materially improve clarity;
- pointer versus value use when it is inherent to the pattern; and
- constraints that keep concrete variants interchangeable.

Preserve a small but important implementation trick even when it needs a tiny
generalization. Simplify incidental demo detail, such as formatting needed only
for printed output. Do not invent structure unsupported by the implementation.

## Create the Skill

Create `go-pattern-<pattern-name>/SKILL.md`. Add
`references/skeleton.md` when the folder shape and file-by-file skeleton would
make the entrypoint harder to scan; otherwise keep a small self-contained
skill. Do not add resource folders, scripts, templates, or examples without a
recurring concrete use.

### Frontmatter description

Use the final skill name in lowercase hyphen-case. Write a short,
discriminating description that says what the skill provides and when it
applies: the matching classifier result and direct requests to implement,
scaffold, or inspect the named Go pattern. Avoid broad adjacent topics or an
exhaustive feature list because descriptions are loaded for every available
skill.

### SKILL.md body

Keep the entrypoint focused on the pattern skeleton. It should contain:

1. a clear title and short purpose;
2. `Non-Obvious Go Notes` with only the decisions a textbook description would
   not reveal; and
3. a contextual link to the folder shape and skeleton when that reference
   exists.

The usual output shape is therefore YAML frontmatter, title, purpose,
non-obvious Go notes, and a skeleton reference. It is not a reason to force a
fixed section or extra file when the pattern is genuinely simpler.

Do not repeat generic Go testing, formatting, linting, validation-scope,
architecture, or error-handling rules. The `go` skill owns them.

### Folder shape and file-by-file skeleton

Use the reference to show the relevant folder shape and a concise code block
for each significant file. Use readable placeholders such as `<module>`,
`<pattern-root>`, `<family>`, `<type>`, `<productName>`, `factory`, `creator`,
`component`, or `adapter` when they improve reuse. Do not use placeholders that
only restate language facts.

Keep names meaningful and use short local variables when they make pattern
roles clearer than a deeply nested call. Show one representative template for
repeated files, then state the corresponding substitutions rather than copying
the same skeleton many times.

Keep the skeleton lightweight: preserve the pattern structure, interfaces,
ownership, and important Go adaptation, but remove business-specific language
and non-essential demo mechanics. Do not generalize until the useful lesson
disappears.

## Non-Obvious Go Notes

This section records only insights another LLM could miss from the textbook
pattern, for example:

- place shared identifiers beside the common contract to avoid circular
  dependencies;
- use a helper for behavior that would live in an abstract base class;
- return a stateless zero-value concrete type by value;
- return an interface from a selector rather than a concrete type;
- separate a client-owned port from adapters;
- embed or compose a shared bridge side instead of simulating inheritance;
- construct a dependent optional part in a recipe or caller; or
- keep the shared type and common contract in one package.

Do not add textbook definitions or notes that merely restate skeleton code.

## Comments

Default to no comment. Keep a comment only when it explains a pattern role,
ownership decision, caller consequence, or non-obvious Go adaptation that code
cannot express. Prefer the important explanation beside the relevant skeleton
over a trailing adaptation section.

Useful comments explain why a shared contract exists, why a helper replaces
abstract-base behavior, why a clone leaves a back-reference unset, why a
selector returns an interface, or why an accessor is the only construction
path. Do not narrate setters, getters, returns, assignments, or signatures.

## Final Check

Before finishing, confirm that the skill:

- is a reusable skeleton derived from a real Go implementation;
- is minimal without becoming vague;
- preserves the non-obvious Go lessons;
- has a narrow, clear trigger;
- does not overlap with `design-pattern-decision` or the generic `go` skill;
- does not make repository-adaptation decisions that belong to the caller; and
- remains understandable to an LLM with no prior repository context.

Run the available lightweight skill validator when it applies.

## Copyable Prompt Template

```text
Based on <path-to-guideline-file>, create a Go design-pattern skill from the
reference implementation below.

Inputs:
- Repository root: <path-to-this-repo>
- Pattern family: <creational|structural|behavioral>
- Pattern name: <pattern-name>
- Reference implementation: <path-to-pattern-implementation-in-this-repo>
- Output draft path: <path-where-the-new-draft-should-be-created>

Read the guideline and the target implementation. Read a supplied reference
skill only if it clarifies the target's structure.

Create a reusable Go pattern skeleton, primarily usable after
`design-pattern-decision` selects this pattern and also for direct requests to
implement, scaffold, or inspect it. Keep the frontmatter description short and
specific to that trigger. Keep SKILL.md compact; put a file-by-file skeleton in
references only when progressive disclosure makes it easier to use.

Preserve only non-obvious Go adaptations from the implementation: package and
interface boundaries, selectors, common identifiers, shared behavior or state,
zero-value types, and constraints that keep variants interchangeable. Generalize
away repository business names and incidental demo mechanics. Do not invent
unsupported structure or duplicate the generic `go` and
`design-pattern-decision` responsibilities.

Use comments only for rationale, ownership, caller consequences, or a Go
adaptation that code cannot express. Do not comment obvious code.

Before finishing, check that the result is concise, reusable, clear to an LLM
with zero context, and does not overlap with the `go` or
`design-pattern-decision` skills.
```

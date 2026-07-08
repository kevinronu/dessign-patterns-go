---
title: Create Go Pattern Skill Brief
---

# Create Go Pattern Skill Brief

Use this document as the full prompt base for an LLM that starts with zero context.

## Goal

Create a new Go design-pattern skill by using:

- a guideline file that defines how these skills should behave
- the pattern implementation that already exists in this repository

The result should be a new skill draft similar in spirit to the existing Go pattern skills, especially:

- lightweight
- file-by-file
- general enough to reuse outside this repo
- explicit enough to preserve the important non-obvious Go adaptations

## How To Use This Brief

Give this brief to the LLM together with these concrete inputs:

- `Guideline file`: `<path-to-guideline-file>`
- `Reference skills`: `<path-to-existing-go-pattern-skill-1>`, `<path-to-existing-go-pattern-skill-2>`
- `Repository root`: `<path-to-this-repo>`
- `Pattern family`: `<creational|structural|behavioral>`
- `Pattern name`: `<pattern-name>`
- `Reference implementation path`: `<path-to-pattern-implementation-in-this-repo>`
- `Output draft path`: `<path-where-the-new-draft-should-be-created>`

## Instructions For The LLM

Read the guideline file first. Then read the existing reference skills to match their level of abstraction, tone, and structure. After that, inspect the target pattern implementation in the repository.

Create the new skill draft from what is actually present in the repository implementation. Do not invent structure that the implementation does not support unless a tiny generalization is necessary to make the skill reusable.

## What The New Skill Must Be

The new skill must be a Go pattern skeleton.

It must:

- provide a reusable pattern skeleton, not a full implementation workflow
- serve primarily as something the `go` skill can load after `design-pattern-decision` selects the pattern
- still remain usable when a user explicitly asks for that Go pattern
- avoid taking orchestration responsibility away from the `go` skill
- avoid taking classification responsibility away from `design-pattern-decision`

## What The New Skill Must Not Do

Do not make the new skill:

- a router
- a decision tree
- a validation guide
- a testing guide
- a repo-specific implementation manual
- a duplicate of the `go` skill

Do not repeat generic Go rules that belong in the `go` skill, such as:

- testing conventions
- formatting steps
- linting steps
- validation scope
- broad architecture rules
- generic error-handling rules

Only keep pattern-specific guidance and the few non-obvious Go adaptations that are easy to lose.

## Required Output Shape

The new skill draft should follow this structure unless the guideline file requires a small adjustment:

1. YAML frontmatter
2. Title
3. Purpose
4. Non-Obvious Go Notes
5. Folder Shape
6. File-By-File Skeleton

Keep it compact.

## Frontmatter Rules

The frontmatter must:

- use the final skill name in lowercase hyphen-case
- describe the skill as a lightweight Go pattern skeleton
- say that it is used primarily after the active `go` skill receives the matching classification from `design-pattern-decision`
- also allow explicit direct use when the user asks to implement, scaffold, or inspect that pattern in Go
- mention that the skill contains only the folder shape, file-by-file code skeleton, and the few non-obvious Go adaptations worth preserving

## Content Rules

The body must:

- stay focused on skeleton and structure
- use placeholders like `<module>`, `<pattern-root>`, `<family>`, `<type>`, `<productName>`, or equivalent when that improves reuse
- keep placeholder names readable for another LLM
- avoid weird placeholder names that repeat language facts, such as suffixes that restate that something is already a struct
- prefer the lightest illustrative code that still preserves the pattern structure
- avoid carrying over non-essential demo detail from the repository when a simpler skeleton communicates the same pattern role
- preserve comments only when they explain something non-obvious
- keep the most important structural lessons inside `Non-Obvious Go Notes` and the code comments themselves, not only in a trailing explanatory section
- remove comments that merely narrate obvious code

## Repository Extraction Rules

While reading the repository implementation, explicitly look for:

- package layout choices
- file naming conventions
- shared types placed in common packages to avoid duplication
- selectors or factory lookup points
- interface boundaries
- any helper that substitutes for abstract base class behavior from other languages
- zero-value structs used as concrete factories, implementers, or holders when no state is needed
- import alias patterns that improve clarity
- any non-obvious adaptation made because Go lacks inheritance or abstract classes

If the repository uses a small but important trick, preserve that idea in the skeleton even if the final skill stays general.

If a repository detail is only incidental demo code, simplify it in the skeleton. For example, prefer a direct illustrative return value over extra formatting imports unless the formatting itself teaches something important about the pattern.

## Generalization Rules

Generalize the skill enough that it no longer depends on this repository's business meaning.

That means:

- remove business-specific names
- remove repo-specific domain language
- keep only the structural lesson
- preserve file and folder shape when that shape is relevant to the pattern

Do not over-generalize to the point that the useful implementation lessons disappear.

## File-By-File Skeleton Rules

For each file:

- show the file path as a heading
- provide a concise code block
- keep names meaningful
- use placeholders only where reuse benefits from them
- prefer short local variables when they make the pattern roles easier to follow for another LLM, instead of inlining everything into one call

When adding comments inside code examples, prefer this order:

1. shared contract comments
2. comments that explain a structural role inside the pattern
3. comments that preserve a non-obvious Go adaptation

Avoid comments that merely restate what a setter, getter, return, or simple assignment already says.

If multiple files follow the same pattern, show one representative template and then explain briefly that the same shape should be repeated with the corresponding placeholder replacements.

## Naming Rules

Prefer stable, general names taken from the pattern itself, for example:

- `family`
- `product`
- `factory`
- `type`
- `creator`
- `component`
- `adapter`

Only use more specific names if the repository pattern clearly teaches a better generic abstraction.

## Non-Obvious Go Notes Rules

This section is important.

Only include insights that another LLM could miss if it only knew the textbook pattern, for example:

- where to place shared type identifiers to avoid circular dependencies
- how to represent behavior that would live in an abstract base class in another language
- when a stateless concrete type can stay as a zero-value struct returned by value
- why a selector returns an interface instead of a concrete type
- why a client-owned port and client package stay separate from adapters
- why embedding or composition carries the shared side of a bridge instead of an abstract base class
- why a dependent optional part is constructed in the recipe or caller before entering a builder chain
- why one package owns both a shared type and the common contract

Do not fill this section with textbook definitions.

## Commenting Rules

Comments in the code skeleton should explain pattern role, ownership, or a Go-specific adaptation.

Good kinds of comments:

- why a shared contract exists
- why a helper stands in for an abstract base class behavior from another language
- why a child clone leaves a back-reference unset
- why an accessor is the only public construction path
- why a selector returns an interface

Avoid comments like:

- "SetPartA sets PartA"
- "Clone clones the item"
- "Return the instance"
- "Build returns the product"

If a section such as `Adaptation Notes` exists, use it only when it adds something that would otherwise duplicate neither `Non-Obvious Go Notes` nor the code comments. Prefer not to create that extra section by default.

## Quality Bar

The new draft should feel like:

- a reusable pattern skeleton
- derived from a real Go implementation
- minimal but not vague
- specific but not repo-bound

## Final Self-Check

Before finishing, verify all of this:

- the skill does not overlap with `design-pattern-decision`
- the skill does not overlap with the generic `go` skill more than necessary
- the skill is not making pattern-selection decisions
- the skill is not making repo-adaptation decisions that belong to the caller
- the skill preserves the non-obvious lessons from the repository implementation
- the skill stays readable for an LLM with zero prior context
- the skill is concise

## Copyable Prompt Template

Use this prompt with the placeholders filled in:

```text
Based on the guidelines in <path-to-guideline-file>, create a new Go design-pattern skill draft.

Read these reference skills first so the new draft matches their style and level of abstraction:
- <path-to-existing-go-pattern-skill-1>
- <path-to-existing-go-pattern-skill-2>

Then inspect this repository pattern implementation:
- Repository root: <path-to-this-repo>
- Pattern family: <creational|structural|behavioral>
- Pattern name: <pattern-name>
- Reference implementation path: <path-to-pattern-implementation-in-this-repo>

Create the draft at:
- <path-where-the-new-draft-should-be-created>

Requirements:
- The result must be a lightweight Go pattern skeleton, not a router, decision tree, validation guide, or repo-specific manual.
- The result must be primarily usable by the `go` skill after `design-pattern-decision` selects this pattern, but it must also remain usable if the user explicitly asks for this pattern in Go.
- Do not duplicate generic Go testing, validation, formatting, linting, or architecture rules that belong to the `go` skill.
- Preserve the non-obvious Go adaptations from the repository implementation.
- Keep the result compact.
- Keep the output structure as:
  1. YAML frontmatter
  2. Title
  3. Purpose
  4. Non-Obvious Go Notes
  5. Folder Shape
  6. File-By-File Skeleton
- Use readable placeholders where reuse benefits from them.
- Keep comments only when they explain pattern role, ownership, or a non-obvious Go adaptation.
- Prefer to place the important explanatory comments next to the relevant skeleton code instead of relying on a trailing `Adaptation Notes` section.

Before finishing, self-check for overlap with `design-pattern-decision` and the generic `go` skill.
```

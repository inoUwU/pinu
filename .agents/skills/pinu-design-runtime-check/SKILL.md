---
name: pinu-design-runtime-check
description: Verify and refresh PINU local design work by starting the frontend dev server, checking client and admin routes in the in-app browser, opening .pen files in Pencil, and reconciling runtime behavior with design.pen. Use when Codex needs to reconnect frontend and Pencil after restart, confirm whether localhost routes actually render, detect blank-screen regressions, or update design.pen and design.preview.pen from local observations.
---

# PINU Design Runtime Check

## Overview

Use this skill to re-establish the local frontend plus Pencil workflow for PINU and record what the running app actually shows before updating wireframes.

## Workflow

### 1. Load the project truth first

- Read `.github/instructions/general.instructions.md` before touching the app.
- Read `docs/spec/list_of_screens.md` and `docs/spec/reverse_engineered_system_spec.md` before trusting route implementations.
- Read route-specific docs when the task mentions QR flow, session flow, checkout, or admin operations.

### 2. Start or reconnect the frontend

- Work from `E:\Github\pinu\frontend`.
- Check whether `pnpm` is available on `PATH`.
- If `pnpm` is missing, use `corepack pnpm`.
- If `corepack pnpm` fails because it needs to write under the user profile, rerun with escalation.
- Start the dev server in the background and log to:
  - `frontend/.playwright/dev-server.log`
  - `frontend/.playwright/dev-server.err.log`
- Wait until the log shows `http://localhost:3000` before opening the browser.

Recommended launch pattern:

```powershell
Start-Process -FilePath 'C:\Program Files\nodejs\corepack.cmd' `
  -ArgumentList 'pnpm','dev' `
  -WorkingDirectory 'E:\Github\pinu\frontend' `
  -WindowStyle Hidden `
  -RedirectStandardOutput 'E:\Github\pinu\frontend\.playwright\dev-server.log' `
  -RedirectStandardError 'E:\Github\pinu\frontend\.playwright\dev-server.err.log' `
  -PassThru
```

Useful checks:

```powershell
pnpm --version
corepack pnpm --version
Get-Content -Tail 60 'E:\Github\pinu\frontend\.playwright\dev-server.log'
```

### 3. Verify runtime in the in-app browser

- Prefer the in-app browser workflow over shell Playwright CLI when the task is about local UI verification.
- If the `playwright-cli` command is missing, treat that as non-blocking and switch to the in-app browser runtime.
- Reuse one browser tab across route checks unless a route specifically needs a fresh tab.
- Capture screenshots for visual truth.
- Capture DOM snapshots only when you need to confirm whether a blank screen is truly empty or only visually hidden.
- Inspect console logs on blank or suspicious routes.

Bootstrap pattern for low-level browser access through `node_repl`:

```js
if (!globalThis.agent) {
  const { setupAtlasRuntime } = await import("C:/Users/ino/.codex/plugins/cache/openai-bundled/browser-use/0.1.0-alpha1/scripts/browser-client.mjs");
  await setupAtlasRuntime({ globals: globalThis, backend: "iab" });
}
await agent.browser.nameSession("🔎 pinu route audit");
if (typeof tab === "undefined") {
  globalThis.tab = await agent.browser.tabs.new();
}
```

Check these route groups unless the task narrows the scope:

- Client:
  - `/`
  - `/category`
  - `/<categoryid>`
  - `/<categoryid>/<mid>`
  - `/order`
  - `/history`
  - `/checkout`
- Admin:
  - `/admin/login`
  - `/admin`
  - `/admin/menu`
  - `/admin/operation`
  - `/admin/bill`
  - `/admin/user`
  - `/admin/qrcode`
  - `/admin/settings`

Record one of these states per route:

- Rendered and matches code intent
- Rendered but scaffold only
- Effectively blank at runtime
- Rendered but mismatched with docs or code intent

Watch for recurring runtime warnings:

- Next.js dynamic params accessed synchronously
- Duplicate React keys
- Script tags rendered inside React trees

### 4. Reconnect Pencil and inspect the .pen file

- Open the target file explicitly with `mcp__pencil__.open_document`.
- Use `get_editor_state`, `batch_get`, and `get_screenshot` before editing.
- Inspect top-level frames and reusable shells before changing layout structure.
- If Pencil shows stale content for a file that already changed on disk, create a same-content copy such as `design.preview.pen`, open that file, and verify there.
- Use `design.preview.pen` only as a cache-bypass artifact. Keep `design.pen` as the canonical file.
- Prefer Pencil edit tools when the document is responsive.
- If Pencil is disconnected or clearly stale, edit the `.pen` JSON directly, then reopen or duplicate the file for verification.

### 5. Update the wireframe from both docs and runtime

- Keep the docs as the specification source.
- Keep current code structure as implementation intent.
- Add explicit runtime annotations with an absolute date, for example `2026-04-29 local runtime`.
- Distinguish these states clearly inside the wireframe:
  - Spec only
  - Scaffold only
  - Implemented
  - Runtime blank
  - Runtime mismatch
- Preserve reusable `mobile-shell` and `admin-shell` patterns unless the task explicitly asks for a redesign.

### 6. Validate before finishing

- If you edited the `.pen` file directly, confirm it still parses as JSON.
- If you edited the `.pen` file directly, confirm there are no duplicate IDs.
- Verify the overview board plus representative client and admin frames with `get_screenshot`.
- If both `design.pen` and `design.preview.pen` exist, confirm whether they are intentionally identical or whether preview is only a temporary cache-bypass copy.

## Decision Rules

- Do not trust code alone when the task is specifically about the current UI.
- Do not trust a blank screenshot alone when DOM snapshot or console logs can tell you whether the page is truly empty.
- Do not assume `pnpm` is available even if the user says local startup works; verify the command path in the current environment.
- Do not assume Pencil has reloaded a changed file just because the file changed on disk.

## Deliverable Pattern

When closing the task:

- State which routes were actually verified at runtime.
- State whether Pencil was editing the canonical file directly or whether `design.preview.pen` was used as a cache-bypass copy.
- State any unresolved mismatch between docs, code intent, and runtime rendering.

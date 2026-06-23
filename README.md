# devstash 🗃️

Your personal dev memory tool. Never Google the same installation command twice.

devstash helps you store and recall library installation commands and usage notes right from your terminal — organized, searchable, and always available.

---

## Installation

**If you have Go installed:**

```bash
go install github.com/RaucousDave/devstash@latest
```

**Mac/Linux (curl):**

```bash
curl -sSL https://github.com/RaucousDave/devstash/releases/latest/download/devstash-linux -o devstash
chmod +x devstash
sudo mv devstash /usr/local/bin/
```

**Windows (PowerShell):**

```powershell
Invoke-WebRequest -Uri https://github.com/RaucousDave/devstash/releases/latest/download/devstash.exe -OutFile devstash.exe
```

Or download the binary for your platform directly from the [releases page](https://github.com/RaucousDave/devstash/releases).

---

## Commands

### `devstash add [library]`

Add a new library to your stash.

```bash
devstash add drizzle
# prompts for install command and a label
```

Add a command to an existing library:

```bash
devstash add drizzle --command
# prompts for a label and the command
```

Update the description of an existing library:

```bash
devstash add drizzle --desc
# prompts for a new description
```

---

### `devstash get [library]`

Browse and retrieve commands for a specific library. Presents an interactive selector so you can pick the exact command you need.

```bash
devstash get drizzle
# shows a selectable list of saved commands

devstash get better-auth
devstash get gorilla/mux
```

---

### `devstash list`

Display all libraries currently saved in your stash.

```bash
devstash list
```

---

## Example Workflow

```bash
# add drizzle for the first time
devstash add drizzle
# Install command: npm install drizzle-orm drizzle-kit
# Label: install

# add more commands to drizzle later
devstash add drizzle --command
# Label: generate migration
# Command: npx drizzle-kit generate

devstash add drizzle --command
# Label: run migration
# Command: npx drizzle-kit migrate

devstash add drizzle --command
# Label: open studio
# Command: npx drizzle-kit studio

# retrieve a command when you need it
devstash get drizzle
# > install
#   generate migration
#   run migration
#   open studio

# see everything in your stash
devstash list
```

---

## How It Works

devstash stores everything locally in `~/.devstash/data.json` on your machine. No account, no internet connection required, no data leaves your computer.

```json
{
  "libraries": {
    "drizzle": {
      "install": "npm install drizzle-orm drizzle-kit",
      "description": "TypeScript ORM for PostgreSQL",
      "commands": [
        { "label": "generate migration", "cmd": "npx drizzle-kit generate" },
        { "label": "run migration", "cmd": "npx drizzle-kit migrate" },
        { "label": "open studio", "cmd": "npx drizzle-kit studio" }
      ]
    }
  }
}
```

---

## Tech Stack

- [Go](https://golang.org)
- [Cobra](https://github.com/spf13/cobra) — CLI framework
- [Huh](https://github.com/charmbracelet/huh) — interactive terminal forms
- [Bubbletea](https://github.com/charmbracelet/bubbletea) — terminal UI

---

## Contributing

Pull requests are welcome. For major changes, open an issue first to discuss what you'd like to change.

---

## License

MIT

# Customizing Your Gemini CLI Interaction

You can customize your interaction with me in a few ways to better suit your workflow.

## Ignoring Files with `.geminiignore`

You can tell me to ignore specific files and directories by creating a `.geminiignore` file in your project's root directory. This file uses the same syntax as `.gitignore`.

This is useful for preventing me from accessing or modifying sensitive information, build artifacts, or large dependency directories.

### Example `.geminiignore`:

```
# Ignore dependency directories
node_modules/
vendor/

# Ignore build output
dist/
build/
*.o

# Ignore log files
*.log

# Ignore local configuration
.env
```

## Saving Preferences with Instructions

You can ask me to remember specific preferences or facts about your projects or coding style. I use a `save_memory` tool for this.

For example, you can say:

- "Remember that I prefer using tabs over spaces."
- "Please remember that the main branch for this project is `develop`."
- "Save this fact: I use `pytest` for all my Python projects."

I will store this information to use in our future interactions, helping me provide more personalized and relevant assistance.

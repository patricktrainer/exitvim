# ExitVim

🚧 Under active construction. As it is, you can exit more than just Vim. 

Welcome to `ExitVim` – the ultimate solution to one of life's greatest mysteries: How do I exit Vim? 

**Are you stuck in Vim?** Do not panic! We've all been there. You opened a file in Vim, and now you feel like you're in a digital escape room. `ExitVim` is here to save your day (and your sanity).

## How it Works

`ExitVim` is a small, but mighty Go program that hunts down every instance of Vim running on your system and sends them a polite (but firm) `SIGTERM`. 

## Features

- **Find and Exit**: Automatically finds all running instances of Vim and terminates them.
- **Safe and Sound**: Sends a `SIGTERM`, which is a safe way to terminate processes. It's like asking nicely rather than pulling the plug.
- **Simple Logging**: Keeps you informed about what's happening under the hood.

## Installation

You don't need to install Vim to use `ExitVim` (that's the problem we're solving, remember?). Just clone this repository, build the program, and you're set to free yourself from Vim's clutches.

```bash
git clone https://github.com/your-repo/exitvim.git
cd exitvim
go build
```

## Usage

```bash
./exitvim
```
And voila! You're free from Vim's grasp. It's like an emergency exit for your terminal.

## Disclaimer

ExitVim is a tongue-in-cheek solution to exiting Vim. Please use it responsibly and remember that [closing Vim normally](https://stackoverflow.com/a/11828573) is always the best practice.

## Contributing

Feel free to contribute to this life-saving project. Fork the repo, make your changes, and open a pull request. Together, we can make the world a safer place for text editors.

Happy Exiting!

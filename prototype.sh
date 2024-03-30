fd --extension bdf | fzf --preview='monobit-banner "The quick brown fox jumps over the lazy dog" --font={} | tr "." " " | tr "@" "#"' --preview-window=up

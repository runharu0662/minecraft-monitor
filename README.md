# minecraft-monitor

Lightweight Go monitor for a private Minecraft Java Edition server over Tailscale, with Discord webhook alerts.

## Features
- Minecraft Java status protocol (not just TCP port opening)
- Check once on startup, then every 60 seconds
- Alert after 3 consecutive failures, if the server was previously reachable
- Notify once on recovery
- No RCON/SSH access required

## Setup (Raspberry Pi)

Requires Go and Tailscale network access to the Minecraft server.

1. `go mod tidy`
2. Create a Discord incoming webhook in your notification channel.
3. Set environment vars in your shell (never commit secrets):
   - `MC_ADDRESS=100.70.61.2:25565`
   - `DISCORD_WEBHOOK_URL=<your webhook URL>`
4. `go run ./cmd/monitor`

When the server is offline at launch, an alert is sent after 3 failed checks. Restarting the process resets in-memory alert state.

## GitHub private repo (from local checkout)

```
git init
git add .
git commit -m 'feat: initial Minecraft monitor'
gh repo create runharu0662/minecraft-monitor --private --source=. --remote=origin --push
```

Do not commit Webhook URLs even in a private repository. Keep them in local environment or protected systemd EnvironmentFile.

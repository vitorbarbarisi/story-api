# Configuração de Rotação e Limpeza de Logs

Este documento explica como configurar e usar os scripts de rotação e limpeza automática de logs para o Story API.

## 📋 Visão Geral

Os scripts gerenciam automaticamente:
- **Rotação de logs**: Evita crescimento descontrolado dos arquivos de log
- **Limpeza automática**: Remove logs antigos e recursos Docker não utilizados
- **Compatibilidade**: Funciona no Linux (systemd/cron) e macOS (launchd)

## 🚀 Instalação

### Linux (VM/Servidor)
```bash
# Executar como root
sudo ./setup_log_rotation
```

### macOS (Desenvolvimento)
```bash
# Executar como usuário normal
./setup_log_rotation
```

## ⚙️ Configurações Aplicadas

### 🐳 **Docker Daemon**
- **Limite por arquivo**: 10MB
- **Arquivos por container**: 3
- **Rotação automática**: Quando atinge o limite

### 📁 **Rotação de Logs**

#### Linux
| Tipo | Localização | Rotação | Retenção |
|------|-------------|---------|----------|
| Docker containers | `/var/lib/docker/containers/` | Diária | 7 dias |
| Sistema/Aplicação | `/var/log/story-api/` | Diária | 30 dias |
| Systemd | `/var/log/story-api-systemd.log` | Diária | 14 dias |

#### macOS
| Tipo | Localização | Rotação | Retenção |
|------|-------------|---------|----------|
| LaunchAgent | `~/Library/Logs/story-api*.log` | 10MB | 7 dias |
| Sistema | Automática | Diária | 7 dias |

### 🔄 **Limpeza Automática**

#### Linux (Cron Jobs)
- **02:00 diariamente**: Limpeza de containers, imagens e volumes
- **03:00 domingos**: Limpeza completa do sistema Docker

#### macOS (LaunchAgent)
- **Diariamente**: Rotação de logs e limpeza Docker
- **Automático**: Baseado em tamanho e idade dos arquivos

## 🛠️ Comandos Úteis

### Verificar Status
```bash
# Linux
sudo systemctl status docker
crontab -l | grep story-api
ls -la /etc/logrotate.d/story-api*

# macOS
launchctl list | grep story-api
ls -la ~/Library/LaunchAgents/com.story-api*
```

### Executar Limpeza Manual
```bash
# Linux
sudo /usr/local/bin/cleanup-story-api-docker
sudo logrotate -f /etc/logrotate.d/story-api-docker

# macOS
~/bin/cleanup-story-api-logs
~/bin/cleanup-story-api-docker
```

### Monitorar Logs
```bash
# Verificar tamanho dos logs
du -sh /var/lib/docker/containers/*/  # Linux
du -sh ~/Library/Logs/story-api*      # macOS

# Ver últimas execuções
journalctl -u docker                  # Linux
cat ~/Library/Logs/story-api.log      # macOS
```

## 📊 Políticas de Retenção

### Configuração Padrão
```bash
# Docker
max-size: 10MB
max-file: 3

# Sistema
rotate: 7-30 dias
compress: sim
```

### Customização
Para alterar as políticas, edite:
- **Linux**: `/etc/logrotate.d/story-api-*`
- **macOS**: `~/bin/cleanup-story-api-*`

## 🔧 Resolução de Problemas

### Docker Daemon não Reinicia
```bash
# Linux
sudo systemctl restart docker
sudo systemctl status docker

# Verificar configuração
sudo docker info | grep -i log
```

### Logrotate não Funciona
```bash
# Testar configuração
sudo logrotate -d /etc/logrotate.d/story-api-docker

# Forçar execução
sudo logrotate -f /etc/logrotate.d/story-api-docker
```

### LaunchAgent não Carrega (macOS)
```bash
# Recarregar
launchctl unload ~/Library/LaunchAgents/com.story-api.logrotate.plist
launchctl load ~/Library/LaunchAgents/com.story-api.logrotate.plist

# Verificar status
launchctl list | grep story-api
```

## 🗑️ Remoção Completa

Para remover todas as configurações:
```bash
# Linux
sudo ./remove_log_rotation

# macOS
./remove_log_rotation
```

## 📈 Monitoramento

### Verificar Espaço em Disco
```bash
# Uso geral
df -h

# Docker específico
docker system df
```

### Logs de Limpeza
```bash
# Linux
grep "story-api" /var/log/cron.log
journalctl -u crond | grep cleanup

# macOS
cat ~/Library/Logs/story-api.log
```

## ⚠️ Notas Importantes

1. **Backup**: Sempre mantenha backups de logs importantes antes da limpeza
2. **Monitoramento**: Monitore regularmente o espaço em disco
3. **Customização**: Ajuste as políticas conforme necessário
4. **Teste**: Teste as configurações após instalação

## 📚 Recursos Adicionais

- [Docker Logging](https://docs.docker.com/config/containers/logging/)
- [Linux Logrotate](https://linux.die.net/man/8/logrotate)
- [macOS LaunchAgent](https://developer.apple.com/library/archive/documentation/MacOSX/Conceptual/BPSystemStartup/Chapters/CreatingLaunchdJobs.html)

---

> **Dica**: Execute os scripts regularmente para manter o sistema otimizado e livre de logs desnecessários! 🚀 
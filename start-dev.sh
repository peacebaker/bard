#!/bin/bash

# constants
DEV_DIR="$HOME/projects/secretSocial"
LOG_DIR="/var/log/secretSocial"

# start up mongodb
sudo systemctl start mongodb
if [[ $? -ne 0 ]]; then
    echo "error: mongodb failed to start"
    exit 1
fi

# create/chown log filedir
sudo mkdir -p "$LOG_DIR"
user=$(whoami)
sudo chown $user:$user "$LOG_DIR"

# move to the dev dir, stop any running services, and clear the bash log
cd "$DEV_DIR"
bash stop-dev.sh
echo -n > "$LOG_DIR/bash.log"

# start network services
go run network/atlas/atlas.go &>> "$LOG_DIR/bash.log" &

# wait for network services to start and then start backend service; in production, this will be handled by systemd
sleep 3
go run backend/backend.go &>> "$LOG_DIR/bash.log" &

# start the alpha neighborhood services
go run neighborhood/guard/guard.go &>> "$LOG_DIR/bash.log" &
go run neighborhood/sus/sus.go &>> "$LOG_DIR/bash.log" &

# finally, start the frontend
cd frontend
npm run dev &>> "$LOG_DIR/bash.log" &
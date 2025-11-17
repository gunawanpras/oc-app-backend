#!/bin/bash
set -e

if [ ! -f /opt/oracle/dbs/initORCLCDB.ora ]; then
  echo "initORCLCDB.ora missing — recreating..."
  export ORACLE_SID=ORCLCDB
  mkdir -p /opt/oracle/dbs
  echo "DB_NAME='ORCLCDB'" > /opt/oracle/dbs/initORCLCDB.ora
  echo "SPFILE='/opt/oracle/oradata/ORCLCDB/spfileORCLCDB.ora'" >> /opt/oracle/dbs/initORCLCDB.ora
else
  echo "initORCLCDB.ora already exists — skipping."
fi

echo "Running user-defined SQL scripts in /opt/oracle/scripts/setup"
for f in /opt/oracle/scripts/setup/*.sql; do
    echo "Executing $f"
    $ORACLE_HOME/bin/sqlplus / as sysdba @"$f"
done

chmod +x startup-scripts/99_create_init_param.sh
# .bash_profile

# Get the aliases and functions
if [ -f ~/.bashrc ]; then
	. ~/.bashrc
fi

# User specific environment and startup programs

PATH=$PATH:$HOME/bin

export PATH
export LANG="en_US.UTF-8"
echo -e "\nOf course it runs on $(uname -s)\n"
CPUTIME=$(ps -eo pcpu | awk 'NR>1' | awk '{tot=tot+$1} END {print tot}')
CPUCORES=$(cat /proc/cpuinfo | grep -c processor)
echo "
    System Summary (collected `date`)
    - CPU Usage (average)               = `echo $CPUTIME / $CPUCORES | bc`%
    - Memory free (real)                = `free -m | head -n 2 | tail -n 1 | awk {'print $4'}` Mb
    - Memory free (cache)               = `free -m | head -n 3 | tail -n 1 | awk {'print $3'}` Mb
    - Swap in use                       = `free -m | tail -n 1 | awk {'print $3'}` Mb
    - System Uptime                     =`uptime`"
for i in `df -P | awk '{print $6}' | grep -v "Mounted"`
    do
    echo "    - Disk Space Used                   = `df -P $i | awk '{ a = $5 } END {print a}'` ("$i")"
done


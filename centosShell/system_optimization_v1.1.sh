#!/bin/bash
# # # # # # # # # # # # # # # # # # # #
####                               ####
####         系 统 初 始 化        ####
####                               ####
####       Write by: Geoffrey      ####
####       Date: 2019-12-04        ####
####       Version: 1.1            ####
# # # # # # # # # # # # # # # # # # # #

cwd="$(cd "$(dirname "$0")"; pwd)"

check_system_version () {
#检查系统版本
if [ -z $(cat /etc/redhat-release | awk '{print $3}' | awk -F '.' '{print $1}' | sed 's/[0-9]//g') ];then
    system_version=`cat /etc/redhat-release | awk '{print $3}' | awk -F '.' '{print $1}'`
elif [ -z $(cat /etc/redhat-release | awk '{print $4}' | awk -F '.' '{print $1}' | sed 's/[0-9]//g') ];then
    system_version=`cat /etc/redhat-release | awk '{print $4}' | awk -F '.' '{print $1}'`
else
    echo "请确认系统是否为CentOS"
    exit
fi
}

replace_repo () {
#替换yum源
repo_dir="/etc/yum.repos.d"
[ ! -d "$repo_dir/bak" ] | mkdir -p $repo_dir/bak

yum install -y yum-utils.noarch  >> /dev/null 2>&1
if [ $? == 0 ];then
    echo "网络测试正常"
else
    echo "请检查网络是否正常"
    exit
fi

repo_dir_file=`ls -l $repo_dir | grep repo | awk '{print $9}'`

for i1 in $repo_dir_file
    do
    mv ${repo_dir}/${i1} ${repo_dir}/bak
    echo "源文件"${i1}"已备份完成"
done

for i2 in `ls ${cwd}/${system_version}/repo/`
    do
    yes | cp ${cwd}/${system_version}/repo/${i2} ${repo_dir}
    echo "源文件"${i2}"已添加"
done

echo "清除缓存"
yum clean all #清除缓存
echo "生成新的缓存"
yum makecache fast #重新生成缓存
}

install_basic_software () {
#安装基础工具
yum -y install gcc gcc-c++ vim vim-enhanced \
    sysstat tree nmap iftop telnet cmake \
    perl perl-devel net-tools kernel-devel \
    htop nmon lrzsz iotop traceroute bc \
    crontabs htop nmon sysbench \
    zlib-devel zlib openssl openssl-devel \
    patch bash-completion lsof unzip zip
}

sync_public_time () {
#配置ntpdate自动对时"
yum -y install ntp ntpdate
echo "* */1 * * * root /usr/sbin/ntpdate -u cn.ntp.org.cn >> /dev/null 2>&1" >> /etc/crontab
ntpdate -u cn.ntp.org.cn
#/etc/init.d/crond restart
}

system_parameter_optimization () {
echo "配置文件的ulimit值"
ulimit -SHn 65535
echo "ulimit -SHn 65535" >> /etc/rc.local
cat >> /etc/security/limits.conf << EOF
*               soft	nofile          65535
*               hard	nofile          65535
EOF
cat >> /etc/pam.d/login << EOF
session    required     pam_limits.so
EOF
echo "基础系统内核优化"
cat >> /etc/sysctl.conf << EOF
net.ipv4.tcp_syncookies = 1
net.ipv4.tcp_syn_retries = 1
net.ipv4.tcp_tw_recycle = 1
net.ipv4.tcp_tw_reuse = 1
net.ipv4.tcp_fin_timeout = 1
net.ipv4.tcp_keepalive_time = 1200
net.ipv4.tcp_max_tw_buckets = 65535
net.ipv4.route.gc_timeout = 100
## nginx优化
net.ipv4.ip_forward = 1
net.ipv4.conf.default.rp_filter = 1
net.ipv4.conf.default.accept_source_route = 0
kernel.sysrq = 0
kernel.core_uses_pid = 1
kernel.msgmnb = 65536
kernel.msgmax = 65536
kernel.shmmax = 68719476736
kernel.shmall = 4294967296
net.ipv4.tcp_sack = 0
net.ipv4.tcp_dsack = 0
net.ipv4.tcp_window_scaling = 1
net.ipv4.tcp_rmem = 32768 87380 8388608
net.ipv4.tcp_wmem = 32768 87380 8388608
net.core.wmem_default = 8388608
net.core.rmem_default = 8388608
net.core.rmem_max = 16777216
net.core.wmem_max = 16777216
net.core.netdev_max_backlog = 262144
net.core.somaxconn = 262144
net.ipv4.tcp_max_orphans = 3276800
net.ipv4.tcp_max_syn_backlog = 262144
net.ipv4.tcp_timestamps = 0
net.ipv4.tcp_synack_retries = 1
net.ipv4.tcp_mem = 94500000 915000000 927000000
net.ipv4.ip_local_port_range = 1024 65535
EOF
/sbin/sysctl -p

#删除多余的用户和组
userdel adm
userdel lp
userdel sync
userdel shutdown
userdel halt
userdel news
userdel uucp
userdel operator
userdel games
userdel gopher
userdel ftp
groupdel lp
groupdel news
groupdel uucp
groupdel games
groupdel dip
groupdel pppusers

#关闭SELinux
sed -i 's@SELINUX=enforcing@SELINUX=disabled@' /etc/selinux/config
#禁止空密码登录
sed -i 's@#PermitEmptyPasswords no@PermitEmptyPasswords no@' /etc/ssh/sshd_config 
#禁止SSH反向解析
sed -i 's@#UseDNS yes@UseDNS no@' /etc/ssh/sshd_config
#解决Linux之间使用SSH远程连接慢的问题
sed -i 's@#GSSAPIAuthentication no@GSSAPIAuthentication no@' /etc/ssh/sshd_config
sed -i 's@GSSAPIAuthentication yes@GSSAPIAuthentication no@' /etc/ssh/sshd_config

#隐藏Linux版本信息显示
> /etc/issue
> /etc/issue.net

/usr/bin/chattr -R +i /etc/cron*
/usr/bin/chattr -R +i /var/spool/cron

chmod +x /etc/rc.local
#
#vim基础语法优化
cat >> /etc/vimrc << EOF
set number
set ruler
set nohlsearch
set shiftwidth=4
set tabstop=4
set expandtab
set cindent 
set autoindent
set mouse=v
syntax on
set cursorline
set softtabstop=4
set statusline=\ %<%F[%1*%M%*%n%R%H]%=\ %y\ %0(%{&fileformat}\ %{&encoding}\ %c:%l/%L%)\
set foldmethod=syntax
set smartindent
EOF
}

update () {
echo "####################################################"
echo "##                                                ##"
echo "Friendly Tips:"
echo "    This system is $(cat /etc/redhat-release)."
echo "##                                                ##"
echo "####################################################"
read -p "Do you want to upgrade this system?
    please input [Y]es or [N]o. " input
case $input in
    Y | y) echo
        echo "This system will be upgrade!"
        echo "Please wait..."
        yum -y upgrade;;
    N | n) echo
        echo "This system will not upgarde."
        ;;
    *)
        echo "Input Error!"
        exit
        ;;
esac
}


centos_6_command () {
/etc/init.d/ntpd restart
sleep 1
/etc/init.d/iptables stop
sleep 1
chkconfig iptables off
sleep 1
chkconfig ip6tables off
sleep 1
/bin/cp -f ${cwd}/${system_version}/bashrc /etc/
/bin/cp -f ${cwd}/${system_version}/motd /etc/
/bin/cp -f ${cwd}/${system_version}/.bash_profile /root/
/bin/cp -f ${cwd}/${system_version}/my.cnf /etc/
}

centos_7_command () {
systemctl stop firewalld
sleep 1
systemctl disable firewalld
sleep 1
systemctl restart ntpd
sleep 1
/usr/bin/cp -f ${cwd}/${system_version}/bashrc /etc/
/usr/bin/cp -f ${cwd}/${system_version}/motd /etc/
/usr/bin/cp -f ${cwd}/${system_version}/.bash_profile /root/
/usr/bin/cp -f ${cwd}/${system_version}/my.cnf /etc/
}

all_judge_system_version_function () {
#判断系统版本
if [ $system_version == 6 ];then
    centos_6_command
elif [ $system_version == 7 ];then
    centos_7_command
else
    exit
fi
}

is_or_no_reboot () {
read -p "Do you want to reboot this system now? [Y/N]" answer
    case $answer in
        Y | y)
            echo "Attention! This system will be reboot!"
            reboot
            ;;
        N | n)
            echo "Please reboot yourself later!"
            exit
            ;;
        *)
            echo "Input Error!"
            exit
            ;;
    esac
}

main () {
check_system_version
sleep 1
replace_repo
sleep 1
install_basic_software
sleep 1
sync_public_time
sleep 1
system_parameter_optimization
sleep 1
all_judge_system_version_function
sleep 1
update
sleep 1
is_or_no_reboot
}

main
#重启服务器
#reboot

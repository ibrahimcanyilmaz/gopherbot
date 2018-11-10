#!/bin/bash -e

# ssh-admin.sh - shell plugin for managing the robot's ssh keypair
[ -z "$GOPHER_INSTALLDIR" ] && { echo "GOPHER_INSTALLDIR not set" >&2; exit 1; }
source $GOPHER_INSTALLDIR/lib/gopherbot_v1.sh

command=$1
shift

configure(){
	cat <<"EOF"
Help:
- Keywords: [ "ssh", "keygen", "key", "replace" "keypair" ]
  Helptext: [ "(bot), generate|replace keypair - create an ssh keypair for the robot" ]
- Keywords: [ "ssh" , "pubkey", "public", "key" ]
  Helptext: [ "(bot), (show) pubkey - dump the robot's public key" ]
CommandMatchers:
- Command: keypair
  Regex: '(?i:(generate|replace) keypair)'
- Command: pubkey
  Regex: '(?i:(show )?pubkey)'
EOF
}

if [ "$command" = "configure" ]
then
	configure
	exit 0
fi

if [ -z "$BOT_SSH_PHRASE" ]
then
	Say "\$BOT_SSH_PHRASE not set; try 'store parameter ssh-init BOT_SSH_PHRASE=<somethingreallylong>'"
	exit 0
fi

case $command in
	"keypair")
		ACTION=$1
		if [ -e $HOME/.ssh/id_rsa ]
		then
			if [ "$ACTION" = "replace" ]
			then
				rm -f $HOME/.ssh/id_rsa $HOME/.ssh/id_rsa.pub
			else
				Say "I've already got an ssh keypair - use 'replace keypair' to replace it"
				exit 0
			fi
		else
			mkdir -p $HOME/.ssh
			chmod 600 $HOME/.ssh
		fi
		BOT=$(GetBotAttribute name)
		/usr/bin/ssh-keygen -q -b 4096 -N "$BOT_SSH_PHRASE" -C "$BOT" -f $HOME/.ssh/id_rsa
		Say "Created"
		;;
	"pubkey")
		if [ ! -e $HOME/.ssh/id_rsa.pub ]
		then
			Say "I don't seem to have an ssh public key - use 'generate keypair' to create one"
		else
			Say -f "$(cat $HOME/.ssh/id_rsa.pub)"
		fi
		;;
esac

##
# Bash setup.
[ -z "$PS1" ] && return

# don't put duplicate lines in the history. See bash(1) for more options
export HISTCONTROL=ignoredups
# ... and ignore same sucessive entries.
export HISTCONTROL=ignoreboth

# check the window size after each command and, if necessary,
# update the values of LINES and COLUMNS.
shopt -s checkwinsize

# make less more friendly for non-text input files, see lesspipe(1)
[ -x /usr/bin/lesspipe ] && eval "$(lesspipe)"

# set variable identifying the chroot you work in (used in the prompt below)
if [ -z "$debian_chroot" ] && [ -r /etc/debian_chroot ]; then
    debian_chroot=$(cat /etc/debian_chroot)
fi

# set a fancy prompt (non-color, unless we know we "want" color)
case "$TERM" in
    xterm-color) color_prompt=yes;;
esac

force_colored_prompt=yes

if [ -n "$force_color_prompt" ]; then
  if [ -x /usr/bin/tput ] && tput setaf 1 >&/dev/null; then
	  color_prompt=yes
  else
	  color_prompt=
  fi
fi

#function get_knife_block_used {
  #knife_used=$(readlink ~/.chef/knife.rb | sed  -e 's/\.rb$//' -e 's/^.*knife-//')
  #echo -n ${knife_used}
#}


function _git_prompt() {
    local git_status="`git status -unormal 2>&1`"
    if ! [[ "$git_status" =~ Not\ a\ git\ repo ]]; then
        if [[ "$git_status" =~ nothing\ to\ commit ]]; then
            local ansi=42
        elif [[ "$git_status" =~ nothing\ added\ to\ commit\ but\ untracked\ files\ present ]]; then
            local ansi=43
        else
            local ansi=45
        fi
        if [[ "$git_status" =~ On\ branch\ ([^[:space:]]+) ]]; then
            branch=${BASH_REMATCH[1]}
            test "$branch" != master || branch='master'
        else
            # Detached HEAD.  (branch=HEAD is a faster alternative.)
            branch="(`git describe --all --contains --abbrev=4 HEAD 2> /dev/null ||
                echo HEAD`)"
        fi
        #echo -n '\[\e[0;37:'"$ansi"';1m\]'"$branch"'\[\033[00m\] '
        echo -en '\e[0;37;'"$ansi"';1m'"$branch"'\e[0m '
        #printf "%${COLUMNS}s\n" "'\[\e[0;37;'"$ansi"';1m\]'"$branch"'\[\e[0m\] '"
    fi
}

function parse_git_branch {
  git branch --no-color 2> /dev/null | sed -e '/^[^*]/d' -e 's/* \(.*\)/\1/'
}

function git-track {
  CURRENT_BRANCH=$(parse_git_branch)
  git config branch.$CURRENT_BRANCH.remote $1
  git config branch.$CURRENT_BRANCH.merge refs/heads/$CURRENT_BRANCH
}

function parse_git_branch_and_add_brackets {
  #if [ -d .git ] ; then
    git branch --no-color 2> /dev/null | sed -e '/^[^*]/d' -e 's/* \(.*\)/\ \[\1\]/'
  #fi
}

if [ "$color_prompt" = yes ]; then
  PS1='${debian_chroot:+($debian_chroot)}\[\033[01;32m\]\u@\h\[\033[00m\]:\[\033[01;34m\]\w\[\033[00m\]\$ '
else
  PS1='${debian_chroot:+($debian_chroot)}\u@\h:\w\$ '
fi
PS1='${debian_chroot:+($debian_chroot)}$(_git_prompt):\w\$ '
unset color_prompt force_color_prompt

# If this is an xterm set the title to user@host:dir
case "$TERM" in
xterm*|rxvt*)
    PROMPT_COMMAND='echo -ne "\033]0;${USER}@${HOSTNAME}: ${PWD/$HOME/~}\007"'
    ;;
*)
    ;;
esac

if [ -f ~/.bash_aliases ]; then
  . ~/.bash_aliases
fi

# enable color support of ls and also add handy aliases
if [ "$TERM" != "dumb" ] && [ -x /usr/bin/dircolors ]; then
  eval "`dircolors -b`"
  alias ls='ls --color=auto'
  alias grep='grep --color=auto'
  alias fgrep='fgrep --color=auto'
  alias egrep='egrep --color=auto'
fi

if [ -f /etc/bash_completion ]; then
  . /etc/bash_completion
fi

for file in `ls ~/.bashrc.d/`
do
	source ~/.bashrc.d/${file}
done

[[ -s "$HOME/.rvm/scripts/rvm" ]] && source "$HOME/.rvm/scripts/rvm" # Load RVM into a shell session *as a function*
source /home/krogebry/.rvm/scripts/rvm
#rvm use 1.9.3

set -o vi
export EDITOR=`which vim`
alias genpass="uuidgen|sed 's/-//g'"
git config --global color.ui true

alias pull="git pull origin master"
alias push="git push origin master"

alias d2="ssh 192.168.1.192"


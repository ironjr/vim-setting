set runtimepath^="/home/ironjr/.vim_go_runtime"

source /home/ironjr/.vim_go_runtime/vimrc/basic.vim
source /home/ironjr/.vim_go_runtime/vimrc/filetypes.vim
source /home/ironjr/.vim_go_runtime/vimrc/plugins.vim
source /home/ironjr/.vim_go_runtime/vimrc/extended.vim

try
  source /home/ironjr/.vim_go_runtime/custom_config.vim
catch
endtry

" Set vim theme
"set t_8f=[38;2;%lu;%lu;%lum
"set t_8b=[48;2;%lu;%lu;%lum
"set termguicolors
set t_Co=256

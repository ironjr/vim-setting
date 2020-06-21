" Startup scripts
filetype plugin on

" Custom indentation for pleasurable coding
set autoindent
set number
set tabstop=4
set shiftwidth=4
set softtabstop=0
set expandtab
set smarttab
set hlsearch

" Vim plugin manager
call plug#begin()
" Essential plugins
Plug 'scrooloose/nerdtree', { 'on': 'NERDTreeToggle' }
Plug 'scrooloose/nerdcommenter'
Plug 'Xuyuanp/nerdtree-git-plugin'
Plug 'vim-airline/vim-airline'
Plug 'vim-airline/vim-airline-themes'
Plug 'octol/vim-cpp-enhanced-highlight'
" Color schemes
Plug 'ayu-theme/ayu-vim'
Plug 'beigebrucewayne/Turtles'
Plug 'baines/vim-colorscheme-thaumaturge'
Plug 'chr4/jellygrass.vim'
Plug 'crusoexia/vim-monokai'
Plug 'sjl/badwolf'
Plug 'tomlion/vim-solidity'
call plug#end()

" NERDtree settings
map <C-n> :NERDTreeToggle<CR>
let g:NERDTreeDirArrowExpandable = "▸"
let g:NERDTreeDirArrowCollapsible = "▾"
let g:NERDTreeIndicatorMapCustom = {
    \ "Modified"  : "✹",
    \ "Staged"    : "✚",
    \ "Untracked" : "✭",
    \ "Renamed"   : "➜",
    \ "Unmerged"  : "═",
    \ "Deleted"   : "✖",
    \ "Dirty"     : "✗",
    \ "Clean"     : "✔︎",
    \ 'Ignored'   : '☒',
    \ "Unknown"   : "?"
    \ }

" NERDCommenter settings
let g:NERDSpaceDelims = 1     " default=0
let g:NERDCompactSexyComs = 1 " default=0
let g:NERDDefaultAlign = 'left'
let g:NERDCustomDelimiters = {
    \ 'c': {
    \   'left': '/**',
    \   'right': '*/'
    \   }
    \ }
let g:NERDCommentEmptyLines = 1
let g:NERDTrimTrailingWhitespace = 1

" Airline settings
let g:airline_powerline_fonts = 1

" Cpp-Enhanced-Highlight settings
let g:cpp_class_scope_highlight = 1      " default = 0
let g:cpp_member_variable_highlight = 1  " default = 0
let g:cpp_class_decl_highlight = 1       " default = 0
"let g:cpp_experimental_simple_template_highlight = 1 " default = 0
"let g:cpp_experimental_template_highlight = 1        " defualt = 0
let g:cpp_concepts_hightlight = 1        " default = 0
"let g:cpp_no_function_highlight = 1      " default = 0
"let c_no_curly_error = 1

" Set vim theme
set t_8f=[38;2;%lu;%lu;%lum
set t_8b=[48;2;%lu;%lu;%lum
"set termguicolors
set t_Co=256

"let ayucolor="light"
"let ayucolor="mirage"
"let ayucolor="dark"
"colorscheme ayu

"colorscheme jellygrass
"colorscheme thaumaturge
"colorscheme turtles
"colorscheme monokai
colorscheme badwolf

" For Windows Terminal
"hi Normal guibg=NONE ctermbg=NONE

" Vundle settings
set nocompatible
filetype off
set rtp+=~/.vim/bundle/Vundle.vim
call vundle#begin()
" Required for running Vundle
Plugin 'VundleVim/Vundle.vim'
" Autocompleter
Plugin 'Valloric/YouCompleteMe'
" GLSL syntax highlighter
Plugin 'tikhomirov/vim-glsl'
call vundle#end()
filetype plugin indent on

" Startup code
autocmd StdinReadPre * let s:std_in=1
"autocmd VimEnter * if argc() == 0 && !exists("s:std_in") | NERDTree | endif
autocmd VimEnter * if argc() == 1 && isdirectory(argv()[0]) && !exists("s:std_in") | exe 'NERDTree' argv()[0] | wincmd p | ene | endif

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
set bg=dark
set backspace=indent,eol,start

" Debug script; if you want to run this, mkdir $HOME/.log/vim first
" set nocompatible
" function! ToggleVerbose()
"     if !&verbose
"         set verbosefile=~/.log/vim/verbose.log
"         set verbose=15
"     else
"         set verbose=0
"         set verbosefile=
"     endif
" endfunction
" call ToggleVerbose()


" Vim plugin manager
call plug#begin()
" Essential plugins
Plug 'neoclide/coc.nvim', {'branch': 'release'}
Plug 'scrooloose/nerdtree', { 'on': 'NERDTreeToggle' }
Plug 'scrooloose/nerdcommenter'
Plug 'Xuyuanp/nerdtree-git-plugin'
" Plug 'tiagofumo/vim-nerdtree-syntax-highlight'
" Plug 'vim-airline/vim-airline'
" Plug 'vim-airline/vim-airline-themes'
Plug 'itchyny/lightline.vim'
Plug 'shinchu/lightline-gruvbox.vim'
Plug 'octol/vim-cpp-enhanced-highlight'
" Color schemes
Plug 'ayu-theme/ayu-vim'
Plug 'beigebrucewayne/Turtles'
Plug 'baines/vim-colorscheme-thaumaturge'
Plug 'chr4/jellygrass.vim'
Plug 'crusoexia/vim-monokai'
Plug 'sjl/badwolf'
Plug 'tomlion/vim-solidity'
Plug 'morhetz/gruvbox'
Plug 'joshdick/onedark.vim'
Plug 'mhartington/oceanic-next'
Plug 'tomasiser/vim-code-dark'
Plug 'jacoborus/tender.vim'
call plug#end()


" NERDtree settings
map <C-n> :NERDTreeToggle<CR>
" let g:NERDTreeDirArrowExpandable = "▸"
" let g:NERDTreeDirArrowCollapsible = "▾"
let g:NERDTreeGitStatusIndicatorMapCustom = {
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
let g:NERDTreeGitStatusPorcelainVersion = 1


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


" Cpp-Enhanced-Highlight settings
let g:cpp_class_scope_highlight = 1      " default = 0
let g:cpp_member_variable_highlight = 1  " default = 0
let g:cpp_class_decl_highlight = 1       " default = 0
"let g:cpp_experimental_simple_template_highlight = 1 " default = 0
"let g:cpp_experimental_template_highlight = 1        " defualt = 0
let g:cpp_concepts_hightlight = 1        " default = 0
"let g:cpp_no_function_highlight = 1      " default = 0
"let c_no_curly_error = 1


"""
" coc.nvim settings
"""
let g:coc_node_path = '/usr/local/bin/node'

" Always show the signcolumn, otherwise it would shift the text each time
" diagnostics appear/become resolved.
if has("nvim-0.5.0") || has("patch-8.1.1564")
    " Recently vim can merge signcolumn and number column into one
    set signcolumn=number
else
    set signcolumn=yes
endif

" Use tab for trigger completion with characters ahead and navigate.
" NOTE: Use command ':verbose imap <tab>' to make sure tab is not mapped
" by other plugin before putting this into your config.
inoremap <silent><expr> <TAB>
    \ pumvisible() ? "\<C-n>" :
    \ <SID>check_back_space() ? "\<TAB>" :
    \ coc#refresh()
inoremap <expr><S-TAB> pumvisible() ? "\<C-p>" : "\<C-h>"

function! s:check_back_space() abort
let col = col('.') - 1
return !col || getline('.')[col - 1]  =~# '\s'
endfunction

" Use <c-space> to trigger completion.
if has('nvim')
    inoremap <silent><expr> <c-space> coc#refresh()
else
    inoremap <silent><expr> <c-@> coc#refresh()
endif

" Make <CR> auto-select the first completion item and notify coc.nvim to
" format on enter, <cr> could be remapped by other vim plugin
inoremap <silent><expr> <cr> pumvisible() ? coc#_select_confirm()
    \: "\<C-g>u\<CR>\<c-r>=coc#on_enter()\<CR>"

" Use `[g` and `]g` to navigate diagnostics
" Use `:CocDiagnostics` to get all diagnostics of current buffer in
" location list.
nmap <silent> [g <Plug>(coc-diagnostic-prev)
nmap <silent> ]g <Plug>(coc-diagnostic-next)

" GoTo code navigation.
nmap <silent> gd <Plug>(coc-definition)
nmap <silent> gy <Plug>(coc-type-definition)
nmap <silent> gi <Plug>(coc-implementation)
nmap <silent> gr <Plug>(coc-references)

" Use K to show documentation in preview window.
nnoremap <silent> K :call <SID>show_documentation()<CR>

function! s:show_documentation()
    if (index(['vim','help'], &filetype) >= 0)
        execute 'h '.expand('<cword>')
    elseif (coc#rpc#ready())
        call CocActionAsync('doHover')
    else
        execute '!' . &keywordprg . " " . expand('<cword>')
    endif
endfunction

" Highlight the symbol and its references when holding the cursor.
autocmd CursorHold * silent call CocActionAsync('highlight')
" Symbol renaming.
nmap <Leader>rn <Plug>(coc-rename)

" Formatting selected code.
xmap <Leader>f  <Plug>(coc-format-selected)
nmap <Leader>f  <Plug>(coc-format-selected)

augroup mygroup
    autocmd!
    " Setup formatexpr specified filetype(s).
    autocmd FileType typescript,json setl formatexpr=CocAction('formatSelected')
    " Update signature help on jump placeholder.
    autocmd User CocJumpPlaceholder call CocActionAsync('showSignatureHelp')
augroup end

" Applying codeAction to the selected region.
" Example: `<Leader>aap` for current paragraph
xmap <Leader>a  <Plug>(coc-codeaction-selected)
nmap <Leader>a  <Plug>(coc-codeaction-selected)

" Remap keys for applying codeAction to the current buffer.
nmap <Leader>ac  <Plug>(coc-codeaction)
" Apply AutoFix to problem on the current line.
nmap <Leader>qf  <Plug>(coc-fix-current)

" Map function and class text objects
" NOTE: Requires 'textDocument.documentSymbol' support from the
" language server.
xmap if <Plug>(coc-funcobj-i)
omap if <Plug>(coc-funcobj-i)
xmap af <Plug>(coc-funcobj-a)
omap af <Plug>(coc-funcobj-a)
xmap ic <Plug>(coc-classobj-i)
omap ic <Plug>(coc-classobj-i)
xmap ac <Plug>(coc-classobj-a)
omap ac <Plug>(coc-classobj-a)

" Remap <C-f> and <C-b> for scroll float windows/popups.
if has('nvim-0.4.0') || has('patch-8.2.0750')
    nnoremap <silent><nowait><expr> <C-f> coc#float#has_scroll() ? coc#float#scroll(1) : "\<C-f>"
    nnoremap <silent><nowait><expr> <C-b> coc#float#has_scroll() ? coc#float#scroll(0) : "\<C-b>"
    inoremap <silent><nowait><expr> <C-f> coc#float#has_scroll() ? "\<c-r>=coc#float#scroll(1)\<cr>" : "\<Right>"
    inoremap <silent><nowait><expr> <C-b> coc#float#has_scroll() ? "\<c-r>=coc#float#scroll(0)\<cr>" : "\<Left>"
    vnoremap <silent><nowait><expr> <C-f> coc#float#has_scroll() ? coc#float#scroll(1) : "\<C-f>"
    vnoremap <silent><nowait><expr> <C-b> coc#float#has_scroll() ? coc#float#scroll(0) : "\<C-b>"
endif

" Use CTRL-S for selections ranges.
" Requires 'textDocument/selectionRange' support of language server.
nmap <silent> <C-s> <Plug>(coc-range-select)
xmap <silent> <C-s> <Plug>(coc-range-select)

" Add `:Format` command to format current buffer.
command! -nargs=0 Format :call CocAction('format')

" Add `:Fold` command to fold current buffer.
command! -nargs=? Fold :call     CocAction('fold', <f-args>)

" Add `:OR` command for organize imports of the current buffer.
command! -nargs=0 OR   :call     CocAction('runCommand', 'editor.action.organizeImport')

" Add (Neo)Vim's native statusline support.
" NOTE: Please see `:h coc-status` for integrations with external plugins that
" provide custom statusline: lightline.vim, vim-airline.
set statusline^=%{coc#status()}%{get(b:,'coc_current_function','')}

" Mappings for CoCList
" Show all diagnostics.
nnoremap <silent><nowait> <space>a  :<C-u>CocList diagnostics<cr>
" Manage extensions.
nnoremap <silent><nowait> <space>e  :<C-u>CocList extensions<cr>
" Show commands.
nnoremap <silent><nowait> <space>c  :<C-u>CocList commands<cr>
" Find symbol of current document.
nnoremap <silent><nowait> <space>o  :<C-u>CocList outline<cr>
" Search workspace symbols.
nnoremap <silent><nowait> <space>s  :<C-u>CocList -I symbols<cr>
" Do default action for next item.
nnoremap <silent><nowait> <space>j  :<C-u>CocNext<CR>
" Do default action for previous item.
nnoremap <silent><nowait> <space>k  :<C-u>CocPrev<CR>
" Resume latest coc list.
nnoremap <silent><nowait> <space>p  :<C-u>CocListResume<CR>


" Airline settings (gruvbox, onedark, oceanicnext, codedark, tender)
let g:airline_powerline_fonts = 1
let g:airline_theme = 'gruvbox'

" Lightline settings (gruvbox, onedark, oceanicnext, codedark, tender)
let g:lightline = {
    \ 'colorscheme': 'gruvbox',
    \ }


" Set vim theme
set t_8f=[38;2;%lu;%lu;%lum
set t_8b=[48;2;%lu;%lu;%lum
" for vim 7
set t_Co=256
set t_ut=
" for vim 8
if (has("termguicolors"))
    set termguicolors
endif

"let ayucolor="light"
"let ayucolor="mirage"
"let ayucolor="dark"
"colorscheme ayu

" colorscheme jellygrass
" colorscheme thaumaturge
" colorscheme turtles
" colorscheme monokai
" colorscheme badwolf

" New colorschemes added 210807
colorscheme gruvbox
" colorscheme onedark
" colorscheme OceanicNext
" colorscheme codedark
" colorscheme tender

" For Windows Terminal
"hi Normal guibg=NONE ctermbg=NONE
"hi LineNr guibg=NONE ctermbg=NONE
""" hi NonText guibg=NONE ctermbg=NONE
"hi EndOfBuffer guibg=NONE ctermbg=NONE


" Chrome browser terminal for GCP
nnoremap <Leader>w <C-w>
nnoremap <Leader>h <C-w>h
nnoremap <Leader>j <C-w>j
nnoremap <Leader>k <C-w>k
nnoremap <Leader>l <C-w>l


" Vundle settings
set nocompatible
filetype off
set rtp+=~/.vim/bundle/Vundle.vim
call vundle#begin()
" Required for running Vundle
Plugin 'VundleVim/Vundle.vim'
" Autocompleter
"Plugin 'Valloric/YouCompleteMe'
" GLSL syntax highlighter
Plugin 'tikhomirov/vim-glsl'
call vundle#end()
filetype plugin indent on


" Startup code
autocmd StdinReadPre * let s:std_in=1
"autocmd VimEnter * if argc() == 0 && !exists("s:std_in") | NERDTree | endif
autocmd VimEnter * if argc() == 1 && isdirectory(argv()[0]) && !exists("s:std_in") | exe 'NERDTree' argv()[0] | wincmd p | ene | endif

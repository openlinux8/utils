syntax on
filetype off
set rtp+=~/.vim/bundle/Vundle.vim
call vundle#begin()
Plugin 'VundleVim/Vundle.vim'
Plugin 'fatih/vim-go'
Plugin 'othree/html5.vim'
"Plugin 'hail2u/vim-css3-syntax'
call vundle#end()
filetype plugin indent on

autocmd FileType go setlocal omnifunc=gocomplete#Complete
autocmd FileType css setlocal omnifunc=csscomplete#CompleteCSS
autocmd FileType html,markdown setlocal omnifunc=htmlcomplete#CompleteTags
autocmd FileType javascript setlocal omnifunc=javascriptcomplete#CompleteJS
autocmd FileType typescript setlocal omnifunc=typescriptcomlete#CompleteTS

autocmd BufRead,BufNewFile *.css,*.js,*.html,*.tpl exec ":call SetIndent()"
func SetIndent()
set expandtab
set tabstop=2
set shiftwidth=2
set softtabstop=2
endfunc

set completeopt=menu
set foldmethod=indent
set nofoldenable

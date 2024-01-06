ctrl+shitp+p
在下拉框中输入settings.json
找到user的那个，default不可改
找到vim相关的配置，貌似其实写在哪里其实都行
加入下面的文本到settings.json
"vim.handleKeys": {
    "<C-d>": false,
    "<C-a>": false,
    "<C-c>": false,
    "<C-v>": false,
    "<C-x>": false,
    "<C-s>": false,
    "<C-z>": false
}